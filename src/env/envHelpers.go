// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/envHelpers.go
// Original timestamp: 2023/12/31 14:50

package env

import (
	"bufio"
	"encoding/json"
	"fmt"
	"nxrmuploader/helpers"
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
	RH     []Repository `json:"YUM"`
	DEBIAN []Repository `json:"APT"`
	//ALPINE []Repository `json:"APK"`
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

func fetchRepoInfo() []Repository {
	var repository []Repository

	for {
		var repo Repository
		repo.Name = getStringVal("Please enter the friendly repo name (ENTER to quit): ")
		if repo.Name == "" {
			break
		}
		repo.URL = getStringVal("Please enter the repo URL: ")
		if !strings.HasSuffix(repo.URL, "/") {
			repo.URL += "/"
		}
		repo.Username = getStringVal("Please enter the username needed to login: ")
		repo.Password = helpers.EncodeString(helpers.GetPassword("Please enter that user's password: "))

		repository = append(repository, repo)
	}

	return repository
}

func getStringVal(prompt string) string {
	fmt.Print(prompt)
	inputVal := bufio.NewReader(os.Stdin)
	input, _ := inputVal.ReadString('\n')

	return strings.TrimSpace(input)
}
