# Quantum Simulator Component

This directory contains the Python-based quantum simulation component for the Quantum Circuit Editor project. This component provides the computational engine for simulating quantum circuits created in the editor.

## Directory Structure

```
simulation/
├── pyproject.toml       # Python project configuration
├── requirements.txt     # Project dependencies
├── src/                 # Source code directory
│   └── quantum_simulator/ # Main package
│       ├── __init__.py  # Package initialization
│       ├── config.py    # Configuration settings
│       └── main.py      # Entry point and core logic
└── tests/              # Test directory
    ├── __init__.py     # Test package initialization
    └── test_basic.py   # Basic tests
```

## Setup Instructions

### Local Development

1. Create and activate a virtual environment:

```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

2. Install dependencies:

```bash
pip install -r requirements.txt
pip install -e .  # Install the package in development mode
```

3. Set up pre-commit hooks:

```bash
pip install pre-commit
pre-commit install
```

### Code Quality Tools

This project uses several code quality tools to maintain a high standard of code:

#### Formatting and Linting

- **Black** - Code formatter
  ```bash
  black src/ tests/
  ```

- **isort** - Import sorter
  ```bash
  isort src/ tests/
  ```

- **flake8** - Linter
  ```bash
  flake8 src/ tests/
  ```

- **pylint** - Static code analyzer
  ```bash
  pylint src/ tests/
  ```

#### Type Checking

- **mypy** - Static type checker
  ```bash
  mypy src/ tests/
  ```

#### Security

- **bandit** - Security linter
  ```bash
  bandit -r src/
  ```

#### Automatic Checks

Pre-commit hooks will run these checks automatically before each commit:
```bash
pre-commit run --all-files
```

### Running Tests

Run tests using pytest:

```bash
pytest tests/
```

With coverage:

```bash
pytest --cov=src tests/
```

## Usage

The simulation component provides a simple API for quantum circuit simulation:

```python
from quantum_simulator import Simulation, SimulationConfig

# Create a simulation with default config
sim = Simulation()

# Or with custom config
config = SimulationConfig(name="my-simulation", iterations=50, debug_mode=True)
sim = Simulation(config)

# Run the simulation
result = sim.run()
print(result)
```

## Future Enhancements

In future iterations, this component will be expanded to include:

1. Support for more quantum gates
2. Circuit validation
3. More detailed measurement outputs
4. Performance optimizations for larger circuits