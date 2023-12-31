// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/addRemoveEnv.go
// Original timestamp: 2023/12/31 14:49

package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RemoveEnvFile(envfile string) error {
	if !strings.HasSuffix(envfile, ".json") {
		envfile += ".json"
	}
	if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "certificatemanager", envfile)); err != nil {
		return err
	}

	fmt.Printf("%s removed succesfully\n", envfile)
	return nil
}

func AddEnvFile(envfile string) error {
	var env RepositoryInfo
	var err error

	if !strings.HasSuffix(envfile, ".json") {
		envfile += ".json"
	}

	if env, err = prompt4EnvironmentValues(); err != nil {
		return err
	} else {
		err = env.SaveEnvironmentFile(envfile)
	}
	return err
}

func prompt4EnvironmentValues() (RepositoryInfo, error) {
	//var env RepositoryInfo
	return RepositoryInfo{}, nil
}
