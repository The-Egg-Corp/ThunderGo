package v1

import (
	"strings"
	"thundergo/util"
	"time"
)

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	date, err := time.Parse(time.RFC3339, str)

	if err != nil {
		return err
	}

	t.Time = date
	return nil
}

func GetAllPackages() ([]PackageListing, error) {
	return util.JsonGetRequest[[]PackageListing]("api/v1/package")
}
