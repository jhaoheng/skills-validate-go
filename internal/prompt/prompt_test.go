package prompt

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestToPrompt(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		result, err := ToPrompt([]string{})
		if err != nil {
			t.Errorf("ToPrompt() error = %v", err)
		}
		if !strings.Contains(result, "<available_skills>") {
			t.Error("ToPrompt() result doesn't contain <available_skills>")
		}
	})

	t.Run("single skill", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "test-skill")
		os.Mkdir(dir, 0755)
		content := "---\nname: test-skill\ndescription: A test skill\n---\nBody"
		os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(content), 0644)

		result, err := ToPrompt([]string{dir})
		if err != nil {
			t.Errorf("ToPrompt() error = %v", err)
		}
		if !strings.Contains(result, "test-skill") {
			t.Error("ToPrompt() result doesn't contain skill name")
		}
	})
}
