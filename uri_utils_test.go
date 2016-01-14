package uriutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
