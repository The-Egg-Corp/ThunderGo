package tests

import (
	"fmt"
	"os"
	"testing"
	TSGO "thundergo/experimental"
	"thundergo/util"
)

// Seems to need Authentication
func Test_ValidateIcon(t *testing.T) {
	var err error
	icon, err := os.ReadFile("../test_icon.png")

	if err != nil {
		t.Fatalf(err.Error())
	}

	str := TSGO.Base64String(icon)
	//fmt.Println(str)

	var res TSGO.ValidatorResponse
	res, err = TSGO.ValidateIcon(str)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.Success == false {
		t.Fatalf("expected success")
	}

	fmt.Println(util.Prettify(res))
}
