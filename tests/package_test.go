package tests

import (
	"fmt"
	"testing"
	TSGOExp "thundergo/experimental"
	"thundergo/util"
	TSGOV1 "thundergo/v1"

	"github.com/samber/lo"
)

func TestPackage(t *testing.T) {
	var err error
	var pkg *TSGOExp.Package

	pkg, err = TSGOExp.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatalf(err.Error())
	}

	cl, _ := pkg.Latest.GetReadme()
	fmt.Println(util.Prettify(cl))
}

func TestAllPackages(t *testing.T) {
	var err error
	var pkgs []TSGOV1.PackageListing

	pkgs, err = TSGOV1.GetCommunityPackages("lethal-company")
	if err != nil {
		t.Fatalf(err.Error())
	}

	pkg := lo.Filter(pkgs, func(item TSGOV1.PackageListing, index int) bool {
		return item.Name == "CSync"
	})

	fmt.Println(util.Prettify(pkg))
}
