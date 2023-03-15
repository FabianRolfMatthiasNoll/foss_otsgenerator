
# FOSSer Tool: foss_otsgenerator

Part of the FOSSer CLI Tool Series to generate full documentation of the SBOM  

This part is responsible to take a sbom in the fosser tool format and fill it for every package with as much metadata as it can, as well as filtering out unimportant libraries



## Features

- Support for Nuget, Npm and Go package managers
- Fetching MetaData like summaries, licenses, author, release date etc.
- Outputs a file with fully mapped libs that can be inserted into a foss document

## Future Updates

- it is planned to support more languages like c++ and generally docker
## Usage/Examples

After building the Go Application it can be used like the following
```golang
fosser_otsgenerator [Path/To/SBOM] [Path/To/Config] [Path/For/Output]
```

The Tool outputs a foss.yml file with the following structure:
```golang
type BuildInfo struct {
	Name      string
	Languages []string
	Modules   []Module
}

type Module struct {
	Name     string
	Source   string
	Version  string
	Hash     string //ID
	Parents  []string
	TopLevel bool
	Info     RepoInfo
}

type RepoInfo struct {
	Author      string
	Description string
	License     string
	Release     time.Time
}

```
`Name` The from the source extracted library Name  
`Source` A generated link to the official package site of that library  
`Version` The exact Version of the package that is used   
`Hash` Hash value of the package  
`Parents` A list of all Parent Dependencies. This list is needed for the foss document
`TopLevel` Important for npm and Docker  

`Author` The official creator from the library  
`Description` A short description for what the library does and what its purpous is  
`Licenses` The License this library uses  
`Release` The Date this library with the specified version was released

 




## Folder Structure

```bash
foss_toolconverter
 ┣ cmd
 ┃ ┗ rootCmd.go
 ┣ internal
 ┃ ┣ models
 ┃ ┃ ┗ dependency.go
 ┃ ┣ manager.go
 ┃ ┣ packageJson.go
 ┃ ┗ syft_convert.go
 ┣ .gitignore
 ┣ go.mod
 ┣ go.sum
 ┗ main.go
```
