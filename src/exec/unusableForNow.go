// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/exec/unusableForNow.go
// Original timestamp: 2024/02/05 12:35

package exec

//func uploadFileOLD(pkg, url, user, passwd string) error {
//	var fqdn, endpoint string
//	var err error
//
//	if fqdn, endpoint, err = parseURL(url); err != nil {
//		return err
//	}
//
//	file, err := os.Open(pkg)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	if strings.HasSuffix(pkg, strings.ToLower(".deb")) {
//		keyval := fmt.Sprintf("@%s", path.Base(pkg))
//		writer.WriteField("apt.asset", keyval)
//		writer.WriteField("type", "application/vnd.debian.binary-package")
//	} else {
//		keyval := fmt.Sprintf("@%s", path.Base(pkg))
//		writer.WriteField("yum.asset", keyval)
//		writer.WriteField("yum.asset.filename", path.Base(pkg))
//		writer.WriteField("type", "application/x-rpm")
//	}
//	part, err := writer.CreateFormFile("file", pkg)
//	if err != nil {
//		return err
//	}
//	_, err = io.Copy(part, file)
//	if err != nil {
//		return err
//	}
//
//	err = writer.Close()
//	if err != nil {
//		return err
//	}
//	targetURL := fqdn + "/service/rest/v1/components?repository=" + endpoint
//	req, err := http.NewRequest("POST", targetURL, body)
//	if err != nil {
//		return err
//	}
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//	req.SetBasicAuth(user, passwd)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// Note : Nexus sends an http 204-NoContent when successful, not 200-OK
//	if resp.StatusCode != http.StatusNoContent {
//		return helpers.CustomError{fmt.Sprintf("%s: %s %s\n", helpers.Red("HTTP error"), helpers.Normal("status="), helpers.Normal(resp.Status))}
//	}
//	return nil
//}
