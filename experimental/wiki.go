package experimental

import (
	"thundergo/util"
)

type WikiList struct {
	Results []PackageWiki `json:"results"`
	Cursor  string        `json:"cursor"`
	HasMore bool          `json:"has_more"`
}

type Wiki struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	DateCreated util.DateTime   `json:"datetime_created"`
	DateUpdated util.DateTime   `json:"datetime_updated"`
	Pages       []WikiPageIndex `json:"pages"`
}

type WikiPage struct {
	Id              string        `json:"id"`
	Title           string        `json:"title"`
	Slug            string        `json:"slug"`
	DateCreated     util.DateTime `json:"datetime_created"`
	DateUpdated     util.DateTime `json:"datetime_updated"`
	MarkdownContent string        `json:"markdown_content"`
}

type WikiPageIndex struct {
	Id          string        `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	DateCreated util.DateTime `json:"datetime_created"`
	DateUpdated util.DateTime `json:"datetime_updated"`
}
