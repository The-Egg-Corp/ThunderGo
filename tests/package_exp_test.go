package tests

import (
	"errors"
	"testing"

	TSGOExp "github.com/the-egg-corp/thundergo/experimental"
)

func TestPackageExp(t *testing.T) {
	var err error
	pkg, err := TSGOExp.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatal(err)
	}

	if pkg != nil {
		_, e := pkg.Latest.Readme()

		if e != nil {
			t.Fatal(errors.New("unable to get the README from the latest version of the specified package"))
		}

		//util.PrettyPrint(cl)
	}
}
