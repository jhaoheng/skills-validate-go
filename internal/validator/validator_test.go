package validator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("valid skill directory", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "test-skill")
		if err := os.Mkdir(dir, 0755); err != nil {
			t.Fatal(err)
		}
		content := "---\nname: test-skill\ndescription: A valid test skill\n---\nBody"
		if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}

		errors := Validate(dir)
		if len(errors) > 0 {
			t.Errorf("Validate() errors = %v, want none", errors)
		}
	})

	t.Run("missing SKILL.md", func(t *testing.T) {
		dir := t.TempDir()
		errors := Validate(dir)
		if len(errors) == 0 {
			t.Error("Validate() errors = none, want error")
		}
	})
}
