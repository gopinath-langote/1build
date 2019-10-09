package fixtures

import "testing"

type setup func(dir string) error
type assertion func(dir string, actualOutput string, t *testing.T) bool
type teardown func(dir string) error

//Test represents all the necessary information to run the test case
type Test struct {
	Feature   string
	Name      string
	CmdArgs   []string
	Setup     setup
	Assertion assertion
	Teardown  teardown
}

// GetFixtures returns all the fixtures to be tested
func GetFixtures() []Test {

	routes := [][]Test{
		featureRootTestData(),
		featureExecuteCmdTestData(),

		featureInitTestsData(),
		featureListTestData(),
		featureSetTestsData(),
		featureUnsetTestsData(),

		featureFlagVersionTestData(),
		featureDeleteTestData(),
	}

	var r1 []Test
	for _, r := range routes {
		r1 = append(r1, r...)
	}
	return r1
}
