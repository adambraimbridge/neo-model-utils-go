package mapper

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

var thingLabels = []string{"Thing"}
var conceptLabels = append(thingLabels, "Concept")
var roleLabels = append(thingLabels, "Role")
var boardRoleLabels = append(roleLabels, "BoardRole")
var brandLabels = append(conceptLabels, "Brand")
var personLabels = append(conceptLabels, "Person")
var organisationLabels = append(conceptLabels, "Organisation")
var membershipLabels = append(conceptLabels, "Membership")
var companyLabels = append(organisationLabels, "Company")
var publicCompanyLabels = append(companyLabels, "PublicCompany")
var privateCompanyLabels = append(companyLabels, "PrivateCompany")
var classificationLabels = append(conceptLabels, "Classification")
var subjectLabels = append(classificationLabels, "Subject")
var sectionLabels = append(classificationLabels, "Section")
var genreLabels = append(classificationLabels, "Genre")
var topicLabels = append(conceptLabels, "Topic")
var locationLabels = append(conceptLabels, "Location")

var baseURI = "http://www.ft.com/ontology/"
var env = "prod"
var envTest = "test"
var thingURIs = []string{baseURI + "core/Thing"}
var conceptURIs = append(thingURIs, baseURI+"concept/Concept")
var roleURIs = append(thingURIs, baseURI+"organisation/Role")
var boardRoleURIs = append(roleURIs, baseURI+"organisation/BoardRole")
var classificationURIs = append(conceptURIs, baseURI+"classification/Classification")
var personURIs = append(conceptURIs, baseURI+"person/Person")
var brandURIs = append(conceptURIs, baseURI+"product/Brand")
var organisationURIs = append(conceptURIs, baseURI+"organisation/Organisation")
var membershipURIs = append(conceptURIs, baseURI+"organisation/Membership")
var companyURIs = append(organisationURIs, baseURI+"company/Company")
var publicCompanyURIs = append(companyURIs, baseURI+"company/PublicCompany")
var privateCompanyURIs = append(companyURIs, baseURI+"company/PrivateCompany")
var subjectURIs = append(classificationURIs, baseURI+"Subject")
var sectionURIs = append(classificationURIs, baseURI+"Section")
var genreURIs = append(classificationURIs, baseURI+"Genre")
var topicURIs = append(conceptURIs, baseURI+"Topic")
var locationURIs = append(conceptURIs, baseURI+"Location")

var uuid = "92f4ec09-436d-4092-a88c-96f54e34007c"

var baseAPIURL = "http://api.ft.com/"
var thingAPIURL = baseAPIURL + "things/" + uuid
var personAPIURL = baseAPIURL + "people/" + uuid
var organisationAPIURL = baseAPIURL + "organisations/" + uuid
var contentAPIURL = baseAPIURL + "content/" + uuid
var brandAPIURL = baseAPIURL + "brands/" + uuid

func TestTypeURIsForPeople(t *testing.T) {
	assert.New(t).EqualValues(personURIs, TypeURIs(personLabels))
}

func TestTypeURIsForOrganisations(t *testing.T) {
	assert.New(t).EqualValues(organisationURIs, TypeURIs(organisationLabels))
}

func TestTypeURIsForMemberships(t *testing.T) {
	assert.New(t).EqualValues(membershipURIs, TypeURIs(membershipLabels))
}

func TestTypeURIsForRoles(t *testing.T) {
	assert.New(t).EqualValues(roleURIs, TypeURIs(roleLabels))
}

