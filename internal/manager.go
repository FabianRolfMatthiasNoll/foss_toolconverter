package internal

import (
	"encoding/json"
	"fmt"
	"foss_toolconverter/internal/models"
	"log"
	"os"
)

type Manager struct {
}

func (Manager) SyftToDep(inputPath, outputPath string, npm bool, npmPath string) {
	var syft Syft
	dependencies, err := syft.Convert(inputPath)
	if err != nil {
		log.Print(err)
	}
	writeFile(dependencies, outputPath)
}

func writeFile(sbom *models.SBOM, outputPath string) {
	jsonData, jsonErr := json.Marshal(sbom)
	if jsonErr != nil {
		log.Print(jsonErr)
	}
	filePath := fmt.Sprintf("%s/dependency.json", outputPath)
	os.WriteFile(filePath, jsonData, 0644)
}
