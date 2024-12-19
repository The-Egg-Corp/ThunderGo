package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	TSGO "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
)

const iconPath = "../test-pkg/icon.png.jpg"
const manifestPath = "../test-pkg/manifest.json"

func TestValidateIcon(t *testing.T) {
	//t.Skip()

	icon, err := os.ReadFile(iconPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, err := TSGO.ValidateIcon(filepath.Base(iconPath), icon)
	if err != nil {
		t.Fatal(err.Error())
	}

	util.PrettyPrint(valid)
}

func TestValidateManifest(t *testing.T) {
	t.Skip()

	var errs []string

	data, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, errs, err := TSGO.ValidateManifest("Owen3H", data)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("Valid: ", valid)

	if len(errs) > 0 {
		util.PrettyPrint(errs)

		if valid {
			t.Fatal("errors were returned, but manifest is still valid")
		}

		return
	}

	if !valid {
		t.Fatal("manifest was marked as invalid despite empty errors array")
	}
}
