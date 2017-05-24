package yum

import (
	"encoding/xml"

	"github.com/cavaliercoder/go-rpm"
)

const otherXmlns = "http://linux.duke.edu/metadata/other"

type Other struct {
	XMLName  xml.Name       `xml:"otherdata"`
	Packages int            `xml:"packages"`
	Package  []OtherPackage `xml:"package"`
	Xmlns    string         `xml:"xmlns,attr"`
}

type OtherPackage struct {
	Changelog []OtherChangelog `xml:"changelog"`
	Package
}

type OtherChangelog struct {
	Author string `xml:"author,attr"`
	Date   string `xml:"date,attr"`
	Value  string `xml:",chardata"`
}

func OtherRender(packages map[string]rpm.PackageFile) Other {
	other := Other{
		Packages: len(packages),
		Xmlns:    otherXmlns,
		Package:  []OtherPackage{},
	}

	for sum, p := range packages {
		pkgversion := Version{
			Epoch:   p.Epoch(),
			Version: p.Version(),
			Release: p.Release(),
		}
		pkgdata := OtherPackage{
			Package: Package{
				Architecture: p.Architecture(),
				Pkgid:        sum,
				Name:         p.Name(),
				Version:      pkgversion,
			},
			Changelog: []OtherChangelog{},
		}
		for _, c := range p.ChangeLog() {
			// TODO: The c string needs to be tokenized do date, author and changelog text
			cl := OtherChangelog{
				Value:  c,
				Author: "Author",
				Date:   "1493812800",
			}
			pkgdata.Changelog = append(pkgdata.Changelog, cl)
		}
		other.Package = append(other.Package, pkgdata)
	}
	return other
}