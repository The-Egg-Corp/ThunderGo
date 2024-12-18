package experimental

import (
	"fmt"

	"github.com/the-egg-corp/thundergo/common"
	"github.com/the-egg-corp/thundergo/util"
)

type Package struct {
	Namespace      string         `json:"namespace"`
	Name           string         `json:"name"`
	FullName       string         `json:"full_name"`
	Owner          string         `json:"owner"`
	PackageURL     string         `json:"package_url"`
	DateCreated    util.DateTime  `json:"date_created"`
	DateUpdated    util.DateTime  `json:"date_updated"`
	Rating         string         `json:"rating_score"`
	Pinned         bool           `json:"is_pinned"`
	Deprecated     bool           `json:"is_deprecated"`
	TotalDownloads string         `json:"total_downloads"`
	Latest         PackageVersion `json:"latest"`
}

type PackageVersion struct {
	Namespace     string        `json:"namespace"`
	Name          string        `json:"name"`
	FullName      string        `json:"full_name"`
	DateCreated   util.DateTime `json:"date_created"`
	VersionNumber string        `json:"version_number"`
	Description   string        `json:"description"`
	Icon          string        `json:"icon"`
	Dependencies  []string      `json:"dependencies"`
	Downloads     int32         `json:"total_downloads"`
	DownloadURL   string        `json:"download_url"`
	WebsiteURL    string        `json:"website_url"`
	Active        bool          `json:"is_active"`
}

func (pkg PackageVersion) Changelog() (string, error) {
	res, err := pkg.getMarkdown("/changelog")
	return res.Markdown, err
}

func (pkg PackageVersion) Readme() (string, error) {
	res, err := pkg.getMarkdown("/readme")
	return res.Markdown, err
}

func (pkg PackageVersion) getMarkdown(file string) (common.MarkdownResponse, error) {
	endpoint := fmt.Sprint("api/experimental/package/", pkg.Namespace, "/", pkg.Name, "/", pkg.VersionNumber, file)
	return util.JsonGetRequest[common.MarkdownResponse](endpoint)
}

type ReviewStatus string

const (
	UNREVIEWED ReviewStatus = "unreviewed"
	APPROVED   ReviewStatus = "approved"
	REJECTED   ReviewStatus = "rejected"
)

func (rs ReviewStatus) Unreviewed() bool { return rs == UNREVIEWED }
func (rs ReviewStatus) Approved() bool   { return rs == APPROVED }
func (rs ReviewStatus) Rejected() bool   { return rs == REJECTED }

type PackageListing struct {
	HasNsfwContent bool         `json:"has_nsfw_content"`
	Categories     string       `json:"categories"`
	Community      string       `json:"community"`
	ReviewStatus   ReviewStatus `json:"review_status"`
}

type PackageWiki struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Wiki      Wiki   `json:"wiki"`
}
