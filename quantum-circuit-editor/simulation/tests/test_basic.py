"""
Basic tests for simulation component.
"""
import pytest
from quantum_simulator.main import Simulation
from quantum_simulator.config import SimulationConfig

def test_simulation_init():
    """Test that simulation can be initialized."""
    sim = Simulation()
    assert isinstance(sim, Simulation)
    assert sim.config is not None

def test_simulation_run():
    """Test that simulation can be run."""
    sim = Simulation()
    result = sim.run()
    assert result["status"] == "success"

def test_custom_config():
    """Test simulation with custom config."""
    config = SimulationConfig(name="test-sim", iterations=10, debug_mode=True)
    sim = Simulation(config)
    assert sim.config.name == "test-sim"
    assert sim.config.iterations == 10
    assert sim.config.debug_mode is True