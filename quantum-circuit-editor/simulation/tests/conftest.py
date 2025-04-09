"""Configure pytest for the quantum_simulator package tests.

This file helps resolve import issues by ensuring the proper path is in sys.path.
"""
import sys
from pathlib import Path

# Add src to path for import resolution
src_path = Path(__file__).parent.parent / "src"
sys.path.insert(0, str(src_path.absolute()))
