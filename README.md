# How to run the project

1) Clone the repository
```bash
git clone https://github.com/ilyes-rhdi/buildit-Gql.git
cd buildit-Gql
```

2) Configure environment variables

Backend
```bash
cd backend
cp .env.example .env
```

Frontend
```bash
cd ../frontend
cp .env.example .env
```

3) Start everything
```bash
docker compose up --build
```

4) Open
- Frontend: http://localhost:5173
- Backend (GraphQL): http://localhost:8080/graphql

---

# NotionQL 

A collaborative workspace application showcasing advanced GraphQL capabilities through a Notion-like interface.

## Table of Contents :

* About the Project
* Why GraphQL?
* Key Features
* Advanced GraphQL Patterns
* Architecture
* Data Models
* Technology Stack
* Getting Started
* GraphQL Examples
* Demo Scenarios
* Project Structure
* Team
## About the Project

NotionQL is a collaborative workspace application inspired by Notion.

- The main objective of this project is to demonstrate how GraphQL can be used
to build a modern web application with good performance and clear data
organization.

- Unlike simple applications that only use GraphQL for basic create, read,
update and delete operations, this project focuses on more advanced but
essential concepts.

* The application supports real-time collaboration, allowing multiple users
to edit the same page and see updates instantly.
It uses nested GraphQL queries to retrieve related data (workspaces, pages
and blocks) in a structured way.

* To ensure good performance, the project implements DataLoader to group database requests and avoid unnecessary queries.
* A role-based permission system is also included to control access to data according to user roles (Owner, Admin, Editor, Viewer).

* Different types of content blocks are handled using GraphQL interface types,which makes the system flexible and easy to extend.
* Finally, cursor-based pagination is used to efficiently load large amounts of data without affecting application performance.

#### This project shows how GraphQL can be applied in a practical and structuredway for real-world collaborative applications.

### Core Functionality
Users can:

* Create and manage workspaces (shared team folders).
*  Build pages with rich content blocks (text, headings, todos, images)
*  Reorder and nest blocks to create hierarchical content
*  Comment on pages or specific blocks
*  Collaborate in real-time with live updates
*  Control access with granular permission system 
 ## Why GraphQL?
* GraphQL transforms how clients interact with   APIs by providing:

       Traditional REST API
      GET /api/page/123          → Returns entirepage with ALL data
      GET /api/page/123/blocks   → Separate request for blocks
      GET /api/blocks/456/author → Another request for each author
#### Result: Multiple round trips, over-fetching, under-fetching
### GraphQL Approach
    graphqlquery {
     page(id: "123") {
    title
    blocks(first: 20) {
      id
      type
      content
      author {
        name
        avatar
      }
    }
    }
    }

#### Result: Single request, exactly the data you need, strongly typed
### Why Notion-like Apps are Perfect for GraphQL:
Notion-style applications are inherently graph-structured:
```mermaid 
graph TD
    W[Workspace] -->|contains| P1[Page 1]
    W -->|contains| P2[Page 2]
    P1 -->|has| B1[Block: Heading]
    P1 -->|has| B2[Block: Paragraph]
    P1 -->|has| B3[Block: Todo]
    B3 -->|nested| B4[Sub-todo 1]
    B3 -->|nested| B5[Sub-todo 2]
    P1 -->|has| C1[Comment]
    B2 -->|has| C2[Comment]
    U1[User: Owner] -->|member of| W
    U2[User: Editor] -->|member of| W
    U1 -->|created| B1
    U2 -->|created| C1
    
    style W fill:#e1f5ff
    style P1 fill:#fff4e1
    style P2 fill:#fff4e1
    style B1 fill:#e8f5e9
    style B2 fill:#e8f5e9
    style B3 fill:#e8f5e9
    style C1 fill:#fce4ec
    style C2 fill:#fce4ec
 `````
##### GraphQL naturally represents these complex relationships, making it the ideal choice for collaborative document applications.
## Key Features

### 1. Workspace Management
Create team workspaces with customizable names
Role-based member management (Owner, Admin, Editor, Viewer)
Transfer ownership and remove members.

### 2. Dynamic Page Building

Rich block-based content editor
Support for multiple block types:

