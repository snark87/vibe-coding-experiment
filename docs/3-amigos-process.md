# 3 Amigos Meeting Process Guidelines

## Overview

This document outlines the structured approach for conducting effective 3 Amigos meetings when using LLM agents as participants. Based on learnings from our first LLM-driven 3 Amigos sessions, these guidelines aim to maximize collaboration effectiveness while addressing the unique dynamics of AI-assisted meetings.

## Meeting Participants

A 3 Amigos meeting typically involves:

1. **Product Owner/Business Analyst (Required)**
   - Represents business needs and user requirements
   - Owns feature prioritization
   - Makes final scope decisions

2. **Developer/Architect (Required)**
   - Provides technical feasibility assessment
   - Identifies implementation approaches
   - Raises architectural concerns

3. **QA/Tester (Required)**
   - Develops testing scenarios
   - Identifies edge cases
   - Ensures quality and requirement testability

4. **Domain Expert (Optional but Recommended)**
   - Provides specific domain knowledge (e.g., Quantum Science)
   - Validates technical accuracy of domain-specific features
   - Ensures adherence to domain standards

5. **UI/UX Designer (Contextual)**
   - Provides user experience perspective
   - Ensures usability and accessibility
   - Represents the voice of the user

6. **Meeting Facilitator (Critical with LLM Participants)**
   - Structures conversation flow
   - Ensures balanced participation
   - Manages time boxing
   - Documents decisions and action items

## Pre-Meeting Process

1. **Requirement Distribution (T-3 days)**
   - Product Owner shares initial user stories/requirements
   - All participants review materials independently
   - Questions and clarifications captured in a shared document

2. **Agenda Setting (T-1 day)**
   - Prioritized list of discussion items
   - Time allocation for each topic
   - Explicit indication of required decisions
   - Clear meeting outcomes/deliverables

3. **Participant Prompting (For LLM Agents)**
   - Specific role assignment with clear boundaries
   - Context sharing with relevant documentation
   - Explicit instructions on contribution focus
   - Prompt templates for each role to ensure consistency

4. **Preparation Checklist**
   - User stories/requirements finalized
   - Acceptance criteria drafted
   - Technical constraints documented
   - Meeting templates prepared
   - Previous related decisions referenced

## Meeting Structure

1. **Opening (10 minutes)**
   - Review meeting purpose and desired outcomes
   - Introduce participants and their roles
   - Set collaboration expectations
   - Review agenda and time allocations

2. **Feature Overview (15 minutes)**
   - Product Owner presents feature requirements
   - Business context and user needs clarification
   - Explicit MVP boundary definition
   - Confirmation of understanding from all participants

3. **Structured Exploration (Time-boxed per feature)**
   - **Round 1: Questions (5-10 min)**
     - All roles ask clarifying questions
     - No solutions proposed yet
     - Focus on understanding requirements

   - **Round 2: Concerns (10-15 min)**
     - Each role raises potential issues from their perspective
     - Documented in shared template
     - No solutions yet, just problem identification

   - **Round 3: Solution Options (15-20 min)**
     - Each role proposes potential approaches
     - Options documented with pros/cons
     - No decision-making yet

   - **Round 4: Decision Framework (10-15 min)**
     - Establish decision criteria matrix
     - Weight criteria based on product priorities
     - Score proposed solutions against criteria

4. **Decision Making (10 minutes per decision point)**
   - Apply structured decision framework
   - Document decision rationale
   - Assign action owners
   - Set completion timeframes

5. **Action Item Capture (10 minutes)**
   - Document specific next steps
   - Assign responsible parties
   - Set due dates
   - Define expected deliverables

6. **Meeting Summary (5 minutes)**
   - Recap decisions made
   - Review action items
   - Confirm next meeting/follow-up

## Structured Decision Framework

To address the "Unclear Decision Ownership" challenge identified in our retrospective, all decisions will follow this framework:

### 1. Decision Matrix Template

| Option | Technical Feasibility (1-5) | User Value (1-5) | MVP Fit (1-5) | Implementation Effort (1-5) | Educational Value (1-5) | Total Score |
|--------|------------------------------|------------------|---------------|------------------------------|-------------------------|-------------|
| Option 1 | | | | | | |
| Option 2 | | | | | | |

### 2. Decision Owner Assignment

Each decision point must have a clear owner based on this guidance:
- **Product decisions**: Product Owner
- **Technical approach decisions**: Lead Developer/Architect
- **Quality standards decisions**: QA Lead
- **Domain accuracy decisions**: Domain Expert
- **User experience decisions**: UI/UX Designer with Product Owner input

### 3. Decision Documentation Standard

All decisions must be documented with:
- Summary of options considered
- Decision criteria and weightings
- Selected approach with rationale
- Rejected alternatives with reasons
- Implementation implications
- Testing considerations