func TestTypeURIsForBoardRoles(t *testing.T) {
	assert.New(t).EqualValues(boardRoleURIs, TypeURIs(boardRoleLabels))
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

func TestTypeURIsForSubject(t *testing.T) {
	assert.New(t).EqualValues(subjectURIs, TypeURIs(subjectLabels))
}

func TestTypeURIsForSection(t *testing.T) {
	assert.New(t).EqualValues(sectionURIs, TypeURIs(sectionLabels))
}

func TestTypeURIsForGenre(t *testing.T) {
	assert.New(t).EqualValues(genreURIs, TypeURIs(genreLabels))
}

func TestTypeURIsForTopic(t *testing.T) {
	assert.New(t).EqualValues(topicURIs, TypeURIs(topicLabels))
}

func TestTypeURIsForLocation(t *testing.T) {
	assert.New(t).EqualValues(locationURIs, TypeURIs(locationLabels))
}

func TestContentAPIURLs(t *testing.T) {
	neoLabels := []string{"Content"}
	assert.New(t).EqualValues(contentAPIURL, APIURL(uuid, neoLabels, env))
	assert.New(t).EqualValues(contentAPIURL, APIURL(uuid, caseMixer(neoLabels), env))
}

func TestCompanyAPIURLs(t *testing.T) {
	neoLabels := []string{"PublicCompany", "Organisation", "Company"}
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, neoLabels, env))
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, caseMixer(neoLabels), env))
}

func TestOrganisationAPIURLs(t *testing.T) {
	neoLabels := []string{"Organisation", "Concept", "Thing"}
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, neoLabels, env))
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, caseMixer(neoLabels), env))
}

func TestPeopleAPIURLs(t *testing.T) {
	neoLabels := []string{"Person"}
	assert.New(t).EqualValues(personAPIURL, APIURL(uuid, neoLabels, env))
	assert.New(t).EqualValues(personAPIURL, APIURL(uuid, caseMixer(neoLabels), env))

}

func TestBrandAPIURLs(t *testing.T) {
	neoLabels := []string{"Brand"}
	assert.New(t).EqualValues(brandAPIURL, APIURL(uuid, neoLabels, env))
	assert.New(t).EqualValues(brandAPIURL, APIURL(uuid, caseMixer(neoLabels), env))

}

func TestDefaultAPIURLs(t *testing.T) {
	neoLabels := []string{"thing", "relationship", "otherPrivateType"}
	assert.New(t).EqualValues(thingAPIURL, APIURL(uuid, neoLabels, env))
	assert.New(t).EqualValues(thingAPIURL, APIURL(uuid, caseMixer(neoLabels), env))

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

func TestMostSpecific(t *testing.T) {
	assert := assert.New(t)

	for _, t := range []struct {
		input    []string
		expected string
		err      error
	}{
		{
			[]string{"Organisation"},
			"Organisation",
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "Company"},
			"PublicCompany",
			nil,
		}, {
			[]string{"PublicCompany", "Organisation"},
			"PublicCompany",
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "PrivateCompany", "Company"},
			"",
			ErrNotHierarchy,
		}, {
			[]string{"zzzzzz", "yyyyyy"},
			"",
			ErrNotHierarchy,
		},
	} {
		ms, err := mostSpecific(t.input)
		assert.Equal(t.expected, ms)
		assert.Equal(t.err, err)
	}
}

func TestTypeSorter(t *testing.T) {
	assert := assert.New(t)

	for _, t := range []struct {
		input    []string
		expected []string
		err      error
	}{
		{
			[]string{"Organisation"},
			[]string{"Organisation"},
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "Company"},
			[]string{"Organisation", "Company", "PublicCompany"},
			nil,
		}, {
			[]string{"PublicCompany", "Organisation"},
			[]string{"Organisation", "PublicCompany"},
			nil,
		}, {
			[]string{"Content", "Thing"},
			[]string{"Thing", "Content"},
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "PrivateCompany", "Company"},
			[]string{"Organisation", "PublicCompany", "PrivateCompany", "Company"},
			ErrNotHierarchy,
		}, {
			[]string{"zzzzzz", "yyyyyy"},
			[]string{"zzzzzz", "yyyyyy"},
			ErrNotHierarchy,
		},
	} {
		sorted, err := SortTypes(t.input)
		assert.Equal(t.expected, sorted)
		assert.Equal(t.err, err)
	}
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