*  Paragraph - Rich text content
*  Heading - Hierarchical titles (H1, H2, H3)
*  Todo - Interactive checkboxes
*  Image - Visual content
Nested block structures for complex layouts

### 3. Real-Time Collaboration

* Live updates using GraphQL subscriptions.
* See changes as they happen.
* Optional presence indicators (who's viewing/editing).

### 4. Comments & Discussion

* Comment on entire pages or specific blocks.
* Threaded discussions.
* Author attribution.

### 5. Security & Permissions

* JWT-based authentication.
* Role-based authorization at field level.
* Secure workspace access control.
 ## Advanced GraphQL Patterns
####  This project showcases enterprise-level GraphQL implementation:
### 1)- Strongly Typed Schema
##### Our GraphQL schema serves as a contract between frontend and backend:

```mermaid

classDiagram
    class Workspace {
        +ID id
        +String name
        +User owner
        +Member[] members
        +Page[] pages
    }
    
    class Page {
        +ID id
        +String title
        +Workspace workspace
        +Block[] blocks
        +Comment[] comments
    }
    
    class Block {
        <<interface>>
        +ID id
        +Int position
        +Block parent
        +User createdBy
    }
    
    class ParagraphBlock {
        +String text
    }
    
    class TodoBlock {
        +String text
        +Boolean checked
    }
    
    class HeadingBlock {
        +String text
        +Int level
    }
    
    class Comment {
        +ID id
        +String content
        +User author
        +Page page
        +Block block
    }
    
    class User {
        +ID id
        +String email
        +String name
        +String avatar
    }
    
    class Member {
        +User user
        +WorkspaceRole role
    }
    
    Workspace "1" --> "*" Page
    Workspace "1" --> "*" Member
    Page "1" --> "*" Block
    Page "1" --> "*" Comment
    Block <|-- ParagraphBlock
    Block <|-- TodoBlock
    Block <|-- HeadingBlock
    Block "1" --> "*" Block : nested
    Comment --> Block : optional
    Member --> User
```

### 2)- GraphQL Schema (Basé sur notre implémentation)
#### Types Principaux
```graphql
# User Type
type User {
  id: ID!
  email: String!
  name: String!
  bio: String
  phone: String
  address: String
  joined: DateTime!
  image: String
  gender: Boolean
}

# Workspace Type
type Workspace {
  id: ID!
  name: String!
  OwnerID: ID!
  Owner: User
  Members: [WorkspaceMember!]
  Pages: [Page!]
}

# Workspace Member
type WorkspaceMember {
  id: ID!
  workspaceId: ID!
  userId: ID!
  role: WorkspaceRole!
  user: User
}

# Page Type
type Page {
  id: ID!
  workspaceId: ID!
  title: String
  archived: Boolean!
}

# Block Type
type Block {
  id: ID!
  pageId: ID!
  parentblockId: ID
  Type: BlockType!
  data: JSON
}

# JSON Scalar pour données dynamiques
scalar JSON
scalar DateTime
```
#### Enums
```graphql
enum WorkspaceRole {
  OWNER
  ADMIN
  MEMBER
}

enum BlockType {
  PARAGRAPH
  HEADING1
  HEADING2
  TODO
  BULLET
  NUMBER
  QUOTE
  CODE
  IMAGE
  DIVIDER
}
```
### Queries Disponibles
```graphql 
type Query {
  # Utilisateur connecté
  me: User
  
  # Workspace
  workspace(id: ID!): Workspace
  workspaces: [Workspace!]!
  
  # Pages
  page(id: ID!): Page
  pages(workspaceId: ID!): [Page!]!
  
  # Blocks
  block(id: ID!): Block
  blocks(pageId: ID!, parentblockId: ID): [Block!]!
}
```
### Mutations Disponibles

```graphql 
type Mutation {
  # Workspace Management
  createWorkspace(name: String!): Workspace!
  addWorkspaceMember(workspaceId: ID!, userId: ID!, role: WorkspaceRole): WorkspaceMember!
  updateWorkspaceMemberRole(workspaceId: ID!, userId: ID!, role: WorkspaceRole!): WorkspaceMember!
  removeWorkspaceMember(workspaceId: ID!, userId: ID!): Boolean!
  transferWorkspaceOwnership(workspaceId: ID!, newOwnerUserId: ID!): Workspace!
  
  # Page Management
  createPage(workspaceId: ID!, title: String): Page!
  updatePage(id: ID!, title: String!): Page!
  archivePage(id: ID!, archived: Boolean!): Page!
  deletePageHard(id: ID!): Boolean!
  
  # Block Management
  createBlock(
    pageId: ID!
    parentblockId: ID
    type: BlockType!
    order: Int
    data: JSON
  ): Block!
  
  updateBlock(id: ID!, type: BlockType!, data: JSON): Block!
  
  moveBlock(
    id: ID!
    pageId: ID!
    parentblockId: ID
    order: Int!
  ): Block!
  
  deleteBlockTree(id: ID!): Boolean!
}
```
#### Exemples de Requêtes Réelles
##### Query 1: Récupérer mes workspaces avec membres
```graphql
query MyWorkspaces {
  me {
    id
    name
    email
    image
  }
  workspaces {
    id
    name
    Owner {
      id
      name
    }
    Members {
      user {
        name
        email
      }
      role
    }
  }
}
```
##### Query 2: Récupérer une page avec tous ses blocks
```graphql
query GetPageWithBlocks($pageId: ID!) {
  page(id: $pageId) {
    id
    title
    workspaceId
    archived
  }
  blocks(pageId: $pageId) {
    id
    Type
    data
    parentblockId
  }
}
```
Variables:
```json
{
  "pageId": "page-uuid-123"
}
```
##### Query 3: Récupérer un workspace complet
```graphql
 query GetWorkspaceDetails($workspaceId: ID!) {
  workspace(id: $workspaceId) {
    id
    name
    Owner {
      id
      name
      email
      image
    }
    Members {
      id
      role
      user {
        name
        email
        image
      }
    }
    Pages {
      id
      title
      archived
    }
  }
}
```
### Variables:
```json
{
  "workspaceId": "workspace-uuid-456"
}
```
#### Exemples de Mutations Réelles
##### Mutation 1: Créer un workspace
```graphql
mutation CreateWorkspace($name: String!) {
  createWorkspace(name: $name) {
    id
    name
    OwnerID
    Owner {
      name
      email
    }
  }
}
```
##### Variables:
```json
{
  "name": "My Team Workspace"
}
```
##### Mutation 2: Créer une page
```graphql
mutation CreatePage($workspaceId: ID!, $title: String) {
  createPage(workspaceId: $workspaceId, title: $title) {
    id
    title
    workspaceId
    archived
  }
}
```
##### Variables:
```json
{
  "workspaceId": "workspace-uuid-123",
  "title": "Project Documentation"
}
```
##### Mutation 3: Créer un block de type TODO
```graphql
mutation CreateTodoBlock($pageId: ID!, $data: JSON) {
  createBlock(
    pageId: $pageId
    type: TODO
    order: 0
    data: $data
  ) {
    id
    Type
    data
    pageId
  }
}
```
##### Variables:
```json
{
  "pageId": "page-uuid-789",
  "data": {
    "text": "Complete GraphQL implementation",
    "checked": false
  }
}
```
##### Mutation 4: Ajouter un membre au workspace
```graphql
mutation AddMember($workspaceId: ID!, $userId: ID!, $role: WorkspaceRole) {
  addWorkspaceMember(
    workspaceId: $workspaceId
    userId: $userId
    role: $role
  ) {
    id
    role
    user {
      name
      email
    }
  }
}
```
##### Variables:
```json
{
  "workspaceId": "workspace-uuid-123",
  "userId": "user-uuid-456",
  "role": "MEMBER"
}
```
##### Mutation 5: Mettre à jour un block
```graphql
mutation UpdateBlock($id: ID!, $type: BlockType!, $data: JSON) {
  updateBlock(id: $id, type: $type, data: $data) {
    id
    Type
    data
  }
}
```
##### Variables:
```json 
{
  "id": "block-uuid-999",
  "type": "PARAGRAPH",
  "data": {
    "text": "Updated content here"
  }
} 
```
##### Mutation 6: Archiver une page
graphqlmutation ArchivePage($id: ID!, $archived: Boolean!) {
  archivePage(id: $id, archived: $archived) {
    id
    title
    archived
  }
}
##### Variables:
json{
  "id": "page-uuid-789",
  "archived": true
}

###  Authentification
* Toutes les queries et mutations nécessitent un JWT token dans le header :
httpAuthorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Si le token est invalide ou absent, vous recevrez une erreur :

```json
{
  "errors": [
    {
      "message": "Not Authorized"
    }
  ]
}
```
##  Architecture : 
```mermaid

graph TB
    subgraph "Client Layer"
        C1[Web Browser]
        C2[React App]
        C3[Apollo Client]
    end
    
    subgraph "Load Balancer"
        LB[Nginx]
    end
    
    subgraph "Backend Layer"
        API[Go GraphQL Server]
        WS[WebSocket Server]
    end
    
    subgraph "Business Logic"
        R[Resolvers]
        S[Services]
        DL[DataLoader]
    end
    
    subgraph "Data Layer"
        DB[(PostgreSQL)]
        CACHE[(Redis Cache)]
    end
    
    C1 --> C2
    C2 --> C3
    C3 -->|HTTP/WS| LB
    LB --> API
    LB --> WS
    API --> R
    WS --> R
    R --> DL
    DL --> S
    S --> DB
    DL --> CACHE
    
    style C2 fill:#61dafb
    style API fill:#00add8
    style DB fill:#336791
    style CACHE fill:#dc382d
``` 
#### Request Flow : 
```mermaid 
sequenceDiagram
    participant Client
    participant Apollo
    participant GraphQL
    participant Resolver
    participant DataLoader
    participant Service
    participant Database
    
    Client->>Apollo: Query page with blocks
    Apollo->>GraphQL: Parse & Validate
    GraphQL->>Resolver: Resolve page field
    Resolver->>Service: GetPage(id)
    Service->>Database: SELECT * FROM pages WHERE id=?
    Database-->>Service: Page data
    Service-->>Resolver: Page object
    
    Resolver->>DataLoader: Load blocks for page
    DataLoader->>DataLoader: Batch requests
    DataLoader->>Service: GetBlocksByPageID([ids])
    Service->>Database: SELECT * FROM blocks WHERE page_id IN (?)
    Database-->>Service: Blocks data
    Service-->>DataLoader: Blocks array
    DataLoader-->>Resolver: Blocks (cached)
    
    Resolver->>DataLoader: Load authors for blocks
    DataLoader->>Service: GetUsersByIDs([unique_ids])
    Service->>Database: SELECT * FROM users WHERE id IN (?)
    Database-->>Service: Users data
    Service-->>DataLoader: Users array
    DataLoader-->>Resolver: Users (cached)
    
    Resolver-->>GraphQL: Complete response
    GraphQL-->>Apollo: JSON response
    Apollo-->>Client: Typed data
```
## Data Models

### Core Entities

```mermaid 
erDiagram
    USER ||--o{ WORKSPACE_MEMBER : "joins"
    USER ||--o{ PAGE : "creates"
    USER ||--o{ BLOCK : "creates"
    
    WORKSPACE ||--o{ WORKSPACE_MEMBER : "has"
    WORKSPACE ||--o{ PAGE : "contains"
    
    PAGE ||--o{ BLOCK : "contains"
    
    BLOCK ||--o{ BLOCK : "nests"
    
    USER {
        uuid id PK
        string email UK
        string password_hash
        string name
        string bio
        string phone
        string address
        string image
        boolean gender
        timestamp created_at
    }
    
    WORKSPACE {
        uuid id PK
        string name
        uuid owner_id FK
        timestamp created_at
        timestamp updated_at
    }
    
    WORKSPACE_MEMBER {
        uuid id PK
        uuid workspace_id FK
        uuid user_id FK
        enum role
        timestamp joined_at
    }
    
    PAGE {
        uuid id PK
        string title
        uuid workspace_id FK
        uuid created_by FK
        boolean archived
        timestamp created_at
        timestamp updated_at
    }
    
    BLOCK {
        uuid id PK
        uuid page_id FK
        uuid parent_block_id FK
        enum type
        jsonb data
        int order_position
        uuid created_by FK
        timestamp created_at
        timestamp updated_at
    }
```

### Role Hierarchy
```mermaid
graph TD
    OWNER[Owner] -->|full control| ADMIN[Admin]
    ADMIN -->|manage members & settings| MEMBER[Member]
    MEMBER -->|can create & edit| NONE[No Access]
    
    style OWNER fill:#ff6b6b
    style ADMIN fill:#ffa500
    style MEMBER fill:#4ecdc4
  ```  

## Technology Stack 
### Backend : 

* Language: Go 1.21+

* GraphQL Library: graphql-go/graphql
* Web Framework: Echo v4
* Database: PostgreSQL 15
* ORM: GORM
* Authentication: JWT tokens
* Middleware: Custom auth middleware

### Frontend : 

* Framework: React 19.2.0
* GraphQL Client: Apollo Client 4.0.11
* Routing: React Router DOM 7.11.0
* Forms: React Hook Form 7.69.0
* Build Tool: Vite 7.2.4
* Language: JavaScript (ES6+)

### DevOps & Infrastructure : 

* Containerization: Docker & Docker Compose
* Reverse Proxy: Nginx
* Hot Reload: Air (Go development)
* Version Control: Git

### Development Tools 

* Code Editor: VS Code (recommended)
* API Testing: GraphQL Playground (built-in)
* Database Client: pgAdmin / DBeaver

## Getting Started 
### Prerequisites :
###### Make sure you have the following installed:

* Docker & Docker Compose (v2.0+)
* Go 1.21+ (for local development)
* Node.js 18+ (for local development)
* PostgreSQL 15+ (if running without Docker)

### Quick Start with Docker :

#### 1)- Clone the repository 
```bash
git clone https://github.com/ilyes-rhdi/buildit-Gql.git
cd buildit-Gql

```
#### 2)- Configure environment variables

##### Create .env file in backend folder:
###### example : 
```bash
cd backend
cp .env.example .env
```
##### Edit backend/.env:
```env
   DATABASE_URL=postgresql://postgres:password@db:5432/notionql
   JWT_SECRET=your-super-secret-key-change-this
   PORT=8080
   REDIS_ADDR=redis:6379
   REDIS_PASSWORD=
   REDIS_DB=0

```
##### Create .env file in frontend folder:
```bash
cd frontend
cp .env.example .env
```

##### Edit frontend/.env:
###### Exemple :
```env
VITE_GRAPHQL_URL=http://localhost:8080/graphql
VITE_WS_URL=ws://localhost:8080/graphql
```
#### 3)- Start all services
###### Exemple :
```bash
# From project root
docker-compose up --build
```
#### 4)-Access the application
###### Exemple :
* Frontend: http://localhost
* GraphQL Playground: http://localhost:8080/graphql

### Local Development Setup
#### Backend Setup
```bash
cd backend

# Install dependencies
go mod download

# Run database migrations (create tables)
go run cmd/migrate/main.go

# Start development server with hot reload
air

```
###### The backend will start on http://localhost:8080
##### Frontend Setup
```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```
###### The frontend will start on http://localhost:5173
##### Database Setup
###### If running PostgreSQL locally:
```sql
-- Create database
CREATE DATABASE notionql;

-- Connect to database
\c notionql

-- Tables will be created automatically by GORM migrations
```
##### Testing the API
###### Once the server is running, open GraphQL Playground at:
``` 
http://localhost:8080/graphql

```
###### Try this test query:
```graphql
query {
  me {
    id
    email
    name
  }
} 
```
###### Note: You'll need to authenticate first to get a JWT token.
### Docker Commands Reference
```bash
# Start services
docker-compose up

# Start in background
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# Rebuild containers
docker-compose up --build

# Remove all containers and volumes
docker-compose down -v
```

## Project Structure 

```notionql/
├── backend/
│   ├── cmd/
│   │   ├── server/
│   │   │   └── main.go              # Application entry point
│   │   └── migrate/
│   │       └── main.go              # Database migrations
│   │
│   ├── internal/
│   │   ├── gql/
│   │   │   ├── mutations/
│   │   │   │   └── root_mutation.go # GraphQL mutations
│   │   │   ├── queries/
│   │   │   │   └── root_query.go    # GraphQL queries
│   │   │   ├── resolvers/
│   │   │   │   ├── user.go          # User resolver
│   │   │   │   ├── workspace.go     # Workspace resolver
│   │   │   │   ├── page.go          # Page resolver
│   │   │   │   ├── block.go         # Block resolver
│   │   │   │   └── helpers.go       # Mapping functions
│   │   │   ├── types/
│   │   │   │   ├── user.go          # User GraphQL type
│   │   │   │   ├── workspace.go     # Workspace GraphQL type
│   │   │   │   ├── page.go          # Page GraphQL type
│   │   │   │   └── block.go         # Block GraphQL type
│   │   │   ├── schema.go            # Schema definition
│   │   │   └── execute.go           # GraphQL handler setup
│   │   │
│   │   ├── models/
│   │   │   ├── user.go              # User data model
│   │   │   ├── workspace.go         # Workspace data model
│   │   │   ├── page.go              # Page data model
│   │   │   └── block.go             # Block data model
│   │   │
│   │   ├── services/
│   │   │   ├── profile.go           # User service
│   │   │   ├── workspace.go         # Workspace service
│   │   │   ├── page.go              # Page service
│   │   │   └── block.go             # Block service
│   │   │
│   │   ├── middlewares/
│   │   │   └── gql/
│   │   │       └── auth.go          # JWT authentication
│   │   │
│   │   └── database/
│   │       └── connection.go        # DB setup
│   │
│   ├── config/
│   │   └── config.go                # App configuration
│   │
│   ├── .env                         # Environment variables
│   ├── .air.toml                    # Hot reload config
│   ├── go.mod                       # Go dependencies
│   ├── go.sum                       # Dependency checksums
│   └── Dockerfile                   # Backend container
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Workspace/           # Workspace components
│   │   │   ├── Page/                # Page components
│   │   │   ├── Block/               # Block components
│   │   │   └── Common/              # Shared components
│   │   │
│   │   ├── pages/
│   │   │   ├── app/
│   │   │   │   ├── AppHome.jsx      # App home page
│   │   │   │   └── WorkspacePage.jsx # Workspace view
│   │   │   ├── auth/
│   │   │   │   ├── login.jsx        # Login page
│   │   │   │   ├── signup.jsx       # Signup page
│   │   │   │   └── verify.jsx       # Email verification
│   │   │   └── errors/
│   │   │       └── ErrorPage.jsx    # 404 page
│   │   │
│   │   ├── context/
│   │   │   ├── AuthContext.jsx      # Auth state management
│   │   │   └── UseAuth.js           # Auth hook
│   │   │
│   │   ├── graphql/
│   │   │   ├── queries.js           # GraphQL queries
│   │   │   ├── mutations.js         # GraphQL mutations
│   │   │   └── client.js            # Apollo Client setup
│   │   │
│   │   ├── App.jsx                  # Main app component
│   │   └── main.jsx                 # React entry point
│   │
│   ├── public/                      # Static assets
│   ├── .env                         # Frontend env vars
│   ├── package.json                 # NPM dependencies
│   ├── vite.config.js               # Vite configuration
│   └── index.html                   # HTML template
│
├── nginx/
│   ├── nginx.conf                   # Nginx configuration
│   └── Dockerfile                   # Nginx container
│
├── docker-compose.yml               # Multi-container orchestration
├── .gitignore                       # Git ignore rules
└── README.md                        # This file
```
### Key Folders Explained

#### Backend Structure:

* ```cmd/ ``` : Entry points for different executables
* ```internal/gql/ ```: All GraphQL-related code (types, resolvers, schema)
* ```internal/models/ ``` : Database models (GORM structs)
* ```internal/services/ ```: Business logic layer
* ```internal/middlewares/ ```: JWT auth and request processing

#### Frontend Structure:

* ```src/components/ ```: Reusable React components
* ```src/pages/ ``` :Route-based page components
* ```src/context/ ```: React Context for global state
* ```src/graphql/ ``` :  Apollo Client configuration and operations
