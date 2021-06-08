package response

type Collection struct {
	Id           string        `json:"_id"`
	LastUpdated  string        `json:"lastUpdated"`
	CreatedOn    string        `json:"createdOn"`
	Name         string        `json:"name"`
	Slug         string        `json:"slug"`
	SingularName string        `json:"singularName"`
	Fields       []interface{} `json:"fields"`
}
