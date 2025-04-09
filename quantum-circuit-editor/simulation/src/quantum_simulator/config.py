"""Configuration settings for simulation component."""
from dataclasses import dataclass


@dataclass
class SimulationConfig:
    """Basic configuration for simulation."""

    name: str = "default-simulation"
    iterations: int = 100
    debug_mode: bool = False

    def __str__(self):
        """Return string representation of the configuration."""
        return (
            f"SimulationConfig(name={self.name}, "
            f"iterations={self.iterations})"
        )
