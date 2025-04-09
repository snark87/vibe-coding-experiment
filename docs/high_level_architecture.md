# Quantum Circuit Editor - High-Level Architecture (Updated)

## 1. Introduction

This document outlines the high-level architecture for the Quantum Circuit Editor MVP, an educational tool for students and quantum computing enthusiasts with little to no prior quantum computing or programming experience. The architecture described here is designed to fulfill the functional requirements specified in the MVP documentation while providing a foundation that can scale for future enhancements.

This is an updated version that addresses the review feedback from the architecture review document.

## 2. Key Use Cases

Before diving into the technical architecture, it's important to understand the primary use cases that drive the system design. These use cases represent the core interactions that users will have with the Quantum Circuit Editor.

```puml
@startuml
left to right direction
actor "Student/Educator" as User

rectangle "Quantum Circuit Editor" {
  usecase "Create New Circuit" as UC1
  usecase "Add/Remove Quantum Gates" as UC2
  usecase "Simulate Circuit" as UC3
  usecase "View Simulation Results" as UC4
  usecase "Save Circuit" as UC5
  usecase "Load Existing Circuit" as UC6
  usecase "Export Circuit to QASM" as UC7
}

User --> UC1
User --> UC2
User --> UC3
User --> UC4
User --> UC5
User --> UC6
User --> UC7

UC2 ..> UC1 : <<extends>>
UC3 ..> UC2 : <<includes>>
UC4 ..> UC3 : <<includes>>
@enduml
```

### 2.1 Primary Use Cases

1. **Create New Circuit**
   - User creates a blank quantum circuit with 2-5 qubits
   - User sets basic circuit parameters (e.g., number of qubits)
   - System displays an empty circuit with the specified number of qubit lines

2. **Add/Remove Quantum Gates**
   - User drags quantum gates (X, Y, Z, H, CNOT) from the gate palette onto qubit lines
   - User removes gates by dragging them off the circuit or using a delete function
   - User repositions gates on the circuit canvas
   - System provides visual feedback during the drag-and-drop operations
   - System validates gate placement and connections

3. **Simulate Circuit**
   - User clicks a "Simulate" button to run the quantum simulation
   - System processes the circuit through the simulation engine
   - System calculates the quantum state and measurement probabilities
   - System indicates simulation progress during calculation

4. **View Simulation Results**
   - System displays the probability distribution of possible measurement outcomes
   - User interprets the simulation results to understand circuit behavior
   - User can toggle between different visualization modes (if available in MVP)

5. **Save Circuit**
   - User saves the current circuit design to their account
   - User provides a name and optional description for the circuit
   - System confirms successful save operation
   - System associates the circuit with the user's account

6. **Load Existing Circuit**
   - User views a list of their previously saved circuits
   - User selects a circuit to load
   - System loads the circuit data and displays it on the canvas
   - User continues editing the loaded circuit

7. **Export Circuit to QASM**
   - User exports the current circuit to QASM format
   - System generates the QASM representation of the circuit
   - User downloads or copies the QASM code for use in other tools

### 2.2 User Flows

The following sequence diagrams illustrate the typical flow for two key user journeys:

#### 2.2.1 Circuit Design and Simulation Flow

```puml
@startuml
actor User
participant "Circuit Canvas" as Canvas
participant "Gate Palette" as Palette
participant "Simulation Engine" as Engine
participant "Results Visualizer" as Results

User -> Canvas: Create new circuit
User -> Palette: Select X gate
User -> Canvas: Place X gate on qubit 1
User -> Palette: Select H gate
User -> Canvas: Place H gate on qubit 2
User -> Palette: Select CNOT gate
User -> Canvas: Connect qubits 1 and 2 with CNOT
User -> Canvas: Request simulation
Canvas -> Engine: Send circuit for simulation
Engine -> Engine: Perform quantum calculation
Engine -> Results: Return probability distribution
Results -> User: Display measurement outcomes
@enduml
```

#### 2.2.2 Save and Export Flow

```puml
@startuml
actor User
participant "Circuit Canvas" as Canvas
participant "Control Panel" as Panel
participant "Backend Services" as Backend
participant "Storage" as Storage

User -> Canvas: Create/modify circuit
User -> Panel: Click "Save Circuit"
Panel -> User: Prompt for circuit name/description
User -> Panel: Enter circuit information
Panel -> Backend: Send save request
Backend -> Storage: Store circuit data
Storage -> Backend: Confirm save
Backend -> Panel: Return success message
Panel -> User: Show success notification
User -> Panel: Click "Export to QASM"
Panel -> Backend: Request QASM conversion
Backend -> Panel: Return QASM code
Panel -> User: Display/download QASM code
@enduml
```

