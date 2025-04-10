# Quantum Circuit Editor - Design System Specification

## 1. Introduction

This document defines the design system for the Quantum Circuit Editor application, providing a comprehensive guide to ensure consistency across the user interface. The design system serves as a reference for all UI development and addresses the improvement needed in the UI requirements review.

## 2. Design Principles

The Quantum Circuit Editor design system is built on the following principles:

### 2.1. Educational Focus
- Clarity over complexity
- Visuals that aid understanding of quantum concepts
- Progressive disclosure of advanced features

### 2.2. Consistency
- Unified visual language across all components
- Predictable patterns for user interaction
- Coherent terminology and iconography

### 2.3. Accessibility
- Inclusive design that works for all users
- Clear visual hierarchies
- Sufficient color contrast

### 2.4. Simplicity
- Minimal UI elements to reduce cognitive load
- Focus on essential functionality
- Clear paths to complete primary tasks

## 3. Color System

### 3.1. Primary Color Palette

| Variable | Hex Value | Usage |
|----------|-----------|-------|
| `--primary-color` | #4361ee | Primary buttons, active states, selected items |
| `--primary-light` | #5a74f1 | Hover states, highlights, secondary elements |
| `--primary-dark` | #304dc9 | Pressed states, borders, focused elements |
| `--secondary-color` | #3a0ca3 | Secondary UI elements, alternative emphasis |
| `--accent-color` | #7209b7 | Accent elements, highlighting important UI elements |

### 3.2. Semantic Colors

| Variable | Hex Value | Usage |
|----------|-----------|-------|
| `--success-color` | #2a9d8f | Success states, confirmations, valid inputs |
| `--warning-color` | #f4a261 | Warning states, alerts that need attention |
| `--danger-color` | #e63946 | Error states, destructive actions, critical alerts |

### 3.3. Neutral Colors

| Variable | Hex Value | Usage |
|----------|-----------|-------|
| `--white` | #ffffff | Backgrounds, text on dark backgrounds |
| `--light-gray` | #f8f9fa | Secondary backgrounds, subtle highlights |
| `--medium-gray` | #dee2e6 | Borders, dividers, disabled states |
| `--dark-gray` | #6c757d | Secondary text, icons, less emphasized content |
| `--black` | #212529 | Primary text, headings |

### 3.4. Quantum Gate Colors

Each quantum gate type has a dedicated color to aid recognition and understanding:

| Variable | Hex Value | Gate |
|----------|-----------|------|
| `--h-gate-color` | #4361ee | Hadamard gate (H) |
| `--x-gate-color` | #e63946 | Pauli-X gate (X) |
| `--y-gate-color` | #2a9d8f | Pauli-Y gate (Y) |
| `--z-gate-color` | #f4a261 | Pauli-Z gate (Z) |
| `--cnot-gate-color` | #7209b7 | Controlled NOT gate (CNOT) |

## 4. Typography

### 4.1. Font Family

The application uses a system font stack with 'Inter' as the preferred font:

```css
font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
```

For monospaced text (e.g., code, gate labels), use:

```css
font-family: 'Courier New', monospace;
```

### 4.2. Font Sizes

A structured type scale for consistent sizing:

| Element | Size | Weight | Line Height | Usage |
|---------|------|--------|-------------|-------|
| h1 | 2rem (32px) | 600 | 1.2 | Main screen titles |
| h2 | 1.5rem (24px) | 600 | 1.2 | Section headers, panel titles |
| h3 | 1.25rem (20px) | 600 | 1.2 | Subsection titles, feature headers |
| Body | 1rem (16px) | 400 | 1.5 | Primary text content |
| Small | 0.875rem (14px) | 400 | 1.5 | Secondary text, labels, helper text |
| XSmall | 0.75rem (12px) | 400 | 1.5 | Metadata, timestamps, small labels |

### 4.3. Font Weights

| Weight | Usage |
|--------|-------|
| 400 | Regular text, body content |
| 500 | Medium emphasis, button text, user name |
| 600 | Headings, strong emphasis |
| 700 | Extra emphasis (use sparingly) |

## 5. Spacing System

A consistent spacing scale for margins, paddings, and layout gaps:

| Size | Value | Usage |
|------|-------|-------|
| Extra small | 0.25rem (4px) | Minimal separation, tight grouping |
| Small | 0.5rem (8px) | Close elements, button padding |
| Medium | 1rem (16px) | Standard spacing, form controls |
| Large | 1.5rem (24px) | Section separation, panel padding |
| Extra large | 2rem (32px) | Major section divisions |
| 2x Extra large | 3rem (48px) | Screen-level padding, major separations |

