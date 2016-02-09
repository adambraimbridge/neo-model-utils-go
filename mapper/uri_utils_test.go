package mapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

var thingLabels = []string{"Thing"}
var conceptLabels = append(thingLabels, "Concept")
var brandLabels = append(thingLabels, "Brand")
var personLabels = append(conceptLabels, "Person")
var organisationLabels = append(conceptLabels, "Organisation")
var companyLabels = append(organisationLabels, "Company")
var publicCompanyLabels = append(companyLabels, "PublicCompany")
var privateCompanyLabels = append(companyLabels, "PrivateCompany")

var baseURI = "http://www.ft.com/ontology/"
var env = "prod"
var envTest = "test"
var personURIs = []string{baseURI + "person/Person"}
var brandURIs = []string{baseURI + "product/Brand"}
var organisationURIs = []string{baseURI + "organisation/Organisation"}
var companyURIs = append(organisationURIs, baseURI+"company/Company")
var publicCompanyURIs = append(companyURIs, baseURI+"company/PublicCompany")
var privateCompanyURIs = append(companyURIs, baseURI+"company/PrivateCompany")

var baseAPIURL = "http://api.ft.com/"
var thingAPIURLRegex = baseAPIURL + "things/[w-]*"
var personAPIURLRegex = baseAPIURL + "people/[w-]*"
var organisationAPIURLRegex = baseAPIURL + "organisations/[w-]*"
var brandAPIURLRegex = baseAPIURL + "brands/[w-]*"

func TestTypeURIsForPeople(t *testing.T) {
	assert.New(t).EqualValues(personURIs, TypeURIs(personLabels))
}

func TestTypeURIsForOrganisations(t *testing.T) {
	assert.New(t).EqualValues(organisationURIs, TypeURIs(organisationLabels))
}

func TestTypeURIsForCompany(t *testing.T) {
	assert.New(t).EqualValues(companyURIs, TypeURIs(companyLabels))
}

func TestTypeURIsForPublicCompany(t *testing.T) {
	assert.New(t).EqualValues(publicCompanyURIs, TypeURIs(publicCompanyLabels))
}

func TestTypeURIsForPrivateCompany(t *testing.T) {
	assert.New(t).EqualValues(privateCompanyURIs, TypeURIs(privateCompanyLabels))
}

func TestTypeURIsForBrand(t *testing.T) {
	assert.New(t).EqualValues(brandURIs, TypeURIs(brandLabels))
}

func TestCompanyAPIURLs(t *testing.T) {
	neoLabels := []string{"PublicCompany", "PrivateCompany", "Organisation", "Company"}
	for _, neoLabel := range neoLabels {
		assert.New(t).Regexp(organisationAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoLabel}, env))
		assert.New(t).Regexp(organisationAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{mixUpCase(neoLabel)}, env))
	}
}

func TestPeopleAPIURLs(t *testing.T) {
	neoLabels := []string{"Person"}
	for _, neoLabel := range neoLabels {
		assert.New(t).Regexp(personAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoLabel}, env))
		assert.New(t).Regexp(personAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{mixUpCase(neoLabel)}, env))
	}
}

func TestBrandAPIURLs(t *testing.T) {
	neoLabels := []string{"Brand"}
	for _, neoLabel := range neoLabels {
		assert.New(t).Regexp(brandAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoLabel}, env))
		assert.New(t).Regexp(brandAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{mixUpCase(neoLabel)}, env))
	}
}

func TestDefaultAPIURLs(t *testing.T) {
	neoLabels := []string{"thing", "relationship", "otherPrivateType"}
	for _, neoLabel := range neoLabels {
		assert.New(t).Regexp(thingAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{neoLabel}, env))
		assert.New(t).Regexp(thingAPIURLRegex, APIURL("92f4ec09-436d-4092-a88c-96f54e34007c", []string{mixUpCase(neoLabel)}, env))
	}
}

func TestInsensitivePersonTypeURIs(t *testing.T) {
	neoLabels := caseMixer(personLabels)
	assert.New(t).EqualValues(personURIs, TypeURIs(neoLabels))
}

func TestInsensitiveOrganisationTypeURIs(t *testing.T) {
	neoLabels := caseMixer(organisationLabels)
	assert.New(t).EqualValues(organisationURIs, TypeURIs(neoLabels))
}

func TestInsensitiveCompanyTypeURIs(t *testing.T) {
	neoLabels := caseMixer(companyLabels)
	assert.New(t).EqualValues(companyURIs, TypeURIs(neoLabels))
}

func TestInsensitivePublicCompanyTypeURIs(t *testing.T) {
	neoLabels := caseMixer(publicCompanyLabels)
	assert.New(t).EqualValues(publicCompanyURIs, TypeURIs(neoLabels))
}

func TestInsensitivePrivateCompanyTypeURIs(t *testing.T) {
	neoLabels := caseMixer(privateCompanyLabels)
	assert.New(t).EqualValues(privateCompanyURIs, TypeURIs(neoLabels))
}

func caseMixer(toMixUp []string) (mixedUp []string) {
	mixedUp = toMixUp
	for idx := range mixedUp {
		mixedUp[idx] = mixUpCase(mixedUp[idx])
	}
	return mixedUp
}

func mixUpCase(toMixUp string) (mixedUp string) {
	for idx, rune := range toMixUp {
		if idx%2 == 0 {
			if unicode.IsUpper(rune) {
				rune = unicode.ToLower(rune)
			} else {
				rune = unicode.ToUpper(rune)
			}
		}
		mixedUp += string(rune)
	}
	return mixedUp
}
