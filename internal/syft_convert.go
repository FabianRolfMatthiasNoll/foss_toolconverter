package internal

import "fmt"

type Syft struct {
	Artifacts []struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Version  string   `json:"version"`
		Licenses []string `json:"licenses"`
		Language string   `json:"language"`
		Purl     string   `json:"purl"`
	} `json:"artifacts"`
	ArtifactRelationships []struct {
		Parent string `json:"parent"`
		Child  string `json:"child"`
		Type   string `json:"type"`
	} `json:"artifactRelationships"`
	Source struct {
		//used to identify source id in ArtifactRelations
		ID string `json:"id"`
	} `json:"source"`
}

func (s *Syft) Convert() {
	fmt.Println("Converting")
}
