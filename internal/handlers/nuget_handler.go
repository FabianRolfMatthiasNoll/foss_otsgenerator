package handlers

import (
	"fmt"
	"foss_otsgenerator/internal/api"
	"foss_otsgenerator/internal/models"
)

type Nuget struct {
}

func (Nuget) GetPackageInfo(dependencys []models.Dependency) (modules []models.Module) {
	fmt.Println("Doing Nuget Stuff times:", len(dependencys))
	var api api.NugetAPI
	api.Init()
	for _, dep := range dependencys {
		module := api.GetPackageData(dep)
		modules = append(modules, module)
	}
	return modules
}
