package response

type Site struct {
	Id            string `json:"_id"`
	CreatedOn     string `json:"createdOn"`
	Name          string `json:"name"`
	ShortName     string `json:"shortName"`
	LastPublished string `json:"lastPublished"`
	PreviewUrl    string `json:"previewUrl"`
	Timezone      string `json:"timezone"`
	Database      string `json:"database"`
}
