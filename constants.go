package spin

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const configDir = ".config/spin"

// ConfigDir is the top-level configuration directory
func ConfigDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(home, configDir)
}
