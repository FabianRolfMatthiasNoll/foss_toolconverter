package models

type SBOM struct {
	ProjectName  string
	Languages    []string
	Dependencies map[string][]Dependency
}

type Dependency struct {
	ID         string //Hash value
	ImportName string //The Name that is imported. In some languages it is a name and some a url
	Version    string
	Licenses   []string
	Language   string
	//This tool should mostly work language independent (except npm)
	//Source     string //constructing a common url for that language
	TopLevel bool //Important for npm and Docker
}