## 6. Shadows and Elevation

### 6.1. Shadow Variables

| Variable | Value | Usage |
|----------|-------|-------|
| `--shadow-sm` | 0 1px 2px rgba(0,0,0,0.05) | Subtle elevation, cards, buttons |
| `--shadow-md` | 0 4px 6px rgba(0,0,0,0.1) | Moderate elevation, popovers, dropdowns |
| `--shadow-lg` | 0 10px 15px rgba(0,0,0,0.1) | High elevation, modals, dialogs |

### 6.2. Elevation Principles

- Use elevation to establish visual hierarchy
- Higher elevation indicates more prominent UI elements
- Avoid excessive nesting of elevated elements

## 7. Border Radius

| Variable | Value | Usage |
|----------|-------|-------|
| `--border-radius-sm` | 4px | Small elements, tags, inputs |
| `--border-radius-md` | 8px | Buttons, cards, panels |
| `--border-radius-lg` | 12px | Modals, larger containers |
| `--border-radius-xl` | 20px | Hero elements, feature highlights |

## 8. Component Library

### 8.1. Button Components

#### Primary Button
- Background: `--primary-color`
- Text color: `--white`
- Padding: 0.5rem 1rem
- Border radius: `--border-radius-md`
- Font weight: 500
- Hover state: Background `--primary-light`, slight elevation
- Active state: No elevation, background `--primary-dark`

#### Secondary Button
- Background: `--medium-gray`
- Text color: `--black`
- Same dimensions as primary button
- Hover state: Background `--dark-gray`, text `--white`

