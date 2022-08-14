package api

type CreateProjectRequest struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"URL,omitempty"`
}

type UpdateProjectRequest struct {
	ID   string `param:"id"`
	Name string `json:"name,omitempty"`
	URL  string `json:"URL,omitempty"`
}

type DeleteProjectRequest struct {
}
