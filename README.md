# neo-model-utils-go
Provides various packages for managing the space between Neo4j and our Models.

## Packages
### mapper
Has the following features:

1. APIURL - Establishing the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
2. IDURL - Adds the appropriate prefix e.g http://api.ft.com/things/
3. TypeURIs - Builds up the type URI based on type e.g http://www.ft.com/ontology/Person