#### Danger Button
- Background: `--danger-color`
- Text color: `--white`
- Same dimensions as primary button
- Hover state: Darker red (#d32f2f)

#### Icon Button
- Size: 40px × 40px
- Border radius: 50% (circular)
- Background: transparent
- Icon color: `--dark-gray`
- Hover state: Background rgba(0,0,0,0.05), icon color `--black`

#### Small Button
- Padding: 0.25rem 0.75rem
- Font size: 0.875rem
- Otherwise same as primary/secondary buttons

### 8.2. Input Components

#### Text Input
- Border: 1px solid `--medium-gray`
- Border radius: `--border-radius-md`
- Padding: 0.5rem 1rem
- Focus state: Border color `--primary-color`

#### Number Input
- Contains increment/decrement buttons
- Input width: 40px
- Text alignment: center
- Border radius for decrement button: `--border-radius-md` on left side only
- Border radius for increment button: `--border-radius-md` on right side only

### 8.3. Panel Components

#### Standard Panel
- Background: `--white`
- Shadow: `--shadow-sm`
- Border radius: `--border-radius-md`
- Header: Padding 0.75rem 1rem, border-bottom 1px solid `--medium-gray`
- Content area: Padding 1rem, overflow-y: auto

#### Collapsible Panel
- Collapsed state: flex-basis 48px
- Transition: flex-basis 0.3s ease
- Header remains visible in collapsed state
- Content hidden when collapsed

### 8.4. Card Components

#### Welcome Card
- Background: `--white`
- Shadow: `--shadow-lg`
- Border radius: `--border-radius-lg`
- Two-panel layout on desktop, stacked on mobile
- Content padding: 3rem (desktop), 1.5rem (mobile)

#### Tutorial Card
- Max width: 600px
- Background: `--white`
- Shadow: `--shadow-lg`
- Border radius: `--border-radius-lg`
- Header with border-bottom
- Content padding: 2rem

### 8.5. Circuit Editor Components

#### Gate Visual
- Size: 40px × 40px
- Border radius: `--border-radius-sm`
- Background color: Gate-specific color
- Font: Bold, 'Courier New', monospace
- Text color: `--white`

#### Circuit Grid
- Background: `--white`
- Border radius: `--border-radius-md`
- Shadow: `--shadow-sm`
- Padding: 1rem 2rem

#### Qubit Line
- Height: 60px
- Wire: 2px height, `--black` color
- Qubit label: monospace, bold

#### Trash Zone
- Circular (border-radius: 50%)
- Size: 100px × 100px
- Border: 2px dashed `--danger-color`
- Background: rgba(230, 57, 70, 0.15)
- Hover state: Background rgba(230, 57, 70, 0.3), scale 1.05

### 8.6. Results Components

#### Probability Bar Chart
- Bar container: height 200px, background `--light-gray`
- Bar: Background `--primary-color`
- Bar border radius: `--border-radius-sm` on top only
- Label font: monospace, 0.75rem

#### Results Summary
- Text color: `--dark-gray`
- Font size: 0.875rem

## 9. Layout System

### 9.1. Screen Layout

The main application interface follows a three-panel layout:
1. Left panel: Gate palette (collapsible)
2. Center panel: Circuit canvas (flexible width)
3. Right panel: Results (collapsible)

Header remains fixed at top with consistent height of 64px.

### 9.2. Grid System

For components requiring grid layouts (e.g., gate palette):
- Use CSS Grid with responsive column counts
- Standard gap: 1rem
- Defaults: 2 columns for gate palette, flexible for circuit management

### 9.3. Responsive Breakpoints

| Breakpoint | Value | Targeting |
|------------|-------|-----------|
| Small | 576px | Mobile phones |
| Medium | 768px | Tablets, small laptops |
| Large | 992px | Laptops |
| Extra large | 1200px | Desktops, large screens |

### 9.4. Responsive Behaviors

- Below 1200px: Welcome card stacks vertically
- Below 992px: Side panels become absolutely positioned with show/hide toggles
- Below 768px: Header becomes multi-row, toolbar wraps
- Below 576px: Feature sections stack vertically

## 10. Motion and Animation

### 10.1. Transitions

- Standard duration: 0.2s - 0.3s
- Easing: ease, ease-in-out
- Properties commonly animated: transform, opacity, background-color

### 10.2. Animation Patterns

#### Hover Animation
- Button hover: translateY(-1px)
- Gate hover: translateY(-2px), add shadow

#### Pulse Animation
- Used for attention-grabbing elements
- Scale 1 → 1.05 → 1 over 2s, infinite
- Apply with `.pulse-animation` class

## 11. Iconography Guidelines

### 11.1. Icon Style
- Clean, simple line icons
- Consistent stroke width
- Use semantic colors based on context
- Size: 24px default, 16px for small contexts

### 11.2. Icon Usage
- Always accompany standalone icons with tooltips or text labels
- Use icons consistently for the same actions throughout the application
- Maintain padding around icons within containers (minimum 8px)

## 12. Accessibility Standards

### 12.1. Color Contrast
- Text and interactive elements should maintain at least AA compliance (4.5:1 ratio)
- Use tools like Wave or Lighthouse to verify contrast compliance

### 12.2. Focus States
- All interactive elements must have visible focus states
- Do not rely solely on color to indicate focus

### 12.3. Text Alternatives
- Provide alt text for all images
- Label all form elements clearly
- Use ARIA attributes appropriately

### 12.4. Keyboard Navigation
- Ensure logical tab order
- Provide keyboard alternatives for drag operations
- Ensure all interactive elements are accessible via keyboard

## 13. Usage Guidelines

### 13.1. Implementation Steps

1. Include the design system CSS file in your project
2. Follow the component patterns defined in this document
3. Use the design system variables for all styling
4. Reference the component examples for proper implementation

### 13.2. Component Extensions

When adding new components:
1. Follow existing design patterns
2. Use the defined color palette, typography, and spacing
3. Document the new component in this specification
4. Ensure accessibility compliance

### 13.3. Design Reviews

All new UI components should be reviewed against this design system for:
1. Visual consistency with existing components
2. Proper use of design tokens (colors, spacing, etc.)
3. Responsive behavior
4. Accessibility compliance

## 14. Version Control

### 14.1. Design System Versioning
- Major version: Significant visual or structural changes
- Minor version: New components or significant component alterations
- Patch version: Bug fixes, minor adjustments

### 14.2. Changelog Management
- Document all changes with rationale
- Provide migration guidance for breaking changes
- Include visual examples of changes where applicable

## 15. Implementation Notes

This design system is currently implemented through CSS custom properties (variables) and can be found in:
- `/ui-design/styles.css` - Main implementation
- `/quantum-circuit-editor/frontend/src/index.css` - Frontend implementation

When updating the design system, ensure both files are kept in sync.

## 16. Future Extensions

Planned enhancements for the design system (post-MVP):
1. Component library implementation in a JavaScript framework
2. Design tokens in a format that can be consumed by multiple platforms
3. Animation library for consistent motion across the application
4. Extended guidance for data visualization components
5. Dark mode theme
