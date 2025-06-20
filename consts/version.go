package consts

// Version can be set at build time with -ldflags
var Version = "v0.1.0"

// GetVersion returns the formatted version string
func GetVersion() string {
	return "Notion Todo CLI " + Version
}
