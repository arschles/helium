package config

type Job struct {
	Image string   `toml:"image"`
	Tasks []string `toml:"tasks"`
}
