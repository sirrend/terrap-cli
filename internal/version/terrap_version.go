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

var Version = "0.0.3" // changes on build

func (t *TerrapVersion) SetVersion() {
	(*t).Product = "Terrap"
	(*t).Version = Version
	(*t).GoVersion = runtime.Version()
	(*t).System = runtime.GOOS
}