## 3. System Overview

The Quantum Circuit Editor is a web-based application that enables users to:
- Create and edit quantum circuits through a drag-and-drop interface
- Simulate small quantum circuits (2-5 qubits)
- Visualize simulation results as probability distributions
- Save, load, and export circuit designs

### 3.1 Architectural Approach: Simplified Monolith for MVP

Based on the architecture review, we're adopting a simplified monolithic approach for the MVP to reduce complexity and accelerate development:

```puml
@startuml
' System layers
rectangle "Presentation Layer" as PresentationLayer
rectangle "Application Layer" as ApplicationLayer
rectangle "Domain Layer" as DomainLayer

' Implementation components
rectangle "Frontend\n(TypeScript/React)" as Frontend
rectangle "Backend API\n(Go)" as Backend
rectangle "Persistence" as Persistence
rectangle "Quantum Simulation Module\n(Python API)" as SimulationEngine

' Connections between layers
PresentationLayer <--> ApplicationLayer
ApplicationLayer <--> DomainLayer

' Connections between implementation components
Frontend <--> Backend
Backend <--> Persistence
Backend <--> SimulationEngine

' Layout
PresentationLayer -[hidden]d- Frontend
ApplicationLayer -[hidden]d- Backend
DomainLayer -[hidden]d- Persistence
@enduml
```

### 3.2 Technology Stack Justification

While we're using three different languages (TypeScript, Go, Python), this is based on the following justifications:

- **Frontend (TypeScript/React)**: Team has strong expertise and the interactive nature of the application requires a robust UI framework
- **Backend (Go)**: Team has expertise in Go, which offers excellent performance for web services and API handling
- **Simulation (Python)**: Leveraging existing Python quantum computing libraries (like Qiskit or Cirq) rather than building simulation capabilities from scratch

### 3.3 Cross-Language Integration Strategy

To mitigate the complexity of the multi-language approach, we'll implement:

- A clean, well-defined API between components
- A simple REST API for Frontend-Backend communication
- Python simulation code exposed as a straightforward HTTP service
- Comprehensive integration tests and documentation

## 4. Component Architecture

### 4.1 Frontend (TypeScript + React)

The frontend is responsible for the user interface and direct user interactions. It is built using TypeScript and React to leverage the team's expertise.

```puml
@startuml
package "Frontend (TypeScript + React)" {
  [Circuit Canvas] as Canvas
  [Gate Palette] as Palette
  [Control Panel] as Panel
  [Results Visualizer] as Results

  database "Circuit State" as State

  Canvas --> State
  Palette --> Canvas
  Panel --> State
  State --> Results
}
@enduml
```

Key components:
- **Circuit Canvas**: The main workspace for circuit design and visualization
- **Gate Palette**: Collection of available quantum gates (X, Y, Z, H, CNOT)
- **Control Panel**: Interface for simulation control, saving, loading, and exporting
- **Results Visualizer**: Component for displaying simulation results as probability distributions

State management will be handled using React's built-in state management capabilities, with a centralized state for the circuit representation.

#### 4.1.1 State Management Strategy

For the MVP, we'll use React Context API combined with useReducer for managing the quantum circuit state:

```typescript
// Simplified state management approach
interface CircuitState {
  qubits: number;
  gates: Gate[];
  simulationResults: SimulationResult | null;
  isDirty: boolean;
  history: CircuitHistoryState[];
  historyIndex: number;
}

// Actions for state management
type CircuitAction =
  | { type: 'ADD_GATE', gate: Gate }
  | { type: 'REMOVE_GATE', gateId: string }
  | { type: 'MOVE_GATE', gateId: string, newPosition: Position }
  | { type: 'SET_SIMULATION_RESULTS', results: SimulationResult }
  | { type: 'UNDO' }
  | { type: 'REDO' };
```

This approach provides:
- Circuit state tracking
- Undo/redo functionality
- Change detection for save prompts
- Centralized state for all components

### 4.2 Backend Services (Go)

The backend services handle business logic, user authentication, and coordination between the frontend and the quantum simulation engine.

For the MVP, we'll implement a simplified monolithic backend in Go:

```puml
@startuml
package "Backend Services (Go)" {
  [HTTP Router] as Router
  [Authentication Service] as Auth
  [Circuit Manager] as CircuitMgr
  [Simulation API Client] as SimClient
  [Database Access Layer] as DAL

  Router --> Auth
  Router --> CircuitMgr
  Router --> SimClient
  CircuitMgr --> DAL
}
@enduml
```

Key components:
- **HTTP Router**: Entry point for frontend requests
- **Authentication Service**: Manages user authentication with Google
- **Circuit Manager**: Handles saving, loading, and validating circuit designs
- **Simulation API Client**: Communicates with the Python simulation service
- **Database Access Layer**: Handles persistence operations

