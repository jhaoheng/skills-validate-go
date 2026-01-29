// Package parser provides YAML frontmatter parsing for SKILL.md files.
package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/jhaoheng/skills-validate-go/internal/errors"
	"github.com/jhaoheng/skills-validate-go/internal/models"
)

func FindSkillMD(skillDir string) (string, error) {
	for _, name := range []string{"SKILL.md", "skill.md"} {
		path := filepath.Join(skillDir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", errors.NewParseError(fmt.Sprintf("SKILL.md not found in %s", skillDir))
}

func ParseFrontmatter(content string) (map[string]interface{}, string, error) {
	if !strings.HasPrefix(content, "---") {
		return nil, "", errors.NewParseError("SKILL.md must start with YAML frontmatter (---)")
	}

	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, "", errors.NewParseError("SKILL.md frontmatter not properly closed with ---")
	}

	frontmatterStr := parts[1]
	body := strings.TrimSpace(parts[2])

	var metadata map[string]interface{}
	if err := yaml.Unmarshal([]byte(frontmatterStr), &metadata); err != nil {
		return nil, "", errors.NewParseError(fmt.Sprintf("Invalid YAML in frontmatter: %v", err))
	}

	if metadata == nil {
		return nil, "", errors.NewParseError("SKILL.md frontmatter must be a YAML mapping")
	}

	return metadata, body, nil
}

func ReadProperties(skillDir string) (*models.SkillProperties, error) {
	skillMDPath, err := FindSkillMD(skillDir)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(skillMDPath)
	if err != nil {
		return nil, errors.NewParseError(fmt.Sprintf("Failed to read SKILL.md: %v", err))
	}

	metadata, _, err := ParseFrontmatter(string(content))
	if err != nil {
		return nil, err
	}

	name, hasName := metadata["name"]
	if !hasName {
		return nil, errors.NewValidationError("Missing required field in frontmatter: name")
	}

	description, hasDesc := metadata["description"]
	if !hasDesc {
		return nil, errors.NewValidationError("Missing required field in frontmatter: description")
	}

	nameStr, ok := name.(string)
	if !ok || strings.TrimSpace(nameStr) == "" {
		return nil, errors.NewValidationError("Field 'name' must be a non-empty string")
	}

	descStr, ok := description.(string)
	if !ok || strings.TrimSpace(descStr) == "" {
		return nil, errors.NewValidationError("Field 'description' must be a non-empty string")
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
