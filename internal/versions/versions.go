package versions

import (
	"github.com/Masterminds/semver/v3"
	validate "golang.org/x/mod/semver"
)

type SemverVersion struct {
	Version             *semver.Version
	Major, Minor, Patch uint64
}

func (v *SemverVersion) Init(ver string) bool {
	if validate.IsValid(ver) {
		parsed := semver.MustParse(ver)
		*v = SemverVersion{
			Version: parsed,
			Major:   parsed.Major(),
			Minor:   parsed.Minor(),
			Patch:   parsed.Patch(),
		}
		return true
	}
	return false
}

/*
@brief: IsOlderThen checks if version is lower than given string
@
@params: new - string - version as string
@
@return: bool - true or false
*/

func (v *SemverVersion) IsOlderThen(new string) bool {
	if validate.IsValid(new) {
		newVer := semver.MustParse(new)

		return v.Version.LessThan(newVer)
	}

	return false
}

/*
@brief: IsNewerThen checks if version is bigger than given string
@
@params: new - string - version as string
@
@return: bool - true or false
*/

func (v *SemverVersion) IsNewerThen(new string) bool {
	if validate.IsValid(new) {
		newVer := semver.MustParse(new)

		return v.Version.GreaterThan(newVer)
	}

	return false
}
