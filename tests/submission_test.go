package tests

import (
	TSGO "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
	"os"
	"testing"
)

func TestValidateIcon(t *testing.T) {
	var err error
	icon, err := os.ReadFile("../test_icon.png")

	if err != nil {
		t.Fatalf(err.Error())
	}

	valid, err := TSGO.ValidateIcon(icon)
	if err != nil {
		t.Fatalf(err.Error())
	}

	util.PrettyPrint(valid)
}
