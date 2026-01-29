// Package models defines the core data structures for Agent Skills.
package models

// SkillProperties represents properties parsed from a skill's SKILL.md frontmatter.
type SkillProperties struct {
	Name          string            `json:"name" yaml:"name"`
	Description   string            `json:"description" yaml:"description"`
	License       string            `json:"license,omitempty" yaml:"license,omitempty"`
	Compatibility string            `json:"compatibility,omitempty" yaml:"compatibility,omitempty"`
	AllowedTools  string            `json:"allowed-tools,omitempty" yaml:"allowed-tools,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// ToMap converts SkillProperties to a map, excluding empty values.
func (s *SkillProperties) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
	}
	if s.License != "" {
		result["license"] = s.License
	}
	if s.Compatibility != "" {
		result["compatibility"] = s.Compatibility
	}
	if s.AllowedTools != "" {
		result["allowed-tools"] = s.AllowedTools
	}
	if len(s.Metadata) > 0 {
		result["metadata"] = s.Metadata
	}
	return result
}
