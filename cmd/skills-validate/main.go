// Package main provides the CLI for Agent Skills validation tool.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jhaoheng/skills-validate-go/pkg/skillsref"
)

var version = "dev"

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "skills-validate",
	Short: "Validate and manage Agent Skills",
	Long:  "A Go implementation of Agent Skills validation tool.",
}

var validateCmd = &cobra.Command{
	Use:   "validate [skill-path]",
	Short: "Validate a skill directory",
	Long: `Validates that the skill has a valid SKILL.md with proper frontmatter,
correct naming conventions, and required fields.

Exit codes:
  0: Valid skill
  1: Validation errors found`,
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		skillPath := args[0]

		errors := skillsref.Validate(skillPath)

		if len(errors) > 0 {
			fmt.Fprintf(os.Stderr, "Validation failed for %s:\n", skillPath)
			for _, err := range errors {
				fmt.Fprintf(os.Stderr, "  - %s\n", err)
			}
			os.Exit(1)
		}
		fmt.Printf("Valid skill: %s\n", skillPath)
	},
}

var readPropertiesCmd = &cobra.Command{
	Use:   "read-properties [skill-path]",
	Short: "Read and print skill properties as JSON",
	Long: `Parses the YAML frontmatter from SKILL.md and outputs the
properties as JSON.

Exit codes:
  0: Success
  1: Parse error`,
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		skillPath := args[0]

		props, err := skillsref.ReadProperties(skillPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		jsonBytes, err := json.MarshalIndent(props.ToMap(), "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(jsonBytes))
	},
}

var toPromptCmd = &cobra.Command{
	Use:   "to-prompt [skill-path...]",
	Short: "Generate <available_skills> XML for agent prompts",
	Long: `Accepts one or more skill directories and generates the
<available_skills> XML block for inclusion in agent prompts.

Exit codes:
  0: Success
  1: Error`,
	Args: cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		output, err := skillsref.ToPrompt(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(output)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("skills-validate %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(readPropertiesCmd)
	rootCmd.AddCommand(toPromptCmd)
	rootCmd.AddCommand(versionCmd)
}
