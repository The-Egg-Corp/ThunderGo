package experimental

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/the-egg-corp/thundergo/util"
)

type WikiList struct {
	Results []PackageWiki `json:"results"`
	Cursor  string        `json:"cursor"`
	HasMore bool          `json:"has_more"`
}

// Represents the wiki section/tab of a Thunderstore package.
type Wiki struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	DateCreated util.DateTime   `json:"datetime_created"`
	DateUpdated util.DateTime   `json:"datetime_updated"`
	Pages       []WikiPageIndex `json:"pages"`
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

// Requests the content of this page.
//
// The received string will usually be formatted as Markdown.
func (index WikiPageIndex) GetContent() (*string, error) {
	endpoint := fmt.Sprint("api/experimental/wiki/page/", index.ID)
	res, err := util.JsonGetRequest[WikiPage](endpoint)

	return lo.Ternary(err == nil, nil, &res.MarkdownContent), nil
}
