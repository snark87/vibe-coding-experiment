from setuptools import setup, find_packages

setup(
    name="quantum_simulator",
    version="0.1.0",
    packages=find_packages(where="src"),
    package_dir={"": "src"},
    install_requires=[
        "pydantic>=2.4.2",
    ],
    python_requires=">=3.9",
)