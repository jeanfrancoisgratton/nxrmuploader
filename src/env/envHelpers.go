// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/envHelpers.go
// Original timestamp: 2023/12/31 14:50

package env

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

var EnvConfigFile string

type Repository struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RepositoryInfo struct {
	YUM []Repository `json:"YUM"`
	APT []Repository `json:"APT"`
}

// Load the JSON environment file in the user's .config/certificatemanager directory, and store it into a data type (struct)
func LoadEnvironmentFile() (RepositoryInfo, error) {
	var payload RepositoryInfo
	var err error

	if !strings.HasSuffix(EnvConfigFile, ".json") {
		EnvConfigFile += ".json"
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "nxrmuploader", EnvConfigFile)
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return RepositoryInfo{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return RepositoryInfo{}, err
	} else {
		return payload, nil
	}
}

// Save the above structure into a JSON file in the user's .config/certificatemanager directory
func (e RepositoryInfo) SaveEnvironmentFile(outputfile string) error {
	if outputfile == "" {
		outputfile = EnvConfigFile
	}
	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "nxrmuploader", outputfile)
	err = os.WriteFile(rcFile, jStream, 0600)

	return err
}
