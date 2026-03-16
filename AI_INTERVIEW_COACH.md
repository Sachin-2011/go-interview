# AI Interview Coach - Feature Documentation

## 🎯 Overview

The **AI Interview Coach** is a comprehensive chatbot-style interview preparation system integrated into the GoInterview platform. It generates personalized interview questions based on job descriptions and provides interactive learning through AI-powered explanations and follow-up conversations.

## 📋 Features

### 1. **Smart Question Generation**
- Generates 10 context-aware interview questions per session
- Question distribution:
  - 40% Technical questions (4 questions)
  - 30% Role-specific questions (3 questions)  
  - 30% Aptitude/Behavioral/Situational questions (3 questions)

### 2. **Rich Question Metadata**
Each generated question includes:
- Question text
- Short explanation (2-3 lines)
- Topic tags (e.g., Docker, CI/CD, SQL, Communication)
- Category classification
- Unique identifier for tracking

### 3. **Interactive Learning Interface**
- **Accordion UI**: Questions displayed in collapsible accordion format
- **Pin Functionality**: Pin important questions to the top
- **Load More**: Generate additional 10 questions for the same role
- **Learn More Panel**: Side drawer with detailed explanations

### 4. **AI-Powered Explanations**
When "Learn More" is clicked:
- Beginner-friendly explanation of the question
- Why interviewers ask this question
- Key concepts to understand
- Structured approach to answer
- Common mistakes to avoid

### 5. **Follow-up Chat System**
- Context-aware conversation with AI coach
- Ask unlimited follow-up questions
- Full conversation history maintained
- Smart, helpful responses in simple language

### 6. **Data Persistence**
All data stored on filesystem:
- User sessions by username
- Questions and metadata
- Pin status
- Conversation history

## 🔧 Technical Architecture

### Backend Components

#### **Models** (`internal/models/interview_coach.go`)
- `InterviewQuestion`: Question data structure
- `InterviewSession`: User session with job details
- `ConversationMessage`: Chat message structure
- `QuestionConversation`: Conversation thread per question
- Request/Response models for all API endpoints

#### **Service** (`internal/services/interview_coach.go`)
- `InterviewCoachService`: Core business logic
- `GenerateQuestions`: AI-powered question generation
- `ExplainQuestion`: Detailed explanations
- `HandleFollowup`: Context-aware chat responses
- `PinQuestion`: Pin/unpin functionality
- File-based persistence system

#### **API Handlers** (`internal/handlers/api.go`)
New endpoints:
- `POST /api/coach/generate` - Generate questions
- `POST /api/coach/explain` - Get detailed explanation
- `POST /api/coach/followup` - Handle follow-up chat
- `POST /api/coach/pin` - Pin/unpin questions

#### **Web Handler** (`internal/handlers/web.go`)
- `InterviewCoachPage`: Renders the coach UI

### Frontend Components

#### **Template** (`templates/interview_coach.html`)
Comprehensive UI with:
- Job role/description input form
- Loading states with spinners
- Questions accordion with Bootstrap
- Side panel for explanations
- Real-time chat interface
- Pin/unpin visual indicators
- Responsive design

#### **Styling Features**
- Gradient hero section
- Color-coded category badges
- Smooth transitions and animations
- Pinned question highlighting (yellow border)
- Side panel slide-in animation
- Chat message bubbles (user vs assistant)

#### **JavaScript Functionality**
- Async API calls with fetch
- Dynamic DOM manipulation
- State management (session ID, question tracking)
- Real-time UI updates
- Keyboard shortcuts (Escape to close panel)
- Auto-scroll in chat

## 🚀 Usage Flow

### 1. **Initial Setup**
User navigates to `/coach` from the main navigation

### 2. **Enter Job Details**
```
Job Role: Backend Developer
Job Description: [Paste full JD from LinkedIn/company website]
```

### 3. **Generate Questions**
Click "Generate Questions" → AI generates 10 mixed questions

### 4. **Review Questions**
- Browse accordion list
- See short explanations
- View topic tags
- Pin important ones

### 5. **Learn More**
Click "Learn More" on any question:
- Side panel opens
- AI provides detailed explanation
- Chat interface becomes available

### 6. **Ask Follow-ups**
Type follow-up questions in chat:
- "Can you explain this concept in simpler terms?"
- "What are real-world examples?"
- "How should I prepare for this?"

### 7. **Load More**
Click "Load 10 More" to generate additional questions for same role

## 📁 Data Storage Structure

```
interview-coach-data/
├── sessions/
│   └── {username}/
│       └── {session-id}.json
└── conversations/
    └── {question-id}.json
```

