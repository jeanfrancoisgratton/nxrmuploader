package main

import (
	"nxrmuploader/cmd"
	"os"
	"path/filepath"
)

func main() {

	os.Mkdir(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmuploader"), os.ModePerm)
	cmd.Execute()
}
