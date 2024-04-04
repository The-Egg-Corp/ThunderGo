package v1

import "thundergo/util"

type PackageListing struct {
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

type PackageCategory struct {
}

type PackageDependency struct {
	CommunityID   *string `json:"community_identifier"`
	CommunityName *string `json:"community_name"`
	Description   string  `json:"description"`
	ImageSource   *string `json:"image_src"`
	Namespace     string  `json:"namespace"`
	PackageName   string  `json:"package_name"`
	VersionNumber string  `json:"version_number"`
}

// type PackageVersion struct {
// 	DateCreated   Time   `json:"date_created"`
// 	Downloads     int32  `json:"download_count"`
// 	DownloadURL   string `json:"download_url"`
// 	InstallURL    string `json:"install_url"`
// 	VersionNumber string `json:"version_number"`
// }

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

type PackageMetrics struct {
	Downloads     uint32 `json:"downloads"`
	Rating        uint16 `json:"rating_score"`
	LatestVersion string `json:"latest_version"`
}

type PackageVersionMetrics struct {
	Downloads uint32 `json:"downloads"`
}
