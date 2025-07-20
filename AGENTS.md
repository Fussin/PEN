# AGENTS.md

This file provides guidance for Jules AI, the autonomous agent developing the AUTONOMOUSPEN AI platform.

## Polyglot Architecture

This project uses a polyglot architecture, combining multiple programming languages to leverage their respective strengths:

-   **Go:** For high-performance, concurrent microservices.
-   **Python:** For AI and machine learning components.
-   **Solidity:** For blockchain-based smart contracts.

## Development Environment Setup

-   **Go:** Go 1.19+
-   **Python:** Python 3.10+
-   **Node.js:** Node.js 16+ (for Solidity development)
-   **Dependency Management:**
    -   Go: Go Modules (`go.mod`)
    -   Python: `pip` and `requirements.txt`
    -   Solidity: `npm` or `yarn` (`package.json`)
-   **Code Formatters:**
    -   Go: `gofmt`
    -   Python: Black
-   **Linters:**
    -   Go: `golangci-lint`
    -   Python: Flake8
    -   Solidity: `solhint`
-   **Testing Frameworks:**
    -   Go: Go's built-in testing package
    -   Python: Pytest
    -   Solidity: Hardhat or Truffle

## Common Commands

-   **Go:**
    -   `go build`
    -   `go test ./...`
-   **Python:**
    -   `pip install -r requirements.txt`
    -   `pytest`
-   **Solidity:**
    -   `npm install`
    -   `npx hardhat compile`
    -   `npx hardhat test`

## Project Structure

-   `go/services/`: Go microservices.
-   `python/ai_components/`: Python AI components.
--   `solidity/`: Solidity smart contracts, scripts, and tests.
-   `k8s/`: Kubernetes manifests.
-   `.github/workflows/`: GitHub Actions CI/CD pipelines.
-   `docs/`: Project documentation.
-   `scripts/`: Miscellaneous scripts.
-   `AGENTS.md`: This file.
-   `README.md`: Project overview and roadmap.
