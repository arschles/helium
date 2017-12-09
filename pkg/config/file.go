package config

type File struct {
	Jobs      map[string]*Job      `toml:"jobs"`
	Workflows map[string]*Workflow `toml:"workflows"`
}
