package version

import (
	"runtime"
)

type TerrapVersion struct {
	Product   string
	Version   string
	GoVersion string
	System    string
}

var Version = "filled on build"

func (tv *TerrapVersion) SetVersion() {
	(*tv).Product = "Terrap"
	(*tv).Version = Version
	(*tv).GoVersion = runtime.Version()
	(*tv).System = runtime.GOOS
}
