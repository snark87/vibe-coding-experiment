---
track: FRONTEND
task_type: EPIC
task_id: FEAT-0001
state: TODO
---
# Guest User Circuit Creation Epic

## Overview

This epic covers the foundational user journey for guest users to create quantum circuits without requiring authentication. This is a core functionality of the Quantum Circuit Editor, enabling users to quickly experiment with quantum computing concepts without the friction of account creation.

## Background

The Quantum Circuit Editor is being developed as an educational tool for students and quantum computing enthusiasts with little to no prior quantum computing or programming experience. Based on our UI and functional requirements, allowing guest users to create circuits is essential for reducing barriers to entry and encouraging experimentation.

## User Story

**As a** guest user without an account,
**I want to** create a basic quantum circuit using a drag-and-drop interface,
**So that** I can learn about quantum computing concepts without requiring authentication.

## Scope

### Included

1. A landing page that allows users to proceed as guests
2. Basic circuit editor workspace with drag-and-drop functionality
3. Limited gate set (H gate and CNOT only) for MVP
4. Session-based circuit preservation (circuit remains available within the same browser session)
5. Minimal inline help (tooltips only)

### Excluded

1. Account creation/authentication flow
2. Persistent saving of circuits between sessions
3. Export functionality to QASM
4. Additional quantum gates (X, Z, etc.)
5. Circuit management features
6. Comprehensive help/tutorial screens
7. Circuit simulation functionality

## Acceptance Criteria

1. Guest users can access the circuit editor directly from the landing page
2. Users can add 2-3 qubits to the workspace
3. Users can drag and drop H gates onto the circuit
4. Users can drag and drop CNOT gates connecting two qubits
5. Gates can be positioned and repositioned on the circuit
6. Circuit validation provides visual feedback for invalid configurations
7. The circuit state persists during the browser session
8. Basic tooltips provide guidance on gate functionality

## Definition of Done

1. Guest users can successfully create a complete quantum circuit without authentication
2. All acceptance criteria are met and verified
3. The implementation is responsive and works in major browsers
4. The feature is accessible with minimal friction (no more than 2 clicks from landing page)
5. Basic error handling is implemented for invalid operations
6. The UI provides clear visual feedback during all operations

## Dependencies

- UI Component Architecture implementation (ARC-0001-01)
- Minimal data model for representing circuits (ARC-0001-03)

## Technical Considerations

1. Circuit data should be stored in browser session storage
2. The implementation should follow the UI simplifications outlined in section 11.2 of the UI requirements
3. Focus on the core user flow: create circuit
4. Initial implementation should limit circuits to 3 qubits maximum as specified in the MVP simplification strategy

## Future Extensions

While out of scope for this epic, the following capabilities would naturally extend this feature:
- Prompting guest users to create accounts to save their work
- Adding export functionality for guest-created circuits
- Expanding available gates and qubit count
- Adding circuit simulation capabilities