### Session File Example
```json
{
  "session_id": "uuid-here",
  "username": "johndoe",
  "job_role": "Backend Developer",
  "job_description": "Full JD text...",
  "questions": [
    {
      "id": "question-uuid",
      "question": "Explain microservices architecture",
      "explanation": "Short explanation...",
      "tags": ["Architecture", "Microservices", "System Design"],
      "category": "technical",
      "is_pinned": false,
      "generated_at": "2026-01-14T..."
    }
  ],
  "created_at": "2026-01-14T...",
  "updated_at": "2026-01-14T..."
}
```

### Conversation File Example
```json
{
  "question_id": "question-uuid",
  "messages": [
    {
      "id": "msg-uuid",
      "role": "assistant",
      "content": "Detailed explanation...",
      "timestamp": "2026-01-14T..."
    },
    {
      "id": "msg-uuid-2",
      "role": "user",
      "content": "Can you explain more about...",
      "timestamp": "2026-01-14T..."
    }
  ]
}
```

## 🔌 API Reference

### Generate Questions
```http
POST /api/coach/generate
Content-Type: application/json

{
  "job_role": "Backend Developer",
  "job_description": "Full JD text...",
  "session_id": "optional-for-load-more"
}
```

**Response:**
```json
{
  "session_id": "uuid",
  "questions": [...],
  "success": true
}
```

### Explain Question
```http
POST /api/coach/explain
Content-Type: application/json

{
  "question_id": "uuid",
  "session_id": "uuid",
  "question": "Question text"
}
```

**Response:**
```json
{
  "explanation": "Detailed explanation...",
  "success": true
}
```

### Follow-up Chat
```http
POST /api/coach/followup
Content-Type: application/json

{
  "question_id": "uuid",
  "session_id": "uuid",
  "original_question": "Question text",
  "followup_message": "User's follow-up question"
}
```

**Response:**
```json
{
  "response": "AI's answer...",
  "success": true
}
```

### Pin Question
```http
POST /api/coach/pin
Content-Type: application/json

{
  "question_id": "uuid",
  "session_id": "uuid",
  "pin": true
}
```

**Response:**
```json
{
  "success": true,
  "message": "Question pinned successfully"
}
```

## 🎨 UI/UX Features

### Visual Indicators
- **Pinned Questions**: Yellow left border + background highlight
- **Category Badges**: 
  - Technical: Blue
  - Role-specific: Purple
  - Aptitude: Green
- **New Feature Badge**: Pulsing yellow indicator in navigation

### Responsive Design
- Mobile-friendly accordion
- Adaptive side panel (600px wide on desktop)
- Touch-optimized buttons
- Smooth animations

### Accessibility
- Keyboard navigation support
- ARIA labels
- Screen reader friendly
- Focus management in modals

## 🧪 AI Integration

### Supported Providers
- **Gemini** (Google) - Default
- **OpenAI** (GPT-4o-mini)
- **Claude** (Anthropic)

### Configuration
Set in environment variables:
```bash
AI_PROVIDER=gemini
GEMINI_API_KEY=your-key-here
```

### Prompt Engineering
The service uses carefully crafted prompts:
- **Question Generation**: Enforces 40/30/30 distribution
- **Explanations**: Beginner-friendly, encouraging tone
- **Follow-ups**: Context-aware with conversation history

## 🔐 Security & Privacy

- Username-based isolation
- Server-side session management
- No sensitive data in client-side storage
- API key secure handling
- Input validation on all endpoints

## 📊 Future Enhancements

Potential additions:
- [ ] Export questions as PDF
- [ ] Share session with others
- [ ] Company-specific question banks
- [ ] Progress tracking across sessions
- [ ] Mock interview recording
- [ ] Voice-based interactions
- [ ] Question difficulty ratings
- [ ] Community-contributed questions

## 🐛 Troubleshooting

### Common Issues

**Questions not generating:**
- Check AI API key is set correctly
- Verify AI provider is configured
- Check browser console for errors

**Side panel not opening:**
- Ensure JavaScript is enabled
- Check for console errors
- Try refreshing the page

**Pin functionality not working:**
- Verify session ID is valid
- Check network tab for failed requests

## 📝 Development Notes

### Adding New Features
1. Update models in `interview_coach.go`
2. Add service methods
3. Create API handlers
4. Update frontend template
5. Test all endpoints

### Testing
Manual testing checklist:
- [ ] Generate questions for different roles
- [ ] Pin/unpin functionality
- [ ] Load more questions
- [ ] Learn more explanations
- [ ] Follow-up chat
- [ ] Multiple sessions
- [ ] Data persistence

## 🤝 Contributing

When adding features to AI Interview Coach:
1. Follow existing code patterns
2. Update this documentation
3. Test with all AI providers
4. Ensure mobile responsiveness
5. Add error handling

---

**Built with ❤️ for the GoInterview Platform**
