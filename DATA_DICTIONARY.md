# Data Dictionary - GoInterview Platform

**Project:** Go Interview Practice Platform  
**Date:** January 14, 2026  
**Version:** 1.0  
**Status:** Active

---

## Table of Contents

1. [Challenge Module](#challenge-module)
2. [Package Module](#package-module)
3. [Interview Coach Module](#interview-coach-module)
4. [Type Aliases & Collections](#type-aliases--collections)
5. [API Request/Response Models](#api-requestresponse-models)

---

## Challenge Module

### 1. Challenge

**Description:** Represents a single coding challenge in the platform

**Location:** `internal/models/challenge.go`

| Field Name | Type | Length | Required | Description | Example |
|---|---|---|---|---|---|
| ID | int | - | Yes | Unique challenge identifier (1-30) | 1 |
| Title | string | 100-200 | Yes | Challenge title/name | "Two Sum" |
| Description | string | 500-2000 | Yes | Full problem statement and requirements | "Given an array of integers nums..." |
| Difficulty | string | 20 | Yes | Challenge difficulty level | "Beginner", "Intermediate", "Advanced" |
| Template | string | 1000-5000 | Yes | Go code template/skeleton for users | "func solution(nums []int) []int { ... }" |
| TestFile | string | 1000-10000 | Yes | Go test file with test cases | "func TestSolution(t *testing.T) { ... }" |
| LearningMaterials | string | 500-5000 | Yes | Markdown content with learning resources | "## Key Concepts\n- Hash Maps\n- Two Pointers" |
| Hints | string | 200-1000 | Yes | Progressive hints for solving the challenge | "Hint 1: Use a hash map\nHint 2: ..." |

**Relationships:**
- One Challenge → Many Submissions
- One Challenge → One Scoreboard Entry

**Sample Data:**
```json
{
  "id": 1,
  "title": "Two Sum",
  "description": "Given an array of integers nums and an integer target...",
  "difficulty": "Beginner",
  "template": "func twoSum(nums []int, target int) []int { ... }",
  "testFile": "func TestTwoSum(t *testing.T) { ... }",
  "learningMaterials": "# Two Sum Solution\n## Approach 1: Hash Map",
  "hints": "Use a hash map to store..."
}
```

---

### 2. Submission

**Description:** Represents a user's solution submission for a challenge

**Location:** `internal/models/challenge.go`

| Field Name | Type | Length | Required | Description | Example |
|---|---|---|---|---|---|
| Username | string | 50 | Yes | GitHub username of submitter | "john-doe" |
| ChallengeID | int | - | Yes | Reference to challenge being solved | 1 |
| Code | string | 1000-10000 | Yes | User's submitted Go code | "func twoSum(nums []int, target int) []int { ... }" |
| SubmittedAt | time.Time | - | Yes | Timestamp of submission | "2026-01-14T10:30:00Z" |
| Passed | bool | - | Yes | Whether all tests passed | true |
| TestOutput | string | 500-5000 | No | Output from test execution | "PASS: TestTwoSum ...\nok command-line-arguments 0.001s" |
| ExecutionMs | int64 | - | No | Execution time in milliseconds | 125 |

**Relationships:**
- Many Submissions → One Challenge
- Many Submissions → One User

**File Storage Path:** `challenge-{id}/submissions/{username}/solution-template.go`

**Sample Data:**
```json
{
  "username": "john-doe",
  "challengeId": 1,
  "code": "func twoSum(nums []int, target int) []int { ... }",
  "submittedAt": "2026-01-14T10:30:00Z",
  "passed": true,
  "testOutput": "PASS: All tests passed",
  "executionMs": 125
}
```

---

### 3. ScoreboardEntry

**Description:** Represents a submission entry on the challenge scoreboard

**Location:** `internal/models/challenge.go`

| Field Name | Type | Length | Required | Description | Example |
|---|---|---|---|---|---|
| Username | string | 50 | Yes | GitHub username | "john-doe" |
| ChallengeID | int | - | Yes | Challenge ID reference | 1 |
| SubmittedAt | time.Time | - | Yes | When the challenge was completed | "2026-01-14T10:30:00Z" |

**Relationships:**
- Many ScoreboardEntries → One Challenge
- Many ScoreboardEntries → One User

**Sample Data:**
```json
{
  "username": "john-doe",
  "challengeId": 1,
  "submittedAt": "2026-01-14T10:30:00Z"
}
```

---

### 4. UserAttemptedChallenges

**Description:** Tracks user's progress and attempts across challenges

**Location:** `internal/models/challenge.go`

| Field Name | Type | Description |
|---|---|---|
| Username | string | Unique user identifier |
| AttemptedIDs | map[int]bool | Map of challenge IDs user has attempted (key=ChallengeID, value=true if attempted) |
| Scores | map[int]int | Map of scores per challenge (0-100 scale) |

**Sample Data:**
```json
{
  "username": "john-doe",
  "attemptedIds": {
    "1": true,
    "2": true,
    "5": true
  },
  "scores": {
    "1": 100,
    "2": 85,
    "5": 100
  }
}
```

---

## Package Module

### 5. Package

**Description:** Represents a Go package with learning challenges (e.g., Gin, GORM, Cobra)

**Location:** `internal/models/package.go`

| Field Name | Type | Description | Example |
|---|---|---|---|
| Name | string | Unique package identifier | "gin" |
| DisplayName | string | User-friendly package name | "Gin Web Framework" |
| Description | string | Package description | "A popular web framework for building APIs..." |
| Version | string | Package version | "1.9.1" |
| GitHubURL | string | GitHub repository URL | "https://github.com/gin-gonic/gin" |
| DocumentationURL | string | Official docs URL | "https://gin-gonic.com/" |
| Stars | int | GitHub stars count | 76300 |
| Category | string | Package category | "Web", "CLI", "Database" |
| Difficulty | string | Learning difficulty | "Beginner", "Intermediate", "Advanced" |
| Prerequisites | []string | Required prior knowledge | ["Go Basics", "HTTP Concepts"] |
| LearningPath | []string | Ordered list of challenge IDs | ["challenge-1-basic-routing", "challenge-2-middleware"] |
| Tags | []string | Technology tags | ["Framework", "REST API", "Go"] |
| EstimatedTime | string | Learning duration | "4-6 hours" |
| RealWorldUsage | []string | Real-world applications | ["API Development", "Microservices"] |
| ChallengeDetails | map[string]*ChallengeInfo | Challenge metadata (optional) | - |

**Relationships:**
- One Package → Many PackageChallenges
- One Package → Many PackageProgresses

**Sample Data:**
```json
{
  "name": "gin",
  "display_name": "Gin Web Framework",
  "description": "A popular web framework for building REST APIs in Go",
  "version": "1.9.1",
  "github_url": "https://github.com/gin-gonic/gin",
  "documentation_url": "https://gin-gonic.com/",
  "stars": 76300,
  "category": "Web",
  "difficulty": "Intermediate",
  "prerequisites": ["Go Basics", "HTTP Concepts"],
  "learning_path": ["challenge-1-basic-routing", "challenge-2-middleware"],
  "tags": ["Framework", "REST API", "Go"],
  "estimated_time": "4-6 hours",
  "real_world_usage": ["API Development", "Microservices"]
}
```

---

### 6. ChallengeInfo

**Description:** Metadata for individual challenges within a package learning path

**Location:** `internal/models/package.go`

| Field Name | Type | Description | Example |
|---|---|---|---|
| ID | string | Unique challenge identifier | "challenge-1-basic-routing" |
| Title | string | Challenge title | "Basic Routing" |
| Description | string | Detailed description | "Learn how to set up basic routes in Gin..." |
| Difficulty | string | Challenge difficulty | "Beginner" |
| EstimatedTime | string | Time to complete | "30-45 minutes" |
| LearningObjectives | []string | What user will learn | ["Create routes", "Handle HTTP methods"] |
| Prerequisites | []string | Required prior knowledge | ["Go Basics"] |
| Tags | []string | Topic tags | ["Routing", "HTTP"] |
| RealWorldConnection | string | Practical application | "Essential for any REST API" |
| Icon | string | Bootstrap icon class | "bi-shuffle" |
| Status | string | Challenge status | "available", "coming-soon" |
| Order | int | Position in learning path | 1 |

---

### 7. PackageChallenge

**Description:** Complete challenge data within a package context

**Location:** `internal/models/package.go`

| Field Name | Type | Description |
|---|---|---|
| ID | string | Unique challenge ID (e.g., "challenge-1-basic-routing") |
| PackageName | string | Parent package name |
| Title | string | Challenge title |
| Description | string | Full problem description |
| ShortDescription | string | Brief description for cards |
| Difficulty | string | Challenge difficulty |
| LearningObjectives | []string | Learning goals |
| Template | string | Go code template |
| TestFile | string | Go test file |
| LearningMaterials | string | Study resources |
| Hints | string | Progressive hints |
| Requirements | []string | Specific requirements |
| BonusPoints | []string | Extra credit opportunities |
| RealWorldConnection | string | Real-world use case |
| EstimatedTime | string | Time estimate |
| Tags | []string | Topic tags |
| Prerequisites | []string | Required knowledge |
| Icon | string | Bootstrap icon |
| Order | int | Position in learning path |
| Status | string | "available", "coming-soon" |

**File Storage Path:** `packages/{packageName}/{challengeId}/solution-template.go`

---

### 8. PackageSubmission

**Description:** User's submission for a package challenge

**Location:** `internal/models/package.go`

| Field Name | Type | Description | Example |
|---|---|---|---|
| Username | string | GitHub username | "john-doe" |
| PackageName | string | Package name | "gin" |
| ChallengeID | string | Challenge ID | "challenge-1-basic-routing" |
| Code | string | Submitted Go code | "func main() { ... }" |
| SubmittedAt | time.Time | Submission timestamp | "2026-01-14T10:30:00Z" |
| Passed | bool | All tests passed? | true |
| TestOutput | string | Test execution output | "PASS: All tests" |
| ExecutionMs | int64 | Execution time (ms) | 150 |
| TestsPassed | int | Number of tests passed | 8 |
| TestsTotal | int | Total test count | 8 |

**File Storage Path:** `packages/{packageName}/{challengeId}/submissions/{username}/solution.go`

---

### 9. PackageProgress

**Description:** Tracks user's progress through a package learning path

**Location:** `internal/models/package.go`

| Field Name | Type | Description |
|---|---|---|
| Username | string | GitHub username |
| PackageName | string | Package name |
| CompletedChallenges | []string | IDs of completed challenges |
| InProgress | string | Currently working challenge ID |
| StartedAt | time.Time | When user started learning this package |
| LastActivity | time.Time | Last submission timestamp |
| TotalTime | time.Duration | Total learning time |
| Achievements | []string | Earned badges/achievements |
| Score | int | Overall score (0-100) |

---

### 10. PackageScoreboardEntry

**Description:** Entry in the package-wide leaderboard

**Location:** `internal/models/package.go`

| Field Name | Type | Description |
|---|---|---|
| Username | string | GitHub username |
| PackageName | string | Package name |
| ChallengeID | string | Challenge ID |
| SubmittedAt | time.Time | Submission time |
| ExecutionMs | int64 | Execution time |
| TestsPassed | int | Tests passed |
| TestsTotal | int | Total tests |
| IsSponsor | bool | Is GitHub sponsor? |

---

## Interview Coach Module

### 11. InterviewQuestion

**Description:** A single interview question generated by AI coach

**Location:** `internal/models/interview_coach.go`

| Field Name | Type | Description | Example |
|---|---|---|---|
| ID | string | Unique question UUID | "550e8400-e29b-41d4-a716-446655440000" |
| Question | string | The interview question | "Explain microservices architecture" |
| Explanation | string | 2-3 line explanation | "Tests understanding of system design and distributed systems" |
| Tags | []string | Topic tags | ["Architecture", "System Design", "Microservices"] |
| Category | string | Question category | "technical", "role-specific", "aptitude" |
| IsPinned | bool | Is question pinned? | true |
| PinnedAt | *time.Time | When was it pinned? | "2026-01-14T11:00:00Z" |
| GeneratedAt | time.Time | When was it generated? | "2026-01-14T10:30:00Z" |

**Categories:**
- `technical`: 40% - Technical/coding questions
- `role-specific`: 30% - Role-specific questions
- `aptitude`: 30% - Behavioral/aptitude/soft skills

**Sample Data:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "question": "Explain the SOLID principles in object-oriented design",
  "explanation": "Tests understanding of design patterns and software architecture principles",
  "tags": ["Architecture", "Design Patterns", "Best Practices"],
  "category": "technical",
  "is_pinned": true,
  "pinned_at": "2026-01-14T11:00:00Z",
  "generated_at": "2026-01-14T10:30:00Z"
}
```

---

### 12. InterviewSession

**Description:** A complete coaching session with job details and generated questions

**Location:** `internal/models/interview_coach.go`

| Field Name | Type | Description |
|---|---|---|
| SessionID | string | Unique session UUID |
| Username | string | GitHub username |
| JobRole | string | Job title (e.g., "Backend Developer") |
| JobDescription | string | Full job description from LinkedIn/JD |
| Questions | []InterviewQuestion | Generated questions (can be 10+) |
| CreatedAt | time.Time | Session creation time |
| UpdatedAt | time.Time | Last modified time |

**File Storage Path:** `interview-coach-data/sessions/{username}/{session-id}.json`

**Sample Data:**
```json
{
  "session_id": "session-uuid-here",
  "username": "john-doe",
  "job_role": "Backend Developer",
  "job_description": "We are looking for an experienced Backend Developer...",
  "questions": [
    { "id": "q1", "question": "...", ... },
    { "id": "q2", "question": "...", ... }
  ],
  "created_at": "2026-01-14T10:30:00Z",
  "updated_at": "2026-01-14T10:35:00Z"
}
```

---

### 13. ConversationMessage

**Description:** A single message in the follow-up chat conversation

**Location:** `internal/models/interview_coach.go`

| Field Name | Type | Description | Example |
|---|---|---|---|
| ID | string | Unique message UUID | "msg-uuid-123" |
| QuestionID | string | Related question UUID | "q1-uuid" |
| Username | string | Message author username | "john-doe" |
| Role | string | Message sender type | "user", "assistant" |
| Content | string | Message text | "Can you explain this concept?" |
| Timestamp | time.Time | When message was sent | "2026-01-14T11:05:00Z" |

**Sample Data:**
```json
{
  "id": "msg-uuid-123",
  "question_id": "q1-uuid",
  "username": "john-doe",
  "role": "user",
  "content": "Can you explain microservices more simply?",
  "timestamp": "2026-01-14T11:05:00Z"
}
```

---

### 14. QuestionConversation

**Description:** Complete conversation thread for a single question

**Location:** `internal/models/interview_coach.go`

| Field Name | Type | Description |
|---|---|---|
| QuestionID | string | The question being discussed |
| Messages | []ConversationMessage | All messages in the conversation |

**File Storage Path:** `interview-coach-data/conversations/{question-id}.json`

**Sample Data:**
```json
{
  "question_id": "q1-uuid",
  "messages": [
    {
      "id": "msg-1",
      "role": "assistant",
      "content": "Here's a detailed explanation..."
    },
    {
      "id": "msg-2",
      "role": "user",
      "content": "Can you give an example?"
    },
    {
      "id": "msg-3",
      "role": "assistant",
      "content": "Sure! Here's a practical example..."
    }
  ]
}
```

---

## Type Aliases & Collections

### Collection Types

**Location:** `internal/models/challenge.go`, `internal/models/package.go`, `internal/models/interview_coach.go`

#### Challenge Module Collections

| Type Alias | Definition | Purpose |
|---|---|---|
| ChallengeMap | map[int]*Challenge | Store all challenges indexed by ID |
| ScoreboardMap | map[int][]ScoreboardEntry | Store scoreboard entries indexed by challenge ID |
| UserAttemptsMap | map[string]*UserAttemptedChallenges | Track all users' attempts |

#### Package Module Collections

| Type Alias | Definition | Purpose |
|---|---|---|
| PackageMap | map[string]*Package | Store all packages indexed by name |
| PackageChallengeMap | map[string]map[string]*PackageChallenge | Store challenges: package → challenge_id → challenge |
| PackageSubmissionMap | map[string][]PackageSubmission | Store submissions indexed by package name |
| PackageProgressMap | map[string]map[string]*PackageProgress | Track progress: username → package → progress |
| PackageScoreboardMap | map[string][]PackageScoreboardEntry | Store leaderboard entries by package |

---

## API Request/Response Models

### Interview Coach API Models

**Location:** `internal/models/interview_coach.go`

#### 15. GenerateQuestionsRequest

**Used By:** `POST /api/coach/generate`

| Field Name | Type | Required | Description |
|---|---|---|---|
| JobRole | string | Yes | Job title/position |
| JobDescription | string | Yes | Full job description |
| SessionID | string | No | Session ID for "Load More" |

---

#### 16. GenerateQuestionsResponse

**Used By:** `POST /api/coach/generate`

| Field Name | Type | Description |
|---|---|---|
| SessionID | string | New or existing session ID |
| Questions | []InterviewQuestion | Generated 10 questions |
| Success | bool | Request succeeded? |
| Message | string | Error message if failed |

---

#### 17. ExplainQuestionRequest

**Used By:** `POST /api/coach/explain`

| Field Name | Type | Required | Description |
|---|---|---|---|
| QuestionID | string | Yes | Question UUID |
| SessionID | string | Yes | Session UUID |
| Question | string | Yes | Full question text |

---

#### 18. ExplainQuestionResponse

**Used By:** `POST /api/coach/explain`

| Field Name | Type | Description |
|---|---|---|
| Explanation | string | Detailed beginner-friendly explanation |
| Success | bool | Request succeeded? |
| Message | string | Error message if failed |

---

#### 19. FollowupQuestionRequest

**Used By:** `POST /api/coach/followup`

| Field Name | Type | Required | Description |
|---|---|---|---|
| QuestionID | string | Yes | Question UUID |
| SessionID | string | Yes | Session UUID |
| FollowupMessage | string | Yes | User's follow-up question |
| OriginalQuestion | string | Yes | Original question text |

---

#### 20. FollowupQuestionResponse

**Used By:** `POST /api/coach/followup`

| Field Name | Type | Description |
|---|---|---|
| Response | string | AI's response to follow-up |
| Success | bool | Request succeeded? |
| Message | string | Error message if failed |

---

#### 21. PinQuestionRequest

**Used By:** `POST /api/coach/pin`

| Field Name | Type | Required | Description |
|---|---|---|---|
| QuestionID | string | Yes | Question UUID |
| SessionID | string | Yes | Session UUID |
| Pin | bool | Yes | true to pin, false to unpin |

---

#### 22. PinQuestionResponse

**Used By:** `POST /api/coach/pin`

| Field Name | Type | Description |
|---|---|---|
| Success | bool | Operation succeeded? |
| Message | string | Status message |

---

## Data Constraints & Validations

### Difficulty Levels
- `Beginner` - Basic problems for learners
- `Intermediate` - Moderate complexity
- `Advanced` - Complex, real-world scenarios

### Question Categories (40/30/30 Distribution)
- `technical` - Coding/technical questions (40%)
- `role-specific` - Job role-specific questions (30%)
- `aptitude` - Behavioral/soft skills (30%)

### Package Categories
- `Web` - Web frameworks (Gin, Fiber, Echo)
- `CLI` - Command-line frameworks (Cobra)
- `Database` - Database ORMs (GORM, MongoDB)

### Status Values
- `available` - Challenge/package is active
- `coming-soon` - Under development

### Message Roles
- `user` - Message from user
- `assistant` - Message from AI coach

### Test Result Values
- `true` - All tests passed
- `false` - Some tests failed

---

## Data Relationships Diagram

```
Challenge
├── 1---M Submission
├── 1---M ScoreboardEntry
└── 1---M UserAttemptedChallenges

Package
├── 1---M PackageChallenge
├── 1---M PackageProgress
└── 1---M PackageScoreboardEntry

InterviewSession
├── 1---M InterviewQuestion
└── 1---M QuestionConversation

QuestionConversation
└── 1---M ConversationMessage
```

---

## Data Storage Locations

### File System Storage

```
project-root/
├── challenge-{1-30}/
│   └── submissions/{username}/
│       └── solution-template.go
│
├── packages/{package-name}/{challenge-id}/
│   ├── submissions/{username}/
│   │   └── solution.go
│   └── {challenge-files}
│
└── interview-coach-data/
    ├── sessions/{username}/
    │   └── {session-id}.json
    └── conversations/
        └── {question-id}.json
```

---

## Version History

| Version | Date | Changes |
|---|---|---|
| 1.0 | 2026-01-14 | Initial data dictionary with all modules |

---

**End of Data Dictionary**
