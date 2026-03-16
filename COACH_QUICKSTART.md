# Quick Start Guide - AI Interview Coach

## 🚀 Getting Started in 3 Steps

### Step 1: Ensure AI Service is Configured

Make sure you have an AI API key set up. Check your `.env` file or set environment variables:

```bash
# Option 1: Gemini (Recommended - Default)
AI_PROVIDER=gemini
GEMINI_API_KEY=your_gemini_api_key_here

# Option 2: OpenAI
AI_PROVIDER=openai
OPENAI_API_KEY=your_openai_api_key_here

# Option 3: Claude
AI_PROVIDER=claude
CLAUDE_API_KEY=your_claude_api_key_here
```

### Step 2: Start the Server

```bash
cd web-ui
go run main.go
```

You should see:
```
Loading challenges...
Loading scoreboards...
Loading packages...
Initializing AI Interview Coach...
Server starting on http://localhost:8080
```

### Step 3: Access the Coach

1. Open your browser to http://localhost:8080
2. Click **"AI Interview Coach"** in the navigation (with yellow "New" badge)
3. Enter a job role (e.g., "Backend Developer")
4. Paste a job description
5. Click **"Generate Questions"**

## 🎯 Example Usage

### Sample Input

**Job Role:**
```
Backend Developer
```

**Job Description:**
```
We are looking for an experienced Backend Developer to join our team.

Responsibilities:
- Design and implement RESTful APIs
- Work with microservices architecture
- Optimize database queries and performance
- Collaborate with frontend team
- Write unit and integration tests

Requirements:
- 3+ years of Go/Python experience
- Strong understanding of Docker and Kubernetes
- Experience with PostgreSQL/MySQL
- Knowledge of CI/CD pipelines
- Excellent problem-solving skills
```

### Expected Output

The AI will generate 10 questions like:

**Technical (40%):**
1. "Explain how you would design a RESTful API for a user management system"
2. "What are the key differences between vertical and horizontal scaling?"
3. "How do you optimize slow database queries in production?"
4. "Describe your approach to implementing authentication in a microservices architecture"

**Role-Specific (30%):**
1. "How do you handle database migrations in a production environment?"
2. "Explain your experience with containerization using Docker"
3. "Describe a challenging bug you fixed in a backend system"

**Aptitude/Behavioral (30%):**
1. "How do you prioritize tasks when working on multiple features?"
2. "Describe a time when you had to collaborate with the frontend team on a complex feature"
3. "How do you approach learning a new technology or framework?"

## 💡 Feature Walkthrough

### Pin Important Questions
Click the **"Pin"** button on any question to move it to the top of your list.

### Get Detailed Explanations
Click **"Learn More"** to open a side panel with:
- Beginner-friendly explanation
- Why interviewers ask this
- Key concepts to know
- How to structure your answer
- Common mistakes

### Ask Follow-up Questions
In the side panel chat:
- "Can you give me a real-world example?"
- "What resources should I study?"
- "How would you answer this question?"

### Load More Questions
Click **"Load 10 More"** to generate additional questions for the same role.

## 🔍 Troubleshooting

### Issue: "Failed to generate questions"

**Solution:**
1. Check AI API key is set correctly
2. Verify internet connection
3. Check server logs for detailed error
4. Try a different AI provider

### Issue: Side panel not showing

**Solution:**
1. Clear browser cache
2. Check browser console for JavaScript errors
3. Refresh the page

### Issue: Chat responses are slow

**Solution:**
1. This is normal - AI responses can take 3-10 seconds
2. Check your AI provider's rate limits
3. Consider switching to a faster model

## 📊 Data Location

All interview coach data is stored in:
```
go-interview/
└── interview-coach-data/
    ├── sessions/
    │   └── {username}/
    │       └── {session-id}.json
    └── conversations/
        └── {question-id}.json
```

You can delete this folder to reset all coach data.

## 🎨 Customization

### Change Question Distribution

Edit `interview_coach.go` line ~80:
```go
Requirements:
- 40%% (4 questions) must be technical/coding questions
- 30%% (3 questions) must be role-specific questions
- 30%% (3 questions) must be aptitude/behavioral/situational questions
```

### Modify AI Response Style

Update the prompts in:
- `generateQuestionsWithAI()` - Question generation
- `ExplainQuestion()` - Explanations
- `HandleFollowup()` - Chat responses

## 🧪 Testing Checklist

- [ ] Generate questions for "Backend Developer"
- [ ] Generate questions for "Data Analyst"  
- [ ] Generate questions for "DevOps Engineer"
- [ ] Pin a question → verify it moves to top
- [ ] Click "Learn More" → verify side panel opens
- [ ] Ask follow-up question → verify response
- [ ] Load 10 more → verify new questions added
- [ ] Refresh page → verify session persists (if username set)

## 📱 Mobile Usage

The interface is fully responsive:
- Accordion works on touch devices
- Side panel is full-width on mobile
- Buttons are touch-optimized
- Text is readable on small screens

## 🔗 Integration Points

The AI Interview Coach integrates with:
- **AI Service**: Uses existing AIService for all providers
- **User Service**: Tracks sessions by username
- **Navigation**: New tab in main menu
- **Base Template**: Inherits styling and layout

## 🎓 Best Practices

1. **Be Specific**: More detailed job descriptions = better questions
2. **Use Real JDs**: Copy actual job postings for best results
3. **Pin Key Questions**: Keep track of challenging topics
4. **Ask Follow-ups**: Don't hesitate to ask for clarification
5. **Load More Wisely**: Quality over quantity - review before loading more

## 🚧 Known Limitations

- Session data is per-username (requires git config or cookie)
- No export functionality yet
- Questions are not saved to database (filesystem only)
- Chat history limited to current session
- No question difficulty ratings

## 🔜 Coming Soon

- Export questions as PDF/Markdown
- Share sessions with others
- Practice mode with timer
- Voice-based interactions
- Company-specific question banks

---

**Need help?** Open an issue on GitHub or check the full documentation in `AI_INTERVIEW_COACH.md`
