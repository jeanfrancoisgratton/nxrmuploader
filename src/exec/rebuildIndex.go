// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// rebuildIndex.go, jfgratton : 2024-02-15

package exec

//
//func main() {
//	// Nexus repository URL
//	repoURL := "https://nexus:1808/repository/aptTest/"
//
//	// Endpoint for rebuilding the index in Nexus
//	endpoint := "service/rest/v1/repositories/deb/index"
//
//	// Full URL for the rebuild index endpoint
//	url := repoURL + endpoint
//
//	// Create an empty request body since rebuilding index may not require any payload
//	requestBody := []byte("{}")
//
//	// Make a POST request to rebuild the index
//	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
//	if err != nil {
//		fmt.Println("Error making the request:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// Check the response status code
//	if resp.StatusCode == http.StatusOK {
//		fmt.Println("Index successfully rebuilt.")
//	} else {
//		fmt.Printf("Failed to rebuild index. Status code: %d\n", resp.StatusCode)
//		// Optionally, you can print the response body for more details
//		// body, _ := ioutil.ReadAll(resp.Body)
//		// fmt.Println(string(body))
//	}
//}
