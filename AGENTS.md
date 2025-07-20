# AGENTS.md

This file provides guidance for Jules AI, the autonomous agent developing the AUTONOMOUSPEN AI platform.

## Development Environment Setup

-   **Language:** Python 3.10+
-   **Dependency Management:** `pip` and `requirements.txt` files in each microservice directory.
-   **Code Formatter:** Black
-   **Linter:** Flake8
-   **Testing Framework:** Pytest

## Common Commands

-   **Install Dependencies:** `pip install -r requirements.txt` (run in each microservice directory)
-   **Run Tests:** `pytest`
-   **Format Code:** `black .`
-   **Lint Code:** `flake8 .`

## Code Style Guidelines

-   Follow PEP 8 standards.
-   All code must be formatted with Black.
-   All code must pass Flake8 linting with zero warnings.
-
## Project Structure

-   `src/`: Contains the source code for each microservice.
    -   `src/<service_name>/`: Each microservice has its own directory.
        -   `src/<service_name>/requirements.txt`: Dependencies for the service.
-   `tests/`: Contains all tests.
-   `docs/`: Contains project documentation.
-   `scripts/`: Contains automation scripts.
-   `AGENTS.md`: This file.
-   `.pre-commit-config.yaml`: Configuration for pre-commit hooks.
-   `pytest.ini`: Configuration for Pytest.
-   `README.md`: Project overview and roadmap.
