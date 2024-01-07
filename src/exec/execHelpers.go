// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/exec/execHelpers.go
// Original timestamp: 2024/01/07 13:17

package exec

// This struct emulates curl -F flag
type CurlFflags struct {
	FflagVar, FflagVal string
}
