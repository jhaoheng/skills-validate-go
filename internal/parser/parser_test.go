package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindSkillMD(t *testing.T) {
	t.Run("finds SKILL.md", func(t *testing.T) {
		dir := t.TempDir()
		os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("test"), 0644)
		_, err := FindSkillMD(dir)
		if err != nil {
			t.Errorf("FindSkillMD() error = %v, want nil", err)
		}
	})

	t.Run("missing SKILL.md", func(t *testing.T) {
		dir := t.TempDir()
		_, err := FindSkillMD(dir)
		if err == nil {
			t.Error("FindSkillMD() error = nil, want error")
		}
	})
}

func TestReadProperties(t *testing.T) {
	t.Run("valid skill", func(t *testing.T) {
		dir := t.TempDir()
		content := "---\nname: my-skill\ndescription: Test description\n---\nBody"
		os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(content), 0644)

		props, err := ReadProperties(dir)
		if err != nil {
			t.Errorf("ReadProperties() error = %v", err)
		}
		if props.Name != "my-skill" {
			t.Errorf("ReadProperties() name = %v, want my-skill", props.Name)
		}
	})
}
