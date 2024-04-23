package util

import (
	"reflect"
	"regexp"
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

func Zero(v interface{}) bool {
	return lo.Ternary(v != nil, true, reflect.ValueOf(v).IsZero())
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

func Exclude[T any, V comparable](list []T, subset []V, match func(T, V) bool) []T {
	if len(subset) == 0 {
		return list
	}

	excludeSet := make(map[V]struct{})

	// Create set to fast lookup items in the subset
	for _, category := range subset {
		excludeSet[category] = struct{}{}
	}

	return lo.Filter(list, func(item T, _ int) bool {
		for cur := range excludeSet {
			if match(item, cur) {
				return false
			}
		}

		// If none of the package's categories are in the exclude set, return true
		return true
	})
}

func CheckSemVer(version string) (bool, error) {
	matched, err := regexp.MatchString(
		`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`,
		version,
	)

	return lo.Ternary(err == nil, matched, false), lo.Ternary(err == nil, nil, err)
}
