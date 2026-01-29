// Package parser provides YAML frontmatter parsing for SKILL.md files.
package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	skillerrors "github.com/jhaoheng/skills-validate-go/internal/errors"
	"github.com/jhaoheng/skills-validate-go/internal/models"
)

// FindSkillMD locates the SKILL.md file in the given directory.
func FindSkillMD(skillDir string) (string, error) {
	for _, name := range []string{"SKILL.md", "skill.md"} {
		path := filepath.Join(skillDir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", skillerrors.NewParseError(fmt.Sprintf("SKILL.md not found in %s", skillDir))
}

// ParseFrontmatter extracts and parses YAML frontmatter from SKILL.md content.
func ParseFrontmatter(content string) (map[string]interface{}, string, error) {
	if !strings.HasPrefix(content, "---") {
		return nil, "", skillerrors.NewParseError("SKILL.md must start with YAML frontmatter (---)")
	}

	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, "", skillerrors.NewParseError("SKILL.md frontmatter not properly closed with ---")
	}

	frontmatterStr := parts[1]
	body := strings.TrimSpace(parts[2])

	var metadata map[string]interface{}
	if err := yaml.Unmarshal([]byte(frontmatterStr), &metadata); err != nil {
		return nil, "", skillerrors.NewParseError(fmt.Sprintf("Invalid YAML in frontmatter: %v", err))
	}

	if metadata == nil {
		return nil, "", skillerrors.NewParseError("SKILL.md frontmatter must be a YAML mapping")
	}

	return metadata, body, nil
}

// ReadProperties reads and parses the skill properties from SKILL.md.
func ReadProperties(skillDir string) (*models.SkillProperties, error) {
	skillMDPath, err := FindSkillMD(skillDir)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(skillMDPath)
	if err != nil {
		return nil, skillerrors.NewParseError(fmt.Sprintf("Failed to read SKILL.md: %v", err))
	}

	metadata, _, err := ParseFrontmatter(string(content))
	if err != nil {
		return nil, err
	}

	name, hasName := metadata["name"]
	if !hasName {
		return nil, skillerrors.NewValidationError("Missing required field in frontmatter: name")
	}

	description, hasDesc := metadata["description"]
	if !hasDesc {
		return nil, skillerrors.NewValidationError("Missing required field in frontmatter: description")
	}

	nameStr, ok := name.(string)
	if !ok || strings.TrimSpace(nameStr) == "" {
		return nil, skillerrors.NewValidationError("Field 'name' must be a non-empty string")
	}

	descStr, ok := description.(string)
	if !ok || strings.TrimSpace(descStr) == "" {
		return nil, skillerrors.NewValidationError("Field 'description' must be a non-empty string")
	}

	props := &models.SkillProperties{
		Name:        strings.TrimSpace(nameStr),
		Description: strings.TrimSpace(descStr),
	}

	if license, ok := metadata["license"].(string); ok {
		props.License = license
	}
	if compat, ok := metadata["compatibility"].(string); ok {
		props.Compatibility = compat
	}
	if allowedTools, ok := metadata["allowed-tools"].(string); ok {
		props.AllowedTools = allowedTools
	}
	if metaMap, ok := metadata["metadata"].(map[string]interface{}); ok {
		props.Metadata = make(map[string]string)
		for k, v := range metaMap {
			if strVal, ok := v.(string); ok {
				props.Metadata[k] = strVal
			}
		}
	}

	return props, nil
}