#### 4.2.1 API Endpoints

The backend will expose the following RESTful endpoints:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/auth/login` | POST | Authenticate user with Google |
| `/api/auth/logout` | POST | Log out current user |
| `/api/circuits` | GET | Get list of user's circuits |
| `/api/circuits` | POST | Create new circuit |
| `/api/circuits/{id}` | GET | Get specific circuit |
| `/api/circuits/{id}` | PUT | Update circuit |
| `/api/circuits/{id}` | DELETE | Delete circuit |
| `/api/simulate` | POST | Run circuit simulation |
| `/api/export/qasm` | POST | Convert circuit to QASM |

### 4.3 Quantum Simulation Module (Python)

The quantum simulation engine is responsible for the scientific calculations required to simulate quantum circuits.

The simulation module will be a lightweight HTTP service exposing Python-based quantum simulation:

```puml
@startuml
package "Quantum Simulation Module (Python)" {
  [HTTP API] as API
  [Circuit Validator] as Validator
  [Circuit Simulator] as Simulator
  [Visualization Helper] as Visualizer
  [QASM Generator] as QASMGen

  API --> Validator
  API --> Simulator
  API --> Visualizer
  API --> QASMGen
}
@enduml
```

#### 4.3.1 Simulation Performance Considerations

For the MVP with 2-5 qubits:
- State vectors are manageable (2^5 = 32 amplitudes)
- Matrix operations use NumPy for optimization
- Precomputed gate operations reduce calculation time
- Response time targets < 500ms for typical circuits

We'll leverage an existing quantum computing library like Qiskit, which is optimized for these types of calculations.

## 5. Data Architecture

### 5.1 Circuit Data Model

```puml
@startuml
class Circuit {
  +id: UUID
  +name: String
  +description: String
  +creationDate: DateTime
  +lastModified: DateTime
  +userId: String
  +qubits: Integer
  +gates: List<Gate>
}

class Gate {
  +id: UUID
  +type: GateType
  +position: Integer
  +targets: List<Integer>
  +controls: List<Integer>
}

enum GateType {
  X
  Y
  Z
  H
  CNOT
}

Circuit "1" *-- "*" Gate
Gate -- GateType
@enduml
```

### 5.2 Database Selection: PostgreSQL

For the MVP, we'll use PostgreSQL as our database for several reasons:
- Strong support for JSON data types (for storing circuit configurations)
- ACID compliance for data integrity
- Excellent support in Go ecosystem
- Familiar to the development team
- Easy to host in cloud environments

## 6. Integration Architecture

### 6.1 Frontend-Backend Integration

```puml
@startuml
participant "React Frontend" as Frontend
participant "Go Backend" as Backend
participant "Python Simulation" as Simulation
database "PostgreSQL" as DB

Frontend -> Backend: HTTP Request (JSON)
Backend -> DB: Query/Update
DB --> Backend: Result
Backend -> Simulation: Simulation Request (JSON)
Simulation --> Backend: Simulation Results (JSON)
Backend --> Frontend: HTTP Response (JSON)
@enduml
```

### 6.2 Authentication Flow

```puml
@startuml
actor User
participant "Frontend" as Frontend
participant "Google OAuth" as OAuth
participant "Backend" as Backend

User -> Frontend: Click "Login with Google"
Frontend -> OAuth: Redirect to Google Auth
User -> OAuth: Enter credentials
OAuth --> Frontend: Auth code
Frontend -> Backend: Auth code
Backend -> OAuth: Verify token
OAuth --> Backend: User info
Backend -> Backend: Create session
Backend --> Frontend: JWT token
Frontend -> Frontend: Store token
@enduml
```

## 7. Deployment Architecture

For the MVP, we'll use a simplified deployment architecture:

```puml
@startuml
cloud "Google Cloud Platform" {
  node "Cloud Run" as FrontendService {
    [Frontend Container]
  }

  node "Cloud Run" as BackendService {
    [Backend Container]
  }

  node "Cloud Run" as SimulationService {
    [Simulation Container]
  }

  database "Cloud SQL" as DB {
    [PostgreSQL]
  }

  FrontendService --> BackendService: API calls
  BackendService --> SimulationService: Simulation requests
  BackendService --> DB: Data persistence
}

actor "User" as User
User --> FrontendService: Web access

node "Auth Provider" as Auth {
  [Google Auth]
}

