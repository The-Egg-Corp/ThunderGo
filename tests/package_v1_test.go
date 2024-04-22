package tests

import (
	"errors"
	"fmt"
	"testing"
	"time"

	//"github.com/samber/lo"

	"github.com/the-egg-corp/thundergo/util"
	TSGOV1 "github.com/the-egg-corp/thundergo/v1"
)

func TestAllPackages(t *testing.T) {
	var err error
	var pkgs TSGOV1.PackageList

	pkgs, err = TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pkgs.Size())
}

func TestPackagesFromList(t *testing.T) {
	var err error
	var pkgs TSGOV1.PackageList

	pkgs, err = TSGOV1.PackagesFromCommunities(TSGOV1.NewCommunityList("riskofrain2", "valheim"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pkgs.Size())
}

func TestCommunityPackages(t *testing.T) {
	comm := TSGOV1.Community{
		Identifier: "riskofrain2",
	}

	pkgs, _ := comm.AllPackages()
	// pkgs = pkgs.Filter(func(pkg TSGOV1.Package) bool {
	// 	return pkg.Owner == "Owen3H"
	// })

	time.Sleep(100 * time.Millisecond)
	util.PrettyPrint(pkgs.Size())
}

func TestPackageVersion(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	ver := pkgs.Get("Owen3H", "CSync").GetVersion("3.0.0")
	if ver == nil {
		t.Fatal(errors.New("could not find specific package version"))
	}
}

func TestPackageGet(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	pkg := pkgs.Get("Owen3H", "CSync")
	if pkg == nil {
		t.Fatal(errors.New("could not get package given the name and author"))
	}

	//util.PrettyPrint(pkg)
}

func TestPackageGetExact(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	pkg := pkgs.GetExact("Owen3H-CSync")
	if pkg == nil {
		t.Fatal(errors.New("could not get package by its full name"))
	}

	//util.PrettyPrint(pkg)
}

func TestPackageGetByUUID(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	pkg := pkgs.GetByUUID("13d217b1-1e90-431a-a826-cd29c9eaea36")
	if pkg == nil {
		t.Fatal(errors.New("could not get package using UUID"))
	}

	//util.PrettyPrint(pkg)
}

func TestMetrics(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	metrics, err := pkgs.Get("Owen3H", "CSync").Metrics()
	if err != nil {
		t.Fatal(err)
	}

	util.PrettyPrint(metrics)
}

func TestPackageDates(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err)
	}

	pkg := pkgs.Get("Owen3H", "CSync")

	if pkg.DateCreated.IsZero() {
		t.Error("DateCreated should be valid and not its zero value.")
	}

	if pkg.DateUpdated.IsZero() {
		t.Error("DateUpdated should be valid and not its zero value.")
	}

	util.PrettyPrint("Created " + pkg.DateCreated.Humanize())
	util.PrettyPrint("Updated " + pkg.DateUpdated.Humanize())
}

func TestPackageFilter(t *testing.T) {
	comm := TSGOV1.Community{
		Identifier: "riskofrain2",
	}

	pkgs, err := comm.AllPackages()

	if err != nil {
		t.Fatal(err)
	}

	filtered := pkgs.ExcludeCategories("modpack", "modpacks")

	fmt.Println(filtered.Size())
}
