package uriutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeURIsForPeople(t *testing.T) {
	typesFromNeo := []string{"Person", "Concept", "Thing"}
	expectedURIs := []string{"http://www.ft.com/ontology/person/Person"}
	actualURIs := TypeURIs(typesFromNeo)
	assert.New(t).EqualValues(expectedURIs, actualURIs)
}

func TestTypeURIsForOrganisations(t *testing.T) {
	typesFromNeo := []string{"Organisation", "Concept", "Thing"}
	expectedURIs := []string{"http://www.ft.com/ontology/organisation/Organisation"}
	actualURIs := TypeURIs(typesFromNeo)
	assert.New(t).EqualValues(expectedURIs, actualURIs)
}

func TestTypeURIsForCompany(t *testing.T) {
	typesFromNeo := []string{"Organisation", "Company", "Concept", "Thing"}
	expectedURIs := []string{"http://www.ft.com/ontology/organisation/Organisation",
		"http://www.ft.com/ontology/company/Company"}
	actualURIs := TypeURIs(typesFromNeo)
	assert.New(t).EqualValues(expectedURIs, actualURIs)
}

func TestTypeURIsForPublicCompany(t *testing.T) {
	typesFromNeo := []string{"PublicCompany", "Organisation", "Company", "Concept", "Thing"}
	expectedURIs := []string{"http://www.ft.com/ontology/company/PublicCompany",
		"http://www.ft.com/ontology/organisation/Organisation",
		"http://www.ft.com/ontology/company/Company"}
	actualURIs := TypeURIs(typesFromNeo)
	assert.New(t).EqualValues(expectedURIs, actualURIs)
}

func TestTypeURIsForPrivateCompany(t *testing.T) {
	typesFromNeo := []string{"PrivateCompany", "Organisation", "Company", "Concept", "Thing"}
	expectedURIs := []string{"http://www.ft.com/ontology/company/PrivateCompany",
		"http://www.ft.com/ontology/organisation/Organisation",
		"http://www.ft.com/ontology/company/Company"}
	actualURIs := TypeURIs(typesFromNeo)
	assert.New(t).EqualValues(expectedURIs, actualURIs)
}

func TestCompanyAPIURLs(t *testing.T) {
	apiSpecificNeoTypes := []string{"PublicCompany", "PrivateCompany", "Organisation", "Company",
		"publiccompany", "PRIVATECompany"}
	expectedAPIURLRegex := "http://api.ft.com/organisations/[w-]*"
	for _, neoType := range apiSpecificNeoTypes {
		assert.New(t).Regexp(expectedAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoType}))
	}
}

func TestPeopleAPIURLs(t *testing.T) {
	apiSpecificNeoTypes := []string{"Person", "person"}
	expectedAPIURLRegex := "http://api.ft.com/people/[w-]*"
	for _, neoType := range apiSpecificNeoTypes {
		assert.New(t).Regexp(expectedAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoType}))
	}
}

func TestDefaultAPIURLs(t *testing.T) {
	apiSpecificNeoTypes := []string{"thing", "relationship", "otherPrivateType"}
	expectedAPIURLRegex := "http://api.ft.com/things/[w-]*"
	for _, neoType := range apiSpecificNeoTypes {
		assert.New(t).Regexp(expectedAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoType}))
	}
}
