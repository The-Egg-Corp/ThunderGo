package util

import "github.com/sanity-io/litter"

func Prettify(i interface{}) string {
	litter.Config.StripPackageNames = true
	return litter.Sdump(i)
}
