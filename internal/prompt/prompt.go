// Package prompt generates <available_skills> XML for agent system prompts.
package prompt

import (
	"fmt"
	"html"
	"strings"

	"github.com/jhaoheng/skills-validate-go/internal/parser"
)

// ToPrompt generates an XML representation of available skills for agent prompts.
func ToPrompt(skillDirs []string) (string, error) {
	if len(skillDirs) == 0 {
		return "<available_skills>\n</available_skills>", nil
	}

	var lines []string
	lines = append(lines, "<available_skills>")

	for _, skillDir := range skillDirs {
		props, err := parser.ReadProperties(skillDir)
		if err != nil {
			return "", fmt.Errorf("failed to read properties for %s: %w", skillDir, err)
		}

		lines = append(lines, "<skill>")
		lines = append(lines, "<name>")
		lines = append(lines, html.EscapeString(props.Name))
		lines = append(lines, "</name>")
		lines = append(lines, "<description>")
		lines = append(lines, html.EscapeString(props.Description))
		lines = append(lines, "</description>")

		skillMDPath, err := parser.FindSkillMD(skillDir)
		if err != nil {
			return "", fmt.Errorf("failed to find SKILL.md for %s: %w", skillDir, err)
		}

		lines = append(lines, "<location>")
		lines = append(lines, skillMDPath)
		lines = append(lines, "</location>")

		lines = append(lines, "</skill>")
	}

	lines = append(lines, "</available_skills>")

	return strings.Join(lines, "\n"), nil
}
