#!/bin/bash

# Script to set up pre-commit hooks for the quantum simulation component

echo "Setting up pre-commit hooks for quantum simulation component..."

# Check if virtual environment is active
if [[ -z "${VIRTUAL_ENV}" ]]; then
  echo "No virtual environment detected."
  echo "Creating a new virtual environment..."
  python -m venv venv

  # Activate virtual environment
  if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    source venv/Scripts/activate
  else
    source venv/bin/activate
  fi
  echo "Virtual environment activated."
else
  echo "Using existing virtual environment: $VIRTUAL_ENV"
fi

# Install dependencies
echo "Installing dependencies..."
pip install -r requirements.txt
pip install -e .

# Install pre-commit
echo "Installing pre-commit..."
pip install pre-commit

# Install the pre-commit hooks
echo "Installing pre-commit hooks..."
pre-commit install

# Run pre-commit hooks on all files to verify installation
echo "Running pre-commit on all files to verify installation..."
pre-commit run --all-files

echo "Pre-commit setup complete!"
echo "You can now use 'git commit' and pre-commit hooks will run automatically."
echo "To manually run pre-commit hooks: pre-commit run --all-files"
