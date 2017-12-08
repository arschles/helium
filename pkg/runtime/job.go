package runtime

type Job struct {
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Tasks []string `json:"tasks"`
}
