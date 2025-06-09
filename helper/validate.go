package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

// ValidateInputFolder validates the folder structure for required files
func ValidateInputFolder(path string) error {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("folder not found: %s", path)
	}

	hasBlend := false
	err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if filepath.Ext(p) == ".blend" {
			hasBlend = true
		}
		return nil
	})
	if err != nil {
		return err
	}
	if !hasBlend {
		return fmt.Errorf("no .blend file found in %s", path)
	}

	texPath := filepath.Join(path, "textures")
	info, err = os.Stat(texPath)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("textures/ folder missing in %s", path)
	}

	return nil
}