User --> Auth: Authentication
Auth --> BackendService: Auth Token
@enduml
```

### 7.1 Resource Estimates

| Component | Resources | Estimated Monthly Cost |
|-----------|-----------|------------------------|
| Frontend | 1 Cloud Run instance | ~$15-30 |
| Backend | 1 Cloud Run instance | ~$15-30 |
| Simulation | 1 Cloud Run instance | ~$20-40 |
| Database | Small Cloud SQL instance | ~$25-50 |
| Storage | 10GB | ~$2-5 |
| **Total** | | **~$77-155** |

This estimate assumes modest usage for an educational MVP. Actual costs will depend on traffic patterns.

## 8. Security Considerations

### 8.1 Authentication and Authorization

- User authentication via Google OAuth 2.0
- JWT tokens for session management
- Role-based access control for circuit ownership
- HTTPS for all communications
- API rate limiting to prevent abuse

### 8.2 Data Protection

- Encryption at rest for database
- Input validation and sanitization for all API endpoints
- CSRF protection for web forms
- Security headers for frontend (Content Security Policy, etc.)

## 9. Error Handling

### 9.1 Frontend Error Handling

- Graceful error displays for users
- Connectivity loss handling with retry mechanisms
- Form validation with clear user feedback
- Client-side validation to catch errors early

### 9.2 Backend Error Handling

- Structured error responses (status code, message, details)
- Comprehensive error logging
- Graceful degradation when simulation service is unavailable
- Rate limiting with appropriate error messages

### 9.3 Simulation Error Handling

- Validation of circuit configurations before simulation
- Timeouts for long-running simulations
- Graceful error messages for invalid quantum operations

## 10. Testing Strategy

### 10.1 Frontend Testing

- Unit tests for React components using Jest and React Testing Library
- Integration tests for state management logic
- End-to-end tests using Cypress for critical user flows

### 10.2 Backend Testing

- Unit tests for business logic and API handlers
- Integration tests for database operations
- API contract tests to verify endpoint behavior

### 10.3 Simulation Testing

- Unit tests for quantum operations
- Validation tests against known quantum circuits with verified outputs
- Performance tests for various circuit sizes

## 11. Monitoring and Observability

### 11.1 Application Monitoring

- Request/response metrics for API endpoints
- Error rate tracking
- User activity metrics (circuits created, simulations run)
- Performance metrics for simulation operations

### 11.2 Infrastructure Monitoring

- CPU/memory usage for services
- Database performance metrics
- API latency tracking
- Error logs aggregation

We'll implement monitoring using Google Cloud Monitoring for the MVP.

## 12. Local Development Setup

For local development, we'll provide:

- Docker Compose configuration for running all services locally
- Development documentation with setup instructions
- Seeded database with sample circuits
- Environment configuration templates
- Local authentication bypass option for easier development

## 13. Technical Trade-offs and Alternatives

### 13.1 Technology Stack Alternatives Considered

| Component | Selected | Alternatives | Rationale |
|-----------|----------|--------------|-----------|
| Frontend | React + TypeScript | Vue.js, Angular | Team expertise, strong ecosystem |
| Backend | Go | Node.js, Python | Team expertise, performance |
| Simulation | Python | Julia, C++ | Existing libraries, scientific computing strength |
| Database | PostgreSQL | MongoDB, MySQL | JSON support, reliability |

### 13.2 Architecture Alternatives

| Aspect | Selected | Alternative | Rationale |
|--------|----------|-------------|-----------|
| Overall Architecture | Simplified Monolith | Microservices | Reduced complexity for MVP |
| Frontend State | Context+Reducer | Redux, MobX | Simpler learning curve, sufficient for MVP |
| API Style | REST | GraphQL | Familiarity, simplicity |
| Deployment | Cloud Run | Kubernetes | Easier operations, automatic scaling |

## 14. Future Extensibility

The architecture is designed to accommodate future enhancements:
- Support for additional quantum gates
- Integration with real quantum hardware
- More advanced visualization features
- Collaboration features

### 14.1 Potential Scaling Considerations

- Separating services for better scalability when needed
- Adding caching for frequently accessed circuits
- Optimizing database queries as the number of users grows
- Adding more sophisticated authentication and sharing options

## 15. Open Questions and Next Steps

### 15.1 Open Questions

- What are typical classroom usage patterns that might create usage spikes?
- What level of detail should be provided in the tutorial system?
- How important is real-time collaboration for educational use cases?

### 15.2 Next Steps

1. Finalize database schema design
2. Create detailed API specifications
3. Develop proof-of-concept for quantum simulation service
4. Implement authentication flow
5. Create UI mockups and component structure

## 16. Conclusion

This updated architecture document provides a balanced approach for the Quantum Circuit Editor MVP, addressing the review concerns while maintaining the core educational objectives of the application. By simplifying the deployment architecture, clarifying cross-service communication, and providing more detailed implementation guidance, we've created a more feasible plan for delivering a valuable educational tool.
