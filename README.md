# skills-validate-go

A Go implementation of Agent Skills validation tool.

> **Note:** This library is intended for demonstration purposes only. It is not meant to be used in production.

## Features

- ✅ Validate skill directories
- ✅ Read and parse skill properties
- ✅ Generate agent prompt XML

## Installation

### Using go install

```bash
go install github.com/jhaoheng/skills-validate-go/cmd/skills-validate@latest
```

### From Source

```bash
git clone https://github.com/jhaoheng/skills-validate-go.git
cd skills-validate-go
make install
```

### Download Binary

Download the latest release from the [releases page](https://github.com/jhaoheng/skills-validate-go/releases).

## Usage

### CLI

```bash
# Validate a skill
skills-validate validate path/to/skill

# Read skill properties (outputs JSON)
skills-validate read-properties path/to/skill

# Generate <available_skills> XML for agent prompts
skills-validate to-prompt path/to/skill-a path/to/skill-b
```

### Go API

```go
package main

import (
    "fmt"
    "github.com/jhaoheng/skills-validate-go/pkg/skillsref"
)

func main() {
    // Validate a skill directory
    problems, err := skillsref.Validate("my-skill")
    if err != nil {
        panic(err)
    }
    if len(problems) > 0 {
        fmt.Println("Validation errors:", problems)
    }

    // Read skill properties
    props, err := skillsref.ReadProperties("my-skill")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Skill: %s - %s\n", props.Name, props.Description)

    // Generate prompt for available skills
    prompt, err := skillsref.ToPrompt([]string{"skill-a", "skill-b"})
    if err != nil {
        panic(err)
    }
    fmt.Println(prompt)
}
```

## Development

### Prerequisites

- Go 1.21 or higher

### Building

```bash
# Build the binary
make build

# Run tests
make test

# Run tests with coverage
make test-coverage

# Run linters
make lint

# Format code
make fmt
```

### Project Structure

```
skills-validate-go/
├── cmd/
│   └── skills-validate/    # CLI entry point
├── internal/
│   ├── parser/            # SKILL.md parser
│   ├── validator/         # Validation logic
│   ├── prompt/            # Prompt generation
│   ├── models/            # Data models
│   └── errors/            # Error definitions
├── pkg/
│   └── skillsref/         # Public API
└── testdata/              # Test fixtures
```

## License

Apache 2.0

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
