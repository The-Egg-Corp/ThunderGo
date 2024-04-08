package v1

import (
	//"fmt"
	"github.com/the-egg-corp/thundergo/common"
	"github.com/the-egg-corp/thundergo/util"
	"strings"

	"github.com/samber/lo"
)

var PackageCache PackageList

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
	pkg, found := lo.Find(list, func(p Package) bool {
		return strings.EqualFold(p.Name, name) && strings.EqualFold(p.Owner, author)
	})

	if !found {
		return nil
	}

	return &pkg
}

// Grab a single package from the list given the package's full name.
//
// A full name would look like so:
//
//	"Owen3H-CSync-2.2.4"
func (list PackageList) GetExact(fullName string) *Package {
	pkg, found := lo.Find(list, func(p Package) bool {
		return strings.EqualFold(p.FullName, fullName)
	})

	if !found {
		return nil
	}

	return &pkg
}

type Package struct {
	common.BasePackageMetadata
	Owner          string           `json:"owner"`
	UUID           string           `json:"uuid4"`
	PackageURL     string           `json:"package_url"`
	DonationLink   string           `json:"donation_link"`
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
	ver, found := lo.Find(pkg.Versions, func(v PackageVersion) bool {
		return strings.EqualFold(v.VersionNumber, strings.Replace(verNumber, "v", "", 1))
	})

	if !found {
		return nil
	}

	return &ver
}

// type CommunityPackage struct {
// 	Community string `json:"community"`
// 	Package
// }

// func (pkg CommunityPackage) Metrics() (PackageMetrics, error) {
// 	endpoint := fmt.Sprint("c/", pkg.Community, "/api/v1/package-metrics/", pkg.Owner, "/", pkg.Name)
// 	return util.JsonGetRequest[PackageMetrics](endpoint)
// }

// func (pkg CommunityPackage) VersionMetrics(version string) (PackageVersionMetrics, error) {
// 	endpoint := fmt.Sprint("c/", pkg.Community, "/api/v1/package-metrics/", pkg.Owner, "/", pkg.Name, "/", pkg.Versions[0])
// 	return util.JsonGetRequest[PackageVersionMetrics](endpoint)
// }

type PackageVersion struct {
	common.BasePackageMetadata
	UUID          string   `json:"uuid4"`
	Dependencies  []string `json:"dependencies"`
	Description   string   `json:"description"`
	DownloadURL   string   `json:"download_url"`
	Downloads     uint32   `json:"downloads"`
	FileSize      uint64   `json:"file_size"`
	Icon          string   `json:"icon"`
	Active        bool     `json:"is_active"`
	VersionNumber string   `json:"version_number"`
	WebsiteURL    string   `json:"website_url"`
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
