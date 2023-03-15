package internal

import (
	"fmt"
	"foss_otsgenerator/internal/handlers"
	"foss_otsgenerator/internal/models"

	"github.com/TwiN/go-color"
)

type Handler_Interface interface {
	GetPackageInfo([]models.Dependency) []models.Module
}

type Manager struct {
}

func (Manager) GenerateOTS(inputPath, configPath, outputPath string) {
	var sbomManager models.SBOM
	var BuildInfo models.BuildInfo
	sbom, err := sbomManager.ReadAndMapSBOM(inputPath)
	if err != nil {
		fmt.Println("[", color.Colorize(color.Red, "Err"), "] ", err)
		return
	}

	var handler Handler_Interface
	for lang, deps := range sbom.Dependencies {
		switch {
		case lang == "nuget":
			handler = handlers.Nuget{}
		case lang == "golang":
			handler = handlers.Go{}
		case lang == "npm":
			handler = handlers.NPM{}
		case lang == "conan":
			handler = handlers.Conan{}
		default:
			fmt.Printf("Unsupported language")
		}

		modules := handler.GetPackageInfo(deps)
		BuildInfo.Modules = append(BuildInfo.Modules, modules...)
	}
	fmt.Println(BuildInfo)
}
