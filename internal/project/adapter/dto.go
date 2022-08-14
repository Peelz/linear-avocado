package adapter

type UpdateProjectRequest struct {
	ID   int    `param:"id"`
	Name string `json:"name,omitempty"`
	URL  string `json:"URL,omitempty"`
}

type CreateProjectRequest struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"URL,omitempty"`
}
