// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/exec/upload.go
// Original timestamp: 2024/01/02 14:48

package exec

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"nxrmuploader/env"
	"nxrmuploader/helpers"
	"os"
	"strings"
)

var IndexNumber int8

func Upload(packages []string) error {
	var url, user, passwd string
	var repoInfo env.RepositoryInfo
	var err error

	if repoInfo, err = env.LoadEnvironmentFile(); err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to load environment file: %s", err.Error())}
	}

	for _, pkg := range packages {
		if strings.HasSuffix(pkg, strings.ToLower(".deb")) {
			url = repoInfo.DEBIAN[IndexNumber].URL
			user = repoInfo.DEBIAN[IndexNumber].Username
			passwd = repoInfo.DEBIAN[IndexNumber].Password
		} else {
			url = repoInfo.RH[IndexNumber].URL
			user = repoInfo.DEBIAN[IndexNumber].Username
			passwd = repoInfo.DEBIAN[IndexNumber].Password
		}
		if err := uploadFile(pkg, url, user, passwd); err != nil {
			fmt.Printf("%s\n", helpers.Red(err.Error()))
		}
	}
	return nil
}

func uploadFile(pkg, url, user, passwd string) error {
	file, err := os.Open(pkg)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", pkg)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(user, passwd)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return helpers.CustomError{fmt.Sprintf("%s: status: %s\n", helpers.Red("HTTP error"), resp.Status)}
	}
	return nil
}
