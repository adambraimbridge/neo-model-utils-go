package mapper

import (
	"strings"
)

var parentTypes = map[string]string{
	"thing":          "", // this is here to enable iterating over map keys to get all types
	"concept":        "thing",
	"classification": "thing", //TODO: Guy isn't sure whether this should be Concept
	"person":         "concept",
	"organisation":   "concept",
	"company":        "organisation",
	"publiccompany":  "company",
	"privatecompany": "company",
	"brand":          "concept",
	"subject":        "classification",
	"section":        "classification",
}

var apiPaths = map[string]string{
	"organisation": "organisations",
	"person":       "people",
	"brand":        "brands",
	"thing":        "things",
}

// ParentType returns the immediate parent type for a given Type
func ParentType(t string) string {
	return parentTypes[t]
}

func isDescendent(descendent, ancestor string) bool {
	for t := descendent; t != ""; t = ParentType(t) {
		if t == ancestor {
			return true
		}
	}
	return false
}

// returns the most specific from a list of types in an hierarchy
// behaviour is undefined if any of the types are siblings. // TODO: make this an error?
func mostSpecific(types []string) (result string) {
	if len(types) == 0 {
		return ""
	}
	result = types[0]
	for _, t := range types[1:] {
		if isDescendent(t, result) {
			result = t
		}
	}
	return
}

// APIURL - Establishes the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
func APIURL(uuid string, labels []string, env string) string {
	base := "http://api.ft.com/"
	if env == "test" {
		base = "http://test.api.ft.com/"
	}

	for _, label := range labels {
		switch strings.ToLower(label) {
		case "person":
			return base + "people/" + uuid
		case "organisation", "company", "publiccompany", "privatecompany":
			return base + "organisations/" + uuid
		case "brand":
			return base + "brands/" + uuid
		}
	}
	return base + "things/" + uuid
}

// IDURL - Adds the appropriate prefix e.g http://api.ft.com/things/
func IDURL(uuid string) string {
	return "http://api.ft.com/things/" + uuid
}

// TypeURIs - Builds up the type URI based on type e.g http://www.ft.com/ontology/Person
func TypeURIs(labels []string) []string {
	var results []string
	base := "http://www.ft.com/ontology/"
	for _, label := range labels {
		switch strings.ToLower(label) {
		case "person":
			results = append(results, base+"person/Person")
		case "organisation":
			results = append(results, base+"organisation/Organisation")
		case "company":
			results = append(results, base+"company/Company")
		case "publiccompany":
			results = append(results, base+"company/PublicCompany")
		case "privatecompany":
			results = append(results, base+"company/PrivateCompany")
		case "brand":
			results = append(results, base+"product/Brand")
		case "subject":
			results = append(results, base+"Subject")
		case "section":
			results = append(results, base+"Section")
		case "topic":
			results = append(results, base+"Topic")
		}
	}
	return results
}
