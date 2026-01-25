package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"web-ui/internal/models"

	"github.com/google/uuid"
)

// InterviewCoachService handles interview coaching operations
type InterviewCoachService struct {
	aiService     *AIService
	sessions      map[string]*models.InterviewSession
	conversations map[string]*models.QuestionConversation
	mutex         sync.RWMutex
	dataDir       string
}

// NewInterviewCoachService creates a new interview coach service
func NewInterviewCoachService(aiService *AIService) *InterviewCoachService {
	dataDir := filepath.Join("..", "interview-coach-data")
	os.MkdirAll(dataDir, 0755)

	service := &InterviewCoachService{
		aiService:     aiService,
		sessions:      make(map[string]*models.InterviewSession),
		conversations: make(map[string]*models.QuestionConversation),
		dataDir:       dataDir,
	}

	// Load existing sessions from disk
	service.loadSessionsFromDisk()

	return service
}

// GenerateQuestions generates interview questions based on job role and description
func (ics *InterviewCoachService) GenerateQuestions(username, jobRole, jobDescription, sessionID string) (*models.GenerateQuestionsResponse, error) {
	var session *models.InterviewSession

	ics.mutex.Lock()
	defer ics.mutex.Unlock()

	// Check if this is a "Load More" request
	if sessionID != "" {
		if existingSession, ok := ics.sessions[sessionID]; ok {
			session = existingSession
		} else {
			return nil, fmt.Errorf("session not found")
		}
	} else {
		// Create new session
		sessionID = uuid.New().String()
		session = &models.InterviewSession{
			SessionID:      sessionID,
			Username:       username,
			JobRole:        jobRole,
			JobDescription: jobDescription,
			Questions:      []models.InterviewQuestion{},
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		ics.sessions[sessionID] = session
	}

	// Generate questions using AI
	questions, err := ics.generateQuestionsWithAI(jobRole, jobDescription)
	if err != nil {
		return nil, fmt.Errorf("failed to generate questions: %v", err)
	}

	// Add questions to session
	session.Questions = append(session.Questions, questions...)
	session.UpdatedAt = time.Now()

	// Save to disk
	ics.saveSessionToDisk(session)

	return &models.GenerateQuestionsResponse{
		SessionID: sessionID,
		Questions: questions,
		Success:   true,
	}, nil
}

// ListSessions returns lightweight cards for a user's saved preparations
func (ics *InterviewCoachService) ListSessions(username string) ([]models.InterviewSessionCard, error) {
	ics.mutex.RLock()
	defer ics.mutex.RUnlock()

	var cards []models.InterviewSessionCard
	for _, session := range ics.sessions {
		if session.Username != username {
			continue
		}

		cards = append(cards, models.InterviewSessionCard{
			SessionID:      session.SessionID,
			JobRole:        session.JobRole,
			JobDescription: session.JobDescription,
			TotalQuestions: len(session.Questions),
			LastUpdatedAt:  session.UpdatedAt,
			CreatedAt:      session.CreatedAt,
		})
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].LastUpdatedAt.After(cards[j].LastUpdatedAt)
	})

	return cards, nil
}

