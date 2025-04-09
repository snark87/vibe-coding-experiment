# Quantum Circuit Editor Frontend

This directory contains the frontend application for the Quantum Circuit Editor project, built with React and TypeScript.

## Directory Structure

```
frontend/
├── src/                  # Source code
│   ├── assets/           # Static assets (images, fonts, etc.)
│   ├── components/       # React components
│   │   ├── circuit/      # Circuit-specific components
│   │   ├── common/       # Shared UI components
│   │   └── results/      # Result visualization components
│   ├── contexts/         # React context providers
│   ├── hooks/            # Custom React hooks
│   ├── services/         # API and service layer
│   ├── types/            # TypeScript type definitions
│   ├── utils/            # Utility functions
│   ├── App.tsx           # Main application component
│   ├── index.css         # Global styles
│   └── main.tsx          # Application entry point
├── tests/                # Test files
│   ├── unit/             # Unit tests
│   └── integration/      # Integration tests
├── public/               # Static files served directly
├── index.html            # HTML entry point
├── package.json          # NPM dependencies and scripts
├── tsconfig.json         # TypeScript configuration
├── vite.config.ts        # Vite build configuration
├── .eslintrc.js          # ESLint configuration
└── .prettierrc           # Prettier configuration
```

## Getting Started

### Prerequisites

- Node.js (v16+)
- npm or yarn

### Installation

```bash
# Install dependencies
npm install
```

### Development

```bash
# Start development server
npm run dev

# Run tests
npm test

# Lint code
npm run lint

# Build for production
npm run build
```

## Docker Development

The project includes Docker configuration for local development:

```bash
# From the root of the monorepo
docker-compose up frontend
```

## Available Scripts

- `npm run dev` - Start the development server
- `npm run build` - Build for production
- `npm run test` - Run tests
- `npm run lint` - Lint code
- `npm run preview` - Preview production build locally

## Technology Stack

- **Framework**: React 18
- **Language**: TypeScript
- **Build Tool**: Vite
- **Styling**: CSS
- **Drag and Drop**: react-dnd
- **HTTP Client**: Axios
