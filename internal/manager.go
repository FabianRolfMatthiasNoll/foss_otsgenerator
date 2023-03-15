package internal

import (
	"fmt"
	"foss_otsgenerator/internal/models"

	"github.com/TwiN/go-color"
)

type Manager struct {
}

func (Manager) GenerateOTS(inputPath, configPath, outputPath string) {
	var sbomManager models.SBOM
	sbom, err := sbomManager.ReadAndMapSBOM(inputPath)
	if err != nil {
		fmt.Println("[", color.Colorize(color.Red, "Err"), "] ", err)
		//return
	}

	for lan, dep := range sbom.Dependencies {
		fmt.Println(lan, "||", dep)
	}
}
