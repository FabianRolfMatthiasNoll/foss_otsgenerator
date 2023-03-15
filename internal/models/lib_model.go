package models

import (
	"fmt"
	"strings"
	"time"
)

type BuildInfo struct {
	Name      string
	Languages []string
	Modules   []Module
}

type Module struct {
	Name     string
	Source   string //if there is one fetchable fetch it else construct it
	Version  string
	Hash     string //ID
	Parents  []string
	TopLevel bool
	Info     PackageInfo
}

type PackageInfo struct {
	Author      string
	Description string
	License     string
	Release     time.Time
}

func (bi BuildInfo) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Build Info:\n")
	fmt.Fprintf(&b, "  Name: %s\n", bi.Name)
	fmt.Fprintf(&b, "  Languages: %s\n", strings.Join(bi.Languages, ", "))

	if len(bi.Modules) > 0 {
		fmt.Fprintf(&b, "  Modules:\n")
		for _, module := range bi.Modules {
			b.WriteString(module.StringIndented(2))
		}
	}

	return b.String()
}

func (m Module) StringIndented(indent int) string {
	var b strings.Builder

	indentStr := strings.Repeat(" ", indent)

	fmt.Fprintf(&b, "%s- Module:\n", indentStr)
	fmt.Fprintf(&b, "%s  Name: %s\n", indentStr, m.Name)
	fmt.Fprintf(&b, "%s  Source: %s\n", indentStr, m.Source)
	fmt.Fprintf(&b, "%s  Version: %s\n", indentStr, m.Version)
	fmt.Fprintf(&b, "%s  Hash: %s\n", indentStr, m.Hash)
	fmt.Fprintf(&b, "%s  Parents: %s\n", indentStr, strings.Join(m.Parents, ", "))
	fmt.Fprintf(&b, "%s  TopLevel: %v\n", indentStr, m.TopLevel)
	b.WriteString(m.Info.StringIndented(indent + 2))

	return b.String()
}

func (pi PackageInfo) StringIndented(indent int) string {
	var b strings.Builder

	indentStr := strings.Repeat(" ", indent)

	fmt.Fprintf(&b, "%s- Package Info:\n", indentStr)
	fmt.Fprintf(&b, "%s  Author: %s\n", indentStr, pi.Author)
	fmt.Fprintf(&b, "%s  Description: %s\n", indentStr, pi.Description)
	fmt.Fprintf(&b, "%s  License: %s\n", indentStr, pi.License)
	fmt.Fprintf(&b, "%s  Release: %s\n", indentStr, pi.Release.String())

	return b.String()
}
