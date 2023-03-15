package handlers

import (
	"fmt"
	"foss_otsgenerator/internal/models"
)

type Nuget struct {
}

func (Nuget) GetPackageInfo(dependencys []models.Dependency) (modules []models.Module) {
	fmt.Println("Doing Nuget Stuff times:", len(dependencys))
	return nil
}
