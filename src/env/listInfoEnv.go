// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/listInfoEnv.go
// Original timestamp: 2023/12/31 14:50

package env

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"nxrmuploader/helpers"
	"os"
	"path/filepath"
	"strings"
)

func ListEnvironments(envdir string) error {
	var err error
	var dirFH *os.File
	var finfo, fileInfos []os.FileInfo

	// list environment files
	if envdir == "" {
		envdir = filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmuploader")
	}
	if dirFH, err = os.Open(envdir); err != nil {
		return helpers.CustomError{Message: "Unable to read config directory: " + err.Error()}
	}

	if fileInfos, err = dirFH.Readdir(0); err != nil {
		return helpers.CustomError{Message: "Unable to read files in config directory: " + err.Error()}
	}

	for _, info := range fileInfos {
		//if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") && !strings.HasPrefix(info.Name(), "sample") {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			finfo = append(finfo, info)
		}
	}
	//	return nil
	//})

	if err != nil {
		return err
	}

	fmt.Printf("Found %s environment files in %s\n", helpers.Green(fmt.Sprintf("%d", len(finfo))),
		helpers.Green(fmt.Sprintf("%s", filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmuploader"))))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Environment file", "File size", "Modification time"})

	for _, fi := range finfo {
		t.AppendRow([]interface{}{helpers.Green(fi.Name()), helpers.Green(helpers.SI(uint64(fi.Size()))),
			helpers.Green(fmt.Sprintf("%v", fi.ModTime().Format("2006/01/02 15:04:05")))})
	}
	t.SortBy([]table.SortBy{
		{Name: "Environment file", Mode: table.Asc},
		{Name: "File size", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatDefault
	t.Render()

	return nil
}

func ExplainEnvFile(envfiles []string) error {
	//oldEnvFile := EnvConfigFile
	//
	//fmt.Println("Paths are relative to Certificate root dir's path")
	//t := table.NewWriter()
	//t.SetOutputMirror(os.Stdout)
	//t.AppendHeader(table.Row{"Environment file", "Certificate root dir", "CA dir", "Server certificates dir", "Certificates config dir"})
	//
	//for _, envfile := range envfiles {
	//	if !strings.HasSuffix(envfile, ".json") {
	//		envfile += ".json"
	//	}
	//	EnvConfigFile = envfile
	//
	//	if e, err := LoadEnvironmentFile(); err != nil {
	//		EnvConfigFile = oldEnvFile
	//		return err
	//	} else {
	//		t.AppendRow([]interface{}{helpers.Green(envfile), helpers.Green(e.CertificateRootDir), helpers.Green(filepath.Base(e.RootCAdir)),
	//			helpers.Green(filepath.Base(e.ServerCertsDir)), helpers.Green(filepath.Base(e.CertificatesConfigDir))})
	//	}
	//
	//}
	//t.SortBy([]table.SortBy{
	//	{Name: "Environment file", Mode: table.Asc},
	//})
	//t.SetStyle(table.StyleBold)
	//t.Style().Format.Header = text.FormatDefault
	//t.Render()
	//
	//EnvConfigFile = oldEnvFile
	return nil
}
