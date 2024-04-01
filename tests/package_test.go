package tests

import (
	"fmt"
	"testing"
	TSGO "thundergo"
	"thundergo/util"
)

func TestPackage(t *testing.T) {
	var err error
	var pkg *TSGO.PackageExperimental

	pkg, err = TSGO.Experimental.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatalf(err.Error())
	}

	cl, _ := pkg.Latest.GetChangelog()
	fmt.Println(util.Prettify(cl))
}
