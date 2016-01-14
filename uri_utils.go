package uriutils

// ApiURL - Establishes the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
func ApiURL(id string, types []string) string {
	base := "http://api.ft.com/"
	for _, t := range types {
		switch t {
		case "Person":
			return base + "people/" + id
		case "Organisation", "Company", "PublicCompany", "PrivateCompany":
			return base + "organisations/" + id
		}
	}
	return base + "things/" + id
}

// IdURL - Adds the appropriate prefix e.g http://api.ft.com/things/
func IdURL(neoID string) string {
	return "http://api.ft.com/things/" + neoID
}

// TypeURIs - Builds up the type URI based on type e.g http://www.ft.com/ontology/Person
func TypeURIs(neoTypes []string) []string {
	var results []string
	base := "http://www.ft.com/ontology/"
	for _, t := range neoTypes {
		switch t {
		case "Person":
			results = append(results, base+"person/Person")
		case "Organisation":
			results = append(results, base+"organisation/"+t)
		case "Company", "PublicCompany", "PrivateCompany":
			results = append(results, base+"company/"+t)
		}
	}
	return results
}
