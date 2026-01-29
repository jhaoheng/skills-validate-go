// Package skillsref provides a public API for validating and working with Agent Skills.
package skillsref

import (
	"encoding/json"

	"github.com/jhaoheng/skills-validate-go/internal/models"
	"github.com/jhaoheng/skills-validate-go/internal/parser"
	"github.com/jhaoheng/skills-validate-go/internal/prompt"
	"github.com/jhaoheng/skills-validate-go/internal/validator"
)

func Validate(skillDir string) []string {
	return validator.Validate(skillDir)
}

func ReadProperties(skillDir string) (*models.SkillProperties, error) {
	return parser.ReadProperties(skillDir)
}

func ReadPropertiesJSON(skillDir string) (string, error) {
	props, err := parser.ReadProperties(skillDir)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.MarshalIndent(props.ToMap(), "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func ToPrompt(skillDirs []string) (string, error) {
	return prompt.ToPrompt(skillDirs)
}
