package handlers

import (
	"fmt"
	"foss_otsgenerator/internal/models"
)

type NPM struct {
}

func (NPM) GetPackageInfo(dependencys []models.Dependency) (modules []models.Module) {
	fmt.Println("Doing NPM Stuff times:", len(dependencys))
	return nil
}
