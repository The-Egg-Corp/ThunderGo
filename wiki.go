package thundergo

type Wiki struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	DateCreated string          `json:"datetime_created"`
	DateUpdated string          `json:"datetime_updated"`
	Pages       []WikiPageIndex `json:"pages"`
}

type WikiPage struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Slug            string `json:"slug"`
	DateCreated     string `json:"datetime_created"`
	DateUpdated     string `json:"datetime_updated"`
	MarkdownContent string `json:"markdown_content"`
}

type WikiPageIndex struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	DateCreated string `json:"datetime_created"`
	DateUpdated string `json:"datetime_updated"`
}
