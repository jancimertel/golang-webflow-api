package response

import "time"

type Item struct {
	Archived       bool      `json:"_archived"`
	Draft          bool      `json:"_draft"`
	Color          string    `json:"color"`
	Featured       bool      `json:"featured"`
	Name           string    `json:"name"`
	PostBody       string    `json:"post-body"`
	PostSummary    string    `json:"post-summary"`
	ThumbnailImage File      `json:"thumbnail-image"`
	MainImage      File      `json:"main-image"`
	Slug           string    `json:"slug"`
	UpdatedOn      time.Time `json:"updated-on"`
	UpdatedBy      string    `json:"updated-by"`
	CreatedOn      time.Time `json:"created-on"`
	CreatedBy      string    `json:"created-by"`
	PublishedOn    time.Time `json:"published-on"`
	PublishedBy    string    `json:"published-by"`
	Author         string    `json:"author"`
	Cid            string    `json:"_cid"`
	Id             string    `json:"_id"`
}

type File struct {
	FileId string `json:"fileId"`
	Url    string `json:"url"`
}

type Items struct {
	Items  []Item `json:"items"`
	Count  int    `json:"count"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
}
