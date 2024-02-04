// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/addRemoveEnv.go
// Original timestamp: 2023/12/31 14:49

package env

import (
	"fmt"
	"nxrmuploader/helpers"
	"os"
	"path/filepath"
	"strings"
)

func RemoveEnvFile(envfiles []string) error {

	for _, envfile := range envfiles {
		if !strings.HasSuffix(envfile, ".json") {
			envfile += ".json"
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmuploader", envfile)); err != nil {
			return err
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmuploader", envfile)); err != nil {
			return err
		}
		fmt.Printf("%s removed succesfully\n", envfile)
	}
	return nil
}

func AddEnvFile(envfile string) error {
	var env RepositoryInfo
	var err error

	if envfile == "" {
		envfile = EnvConfigFile
	}
	if !strings.HasSuffix(envfile, ".json") {
		envfile += ".json"
	}

	env = prompt4EnvironmentValues()

	if err = env.SaveEnvironmentFile(envfile); err != nil {
		return err
	}
	return nil
}

func prompt4EnvironmentValues() RepositoryInfo {
	var env RepositoryInfo
	var YUM, APT []Repository

	// Fetch YUM repo(s) info
	fmt.Printf("%s\n", helpers.White("YUM repositories"))
	YUM = fetchRepoInfo()

	// Fetch APT repo(s) info
	fmt.Printf("%s\n", helpers.White("APT repositories"))
	APT = fetchRepoInfo()

	// We now add those repos into the super-struct
	env.RH = YUM
	env.DEBIAN = APT
	return env
}
