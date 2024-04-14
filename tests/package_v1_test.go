package tests

import (
	"errors"
	"testing"

	"github.com/samber/lo"
	"github.com/the-egg-corp/thundergo/util"
	TSGOV1 "github.com/the-egg-corp/thundergo/v1"
)

func TestAllPackages(t *testing.T) {
	var err error
	var pkgs TSGOV1.PackageList

	pkgs, err = TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err.Error())
	}

	pkgsMap := lo.Associate(pkgs, func(pkg TSGOV1.Package) (string, string) {
		return pkg.FullName, pkg.UUID
	})

	util.PrettyPrint(pkgsMap)
}

func TestCommunityPackages(t *testing.T) {
	comm := TSGOV1.Community{
		Identifier: "lethal-company",
	}

	pkgs, _ := comm.AllPackages()
	pkgs = pkgs.Filter(func(pkg TSGOV1.Package) bool {
		return pkg.Owner == "Owen3H"
	})

	util.PrettyPrint(pkgs)
}

func TestPackageVersion(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err.Error())
	}

	pkg := pkgs.Get("Owen3H", "CSync").GetVersion("2.0.0")
	util.PrettyPrint(pkg)
}

func TestPackageGet(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err.Error())
	}

	pkg := pkgs.Get("Owen3H", "CSync")
	if pkg == nil {
		t.Fatal(errors.New("could not get package given the name and author"))
	}

	util.PrettyPrint(pkg)
}

func TestPackageGetExact(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err.Error())
	}

	pkg := pkgs.GetExact("Owen3H-CSync")
	if pkg == nil {
		t.Fatal(errors.New("could not get package by its full name"))
	}

	util.PrettyPrint(pkg)
}

func TestPackageGetByUUID(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err.Error())
	}

	pkg := pkgs.GetByUUID("13d217b1-1e90-431a-a826-cd29c9eaea36")
	if pkg == nil {
		t.Fatal(errors.New("could not get package using UUID"))
	}

	util.PrettyPrint(pkg)
}

func TestMetrics(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatal(err.Error())
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
		t.Fatal(err.Error())
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
