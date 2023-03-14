package internal

import (
	"encoding/json"
	"foss_toolconverter/internal/models"
	"os"
	"strings"
)

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

func (syft *Syft) readJson(path string) (*Syft, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, syft)
	if err != nil {
		return nil, err
	}
	return syft, err
}

func (s *Syft) Convert(inputPath string) (*models.SBOM, error) {
	syftReader := Syft{}
	syft, syftErr := syftReader.readJson(inputPath)
	if syftErr != nil {
		return nil, syftErr
	}

	Dependencies := &models.SBOM{}

	if Dependencies.Dependencies == nil {
		Dependencies.Dependencies = make(map[string][]models.Dependency)
	}

	for _, dep := range syft.Artifacts {
		var toplevel = true

		//Generating Language out of the start of the purl
		language := strings.Split(dep.Purl, "/")[0][4:]

		if !contains(Dependencies.Languages, language) {
			Dependencies.Languages = append(Dependencies.Languages, language)
		}

		dependency := models.Dependency{
			ImportName: dep.Name,
			Language:   language,
			Version:    dep.Version,
			Licenses:   dep.Licenses,
			ID:         dep.ID,
			TopLevel:   toplevel,
		}

		if _, exists := Dependencies.Dependencies[language]; !exists {
			Dependencies.Dependencies[language] = []models.Dependency{}
		}

		Dependencies.Dependencies[language] = append(Dependencies.Dependencies[language], dependency)
	}
	return Dependencies, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
