package models

type SBOM struct {
	ProjectName  string
	Languages    []string
	Dependencies map[string][]Dependency
}

type Dependency struct {
	ID         string
	ImportName string
	Version    string
	Licenses   []string
	Language   string
	Source     string
	TopLevel   bool
}
