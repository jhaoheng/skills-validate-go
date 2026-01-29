// Package validator provides skill validation logic.
package validator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"

	"github.com/jhaoheng/skills-validate-go/internal/parser"
)

const (
	maxSkillNameLength     = 64
	maxDescriptionLength   = 1024
	maxCompatibilityLength = 500
)

var allowedFields = map[string]bool{
	"name":          true,
	"description":   true,
	"license":       true,
	"allowed-tools": true,
	"metadata":      true,
	"compatibility": true,
}

func validateName(name string, skillDir string) []string {
	var errors []string

	if name == "" || strings.TrimSpace(name) == "" {
		errors = append(errors, "Field 'name' must be a non-empty string")
		return errors
	}

	name = norm.NFKC.String(strings.TrimSpace(name))

	if len(name) > maxSkillNameLength {
		errors = append(errors, fmt.Sprintf("Skill name '%s' exceeds %d character limit (%d chars)",
			name, maxSkillNameLength, len(name)))
	}

	if name != strings.ToLower(name) {
		errors = append(errors, fmt.Sprintf("Skill name '%s' must be lowercase", name))
	}

	if strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") {
		errors = append(errors, "Skill name cannot start or end with a hyphen")
	}

	if strings.Contains(name, "--") {
		errors = append(errors, "Skill name cannot contain consecutive hyphens")
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' {
			errors = append(errors, fmt.Sprintf("Skill name '%s' contains invalid characters. Only letters, digits, and hyphens are allowed.", name))
			break
		}
	}

	if skillDir != "" {
		dirName := filepath.Base(skillDir)
		dirName = norm.NFKC.String(dirName)
		if dirName != name {
			errors = append(errors, fmt.Sprintf("Directory name '%s' must match skill name '%s'", dirName, name))
		}
	}

	return errors
}

func validateDescription(description string) []string {
	var errors []string

	if description == "" || strings.TrimSpace(description) == "" {
		errors = append(errors, "Field 'description' must be a non-empty string")
		return errors
	}

	if len(description) > maxDescriptionLength {
		errors = append(errors, fmt.Sprintf("Description exceeds %d character limit (%d chars)",
			maxDescriptionLength, len(description)))
	}

	return errors
}

func validateCompatibility(compatibility string) []string {
	var errors []string

	if len(compatibility) > maxCompatibilityLength {
		errors = append(errors, fmt.Sprintf("Compatibility exceeds %d character limit (%d chars)",
			maxCompatibilityLength, len(compatibility)))
	}

	return errors
}

func validateMetadataFields(metadata map[string]interface{}) []string {
	var errors []string
	var extraFields []string

	for field := range metadata {
		if !allowedFields[field] {
			extraFields = append(extraFields, field)
		}
	}

	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("Unexpected fields in frontmatter: %s. Only %v are allowed.",
			strings.Join(extraFields, ", "), getAllowedFieldsList()))
	}

	return errors
}

func getAllowedFieldsList() []string {
	fields := make([]string, 0, len(allowedFields))
	for field := range allowedFields {
		fields = append(fields, field)
	}
	return fields
}

// ValidateMetadata validates the YAML frontmatter metadata of a skill.
func ValidateMetadata(metadata map[string]interface{}, skillDir string) []string {
	var errors []string

	errors = append(errors, validateMetadataFields(metadata)...)

	if name, ok := metadata["name"]; !ok {
		errors = append(errors, "Missing required field in frontmatter: name")
	} else if nameStr, ok := name.(string); ok {
		errors = append(errors, validateName(nameStr, skillDir)...)
	} else {
		errors = append(errors, "Field 'name' must be a string")
	}

	if description, ok := metadata["description"]; !ok {
		errors = append(errors, "Missing required field in frontmatter: description")
	} else if descStr, ok := description.(string); ok {
		errors = append(errors, validateDescription(descStr)...)
	} else {
		errors = append(errors, "Field 'description' must be a string")
	}

	if compatibility, ok := metadata["compatibility"]; ok {
		if compatStr, ok := compatibility.(string); ok {
			errors = append(errors, validateCompatibility(compatStr)...)
		} else {
			errors = append(errors, "Field 'compatibility' must be a string")
		}
	}

	return errors
}

// Validate performs complete validation of a skill directory.
func Validate(skillDir string) []string {
	info, err := os.Stat(skillDir)
	if os.IsNotExist(err) {
		return []string{fmt.Sprintf("Path does not exist: %s", skillDir)}
	}
	if err != nil {
		return []string{fmt.Sprintf("Error accessing path: %v", err)}
	}

	if !info.IsDir() {
		return []string{fmt.Sprintf("Not a directory: %s", skillDir)}
	}

	skillMDPath, err := parser.FindSkillMD(skillDir)
	if err != nil {
		return []string{"Missing required file: SKILL.md"}
	}

	content, err := os.ReadFile(skillMDPath)
	if err != nil {
		return []string{fmt.Sprintf("Failed to read SKILL.md: %v", err)}
	}

	metadata, _, err := parser.ParseFrontmatter(string(content))
	if err != nil {
		return []string{err.Error()}
	}

	return ValidateMetadata(metadata, skillDir)
}
