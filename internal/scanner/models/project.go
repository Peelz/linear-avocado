package models

type Project struct {
	// UUID: any type of unique id
	UUID string `json:"uuid,omitempty"`
	// RepositoryName: string
	Name string `json:"name,omitempty"`
	// RepositoryUrl: string
	URL string `json:"url,omitempty"`
}
