package fixtures

import "testing"

type setup func(dir string) error
type assertion func(dir string, actualOutput string, t *testing.T) bool
type teardown func(dir string) error

type Test struct {
	Feature   string
	Name      string
	CmdArgs   []string
	Setup     setup
	Assertion assertion
	Teardown  teardown
}

func getFixtures() []Test {

	routes := [][]Test{
		featureRootTestData(),
		featureExecuteCmdTestData(),

		featureInitTestsData(),
		featureListTestData(),
		featureSetTestsData(),
		featureUnsetTestsData(),

		featureFlagVersionTestData(),
	}

	var r1 []Test
	for _, r := range routes {
		r1 = append(r1, r...)
	}
	return r1
}
