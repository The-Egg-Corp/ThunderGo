package v1

import (
	//"fmt"
	"github.com/the-egg-corp/thundergo/util"
	"strings"

	"github.com/samber/lo"
)

var PackageCache PackageList

type PackageList []Package

func (list PackageList) Size() int { return len(list) }
func (list PackageList) Filter(predicate func(pkg Package) bool) PackageList {
	arr := make(PackageList, 0, len(list))

	for _, v := range list {
		if predicate(v) {
			arr = append(arr, v)
		}
	}

	return arr
}

func (list PackageList) Get(author string, name string) *Package {
	pkg, found := lo.Find(list, func(p Package) bool {
		return strings.EqualFold(p.Name, name) && strings.EqualFold(p.Owner, author)
	})

	if !found {
		return nil
	}

	return &pkg
}

type Package struct {
	Name           string           `json:"name"`
	FullName       string           `json:"full_name"`
	Owner          string           `json:"owner"`
	PackageURL     string           `json:"package_url"`
	DonationLink   string           `json:"donation_link"`
	DateCreated    util.DateTime    `json:"date_created"`
	DateUpdated    util.DateTime    `json:"date_updated"`
	UUID           string           `json:"uuid4"`
	Rating         uint16           `json:"rating_score"`
	Pinned         bool             `json:"is_pinned"`
	Deprecated     bool             `json:"is_deprecated"`
	HasNsfwContent bool             `json:"has_nsfw_content"`
	Categories     []string         `json:"categories"`
	Versions       []PackageVersion `json:"versions"`
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

type PackageDependency struct {
	CommunityID   *string `json:"community_identifier"`
	CommunityName *string `json:"community_name"`
	Description   string  `json:"description"`
	ImageSource   *string `json:"image_src"`
	Namespace     string  `json:"namespace"`
	PackageName   string  `json:"package_name"`
	VersionNumber string  `json:"version_number"`
}

type PackageVersion struct {
	DateCreated   util.DateTime `json:"date_created"`
	Dependencies  []string      `json:"dependencies"`
	Description   string        `json:"description"`
	DownloadURL   string        `json:"download_url"`
	Downloads     uint32        `json:"downloads"`
	FileSize      uint64        `json:"file_size"`
	Name          string        `json:"name"`
	FullName      string        `json:"full_name"`
	Icon          string        `json:"icon"`
	Active        bool          `json:"is_active"`
	VersionNumber string        `json:"version_number"`
	UUID          string        `json:"uuid4"`
	WebsiteURL    string        `json:"website_url"`
}

// type PackageVersion struct {
// 	DateCreated   Time   `json:"date_created"`
// 	Downloads     int32  `json:"download_count"`
// 	DownloadURL   string `json:"download_url"`
// 	InstallURL    string `json:"install_url"`
// 	VersionNumber string `json:"version_number"`
// }

type PackageMetrics struct {
	Downloads     uint32 `json:"downloads"`
	Rating        uint16 `json:"rating_score"`
	LatestVersion string `json:"latest_version"`
}

type PackageVersionMetrics struct {
	Downloads uint32 `json:"downloads"`
}
