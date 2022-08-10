package unpackers

type Generic struct {
	Parallelism int64       `json:"parallelism" validate:"required,int"`
	BaseURL     string      `json:"base_url" validate:"required,string"`
	Path        string      `json:"path" validate:"required,string"`
	Method      string      `json:"method" validate:"required,string"`
	Payload     interface{} `json:"payload" validate:"required"`
}

// RequiredFormat returns the applicable JSON format for the point data structure
func (m *Generic) RequiredFormat() string {

	return `
	{
        "parallelism": <int>,
		"base_url": <string>,
		"path": <string>,
		"payload": <json>
    }
	`
}
