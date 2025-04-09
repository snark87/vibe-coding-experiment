# Python Coding Guidelines

## Introduction

This document outlines our coding standards and best practices for Python development in the Quantum Circuit Editor project. Following these guidelines will help ensure our codebase remains maintainable, performant, and consistent.

## Code Style and Formatting

### Formatting

- Use [Black](https://github.com/psf/black) for code formatting with a line length of 88 characters.
- Use [isort](https://pycqa.github.io/isort/) for organizing imports.
- Configure your IDE to format code on save.
- Use 4 spaces for indentation (no tabs).

### Naming Conventions

- Follow PEP 8 naming conventions:
  - `snake_case` for functions, variables, and module names
  - `PascalCase` for classes
  - `UPPER_SNAKE_CASE` for constants
  - `_leading_underscore` for private attributes/methods
  - `__double_leading_underscore` for name mangling
- Use meaningful, descriptive names for variables, functions, and classes.
- Avoid single-letter variable names except in very short blocks or as counters in loops.

```python
# Good
class QuantumCircuit:
    def __init__(self, num_qubits):
        self.num_qubits = num_qubits
        self._state = self._initialize_state()

    def _initialize_state(self):
        # Private helper method
        return [0] * self.num_qubits

# Bad
class qc:
    def __init__(self, n):
        self.n = n
        self.s = self.init_s()

    def init_s(self):
        return [0] * self.n
```

### Comments and Documentation

- Write docstrings for all public modules, functions, classes, and methods.
- Use [Google-style docstrings](https://google.github.io/styleguide/pyguide.html#38-comments-and-docstrings) format.
- Document parameters, return values, and exceptions.
- Include examples in docstrings for complex functions.
- Add type hints using Python's typing module.

```python
def simulate_circuit(circuit: QuantumCircuit, shots: int = 1000) -> dict:
    """Simulate a quantum circuit and return measurement results.

    Args:
        circuit: The quantum circuit to simulate.
        shots: Number of simulation shots to run.

    Returns:
        A dictionary mapping measurement outcomes to their counts.

    Raises:
        SimulationError: If the simulation fails.

    Example:
        >>> circuit = QuantumCircuit(2)
        >>> circuit.h(0)
        >>> circuit.cx(0, 1)
        >>> simulate_circuit(circuit)
        {'00': 500, '11': 500}
    """
```

### Code Organization

- Organize imports in the following order:
  1. Standard library imports
  2. Related third-party imports
  3. Local application/library imports
  4. Separate each group with a blank line
- Keep functions and methods focused and small.
- Follow the Single Responsibility Principle.
- Use classes to encapsulate related functionality.

## Development Practices

### Error Handling

- Use exceptions for error handling, not return codes.
- Create custom exception classes for specific error types.
- Handle exceptions at the appropriate level.
- Be specific with exception types (avoid catching Exception).
- Use context managers (with statement) for resource management.

```python
# Good
try:
    result = perform_calculation(value)
except ValueError as e:
    logger.error("Invalid input value: %s", e)
    raise InvalidInputError(f"Could not perform calculation: {e}") from e

# Bad
if value < 0:
    return None  # No explanation of what went wrong
result = perform_calculation(value)
return result
```

### Logging

- Use Python's logging module instead of print statements.
- Configure logging early in the application lifecycle.
- Use appropriate log levels:
  - DEBUG: Detailed information for debugging
  - INFO: Confirmation that things are working as expected
  - WARNING: Something unexpected but not an error
  - ERROR: An error occurred but execution continues
  - CRITICAL: A serious error that may prevent program execution
- Include context in log messages.
- Avoid logging sensitive information.

```python
# Good
import logging
logger = logging.getLogger(__name__)

def process_user_data(user_id, data):
    logger.info("Processing data for user %s", user_id)
    try:
        result = perform_processing(data)
        logger.debug("Processing result: %s", result)
        return result
    except ProcessingError as e:
        logger.error("Failed to process data for user %s: %s", user_id, e)
        raise
```

### Concurrency

- Use `asyncio` for I/O-bound tasks.
- Use `concurrent.futures` for simple parallelism.
- Use `multiprocessing` for CPU-bound tasks.
- Be aware of the Global Interpreter Lock (GIL) limitations.
- Consider thread safety when accessing shared data.
- Use proper synchronization primitives when needed.
- Design for non-blocking operations where appropriate.

### Testing

- Write tests for all new code.
- Use pytest as the testing framework.
- Organize tests to mirror the structure of the module being tested.
- Write both unit and integration tests.
- Use fixtures for test setup and teardown.
- Use parameterized tests for testing similar functionality with different inputs.
- Mock external dependencies in unit tests.
- Name test functions with `test_` prefix followed by the function name and scenario.

```python
# test_circuit.py
import pytest
from quantum.circuit import QuantumCircuit

def test_create_circuit_with_valid_qubits():
    circuit = QuantumCircuit(2)
    assert circuit.num_qubits == 2
    assert len(circuit.state) == 2

def test_create_circuit_with_invalid_qubits():
    with pytest.raises(ValueError):
        QuantumCircuit(-1)

@pytest.mark.parametrize("gate,expected", [
    ("X", [[0, 1], [1, 0]]),
    ("H", [[1/np.sqrt(2), 1/np.sqrt(2)], [1/np.sqrt(2), -1/np.sqrt(2)]]),
])
def test_gate_matrix(gate, expected):
    assert np.allclose(get_gate_matrix(gate), expected)
```

### Dependencies

- Minimize external dependencies.
- Use `requirements.txt` and optionally `setup.py` to declare dependencies.
- Pin dependency versions for reproducible environments.
- Consider using virtual environments (venv) or Docker for isolation.
- Regularly update dependencies and check for security vulnerabilities.
- Document why dependencies are needed.

## Tooling and Linting

### Required Tools

- **Black**: Code formatting
- **isort**: Import sorting
- **mypy**: Static type checking
- **flake8**: Linting
- **pytest**: Testing
- **bandit**: Security checks
- **pylint**: Advanced static analysis

### Configuration

Our project uses these configurations:

```toml
# pyproject.toml
[tool.black]
line-length = 88
target-version = ['py39']
include = '\.pyi?$'

[tool.isort]
profile = "black"
line_length = 88

[tool.mypy]
python_version = "3.9"
disallow_untyped_defs = true
disallow_incomplete_defs = true
check_untyped_defs = true
disallow_untyped_decorators = true
no_implicit_optional = true
warn_redundant_casts = true
warn_return_any = true
warn_unused_ignores = true
```

```ini
# setup.cfg
[flake8]
max-line-length = 88
extend-ignore = E203, W503
exclude = .git,__pycache__,build,dist
```

### IDE Integration

Configure your IDE to:
- Run Black on save
- Run isort on save
- Enable type checking with mypy
- Show linting errors in real time

Recommended VS Code extensions:
- Python (official Microsoft extension)
- Pylance
- Python Docstring Generator
- Python Test Explorer

### CI Integration

Our CI pipeline will:
- Run all tests
- Enforce linting rules
- Check formatting with Black
- Run static type checking with mypy
- Generate and check code coverage
- Perform security checks with bandit

## Project-Specific Conventions

### Directory Structure

```
simulation/
├── pyproject.toml       # Project configuration
├── setup.py             # Package setup
├── requirements.txt     # Dependencies
├── README.md            # Documentation
├── src/                 # Source code
│   └── quantum_simulator/
│       ├── __init__.py
│       ├── circuit.py   # Circuit definition
│       ├── gates.py     # Quantum gate definitions
│       ├── simulator.py # Simulation engine
│       └── utils.py     # Utility functions
└── tests/               # Test code
    ├── __init__.py
    ├── conftest.py      # Test configuration
    ├── test_circuit.py  # Circuit tests
    ├── test_gates.py    # Gate tests
    └── test_simulator.py # Simulator tests
```

### Module Organization

- Keep modules focused on a single responsibility.
- Use `__init__.py` to expose the public API of a package.
- Place implementation details in separate modules.
- Use relative imports within a package.

### Package Distribution

- Use `setup.py` or `pyproject.toml` for package metadata.
- Follow semantic versioning.
- Include a README.md with usage examples.
- Provide clear installation instructions.
- Document dependencies and compatibility.

## Best Practices

### Performance

- Profile before optimizing.
- Use appropriate data structures for the task.
- Leverage NumPy and other optimized libraries for numerical operations.
- Consider memory usage, especially for large matrices or vectors.
- Use generators for large data processing.
- Minimize object creation in tight loops.

### Security

- Validate all user input.
- Use proper authentication and authorization.
- Avoid shell=True in subprocess calls.
- Be careful with deserialization of untrusted data.
- Keep dependencies updated.
- Use secure defaults.

### Maintainability

- Write code that is easy to understand and modify.
- Prefer readability over cleverness.
- Follow the DRY (Don't Repeat Yourself) principle.
- Break down complex functions into smaller, focused ones.
- Use appropriate abstractions.
- Write code for humans, not for computers.

### Python-Specific Best Practices

- Use list comprehensions and generator expressions for clarity.
- Leverage Python's standard library.
- Use context managers with the `with` statement.
- Prefer `is` for comparison with `None`, `True`, and `False`.
- Use `isinstance()` for type checking, not `type()`.
- Leverage duck typing and EAFP (Easier to Ask Forgiveness than Permission).
- Use f-strings for string formatting.

```python
# Good - EAFP style
try:
    value = data["key"]
except KeyError:
    value = default_value

# Good - List comprehension
squares = [x**2 for x in range(10)]

# Good - Context manager
with open("file.txt", "r") as f:
    content = f.read()

# Good - Type checking
if isinstance(value, dict):
    process_dict(value)
```

### Accessibility for New Developers

- Document setup steps clearly.
- Provide examples for common tasks.
- Comment complex algorithms or business logic.
- Use type hints to make code more self-documenting.
- Create clear error messages.

## Quantum Computing Specific Guidelines

### Matrix Representation

- Use NumPy for matrix operations.
- Define standard gates as constants.
- Use established conventions for qubit ordering.
- Document the mathematical representation of quantum operations.

### Simulation Performance

- Consider using specialized libraries (Qiskit, Cirq, etc.) for heavy calculations.
- Use sparse matrices for large qubit counts.
- Implement optimizations for common patterns.
- Document performance limitations.

### Numerical Stability

- Be aware of floating-point precision issues.
- Use appropriate tolerances for equality comparisons.
- Consider using complex128 for better precision in quantum states.
- Normalize quantum states when necessary.

## Recommended Resources

- [PEP 8 - Style Guide for Python Code](https://www.python.org/dev/peps/pep-0008/)
- [Google Python Style Guide](https://google.github.io/styleguide/pyguide.html)
- [The Hitchhiker's Guide to Python](https://docs.python-guide.org/)
- [Effective Python](https://effectivepython.com/) by Brett Slatkin
- [Real Python](https://realpython.com/) tutorials
- [Python Design Patterns](https://python-patterns.guide/)
- [Python Testing with pytest](https://pragprog.com/titles/bopytest/python-testing-with-pytest/) by Brian Okken

## Conclusion

These guidelines should be treated as a living document. As we learn and evolve as a team, we'll update and improve our practices. The primary goal is to create maintainable, reliable, and performant code that fulfills our product requirements and provides value to our users.
