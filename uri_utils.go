package uriutils

import (
	"strings"
)

// APIURL - Establishes the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
func APIURL(id string, types []string) string {
	base := "http://api.ft.com/"
	for _, t := range types {
		switch strings.ToLower(t) {
		case "person":
			return base + "people/" + id
		case "organisation", "company", "publiccompany", "privatecompany":
			return base + "organisations/" + id
		}
	}
	return base + "things/" + id
}

// IDURL - Adds the appropriate prefix e.g http://api.ft.com/things/
func IDURL(neoID string) string {
	return "http://api.ft.com/things/" + neoID
}

// TypeURIs - Builds up the type URI based on type e.g http://www.ft.com/ontology/Person
func TypeURIs(neoTypes []string) []string {
	var results []string
	base := "http://www.ft.com/ontology/"
	for _, t := range neoTypes {
		switch strings.ToLower(t) {
		case "person":
			results = append(results, base+"person/Person")
		case "organisation":
			results = append(results, base+"organisation/"+t)
		case "company", "publiccompany", "privatecompany":
			results = append(results, base+"company/"+t)
		}
	}
	return results
}
