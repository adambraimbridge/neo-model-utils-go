package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func allLabelsFor(label string) []string {
	allLabels := []string{}
	for t := label; t != ""; t = ParentType(t) {
		allLabels = append(allLabels, t)
	}
	return allLabels
}

var thingLabels = allLabelsFor("Thing")
var conceptLabels = allLabelsFor("Concept")
var roleLabels = allLabelsFor("Role")
var boardRoleLabels = allLabelsFor("BoardRole")
var brandLabels = allLabelsFor("Brand")
var personLabels = allLabelsFor("Person")
var organisationLabels = allLabelsFor("Organisation")
var membershipLabels = allLabelsFor("Membership")
var companyLabels = allLabelsFor("Company")
var publicCompanyLabels = allLabelsFor("PublicCompany")
var privateCompanyLabels = allLabelsFor("PrivateCompany")
var classificationLabels = allLabelsFor("Classification")
var subjectLabels = allLabelsFor("Subject")
var sectionLabels = allLabelsFor("Section")
var topicLabels = allLabelsFor("Topic")

var thingURI = "http://www.ft.com/ontology/core/Thing"
var conceptURI = "http://www.ft.com/ontology/concept/Concept"
var roleURI = "http://www.ft.com/ontology/organisation/Role"
var boardRoleURI = "http://www.ft.com/ontology/organisation/BoardRole"
var classificationURI = "http://www.ft.com/ontology/classification/Classification"
var personURI = "http://www.ft.com/ontology/person/Person"
var brandURI = "http://www.ft.com/ontology/product/Brand"
var organisationURI = "http://www.ft.com/ontology/organisation/Organisation"
var membershipURI = "http://www.ft.com/ontology/organisation/Membership"
var companyURI = "http://www.ft.com/ontology/company/Company"
var publicCompanyURI = "http://www.ft.com/ontology/company/PublicCompany"
var privateCompanyURI = "http://www.ft.com/ontology/company/PrivateCompany"
var subjectURI = "http://www.ft.com/ontology/Subject"
var sectionURI = "http://www.ft.com/ontology/Section"
var topicURI = "http://www.ft.com/ontology/Topic"

var uuid = "92f4ec09-436d-4092-a88c-96f54e34007c"

var env = "prod"
var envTest = "test"

var baseAPIURL = "http://api.ft.com/"
var thingAPIURL = baseAPIURL + "things/" + uuid
var personAPIURL = baseAPIURL + "people/" + uuid
var organisationAPIURL = baseAPIURL + "organisations/" + uuid
var contentAPIURL = baseAPIURL + "content/" + uuid
var brandAPIURL = baseAPIURL + "brands/" + uuid

func TestTypeURIsForPeople(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, personURI}, TypeURIs(personLabels))
}

func TestTypeURIsForOrganisations(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI}, TypeURIs(organisationLabels))
}

func TestTypeURIsForMemberships(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, membershipURI}, TypeURIs(membershipLabels))
}

func TestTypeURIsForRoles(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, roleURI}, TypeURIs(roleLabels))
}

func TestTypeURIsForBoardRoles(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, roleURI, boardRoleURI}, TypeURIs(boardRoleLabels))
}

func TestTypeURIsForCompany(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI, companyURI}, TypeURIs(companyLabels))
}

func TestTypeURIsForPublicCompany(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI, companyURI, publicCompanyURI}, TypeURIs(publicCompanyLabels))
}

func TestTypeURIsForPrivateCompany(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI, companyURI, privateCompanyURI}, TypeURIs(privateCompanyLabels))
}

func TestTypeURIsForBrand(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, brandURI}, TypeURIs(brandLabels))
}

func TestTypeURIsForSubject(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, subjectURI}, TypeURIs(subjectLabels))
}

func TestTypeURIsForSection(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, sectionURI}, TypeURIs(sectionLabels))
}

func TestTypeURIsForTopic(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, topicURI}, TypeURIs(topicLabels))
}

func TestContentAPIURLs(t *testing.T) {
	neoLabels := []string{"Content"}
	assert.New(t).EqualValues(contentAPIURL, APIURL(uuid, neoLabels, env))
}

func TestCompanyAPIURLs(t *testing.T) {
	neoLabels := []string{"PublicCompany", "Organisation", "Company"}
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, neoLabels, env))
}

func TestOrganisationAPIURLs(t *testing.T) {
	neoLabels := []string{"Organisation", "Concept", "Thing"}
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, neoLabels, env))
}

func TestPeopleAPIURLs(t *testing.T) {
	neoLabels := []string{"Person"}
	assert.New(t).EqualValues(personAPIURL, APIURL(uuid, neoLabels, env))

}

func TestBrandAPIURLs(t *testing.T) {
	neoLabels := []string{"Brand"}
	assert.New(t).EqualValues(brandAPIURL, APIURL(uuid, neoLabels, env))

}

func TestDefaultAPIURLs(t *testing.T) {
	neoLabels := []string{"thing", "relationship", "otherPrivateType"}
	assert.New(t).EqualValues(thingAPIURL, APIURL(uuid, neoLabels, env))

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
