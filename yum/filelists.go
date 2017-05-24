package yum

import (
	"encoding/xml"

	"github.com/cavaliercoder/go-rpm"
)

const filelistsXmlns = "http://linux.duke.edu/metadata/filelists"

type Filelists struct {
	XMLName  xml.Name           `xml:"filelists"`
	Packages int                `xml:"packages"`
	Package  []FilelistsPackage `xml:"package"`
	Xmlns    string             `xml:"xmlns,attr"`
}

type FilelistsPackage struct {
	File []FilelistsFile `xml:"file"`
	Package
}

type FilelistsFile struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

func FilelistsRender(packages map[string]rpm.PackageFile) Filelists {
	filelists := Filelists{
		Packages: len(packages),
		Xmlns:    filelistsXmlns,
		Package:  []FilelistsPackage{},
	}

	for sum, p := range packages {
		pkgversion := Version{
			Epoch:   p.Epoch(),
			Version: p.Version(),
			Release: p.Release(),
		}
		pkgdata := FilelistsPackage{
			Package: Package{
				Architecture: p.Architecture(),
				Pkgid:        sum,
				Name:         p.Name(),
				Version:      pkgversion,
			},
			File: []FilelistsFile{},
		}
		for _, f := range p.Files() {
			file := FilelistsFile{
				Value: f.Name(),
			}
			if f.IsDir() {
				// TODO: The if does not quite work
				file.Type = "dir"
			}
			pkgdata.File = append(pkgdata.File, file)
		}
		filelists.Package = append(filelists.Package, pkgdata)
	}
	return filelists
}
