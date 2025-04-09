import React from 'react';
import { render, screen } from '@testing-library/react';
import App from '../../src/App';

describe('App Component', () => {
  test('renders Quantum Circuit Editor heading', () => {
    render(<App />);
    // Using a more specific selector with h1 tag
    const headingElement = screen.getByRole('heading', { level: 1, name: /Quantum Circuit Editor/i });
    expect(headingElement).toBeInTheDocument();
  });

  test('renders welcome message', () => {
    render(<App />);
    // Using a more specific selector
    const welcomeElement = screen.getByText(/Welcome to the Quantum Circuit Editor application/i);
    expect(welcomeElement).toBeInTheDocument();
  });
});
