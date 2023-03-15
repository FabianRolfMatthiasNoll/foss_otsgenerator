package api

import (
	"encoding/json"
	"fmt"
	"foss_otsgenerator/internal/models"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

type Package struct {
	Authors           string    `json:"authors"`
	Copyright         string    `json:"copyright"`
	Created           time.Time `json:"created"`
	Description       string    `json:"description"`
	Name              string    `json:"id"`
	LicenseExpression string    `json:"licenseExpression"`
	LicenseURL        string    `json:"licenseUrl"`
	ProjectURL        string    `json:"projectUrl"`
	Published         time.Time `json:"published"`
	Version           string    `json:"version"`
}

type MetaEntry struct {
	ID           string    `json:"@id"`
	Type         []string  `json:"@type"`
	CatalogEntry string    `json:"catalogEntry"`
	Published    time.Time `json:"published"`
}

type NugetAPI struct {
	Version   string `json:"version"`
	Resources []struct {
		ID            string `json:"@id"`
		Type          string `json:"@type"`
		Comment       string `json:"comment,omitempty"`
		ClientVersion string `json:"clientVersion,omitempty"`
	} `json:"resources"`
	BasicAPI string
}

func (nuget *NugetAPI) Init() {
	entryData := nuget.GetDataFromApi("https://api.nuget.org/v3/index.json")
	err := json.Unmarshal(entryData, &nuget)
	if err != nil {
		fmt.Println("[", color.Colorize(color.Red, "Err"), "] ", err)
	}
	for _, res := range nuget.Resources {
		if res.Type == "RegistrationsBaseUrl" {
			nuget.BasicAPI = res.ID
		}
	}
}

func (nuget *NugetAPI) GetPackageData(dependency models.Dependency) models.Module {
	pkgId := dependency.ImportName
	pkgVersion := dependency.Version
	var meta MetaEntry
	var pkg Package
	baseUrl := fmt.Sprintf("%s%s/%s.json", nuget.BasicAPI, strings.ToLower(pkgId), pkgVersion)
	metaEntryData := nuget.GetDataFromApi(baseUrl)
	err := json.Unmarshal(metaEntryData, &meta)
	if err != nil {
		fmt.Println("[", color.Colorize(color.Red, "Err"), "] ", err)
	}
	pkgUrl := meta.CatalogEntry
	pkgData := nuget.GetDataFromApi(pkgUrl)
	err = json.Unmarshal(pkgData, &pkg)
	if err != nil {
		fmt.Println("[", color.Colorize(color.Red, "Err"), "] ", err)
	}
	return nuget.CreateModule(pkg, dependency)
}

func (NugetAPI) CreateModule(pkg Package, dep models.Dependency) models.Module {
	if !(pkg.ProjectURL != "" && pkg.ProjectURL != "https://dot.net/") {
		pkg.ProjectURL = fmt.Sprintf("https://www.nuget.org/packages/%s", pkg.Name)
	}

	//Microsofts package descriptions start with a summary and then have a ton of
	//unimportant information. Thankfully Microsoft has a \n after the summary.
	if strings.Contains(pkg.Description, "\n") {
		pkg.Description = pkg.Description[:strings.Index(pkg.Description, "\n")]
	}

	pkgInfo := models.PackageInfo{
		Author:      pkg.Authors,
		Description: pkg.Description,
		License:     pkg.LicenseExpression,
		Release:     pkg.Published,
	}

	return models.Module{
		Name:     dep.ImportName,
		Source:   pkg.ProjectURL,
		Version:  dep.Version,
		Hash:     dep.ID,
		TopLevel: dep.TopLevel,
		Info:     pkgInfo,
	}
}

func (NugetAPI) GetDataFromApi(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return nil
	}
	return body
}