## Managing LLM Participants Effectively

Based on our retrospective findings, these specific considerations will help maximize LLM agent effectiveness:

### 1. Context Management

- Provide comprehensive but focused context to each LLM role
- Use consistent terminology across all communications
- Reference specific documents rather than summarizing them
- Maintain a shared knowledge base accessible to all LLMs

### 2. Role Clarity

- Clearly define the perspective each LLM should adopt
- Specify the types of contributions expected from each role
- Set boundaries on discussion scope for each participant
- Use role-specific prompt templates for consistency

### 3. Collaborative Documentation

- Use real-time shared documentation during meetings
- Implement standardized templates for all output types
- Structure discussions to build on previous contributions
- Maintain a single source of truth for all decisions

### 4. Preventing Over-Engineering

- Explicitly define MVP boundaries before solution discussions
- Implement a "future considerations" parking lot
- Use complexity budgets for each feature component
- Require justification for any solution exceeding basic implementation

### 5. Balanced Participation

- Implement structured rounds where each role contributes in turn
- Set time limits for each role's input
- Use explicit prompting to bring in quieter perspectives
- Track participation metrics to ensure balance

## Output Templates

### 1. Meeting Summary Template

```
# 3 Amigos Meeting Summary - [Feature ID]

## Participants
- Product: [Name]
- Development: [Name]
- QA: [Name]
- Domain Expert: [Name]
- UI/UX: [Name]
- Facilitator: [Name]

## Decisions Made
1. [Decision 1] - Owner: [Name], Due: [Date]
2. [Decision 2] - Owner: [Name], Due: [Date]

## Action Items
1. [Action 1] - Owner: [Name], Due: [Date]
2. [Action 2] - Owner: [Name], Due: [Date]

## Open Questions
1. [Question 1] - Owner: [Name], Resolution Plan: [Plan]
2. [Question 2] - Owner: [Name], Resolution Plan: [Plan]

## Next Meeting
Date: [Date]
Focus: [Topic]
```

### 2. User Story Refinement Template

```
# User Story: [ID] - [Title]

## Original Requirement
[Paste original story]

## Refined Requirement
[Updated story with clarifications]

## Acceptance Criteria
1. [Criterion 1]
2. [Criterion 2]

## Technical Approach
[Summary of implementation approach]

## Testing Approach
[Summary of testing strategy]

## MVP Boundary
In scope:
- [Feature aspect 1]
- [Feature aspect 2]

Out of scope:
- [Non-MVP aspect 1]
- [Non-MVP aspect 2]
```

### 3. Decision Record Template

```
# Decision Record: [Feature]-[Component]-[Number]

## Decision Required
[Clear statement of the decision needed]

## Options Considered
1. [Option 1]
   - Pros: [List]
   - Cons: [List]
2. [Option 2]
   - Pros: [List]
   - Cons: [List]

## Decision Criteria
1. [Criterion 1] - Weight: [1-5]
2. [Criterion 2] - Weight: [1-5]

## Decision Matrix
[Insert decision matrix with scores]

## Selected Option
[Option chosen]

## Rationale
[Explanation of why this option was selected]

## Implementation Implications
[What this means for development]

## Testing Implications
[What this means for QA]
```

## Continuous Improvement

After each 3 Amigos session:
1. Conduct a brief retrospective (max 15 minutes)
2. Capture at least one improvement to the process
3. Update this guidelines document accordingly
4. Share learnings with the broader team

## Special Considerations for LLM Agent Collaboration

1. **Prompt Engineering**
   - Maintain a library of effective prompts for each role
   - Continuously refine prompts based on meeting effectiveness
   - Include key terminology and constraints in all prompts
   - Use consistent formatting across all prompts

2. **Context Windows**
   - Be mindful of LLM context limitations
   - Provide essential information only to maintain focus
   - Use references to documentation rather than full inclusion
   - Break complex features into manageable discussion units

3. **Iterative Interaction**
   - Build discussions progressively through multiple exchanges
   - Use specific questions to guide LLM contributions
   - Request reformulations when responses lack clarity
   - Explicitly connect new information to previous context

4. **Human Oversight**
   - Designate a human facilitator for all LLM-driven meetings
   - Implement regular checkpoint reviews of LLM outputs
   - Maintain ability to override or redirect LLM contributions
   - Validate critical decisions with human expertise

## Conclusion

These guidelines aim to address the specific challenges identified in our retrospective while building on successful aspects of our LLM-driven 3 Amigos process. By implementing structured decision-making, clear role boundaries, and standardized documentation, we can improve meeting efficiency and outcome quality.

The process should remain flexible enough to adapt to different feature complexities while providing sufficient structure to ensure consistent, high-quality results. Regular retrospectives will continue to refine these guidelines based on ongoing experience.
