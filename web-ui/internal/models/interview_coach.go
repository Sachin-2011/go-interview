package models

import (
	"time"
)

// InterviewQuestion represents a single interview question with metadata
type InterviewQuestion struct {
	ID          string     `json:"id"`
	Question    string     `json:"question"`
	Explanation string     `json:"explanation"`
	Tags        []string   `json:"tags"`
	Category    string     `json:"category"` // "technical", "role-specific", "aptitude"
	IsPinned    bool       `json:"is_pinned"`
	PinnedAt    *time.Time `json:"pinned_at,omitempty"`
	GeneratedAt time.Time  `json:"generated_at"`
}

// InterviewSession represents a coaching session for a specific job
type InterviewSession struct {
	SessionID      string              `json:"session_id"`
	Username       string              `json:"username"`
	JobRole        string              `json:"job_role"`
	JobDescription string              `json:"job_description"`
	Questions      []InterviewQuestion `json:"questions"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}

// InterviewSessionCard is a lightweight summary used for dashboard cards
type InterviewSessionCard struct {
	SessionID      string    `json:"session_id"`
	JobRole        string    `json:"job_role"`
	JobDescription string    `json:"job_description"`
	TotalQuestions int       `json:"total_questions"`
	LastUpdatedAt  time.Time `json:"last_updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// ConversationMessage represents a message in the follow-up chat
type ConversationMessage struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"question_id"`
	Username   string    `json:"username"`
	Role       string    `json:"role"` // "user" or "assistant"
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
}

// QuestionConversation holds the conversation history for a question
type QuestionConversation struct {
	QuestionID string                `json:"question_id"`
	Messages   []ConversationMessage `json:"messages"`
}

// GenerateQuestionsRequest represents the request to generate questions
type GenerateQuestionsRequest struct {
	JobRole        string `json:"job_role"`
	JobDescription string `json:"job_description"`
	SessionID      string `json:"session_id,omitempty"` // For "Load More"
}

// GenerateQuestionsResponse represents the response with generated questions
type GenerateQuestionsResponse struct {
	SessionID string              `json:"session_id"`
	Questions []InterviewQuestion `json:"questions"`
	Success   bool                `json:"success"`
	Message   string              `json:"message,omitempty"`
}

// ListSessionsResponse represents a list of saved sessions for a user
type ListSessionsResponse struct {
	Sessions []InterviewSessionCard `json:"sessions"`
	Success  bool                   `json:"success"`
	Message  string                 `json:"message,omitempty"`
}

// SessionDetailsResponse returns a full session with questions
type SessionDetailsResponse struct {
	Session *InterviewSession `json:"session"`
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
}

// DeleteSessionResponse indicates deletion status
type DeleteSessionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// ExplainQuestionRequest represents the request for detailed explanation
type ExplainQuestionRequest struct {
	QuestionID string `json:"question_id"`
	SessionID  string `json:"session_id"`
	Question   string `json:"question"`
}

// ExplainQuestionResponse represents the detailed explanation response
type ExplainQuestionResponse struct {
	Explanation string `json:"explanation"`
	Success     bool   `json:"success"`
	Message     string `json:"message,omitempty"`
}

// FollowupQuestionRequest represents a follow-up question in the chat
type FollowupQuestionRequest struct {
	QuestionID       string `json:"question_id"`
	SessionID        string `json:"session_id"`
	FollowupMessage  string `json:"followup_message"`
	OriginalQuestion string `json:"original_question"`
}

// FollowupQuestionResponse represents the AI's response to follow-up
type FollowupQuestionResponse struct {
	Response string `json:"response"`
	Success  bool   `json:"success"`
	Message  string `json:"message,omitempty"`
}

// PinQuestionRequest represents request to pin/unpin a question
type PinQuestionRequest struct {
	QuestionID string `json:"question_id"`
	SessionID  string `json:"session_id"`
	Pin        bool   `json:"pin"` // true to pin, false to unpin
}

// PinQuestionResponse represents the pin operation response
type PinQuestionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
