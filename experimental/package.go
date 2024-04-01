package experimental

import (
	"fmt"
	"thundergo/util"
)

type Package struct {
	Namespace      string         `json:"namespace"`
	Name           string         `json:"name"`
	FullName       string         `json:"full_name"`
	Owner          string         `json:"owner"`
	PackageURL     string         `json:"package_url"`
	DateCreated    string         `json:"date_created"`
	DateUpdated    string         `json:"date_updated"`
	Rating         string         `json:"rating_score"`
	Pinned         bool           `json:"is_pinned"`
	Deprecated     bool           `json:"is_deprecated"`
	TotalDownloads string         `json:"total_downloads"`
	Latest         PackageVersion `json:"latest"`
}

type PackageVersion struct {
	Namespace     string `json:"namespace"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	VersionNumber string `json:"version_number"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	Dependencies  string `json:"dependencies"`
	DateCreated   string `json:"date_created"`
	Downloads     int32  `json:"total_downloads"`
	DownloadURL   string `json:"download_url"`
	WebsiteURL    string `json:"website_url"`
	Active        bool   `json:"is_active"`
}

func (pkg PackageVersion) GetChangelog() (MarkdownResponse, error) {
	endpoint := fmt.Sprint("api/experimental/package/", pkg.Namespace, "/", pkg.Name, "/", pkg.VersionNumber, "/changelog")
	return util.JsonRequest[MarkdownResponse](endpoint)
}

// region ReviewStatus Enum
type ReviewStatus string

const (
	Unreviewed ReviewStatus = "unreviewed"
	Approved   ReviewStatus = "approved"
	Rejected   ReviewStatus = "rejected"
)

func (rs ReviewStatus) Unreviewed() bool {
	return rs == Unreviewed
}

func (rs ReviewStatus) Approved() bool {
	return rs == Approved
}

func (rs ReviewStatus) Rejected() bool {
	return rs == Rejected
}

//endregion

type PackageListing struct {
	HasNsfwContent bool         `json:"has_nsfw_content"`
	Categories     string       `json:"categories"`
	Community      string       `json:"community"`
	ReviewStatus   ReviewStatus `json:"review_status"`
}

type PackageCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Response received for markdown files like README etc.
type MarkdownResponse struct {
	Markdown string `json:"markdown"`
}
