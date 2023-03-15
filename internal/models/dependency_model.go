package models

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	TopLevel   bool
}

func (sbom *SBOM) ReadAndMapSBOM(inputPath string) (*SBOM, error) {
	file, err := os.ReadFile(fmt.Sprintf("%s/dependency.json", inputPath))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, sbom)
	if err != nil {
		return nil, err
	}
	return sbom, nil
}
