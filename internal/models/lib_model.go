package models

import "time"

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
