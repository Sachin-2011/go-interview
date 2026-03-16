# Program Description, Naming Conventions & Validations

**Project:** GoInterview - Interactive Go Learning Platform  
**Date:** January 14, 2026  
**Version:** 1.0

---

## Table of Contents

1. [Program Description](#program-description)
2. [System Architecture](#system-architecture)
3. [Naming Conventions](#naming-conventions)
4. [Validation Rules](#validation-rules)
5. [Error Handling](#error-handling)

---

## Program Description

### Overview

**GoInterview** is a comprehensive, interactive Go learning platform designed to help developers master the Go programming language through practical coding challenges, real-world package tutorials, and AI-powered interview preparation. The platform combines competitive programming challenges with guided learning paths and modern interview coaching to provide a complete learning ecosystem.

### Purpose & Goals

1. **Learning Through Practice** - Users solve 30 increasingly difficult Go challenges covering fundamentals to advanced topics
2. **Package Mastery** - Learn popular Go packages (Gin, GORM, Cobra, etc.) through structured, guided learning paths with real-world context
3. **Interview Preparation** - AI-powered interview coach generates role-specific questions with follow-up explanations and practice
4. **Community Engagement** - Leaderboards, scoreboards, and social features encourage healthy competition and knowledge sharing
5. **Personalized Learning** - Track progress, maintain attempt history, and earn achievements and badges

### Key Features

#### 1. Coding Challenges System
- **30 Challenges** - Categorized by difficulty (Beginner/Intermediate/Advanced)
- **Progressive Learning** - Build from basics to advanced patterns
- **Instant Feedback** - Automated testing with detailed output and execution times
- **Learning Resources** - Study materials, hints, and solution explanations included
- **Scoreboard Tracking** - Global leaderboards for competition and motivation

#### 2. Package Learning Paths
- **Structured Curriculum** - 6+ Go packages with guided learning paths
- **Real-World Context** - Learn when and why to use each package
- **Practical Projects** - Challenges within each package teach real-world usage
- **Package Details** - GitHub stars, documentation, categories, and prerequisites
- **Progress Tracking** - Monitor completion percentage and time spent

#### 3. AI Interview Coach
- **Smart Question Generation** - Creates 10 tailored interview questions based on job role and description
- **Balanced Distribution** - 40% technical, 30% role-specific, 30% aptitude questions
- **Multi-Provider AI** - Supports Gemini, OpenAI, and Claude with configurable selection
- **Interactive Learning** - Follow-up questions with full conversation context
- **Smart Pinning** - Pin important questions for focused practice

#### 4. User Management
- **GitHub Integration** - Users identified by GitHub username
- **Persistent Progress** - All submissions and progress saved in filesystem
- **Attempt History** - Track all submission attempts with timestamps
- **Achievement System** - Earn badges and achievements for milestones

### Technology Stack

**Backend:**
- Language: Go 1.21+
- Framework: Standard `net/http` with `http.ServeMux` routing
- Template Engine: Go `html/template`
- AI Integration: Multi-provider abstraction layer (Gemini, OpenAI, Claude)
- Data Persistence: Filesystem-based JSON storage
- External: `github.com/google/uuid` for unique IDs

**Frontend:**
- Framework: Bootstrap 5.3
- Editor: Monaco Editor (VS Code-like editing)
- Scripting: Vanilla JavaScript (no framework dependencies)
- Styling: CSS3 with Bootstrap utilities

**Deployment:**
- Platform: Railway.app (via template configuration)
- Docker: Containerized Go application
- Environment: Configured via `.env` file

### User Workflows

#### Challenge Workflow
1. User selects a challenge from list
2. Reads problem description, hints, and learning materials
3. Writes Go solution using provided template
4. Submits code for automated testing
5. Receives instant feedback (passed/failed tests, execution time)
6. Views position on scoreboard

#### Package Learning Workflow
1. User selects a package to learn
2. Views package overview (stars, docs, prerequisites)
3. Follows structured learning path with challenges
4. Completes challenges sequentially
5. Tracks progress and earns achievements
6. Competes on package-specific leaderboards

#### Interview Preparation Workflow
1. User enters job role and job description
2. AI generates 10 tailored interview questions (balanced across categories)
3. User reads each question with short explanation
4. User clicks "Learn More" to see detailed AI explanation
5. User engages in follow-up chat with full context
6. User pins important questions for later review
7. User generates 10 more questions ("Load More") for extended practice

---

## System Architecture

### Component Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                   FRONTEND (Browser)                        │
│  ┌──────────────────────────────────────────────────────┐   │
│  │           HTML Templates + Vanilla JS                │   │
│  │  - Challenges UI    - Package UI    - Coach UI       │   │
│  │  - Monaco Editor    - Leaderboards  - Chat Interface │   │
│  └──────────────────────────────────────────────────────┘   │
└────────────────┬─────────────────────────────────────────────┘
                 │ HTTP/REST API
┌────────────────▼─────────────────────────────────────────────┐
│                   BACKEND (Go Server)                        │
│  ┌──────────────────────────────────────────────────────┐   │
│  │               HTTP Handlers Layer                     │   │
│  │  - Challenge Handlers    - Package Handlers         │   │
│  │  - Interview Coach       - User Management          │   │
│  └──────────────────────────────────────────────────────┘   │
│  ┌──────────────────────────────────────────────────────┐   │
│  │               Services Layer                         │   │
│  │  - Challenge Service   - Package Service            │   │
│  │  - Coach Service       - AI Service                 │   │
│  └──────────────────────────────────────────────────────┘   │
│  ┌──────────────────────────────────────────────────────┐   │
│  │               Models Layer                           │   │
│  │  - Challenge  - Package  - Interview Coach          │   │
│  └──────────────────────────────────────────────────────┘   │
│  ┌──────────────────────────────────────────────────────┐   │
│  │          Filesystem Storage Layer                    │   │
│  │  - JSON Persistence  - File I/O Operations          │   │
│  └──────────────────────────────────────────────────────┘   │
└────────────────┬─────────────────────────────────────────────┘
                 │
        ┌────────┴────────┬──────────────┬──────────────┐
        │                 │              │              │
   ┌────▼────┐    ┌──────▼──────┐  ┌───▼──────┐  ┌────▼─────┐
   │ Local   │    │  Google     │  │ OpenAI  │  │  Claude  │
   │ Storage │    │  Gemini API │  │  API    │  │   API    │
   │         │    │             │  │         │  │          │
   └─────────┘    └─────────────┘  └─────────┘  └──────────┘
```

### Layer Responsibility

| Layer | Responsibility | Examples |
|---|---|---|
| **Handler** | HTTP request/response, routing | `api.go`, `web.go` |
| **Service** | Business logic, data transformation | Challenge logic, Coach AI calls |
| **Model** | Data structures, validation | Challenge, Package, Interview structs |
| **Storage** | Persistence, file I/O | JSON serialization, filesystem ops |

---

## Naming Conventions

### File & Directory Naming

#### Go Source Files
- **Package Structure**: `internal/models/`, `internal/services/`, `internal/handlers/`, `internal/server/`
- **File Naming**: `snake_case.go` - descriptive, lowercase with underscores
- **Examples**:
  - `challenge.go` - Challenge data models
  - `interview_coach.go` - Interview coaching system models and services
  - `api.go` - API endpoint handlers
  - `web.go` - Web page handlers

#### Template Files
- **Directory**: `templates/`
- **Naming**: `kebab-case.html` - lowercase with hyphens
- **Examples**:
  - `challenges.html` - Challenge listing page
  - `interview-coach.html` - Interview coach UI
  - `base.html` - Base layout template

#### Data Storage Directories
- **Format**: `kebab-case/` or `lowercase/`
- **Examples**:
  - `challenge-{1-30}/` - Challenge directories
  - `packages/` - Package data
  - `interview-coach-data/` - AI coach data
  - `badges/` - User badges

### Package & Variable Naming

#### Package Names
- **Convention**: Single word, lowercase, no underscores
- **Examples**:
  - `models` - Data structures
  - `services` - Business logic
  - `handlers` - HTTP request handlers
  - `server` - Server configuration

#### Function Names
- **Convention**: `PascalCase` (exported), `camelCase` (unexported)
- **Prefix Conventions**:
  - `Get*` - Retrieves data
  - `Load*` - Reads from storage
  - `Save*` - Writes to storage
  - `Handle*` - HTTP handlers
  - `Generate*` - Creates new data
  - `Validate*` - Validates data
  - `Parse*` - Parses data from bytes

**Examples**:
```go
// Exported (Public) - PascalCase
func (s *Service) GenerateQuestions(jobRole, jobDesc string) error
func (h *Handler) GetChallenge(w http.ResponseWriter, r *http.Request)

// Unexported (Private) - camelCase
func (s *Service) generateQuestionsWithAI(jobRole string) ([]Question, error)
func (h *Handler) parseSubmissionRequest(r *http.Request) (*Submission, error)
```

#### Variable Names
- **Convention**: `camelCase` for local variables, short but descriptive
- **Single Letter**: Only for loop counters or very obvious contexts
- **Examples**:
  ```go
  var (
      sessionID     string
      jobRole       string
      userResponses []string
      err          error
  )
  
  for i := 0; i < len(questions); i++ { // OK: loop counter
      q := questions[i]  // OK: obvious from context
  }
  ```

#### Interface Names
- **Convention**: `{Action}er` suffix for single-method interfaces
- **Examples**:
  - `Reader` - implements Read()
  - `Writer` - implements Write()
  - `Handler` - implements Handle()
  - `Generator` - implements Generate()

#### Struct Field Names
- **Convention**: `PascalCase` with JSON tags for serialization
- **Examples**:
  ```go
  type Challenge struct {
      ID                 int       `json:"id"`
      Title              string    `json:"title"`
      Description        string    `json:"description"`
      Difficulty         string    `json:"difficulty"`
      LearningMaterials  string    `json:"learningMaterials"`
  }
  ```

#### Constant Names
- **Convention**: `PascalCase` for exported, usually UPPERCASE for constants
- **Examples**:
  ```go
  const (
      MaxChallengeID      = 30
      MinDifficultyLevel  = 1
      MaxDifficultyLevel  = 3
      DefaultTimeoutMs    = 5000
  )
  ```

#### Error Naming
- **Convention**: `Err{Descriptor}` prefix for error variables
- **Examples**:
  ```go
  var (
      ErrNotFound        = errors.New("not found")
      ErrInvalidInput    = errors.New("invalid input")
      ErrUnauthorized    = errors.New("unauthorized")
  )
  ```

### API Naming

#### Endpoint Paths
- **Convention**: RESTful, lowercase with hyphens, noun-based
- **Format**: `/api/{resource}/{action}` or `/api/{resource}/{id}/{action}`
- **Examples**:
  ```
  GET  /api/challenges              - List all challenges
  GET  /api/challenges/{id}         - Get specific challenge
  POST /api/challenges/{id}/submit  - Submit solution
  GET  /api/leaderboard            - Global leaderboard
  POST /api/coach/generate          - Generate interview questions
  POST /api/coach/followup          - Follow-up question
  ```

#### Query Parameters
- **Convention**: `camelCase`, descriptive names
- **Examples**:
  ```
  ?difficulty=Intermediate
  ?sortBy=submissionTime
  ?filterBy=package-name
  ?sessionId=uuid-here
  ```

#### JSON Field Names
- **Convention**: `camelCase` in JSON (via struct tags)
- **Go Struct**: `PascalCase` internally
- **Examples**:
  ```go
  type ChallengeResponse struct {
      ChallengeID    int    `json:"challengeId"`
      Title          string `json:"title"`
      SubmittedAt    string `json:"submittedAt"`
      PassedTests    int    `json:"passedTests"`
  }
  ```

### Database/Storage File Names

#### Session Files
- **Pattern**: `sessions/{username}/{session-uuid}.json`
- **Example**: `sessions/john-doe/550e8400-e29b-41d4-a716-446655440000.json`

#### Challenge Submissions
- **Pattern**: `challenge-{id}/submissions/{username}/solution-template.go`
- **Example**: `challenge-1/submissions/john-doe/solution-template.go`

#### Package Submissions
- **Pattern**: `packages/{package-name}/{challenge-id}/submissions/{username}/solution.go`
- **Example**: `packages/gin/challenge-1-basic-routing/submissions/john-doe/solution.go`

#### Conversation Data
- **Pattern**: `interview-coach-data/conversations/{question-uuid}.json`
- **Example**: `interview-coach-data/conversations/550e8400-e29b-41d4-a716-446655440000.json`

---

## Validation Rules

### Input Validations

#### Challenge Submission
| Field | Type | Min Length | Max Length | Pattern | Rule |
|---|---|---|---|---|---|
| Username | string | 1 | 50 | `[a-zA-Z0-9-]` | GitHub username format |
| ChallengeID | int | - | - | `1-30` | Must be between 1 and 30 |
| Code | string | 50 | 10000 | Go syntax | Valid Go code that compiles |
| ExecutionMs | int64 | 0 | 30000 | Numeric | Execution time not exceeding 30 seconds |

#### Interview Session Input
| Field | Type | Min Length | Max Length | Required | Rule |
|---|---|---|---|---|---|
| JobRole | string | 2 | 100 | Yes | Non-empty job title |
| JobDescription | string | 50 | 10000 | Yes | Substantial job description |
| SessionID | string | 36 | 36 | No | Valid UUID format if provided |

#### Package Challenge Submission
| Field | Type | Min Length | Max Length | Rule |
|---|---|---|---|---|
| Username | string | 1 | 50 | Valid username |
| PackageName | string | 1 | 50 | Registered package |
| ChallengeID | string | 5 | 50 | Valid challenge ID for package |
| Code | string | 50 | 10000 | Valid Go code |

### Data Format Validations

#### Difficulty Levels
- **Valid Values**: `"Beginner"`, `"Intermediate"`, `"Advanced"`
- **Case Sensitive**: Yes
- **Validation**: Must be exact match

#### Question Categories
- **Valid Values**: `"technical"`, `"role-specific"`, `"aptitude"`
- **Distribution**: Generation should maintain 40/30/30 split
- **Validation**: Must be one of three values

#### Package Categories
- **Valid Values**: `"Web"`, `"CLI"`, `"Database"`, etc.
- **Validation**: Must be registered category

#### Status Fields
- **Valid Values**: `"available"`, `"coming-soon"`
- **Validation**: Status must be one of defined values

#### Message Roles
- **Valid Values**: `"user"`, `"assistant"`
- **Validation**: Must be exactly one of two values

### Business Logic Validations

#### Challenge Validation
```
✓ ChallengeID must exist (1-30)
✓ Submitted code must not be empty
✓ Code must be valid Go syntax
✓ All required fields present
✓ Username follows GitHub naming conventions
✓ Execution time < 30 seconds
```

#### Interview Coach Validation
```
✓ JobRole not empty (minimum 2 chars)
✓ JobDescription substantial (minimum 50 chars)
✓ Generated questions = 10 (or multiples for "Load More")
✓ Category distribution: 40% technical, 30% role-specific, 30% aptitude
✓ Each question has: ID, Question, Explanation, Tags, Category
✓ All UUIDs valid format (RFC 4122)
✓ Timestamps in UTC timezone
```

#### Package Learning Validation
```
✓ Package name must be registered
✓ Challenge ID must exist in package
✓ User can only complete challenges in order
✓ Score calculations accurate (0-100 scale)
✓ Total tests > 0 and TestsPassed <= TestsTotal
```

### File Path Validations

#### Valid Path Characters
- **Allowed**: `[a-zA-Z0-9-_/.]`
- **Not Allowed**: Special characters `*?\"<>|` and spaces
- **Username in Path**: Replace spaces with hyphens if needed

#### Path Traversal Protection
```go
// Validate no ../ or .\ in paths
// Ensure paths stay within designated directories
// Sanitize user input before file operations
```

### JSON Schema Validations

#### Challenge Response
```json
{
  "id": integer (1-30),
  "title": string (required),
  "description": string (required),
  "difficulty": string (one of: Beginner, Intermediate, Advanced),
  "template": string (required),
  "testFile": string (required),
  "learningMaterials": string (required),
  "hints": string (required)
}
```

#### Interview Question Response
```json
{
  "id": string (UUID, required),
  "question": string (required, non-empty),
  "explanation": string (required, 2-3 lines),
  "tags": [string] (required, 2-5 tags),
  "category": string (one of: technical, role-specific, aptitude),
  "is_pinned": boolean (default: false),
  "generated_at": string (RFC3339 timestamp)
}
```

---

## Error Handling

### HTTP Status Codes

| Code | Situation | Example |
|---|---|---|
| 200 OK | Successful request | Challenge retrieved, submission processed |
| 201 Created | Resource created | New session created |
| 400 Bad Request | Invalid input | Missing required fields, invalid JSON |
| 404 Not Found | Resource not found | Challenge ID doesn't exist |
| 500 Internal Error | Server error | AI API failure, file I/O error |

### Error Response Format
```json
{
  "success": false,
  "message": "Descriptive error message",
  "code": "ERROR_CODE",
  "details": {
    "field": "Additional context"
  }
}
```

### Validation Error Messages

#### Challenge Submission
| Validation | Error Message |
|---|---|
| ChallengeID out of range | "Challenge ID must be between 1 and 30" |
| Code empty | "Code cannot be empty" |
| Invalid syntax | "Submitted code contains syntax errors: {details}" |
| Username invalid | "Invalid username format" |

#### Interview Coach
| Validation | Error Message |
|---|---|
| JobRole empty | "Job role is required and must not be empty" |
| JobDescription too short | "Job description must be at least 50 characters" |
| AI request failed | "Failed to generate questions: {ai_error}" |
| Invalid session | "Session not found or invalid session ID" |

### Common Error Scenarios

#### 400 Bad Request
```go
// Missing required field
{
  "success": false,
  "message": "Missing required field: jobRole",
  "code": "MISSING_FIELD"
}

// Invalid JSON format
{
  "success": false,
  "message": "Invalid JSON in request body",
  "code": "INVALID_JSON"
}
```

#### 404 Not Found
```go
{
  "success": false,
  "message": "Challenge with ID 99 not found",
  "code": "CHALLENGE_NOT_FOUND"
}
```

#### 500 Internal Server Error
```go
{
  "success": false,
  "message": "Failed to process request",
  "code": "INTERNAL_ERROR",
  "details": {
    "reason": "AI API temporarily unavailable"
  }
}
```

### Logging Standards

#### Log Levels
- **INFO**: Successful operations, important milestones
- **WARN**: Recoverable issues, deprecated usage
- **ERROR**: Failed operations, validation failures

#### Log Message Format
```
[LEVEL] [Timestamp] [Component] - Action: Descriptive message
```

**Examples**:
```
[INFO] 2026-01-14T10:30:00Z [CoachService] - Generated questions for session: session-uuid
[WARN] 2026-01-14T10:31:00Z [AIService] - Gemini API slow response, retrying
[ERROR] 2026-01-14T10:32:00Z [ChallengeService] - Failed to parse submission: invalid JSON
```

---

## Configuration & Environment Variables

### Required Environment Variables

| Variable | Type | Default | Description |
|---|---|---|---|
| `AI_PROVIDER` | string | "gemini" | AI provider: gemini, openai, claude |
| `GEMINI_API_KEY` | string | - | Google Gemini API key |
| `OPENAI_API_KEY` | string | - | OpenAI API key |
| `CLAUDE_API_KEY` | string | - | Anthropic Claude API key |
| `PORT` | string | "8080" | Server port |
| `ENVIRONMENT` | string | "development" | Environment: development, production |

### File Structure Standards

```
project-root/
├── main.go                          # Entry point
├── internal/
│   ├── models/
│   │   ├── challenge.go            # Challenge models
│   │   ├── package.go              # Package models
│   │   └── interview_coach.go       # Coach models
│   ├── services/
│   │   ├── challenge.go            # Challenge service
│   │   ├── package.go              # Package service
│   │   ├── interview_coach.go       # Coach service
│   │   └── ai.go                   # AI abstraction layer
│   ├── handlers/
│   │   ├── api.go                  # API endpoints
│   │   └── web.go                  # Web pages
│   └── server/
│       └── server.go               # Server setup & routing
├── templates/
│   ├── base.html                   # Base layout
│   ├── challenges.html
│   ├── packages.html
│   └── interview-coach.html
├── assets/                         # Static files (CSS, JS)
├── .env                            # Environment configuration
├── .env.example                    # Environment template
├── go.mod                          # Go modules
└── go.sum                          # Module checksums
```

---

## Version History

| Version | Date | Changes |
|---|---|---|
| 1.0 | 2026-01-14 | Initial documentation |

---

**End of Program Description, Naming Conventions & Validations**
