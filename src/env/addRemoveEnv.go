// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/addRemoveEnv.go
// Original timestamp: 2023/12/31 14:49

package env

import (
	"encoding/json"
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
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "nxrmuploader", envfile)); err != nil {
			return err
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "certificatemanager", envfile)); err != nil {
			return err
		}
		fmt.Printf("%s removed succesfully\n", envfile)
	}
	return nil
}

func AddEnvFile(envfile string) error {
	//var env RepositoryInfo
	var err error

	cfgDir := filepath.Join(os.Getenv("HOME"), ".config", "nxrmuploader")
	fmt.Printf("%s: currently we're only creating a sample file in %s\n",
		helpers.Yellow("PLEASE NOTE"), cfgDir)

	return createSampleFiles(cfgDir)

	//if !strings.HasSuffix(envfile, ".json") {
	//	envfile += ".json"
	//}
	//
	//if env, err = prompt4EnvironmentValues(); err != nil {
	//	return err
	//} else {
	//	err = env.SaveEnvironmentFile(envfile)
	//}
	return err
}

func prompt4EnvironmentValues() (RepositoryInfo, error) {
	//var env RepositoryInfo
	return RepositoryInfo{}, nil
}

func createSampleFiles(configdir string) error {
	var yum, apt []Repository
	//var repos RepositoryInfo

	yum = append(yum, Repository{Name: "Yum repo 1", URL: "https://nexus/repo/yum1", Username: "yum_repo_user1", Password: "yum_repo_passwd1"})
	yum = append(yum, Repository{Name: "Yum repo 2", URL: "https://nexus/repo/yum2", Username: "yum_repo_user2", Password: "yum_repo_passwd2"})

	apt = append(apt, Repository{Name: "Apt repo 1", URL: "https://nexus/repo/apt1", Username: "apt_repo_user1", Password: "apt_repo_passwd1"})
	apt = append(apt, Repository{Name: "Apt repo 2", URL: "https://nexus/repo/apt2", Username: "apt_repo_user2", Password: "apt_repo_passwd2"})
	apt = append(apt, Repository{Name: "Apt repo 3", URL: "https://nexus/repo/apt3", Username: "apt_repo_user3", Password: "apt_repo_passwd3"})

	repos := RepositoryInfo{YUM: yum, APT: apt}

	file, err := os.Create(filepath.Join(configdir, "sample.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(repos, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(jsonData); err != nil {
		return err
	}
	return nil
}
