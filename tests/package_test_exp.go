package tests

import (
	"testing"

	TSGOExp "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
)

func TestPackageExp(t *testing.T) {
	var err error
	var pkg *TSGOExp.Package

	pkg, err = TSGOExp.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatal(err.Error())
	}

	cl, _ := pkg.Latest.Readme()
	util.PrettyPrint(cl)
}
