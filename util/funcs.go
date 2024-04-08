package util

import (
	"strings"
	"time"

	"github.com/sanity-io/litter"
)

// An alias for [time.Time] that is correctly unmarshalled from JSON.
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

// Prints the interface to STDOUT in a readable way.
func PrettyPrint(i interface{}) {
	litter.Config.StripPackageNames = true
	litter.Dump(i)
}
