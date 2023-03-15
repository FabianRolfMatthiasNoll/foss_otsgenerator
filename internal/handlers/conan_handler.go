package handlers

import (
	"fmt"
	"foss_otsgenerator/internal/models"
)

type Conan struct {
}

func (Conan) GetPackageInfo(dependencys []models.Dependency) (modules []models.Module) {
	fmt.Println("Doing Conan Stuff times:", len(dependencys))
	return nil
}
