package main

import (
	"nxrmuploader/cmd"
	"os"
	"path/filepath"
)

func main() {

	os.Mkdir(filepath.Join(os.Getenv("HOME"), ".config", "nxrmuploader"), os.ModePerm)
	cmd.Execute()
}
