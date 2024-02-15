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
	"os/exec"
	"path"
	"strings"
)

const Headers = "-H 'accept: application/json' -H 'Content-Type: multipart/form-data'"

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
			passwd = helpers.DecodeString(repoInfo.DEBIAN[IndexNumber].Password)

		} else {
			url = repoInfo.RH[IndexNumber].URL
			user = repoInfo.RH[IndexNumber].Username
			passwd = helpers.DecodeString(repoInfo.RH[IndexNumber].Password)
		}
		if err := uploadFile(pkg, url, user, passwd); err != nil {
			fmt.Printf("%s\n", helpers.Red(err.Error()))
		}
	}
	return nil
}

func uploadFile(pkg, url, user, passwd string) error {
	var fqdn, endpoint string
	var err error

	if fqdn, endpoint, err = parseURL(url); err != nil {
		return err
	}

	// This is not a "by the book use of exec.Command in the sense that I pass *most of the* arguments in a
	// Single string, here, instead of using params to Command().
	// It's my quick and dirty way of doing things, which is not "incorrect", stricly speaking
	usernamePassword := fmt.Sprintf(" -X POST -u %s:%s", user, passwd)
	endpointURL := fmt.Sprintf(" %s/service/rest/v1/components?repository=%s", fqdn, endpoint)
	var formKeyVal string

	bname := path.Base(pkg)
	if strings.HasSuffix(pkg, strings.ToLower(".deb")) {
		formKeyVal = fmt.Sprintf(" -F 'apt.asset=@%s;type=application/vnd.debian.binary-package'", bname)
	} else {
		formKeyVal = fmt.Sprintf(" -F 'yum.asset=@%s' -F 'yum.asset.filename=%s'", bname, bname)
	}
	//cmd := exec.Command("curl", firstPart, thirdPart)
	fmt.Printf("curl%s%s%s\n", usernamePassword, formKeyVal, endpointURL)
	cmd := exec.Command("curl", usernamePassword, formKeyVal, endpointURL)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return helpers.CustomError{Message: "curl command failed: " + err.Error()}
	}
	return nil
}
func uploadFile3(pkg, url, user, passwd string) error {
	var fqdn, endpoint string
	var err error

	if fqdn, endpoint, err = parseURL(url); err != nil {
		return err
	}

	file, err := os.Open(pkg)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if strings.HasSuffix(pkg, strings.ToLower(".deb")) {
		keyval := fmt.Sprintf("@%s", path.Base(pkg))
		writer.WriteField("apt.asset", keyval)
		writer.WriteField("type", "application/vnd.debian.binary-package")
	} else {
		keyval := fmt.Sprintf("@%s", path.Base(pkg))
		writer.WriteField("yum.asset", keyval)
		writer.WriteField("yum.asset.filename", path.Base(pkg))
		writer.WriteField("type", "application/x-rpm")
	}
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
	targetURL := fqdn + "/service/rest/v1/components?repository=" + endpoint
	req, err := http.NewRequest("POST", targetURL, body)
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

	// Note : Nexus sends an http 204-NoContent when successful, not 200-OK
	if resp.StatusCode != http.StatusNoContent {
		return helpers.CustomError{fmt.Sprintf("%s: %s %s\n", helpers.Red("HTTP error"), helpers.Normal("status="), helpers.Normal(resp.Status))}
	}
	return nil
}

/*
cmd := exec.Command("keytool", "-importkeystore", "-srcstorepass", certPasswd,
		"-deststorepass", certPasswd,
		"-destkeystore", filepath.Join(e.ServerCertsDir, "java", c.CertificateName+".jks"),
		"-srckeystore", filepath.Join(e.ServerCertsDir, "java", c.CertificateName+".p12"),
		"-srcstoretype", "PKCS12")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return helpers.CustomError{Message: "Keytool command failed: " + err.Error()}
	}
*/
