package mapper

import (
	"strings"
)

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
