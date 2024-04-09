package util

import (
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/samber/lo"
	"github.com/sanity-io/litter"
)

// An alias for [time.Time] that is correctly unmarshalled from JSON.
type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	date, err := time.Parse(time.RFC3339, str)

	if err != nil {
		return err
	}

	dt.Time = date
	return nil
}

// Formats this date into a relative string like so:
//
//	"6 seconds ago", "2 months ago" or "3 days from now"
func (dt DateTime) Humanize() string {
	return humanize.Time(dt.Time)
}

// Prints the interface to STDOUT in a readable way.
func PrettyPrint(i interface{}) {
	litter.Config.StripPackageNames = true
	litter.Dump(i)
}

func TryFind[T any](arr []T, pred func(pkg T) bool) *T {
	pkg, found := lo.Find(arr, pred)
	return lo.Ternary(found, &pkg, nil)
}
