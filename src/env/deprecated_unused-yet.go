// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/deprecated_unused-yet.go
// Original timestamp: 2024/01/05 13:15

package env

// I will put in this file commented functions, structs, vars that are of no use NOW, but might be some day.
// Historical stuff is stored here as well

//func createSampleFiles(configdir string) error {
//	//var yum, apt, apk []Repository
//	var yum, apt []Repository
//	//var repos RepositoryInfo
//
//	yum = append(yum, Repository{Name: "Yum repo 1", URL: "https://nexus/repo/yum1", Username: "yum_repo_user1", Password: "yum_repo_passwd1"})
//	yum = append(yum, Repository{Name: "Yum repo 2", URL: "https://nexus/repo/yum2", Username: "yum_repo_user2", Password: "yum_repo_passwd2"})
//	apt = append(apt, Repository{Name: "Apt repo 1", URL: "https://nexus/repo/apt1", Username: "apt_repo_user1", Password: "apt_repo_passwd1"})
//	apt = append(apt, Repository{Name: "Apt repo 2", URL: "https://nexus/repo/apt2", Username: "apt_repo_user2", Password: "apt_repo_passwd2"})
//	apt = append(apt, Repository{Name: "Apt repo 3", URL: "https://nexus/repo/apt3", Username: "apt_repo_user3", Password: "apt_repo_passwd3"})
//	//apk = append(apk, Repository{Name: "Apk repo 1", URL: "https://nexus/repo/apk1", Username: "apl_repo_user1", Password: "apk_repo_passwd1"})
//	//repos := RepositoryInfo{RH: yum, DEBIAN: apt, ALPINE: apk}
//	repos := RepositoryInfo{RH: yum, DEBIAN: apt}
//
//	file, err := os.Create(filepath.Join(configdir, "sample.json"))
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	jsonData, err := json.MarshalIndent(repos, "", "  ")
//	if err != nil {
//		return err
//	}
//
//	if _, err := file.Write(jsonData); err != nil {
//		return err
//	}
//	return nil
//}
