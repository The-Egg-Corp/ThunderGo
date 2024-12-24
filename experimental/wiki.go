package experimental

import (
	"fmt"

	"github.com/the-egg-corp/thundergo/util"
)

type WikiList struct {
	Results []PackageWiki `json:"results"`
	Cursor  string        `json:"cursor"`
	HasMore bool          `json:"has_more"`
}

// Represents the wiki section/tab of a Thunderstore package.
type Wiki struct {
	WikiPage
	Pages []WikiPageIndex `json:"pages"`
}

type WikiPage struct {
	WikiPageIndex
	MarkdownContent string `json:"markdown_content"`
}

type WikiPageIndex struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	DateCreated util.DateTime `json:"datetime_created"`
	DateUpdated util.DateTime `json:"datetime_updated"`
}

type WikiPageUpsert struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	MarkdownContent string `json:"markdown_content"`
}

type WikiPageDelete struct {
	ID string `json:"id"`
}

// Requests the content of this page. The received string will usually be formatted as Markdown.
//
// # API Reference: experimental_package_wiki_write
func (index WikiPageIndex) GetContent() (*string, error) {
	endpoint := fmt.Sprint("api/experimental/wiki/page/", index.ID)

	res, _, err := util.JsonGetRequest[WikiPage](endpoint)
	if res == nil {
		return nil, err
	}

	return &res.MarkdownContent, nil
}

// Dummy description
//
// # API Reference: experimental_package_wiki_read
func GetWiki(namespace, name string) (*Wiki, *int, error) {
	endpoint := fmt.Sprintf("api/experimental/wiki/page/%s/%s", namespace, name)
	return util.JsonGetRequest[Wiki](endpoint)
}

// Dummy description
//
// # API Reference: experimental_package_wiki_delete
func DeleteWikiPage() {

}

// Dummy description
//
// # API Reference: experimental_package_wiki_write
func CreateWikiPage(title, markdownContent string) {

}

// Dummy description
//
// # API Reference: experimental_package_wiki_write
func UpdateWikiPage() {

}
