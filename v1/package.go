package v1

import (
	"fmt"
	"strings"

	"github.com/the-egg-corp/thundergo/util"
)

var pkgCache PackageList

// An alias for a [Package] array with helper functions attached.
type PackageList []Package

// The amount of packages in the list.
//
// Equivalent to len(list) casted to a uint.
func (list PackageList) Size() uint { return uint(len(list)) }

// Performs a filter on the list, returning a new list containing only packages that satisfy the predicate.
func (list PackageList) Filter(predicate func(pkg Package) bool) PackageList {
	arr := make(PackageList, 0, list.Size())

	for _, v := range list {
		if predicate(v) {
			arr = append(arr, v)
		}
	}

	return arr
}

// Grab a single package from the list given the package owner's name and the package's short name.
func (list PackageList) Get(author string, name string) *Package {
	return util.TryFind(list, func(p Package) bool {
		return strings.EqualFold(p.Name, name) && strings.EqualFold(p.Owner, author)
	})
}

// Grab a single package from the list given the package owner's name and the package's short name.
func (list PackageList) GetByUUID(uuid string) *Package {
	return util.TryFind(list, func(p Package) bool {
		return strings.EqualFold(p.UUID, uuid)
	})
}

// Grab a single package from the list given the package's full name.
//
// A full name would look like so:
//
//	"Owen3H-CSync"
func (list PackageList) GetExact(fullName string) *Package {
	return util.TryFind(list, func(p Package) bool {
		return strings.EqualFold(p.FullName, fullName)
	})
}

// Represents a package/mod on Thunderstore that is global and not specific to any community.
//
// To easily find a version from Versions, use [Package.GetVersion].
type Package struct {
	Name           string           `json:"name"`
	FullName       string           `json:"full_name"`
	Owner          string           `json:"owner"`
	UUID           string           `json:"uuid4"`
	PackageURL     string           `json:"package_url"`
	DonationLink   string           `json:"donation_link"`
	DateCreated    util.DateTime    `json:"date_created"`
	DateUpdated    util.DateTime    `json:"date_updated"`
	Rating         uint16           `json:"rating_score"`
	Pinned         bool             `json:"is_pinned"`
	Deprecated     bool             `json:"is_deprecated"`
	HasNsfwContent bool             `json:"has_nsfw_content"`
	Categories     []string         `json:"categories"`
	Versions       []PackageVersion `json:"versions"`
}

// Gets a specific [PackageVersion] from this package's list of versions.
//
// verNumber should be specified in the format: major.minor.patch
//
// Good:
//
//	"v3.0.0", "2.1.1", "1.0.0-beta.1"
//
// Bad:
//
//	"v3.1", "v2", "1.0"
func (pkg Package) GetVersion(verNumber string) *PackageVersion {
	return util.TryFind(pkg.Versions, func(v PackageVersion) bool {
		return strings.EqualFold(v.VersionNumber, strings.Replace(verNumber, "v", "", 1))
	})
}

// Gets this package's statistics such as downloads and likes.
func (pkg Package) Metrics(version ...string) (PackageMetrics, error) {
	endpoint := fmt.Sprint("api/v1/package-metrics/", pkg.Owner, "/", pkg.Name)
	return util.JsonGetRequest[PackageMetrics](endpoint)
}

// A specific version of a package.
//
// Note: This is NOT equivalent to [Package] as its fields differ.
type PackageVersion struct {
	Name          string        `json:"name"`
	FullName      string        `json:"full_name"`
	UUID          string        `json:"uuid4"`
	Dependencies  []string      `json:"dependencies"`
	Description   string        `json:"description"`
	DownloadURL   string        `json:"download_url"`
	Downloads     uint32        `json:"downloads"`
	DateCreated   util.DateTime `json:"date_created"`
	FileSize      uint64        `json:"file_size"`
	Icon          string        `json:"icon"`
	Active        bool          `json:"is_active"`
	VersionNumber string        `json:"version_number"`
	WebsiteURL    string        `json:"website_url"`
}

// type PackageVersion struct {
// 	DateCreated   Time   `json:"date_created"`
// 	Downloads     int32  `json:"download_count"`
// 	DownloadURL   string `json:"download_url"`
// 	InstallURL    string `json:"install_url"`
// 	VersionNumber string `json:"version_number"`
// }

type PackageDependency struct {
	CommunityID   *string `json:"community_identifier"`
	CommunityName *string `json:"community_name"`
	Description   string  `json:"description"`
	ImageSource   *string `json:"image_src"`
	Namespace     string  `json:"namespace"`
	PackageName   string  `json:"package_name"`
	VersionNumber string  `json:"version_number"`
}

type PackageMetrics struct {
	Downloads     uint32 `json:"downloads"`
	Rating        uint16 `json:"rating_score"`
	LatestVersion string `json:"latest_version"`
}

type PackageVersionMetrics struct {
	Downloads uint32 `json:"downloads"`
}
