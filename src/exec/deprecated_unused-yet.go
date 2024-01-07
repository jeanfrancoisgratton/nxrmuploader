// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/exec/deprecated_unused-yet.go
// Original timestamp: 2024/01/07 13:21

package exec

// See https://help.sonatype.com/repomanager3/integrations/rest-and-integration-api/components-api?&_ga=2.260381966.209045506.1704648175-1169652860.1704484669#ComponentsAPI-UploadComponent

/*
func uploadFile(repoURL, filePath, myvar1, myval1, myvar2, myval2 string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    // Add form fields
    _ = writer.WriteField(myvar1, myval1)
    _ = writer.WriteField(myvar2, myval2)

    // Create a form field for the file
    part, err := writer.CreateFormFile("file", filepath.Base(filePath))
    if err != nil {
        return err
    }

    // Copy file content to the part
    _, err = io.Copy(part, file)
    if err != nil {
        return err
    }

    // Close the multipart writer before sending the request
    err = writer.Close()
    if err != nil {
        return err
    }

    // Create HTTP request
    req, err := http.NewRequest("POST", repoURL, body)
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    // Replace 'username' and 'password' with your Nexus credentials
    req.SetBasicAuth("username", "password")

    // Perform the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status: %s", resp.Status)
    }

    return nil
}
*/
