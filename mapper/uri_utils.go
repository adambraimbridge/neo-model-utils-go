package mapper

import (
	"strings"
)

var apiPaths = map[string]string{
	"organisation": "organisations",
	"person":       "people",
	"brand":        "brands",
	"thing":        "things",
}

var typeURIs = map[string]string{
	"person":         "http://www.ft.com/ontology/person/Person",
	"organisation":   "http://www.ft.com/ontology/organisation/Organisation",
	"company":        "http://www.ft.com/ontology/company/Company",
	"publiccompany":  "http://www.ft.com/ontology/company/PublicCompany",
	"privatecompany": "http://www.ft.com/ontology/company/PrivateCompany",
	"brand":          "http://www.ft.com/ontology/product/Brand",
	"subject":        "http://www.ft.com/ontology/Subject",
	"section":        "http://www.ft.com/ontology/Section",
	"topic":          "http://www.ft.com/ontology/Topic",
}

// APIURL - Establishes the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
func APIURL(uuid string, labels []string, env string) string {
	base := "http://api.ft.com/"
	if env == "test" {
		base = "http://test.api.ft.com/"
	}

	allLower(labels)

	path := ""
	mostSpecific, err := MostSpecific(labels)
	if err == nil {
		for t := mostSpecific; t != "" && path == ""; t = ParentType(t) {
			path = apiPaths[t]
		}
	}
	if path == "" {
		// TODO: I don't thing we should default to this, but I kept it
		// for compatability and because this function can't return an error
		path = "things"
	}
	return base + path + "/"
}

// IDURL - Adds the appropriate prefix e.g http://api.ft.com/things/
func IDURL(uuid string) string {
	return "http://api.ft.com/things/" + uuid
}

// TypeURIs - Builds up the type URI based on type e.g http://www.ft.com/ontology/Person
func TypeURIs(labels []string) []string {
	var results []string
	for _, label := range labels {
		uri := typeURIs[strings.ToLower(label)]
		if uri != "" {
			results = append(results, uri)
		}
	}
	return results
}
