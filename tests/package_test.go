package tests

import (
	"fmt"
	"testing"
	TSGO "thundergo/experimental"
	"thundergo/util"
)

func TestPackage(t *testing.T) {
	var err error
	var pkg *TSGO.Package

	pkg, err = TSGO.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatalf(err.Error())
	}

	cl, _ := pkg.Latest.GetChangelog()
	fmt.Println(util.Prettify(cl))
}