// generateQuestionsWithAI uses AI to generate 10 interview questions
func (ics *InterviewCoachService) generateQuestionsWithAI(jobRole, jobDescription string) ([]models.InterviewQuestion, error) {
	prompt := fmt.Sprintf(`You are an expert interview coach. Generate exactly 10 interview questions for the following job role.

Job Role: %s

Job Description:
%s

Requirements:
- 40%% (4 questions) must be technical/coding questions
- 30%% (3 questions) must be role-specific questions
- 30%% (3 questions) must be aptitude/behavioral/situational questions

For each question, provide:
1. The question itself
2. A short 2-3 line explanation of what it tests
3. 2-4 topic tags (e.g., Docker, CI/CD, SQL, Communication, Problem Solving)
4. The category (technical, role-specific, or aptitude)

IMPORTANT: Return ONLY a valid JSON array. Do not add any markdown, code blocks, or explanatory text. The response must be valid JSON that can be parsed.

Example format:
[
  {
    "question": "Explain the SOLID principles in object-oriented design",
    "explanation": "Tests understanding of design patterns and software architecture principles",
    "tags": ["Architecture", "Design Patterns", "Best Practices"],
    "category": "technical"
  }
]

Generate 10 questions following this exact format.`, jobRole, jobDescription)

	// Call AI service with longer timeout
	response, err := ics.aiService.GenerateText(prompt, 4000)
	if err != nil {
		return nil, fmt.Errorf("AI generation failed: %v", err)
	}

	// Log full response for debugging
	fmt.Printf("\n=== RAW AI RESPONSE (length: %d) ===\n", len(response))
	fmt.Printf("%s\n", response)
	fmt.Printf("=== END RAW RESPONSE ===\n\n")

	// Parse JSON response
	response = strings.TrimSpace(response)

	// Clean up the response - remove markdown code blocks if present
	if strings.HasPrefix(response, "```json") {
		response = strings.TrimPrefix(response, "```json")
		response = strings.TrimSpace(response)
	}
	if strings.HasPrefix(response, "```") {
		// Remove opening backticks
		lines := strings.Split(response, "\n")
		if len(lines) > 1 {
			response = strings.Join(lines[1:], "\n")
		}
	}
	if strings.HasSuffix(response, "```") {
		response = strings.TrimSuffix(response, "```")
	}

	response = strings.TrimSpace(response)

	// Remove any leading/trailing quotes
	response = strings.Trim(response, "\"'")

	// Find the actual JSON array - handle cases where AI adds extra text
	if !strings.HasPrefix(response, "[") {
		// Find the start of JSON array
		idx := strings.Index(response, "[")
		if idx != -1 {
			response = response[idx:]
		} else {
			return nil, fmt.Errorf("no JSON array found in response")
		}
	}

	// Find the end of the JSON array
	if strings.Contains(response, "]") {
		lastIdx := strings.LastIndex(response, "]")
		response = response[:lastIdx+1]
	} else {
		return nil, fmt.Errorf("incomplete JSON array in response - response may have been cut off")
	}

	fmt.Printf("\n=== CLEANED JSON (length: %d) ===\n", len(response))
	fmt.Printf("%s\n", response)
	fmt.Printf("=== END CLEANED JSON ===\n\n")

	var rawQuestions []struct {
		Question    string   `json:"question"`
		Explanation string   `json:"explanation"`
		Tags        []string `json:"tags"`
		Category    string   `json:"category"`
	}

	err = json.Unmarshal([]byte(response), &rawQuestions)
	if err != nil {
		// Provide detailed error message
		fmt.Printf("\n!!! JSON PARSE ERROR !!!\n")
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Response length: %d\n", len(response))
		if len(response) > 0 {
			fmt.Printf("First 300 chars: %s\n", response[:min(300, len(response))])
			if len(response) > 300 {
				fmt.Printf("Last 300 chars: %s\n", response[len(response)-min(300, len(response)):])
			}
		}
		return nil, fmt.Errorf("failed to parse AI response as JSON: %v. This may be due to API rate limits or the response being cut off. Please try again", err)
	}

	if len(rawQuestions) == 0 {
		return nil, fmt.Errorf("AI returned empty question list - please try again")
	}

	// Convert to InterviewQuestion models
	questions := make([]models.InterviewQuestion, 0, len(rawQuestions))
	for _, rq := range rawQuestions {
		questions = append(questions, models.InterviewQuestion{
			ID:          uuid.New().String(),
			Question:    rq.Question,
			Explanation: rq.Explanation,
			Tags:        rq.Tags,
			Category:    rq.Category,
			IsPinned:    false,
			GeneratedAt: time.Now(),
		})
	}

	fmt.Printf("\n✓ Successfully generated %d questions\n\n", len(questions))
	return questions, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ExplainQuestion provides a detailed, beginner-friendly explanation
func (ics *InterviewCoachService) ExplainQuestion(sessionID, questionID, question string) (*models.ExplainQuestionResponse, error) {
	prompt := fmt.Sprintf(`You are a friendly and patient interview coach. A candidate needs help understanding this interview question:

Question: %s

Provide a simple, beginner-friendly explanation that covers:
1. What the question is really asking
2. Why interviewers ask this question
3. Key concepts to know
4. A structured approach to answer it
5. Common mistakes to avoid

Keep it conversational, clear, and encouraging. Use simple language suitable for beginners.`, question)

	response, err := ics.aiService.GenerateText(prompt, 1500)
	if err != nil {
		return nil, fmt.Errorf("failed to generate explanation: %v", err)
	}

	// Initialize conversation for this question
	ics.mutex.Lock()
	if _, exists := ics.conversations[questionID]; !exists {
		ics.conversations[questionID] = &models.QuestionConversation{
			QuestionID: questionID,
			Messages:   []models.ConversationMessage{},
		}
	}

	// Store the initial explanation
	ics.conversations[questionID].Messages = append(ics.conversations[questionID].Messages, models.ConversationMessage{
		ID:         uuid.New().String(),
		QuestionID: questionID,
		Username:   "",
		Role:       "assistant",
		Content:    response,
		Timestamp:  time.Now(),
	})
	ics.mutex.Unlock()

	// Save conversation
	ics.saveConversationToDisk(questionID)

	return &models.ExplainQuestionResponse{
		Explanation: response,
		Success:     true,
	}, nil
}

// HandleFollowup handles follow-up questions in the conversation
func (ics *InterviewCoachService) HandleFollowup(username, sessionID, questionID, originalQuestion, followupMessage string) (*models.FollowupQuestionResponse, error) {
	ics.mutex.Lock()
	conversation, exists := ics.conversations[questionID]
	if !exists {
		conversation = &models.QuestionConversation{
			QuestionID: questionID,
			Messages:   []models.ConversationMessage{},
		}
		ics.conversations[questionID] = conversation
	}

	// Add user message to conversation
	userMessage := models.ConversationMessage{
		ID:         uuid.New().String(),
		QuestionID: questionID,
		Username:   username,
		Role:       "user",
		Content:    followupMessage,
		Timestamp:  time.Now(),
	}
	conversation.Messages = append(conversation.Messages, userMessage)
	ics.mutex.Unlock()

	// Build context-aware prompt
	contextMessages := ics.buildConversationContext(conversation.Messages)
	prompt := fmt.Sprintf(`You are a helpful interview coach having a conversation with a candidate about this interview question:

Original Question: %s

Conversation history:
%s

Candidate's follow-up question: %s

Provide a clear, helpful response that:
1. Directly addresses their question
2. Uses simple, beginner-friendly language
3. Provides practical examples when relevant
4. Encourages further learning

Keep your response concise but thorough.`, originalQuestion, contextMessages, followupMessage)

	response, err := ics.aiService.GenerateText(prompt, 1000)
	if err != nil {
		return nil, fmt.Errorf("failed to generate followup response: %v", err)
	}

	// Add assistant response to conversation
	ics.mutex.Lock()
	assistantMessage := models.ConversationMessage{
		ID:         uuid.New().String(),
		QuestionID: questionID,
		Username:   "",
		Role:       "assistant",
		Content:    response,
		Timestamp:  time.Now(),
	}
	conversation.Messages = append(conversation.Messages, assistantMessage)
	ics.mutex.Unlock()

	// Save conversation
	ics.saveConversationToDisk(questionID)

	return &models.FollowupQuestionResponse{
		Response: response,
		Success:  true,
	}, nil
}

// buildConversationContext builds a text representation of conversation history
func (ics *InterviewCoachService) buildConversationContext(messages []models.ConversationMessage) string {
	var context strings.Builder
	for _, msg := range messages {
		if msg.Role == "user" {
			context.WriteString(fmt.Sprintf("Candidate: %s\n", msg.Content))
		} else {
			context.WriteString(fmt.Sprintf("Coach: %s\n", msg.Content))
		}
	}
	return context.String()
}

// PinQuestion pins or unpins a question
func (ics *InterviewCoachService) PinQuestion(sessionID, questionID string, pin bool) (*models.PinQuestionResponse, error) {
	ics.mutex.Lock()
	defer ics.mutex.Unlock()

	session, exists := ics.sessions[sessionID]
	if !exists {
		return nil, fmt.Errorf("session not found")
	}

	// Find and update the question
	found := false
	for i := range session.Questions {
		if session.Questions[i].ID == questionID {
			session.Questions[i].IsPinned = pin
			if pin {
				now := time.Now()
				session.Questions[i].PinnedAt = &now
			} else {
				session.Questions[i].PinnedAt = nil
			}
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("question not found")
	}

	// Sort questions: pinned ones first, then by generation time
	ics.sortSessionQuestions(session)

	session.UpdatedAt = time.Now()
	ics.saveSessionToDisk(session)

	return &models.PinQuestionResponse{
		Success: true,
		Message: "Question pinned successfully",
	}, nil
}

// sortSessionQuestions sorts questions with pinned ones first
func (ics *InterviewCoachService) sortSessionQuestions(session *models.InterviewSession) {
	// Simple bubble sort - pinned items to top
	for i := 0; i < len(session.Questions); i++ {
		for j := i + 1; j < len(session.Questions); j++ {
			// If j is pinned and i is not, swap
			if session.Questions[j].IsPinned && !session.Questions[i].IsPinned {
				session.Questions[i], session.Questions[j] = session.Questions[j], session.Questions[i]
			}
		}
	}
}

// GetSession retrieves a session by ID
func (ics *InterviewCoachService) GetSession(sessionID string) (*models.InterviewSession, error) {
	ics.mutex.RLock()
	defer ics.mutex.RUnlock()

	session, exists := ics.sessions[sessionID]
	if !exists {
		return nil, fmt.Errorf("session not found")
	}

	return session, nil
}

// DeleteSession removes a session permanently for the given user
func (ics *InterviewCoachService) DeleteSession(sessionID, username string) error {
	ics.mutex.Lock()
	defer ics.mutex.Unlock()

	session, exists := ics.sessions[sessionID]
	if !exists {
		return fmt.Errorf("session not found")
	}

	if session.Username != username {
		return fmt.Errorf("unauthorized")
	}

	delete(ics.sessions, sessionID)

	// Remove session file from disk
	sessionFile := filepath.Join(ics.dataDir, "sessions", username, fmt.Sprintf("%s.json", sessionID))
	if err := os.Remove(sessionFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete session file: %v", err)
	}

	return nil
}

// GetConversation retrieves conversation history for a question
func (ics *InterviewCoachService) GetConversation(questionID string) (*models.QuestionConversation, error) {
	ics.mutex.RLock()
	defer ics.mutex.RUnlock()

	conversation, exists := ics.conversations[questionID]
	if !exists {
		return &models.QuestionConversation{
			QuestionID: questionID,
			Messages:   []models.ConversationMessage{},
		}, nil
	}

	return conversation, nil
}

// File persistence methods

func (ics *InterviewCoachService) saveSessionToDisk(session *models.InterviewSession) error {
	sessionFile := filepath.Join(ics.dataDir, "sessions", session.Username, fmt.Sprintf("%s.json", session.SessionID))
	os.MkdirAll(filepath.Dir(sessionFile), 0755)

	data, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(sessionFile, data, 0644)
}

func (ics *InterviewCoachService) saveConversationToDisk(questionID string) error {
	conversation, exists := ics.conversations[questionID]
	if !exists {
		return nil
	}

	convFile := filepath.Join(ics.dataDir, "conversations", fmt.Sprintf("%s.json", questionID))
	os.MkdirAll(filepath.Dir(convFile), 0755)

	data, err := json.MarshalIndent(conversation, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(convFile, data, 0644)
}

func (ics *InterviewCoachService) loadSessionsFromDisk() {
	sessionsDir := filepath.Join(ics.dataDir, "sessions")
	filepath.Walk(sessionsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		var session models.InterviewSession
		if err := json.Unmarshal(data, &session); err != nil {
			return nil
		}

		ics.sessions[session.SessionID] = &session
		return nil
	})

	// Load conversations
	conversationsDir := filepath.Join(ics.dataDir, "conversations")
	filepath.Walk(conversationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		var conversation models.QuestionConversation
		if err := json.Unmarshal(data, &conversation); err != nil {
			return nil
		}

		ics.conversations[conversation.QuestionID] = &conversation
		return nil
	})
}
