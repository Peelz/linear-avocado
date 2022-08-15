package entity

type Project struct {
	// ID: any type of unique id
	ID int `json:"id,omitempty"`
	// UUID: any type of unique id
	UUID string `json:"uuid,omitempty"`
	// RepositoryName: string
	Name string `json:"name,omitempty"`
	// RepositoryUrl: string
	URL string `json:"url,omitempty"`
	// Jobs in
	Jobs []Job `json:"jobs,omitempty"`
}
