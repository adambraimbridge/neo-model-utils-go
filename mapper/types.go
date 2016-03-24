package mapper

import (
	"errors"
	"sort"
	"strings"
)

var parentTypes = map[string]string{
	"thing":                  "", // this is here to enable iterating over map keys to get all types
	"concept":                "thing",
	"role":                   "thing",
	"boardrole":              "role",
	"classification":         "concept",
	"person":                 "concept",
	"organisation":           "concept",
	"membership":             "concept",
	"company":                "organisation",
	"publiccompany":          "company",
	"privatecompany":         "company",
	"brand":                  "concept",
	"subject":                "classification",
	"section":                "classification",
	"genre":                  "classification",
	"location":               "concept",
	"topic":                  "concept",
	"industryclassification": "classification",
	"content":                "thing",
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

// MostSpecific returns the most specific from a list of types in an hierarchy
// behaviour is undefined if any of the types are siblings.
func mostSpecific(types []string) (string, error) {
	if len(types) == 0 {
		return "", errors.New("no types supplied")
	}
	sorted, err := SortTypes(types)
	if err != nil {
		return "", err
	}
	return sorted[len(sorted)-1], nil
}

var ErrNotHierarchy = errors.New("provided types are not a consistent hierarchy")

// SortTypes sorts the given types from least specific to most specific
func SortTypes(types []string) ([]string, error) {
	ts := &typeSorter{types: make([]string, len(types))}
	copy(ts.types, types)
	sort.Sort(ts)
	if ts.invalid {
		return types, ErrNotHierarchy
	}
	return ts.types, nil
}

type typeSorter struct {
	types   []string
	invalid bool
}

func (ts *typeSorter) Len() int {
	return len(ts.types)
}

func (ts *typeSorter) Less(a, b int) bool {
	at := strings.ToLower(ts.types[a])
	bt := strings.ToLower(ts.types[b])
	if isDescendent(bt, at) {
		return true
	}
	if !isDescendent(at, bt) {
		ts.invalid = true
	}
	return false
}

func (ts *typeSorter) Swap(a, b int) {
	ts.types[a], ts.types[b] = ts.types[b], ts.types[a]
}

func allLower(all []string) {
	for i, s := range all {
		all[i] = strings.ToLower(s)
	}
}
