package mapper

import (
	"log"
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

var (
	thingLabels                  = allLabelsFor("Thing")
	conceptLabels                = allLabelsFor("Concept")
	roleLabels                   = allLabelsFor("Role")
	boardRoleLabels              = allLabelsFor("BoardRole")
	brandLabels                  = allLabelsFor("Brand")
	personLabels                 = allLabelsFor("Person")
	organisationLabels           = allLabelsFor("Organisation")
	membershipLabels             = allLabelsFor("Membership")
	companyLabels                = allLabelsFor("Company")
	publicCompanyLabels          = allLabelsFor("PublicCompany")
	privateCompanyLabels         = allLabelsFor("PrivateCompany")
	classificationLabels         = allLabelsFor("Classification")
	industryClassificationLabels = allLabelsFor("IndustryClassification")
	subjectLabels                = allLabelsFor("Subject")
	sectionLabels                = allLabelsFor("Section")
	genreLabels                  = allLabelsFor("Genre")
	topicLabels                  = allLabelsFor("Topic")
	locationLabels               = allLabelsFor("Location")
	specialReportLabels          = allLabelsFor("SpecialReport")
	alphavilleSeriesLabels       = allLabelsFor("AlphavilleSeries")
	financialInstrumentsLabels   = allLabelsFor("FinancialInstrument")

	thingURI                  = "http://www.ft.com/ontology/core/Thing"
	conceptURI                = "http://www.ft.com/ontology/concept/Concept"
	roleURI                   = "http://www.ft.com/ontology/organisation/Role"
	boardRoleURI              = "http://www.ft.com/ontology/organisation/BoardRole"
	classificationURI         = "http://www.ft.com/ontology/classification/Classification"
	industryClassificationURI = "http://www.ft.com/ontology/industry/IndustryClassification"
	personURI                 = "http://www.ft.com/ontology/person/Person"
	brandURI                  = "http://www.ft.com/ontology/product/Brand"
	organisationURI           = "http://www.ft.com/ontology/organisation/Organisation"
	membershipURI             = "http://www.ft.com/ontology/organisation/Membership"
	companyURI                = "http://www.ft.com/ontology/company/Company"
	publicCompanyURI          = "http://www.ft.com/ontology/company/PublicCompany"
	privateCompanyURI         = "http://www.ft.com/ontology/company/PrivateCompany"
	subjectURI                = "http://www.ft.com/ontology/Subject"
	sectionURI                = "http://www.ft.com/ontology/Section"
	genreURI                  = "http://www.ft.com/ontology/Genre"
	topicURI                  = "http://www.ft.com/ontology/Topic"
	locationURI               = "http://www.ft.com/ontology/Location"
	specialReportURI          = "http://www.ft.com/ontology/SpecialReport"
	alphavilleSeriesURI       = "http://www.ft.com/ontology/AlphavilleSeries"
	financialInstrumentURI    = "http://www.ft.com/ontology/FinancialInstrument"

	uuid = "92f4ec09-436d-4092-a88c-96f54e34007c"

	env     = "prod"
	envTest = "test"

	baseAPIURL         = "http://api.ft.com/"
	thingAPIURL        = baseAPIURL + "things/" + uuid
	personAPIURL       = baseAPIURL + "people/" + uuid
	organisationAPIURL = baseAPIURL + "organisations/" + uuid
	contentAPIURL      = baseAPIURL + "content/" + uuid
	brandAPIURL        = baseAPIURL + "brands/" + uuid
)

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
	assert.New(t).EqualValues([]string{thingURI, conceptURI, roleURI}, TypeURIs(roleLabels))
}

func TestTypeURIsForBoardRoles(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, roleURI, boardRoleURI}, TypeURIs(boardRoleLabels))
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
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, brandURI}, TypeURIs(brandLabels))
}

func TestTypeURIsForSubject(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, subjectURI}, TypeURIs(subjectLabels))
}

func TestTypeURIsForSection(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, sectionURI}, TypeURIs(sectionLabels))
}

func TestTypeURIsForGenre(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, genreURI}, TypeURIs(genreLabels))
}

func TestTypeURIsForTopic(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, topicURI}, TypeURIs(topicLabels))
}

func TestTypeURIsForLocation(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, locationURI}, TypeURIs(locationLabels))
}

func TestTypeURIsForSpecialReport(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, specialReportURI}, TypeURIs(specialReportLabels))
}

func TestTypeURIsForIndustryClassification(t *testing.T) {
	log.Printf("labels =%v", industryClassificationLabels)
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, industryClassificationURI}, TypeURIs(industryClassificationLabels))
}

func TestTypeURIsForAlphavilleSeries(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, alphavilleSeriesURI}, TypeURIs(alphavilleSeriesLabels))
}

func TestTypeURIForFinancialInstruments(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, financialInstrumentURI}, TypeURIs(financialInstrumentsLabels))
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
		ms, err := MostSpecificType(t.input)
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
