package util

import (
	"strings"
	"time"

	"github.com/sanity-io/litter"
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

func PrettyPrint(i interface{}) {
	litter.Config.StripPackageNames = true
	litter.Dump(i)
}
