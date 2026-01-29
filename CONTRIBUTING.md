# Contributing to skills-validate-go

Thank you for your interest in contributing!

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Getting Started

1. Fork and clone the repository:
```bash
git clone https://github.com/jhaoheng/skills-validate-go.git
cd skills-validate-go
```

2. Install dependencies:
```bash
go mod download
```

3. Run tests:
```bash
make test
```

4. Build the project:
```bash
make build
```

## Development Workflow

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific package tests
go test ./internal/parser/...
```

### Code Quality

Before submitting a pull request:

```bash
# Format code
make fmt

# Run linters
make lint

# Run tests
make test
```

### Making Changes

1. Create a new branch for your feature/fix
2. Make your changes
3. Add tests for new functionality
4. Ensure all tests pass
5. Run formatters and linters
6. Commit with clear messages
7. Push and create a pull request

## Code Style

- Follow standard Go conventions
- Use `gofmt` and `goimports`
- Add comments for exported functions
- Write clear commit messages

## Questions?

Feel free to open an issue for any questions or concerns.
