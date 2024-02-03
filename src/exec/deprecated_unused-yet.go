// // nxrmuploader
// // Written by J.F. Gratton <jean-francois@famillegratton.net>
// // Original filename: src/exec/deprecated_unused-yet.go
// // Original timestamp: 2024/01/07 13:21
package exec

//
//import (
//	"bytes"
//	"encoding/base64"
//	"fmt"
//	"io"
//	"mime/multipart"
//	"net/http"
//	"os"
//)
//
//// See https://help.sonatype.com/repomanager3/integrations/rest-and-integration-api/components-api?&_ga=2.260381966.209045506.1704648175-1169652860.1704484669#ComponentsAPI-UploadComponent
//
//package main
//
//import (
//"bytes"
//"encoding/base64"
//"fmt"
//"io"
//"mime/multipart"
//"net/http"
//"os"
//)
//
//const (
//	uploadURL = "https://example.com/upload" // replace with your actual upload URL
//)
//
//func uploadFile2(filePath, myVar, myVal, username, password string) error {
//	// Open the file to be uploaded
//	file, err := os.Open(filePath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	// Create a buffer to store the multipart form data
//	var requestBody bytes.Buffer
//	writer := multipart.NewWriter(&requestBody)
//
//	// If key-value pair is provided, add it to the request
//	if myVar != "" {
//		writer.WriteField(myVar, myVal)
//	}
//
//	// Create a form file field and add it to the request
//	fileWriter, err := writer.CreateFormFile("file", filePath)
//	if err != nil {
//		return err
//	}
//
//	// Copy the file content to the form file field
//	_, err = io.Copy(fileWriter, file)
//	if err != nil {
//		return err
//	}
//
//	// Close the multipart writer to finish the request
//	writer.Close()
//
//	// Create an HTTP request with the multipart form data
//	request, err := http.NewRequest("POST", uploadURL, &requestBody)
//	if err != nil {
//		return err
//	}
//
//	// Set the content type header for multipart form data
//	request.Header.Set("Content-Type", writer.FormDataContentType())
//
//	// If basic authentication is provided, set the Authorization header
//	if username != "" && password != "" {
//		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
//		request.Header.Set("Authorization", authHeader)
//	}
//
//	// Perform the HTTP request
//	client := &http.Client{}
//	response, err := client.Do(request)
//	if err != nil {
//		return err
//	}
//	defer response.Body.Close()
//
//	// Handle the response as needed (print it for now)
//	fmt.Println("Upload response:", response.Status)
//
//	return nil
//}
//
//func main2() {
//	filePath := "path/to/your/file.txt" // replace with the actual file path
//
//	// First upload with key-value pair and basic authentication
//	err := uploadFile2(filePath, "myVar", "myVal", "myUser", "myPasswd")
//	if err != nil {
//		fmt.Println("Error uploading file with key-value and basic authentication:", err)
//		return
//	}
//
//	// Second upload without key-value pair and basic authentication
//	err = uploadFile2(filePath, "", "", "myUser", "myPasswd")
//	if err != nil {
//		fmt.Println("Error uploading file with basic authentication:", err)
//		return
//	}
//
//	// Third upload without key-value pair and without basic authentication
//	err = uploadFile2(filePath, "", "", "", "")
//	if err != nil {
//		fmt.Println("Error uploading file:", err)
//		return
//	}
//}
