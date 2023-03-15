package handlers

import (
	"fmt"
	"foss_otsgenerator/internal/models"
)

type Go struct {
}

func (Go) GetPackageInfo(dependencys []models.Dependency) (modules []models.Module) {
	fmt.Println("Doing Golang Stuff times:", len(dependencys))
	return nil
}
