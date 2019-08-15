package fixtures

import "testing"

type Setup func(dir string) error
type Assertion func(dir string, actualOutput string, t *testing.T) bool
type Teardown func(dir string) error

type Test struct {
	Feature   string
	Name      string
	CmdArgs   []string
	Setup     Setup
	Assertion Assertion
	Teardown  Teardown
}

func GetFixtures() []Test {

	routes := [][]Test{
		FeatureRootTestData(),
		FeatureExecuteCmdTestData(),

		FeatureInitTestsData(),
		FeatureListTestData(),
		FeatureSetTestsData(),
		FeatureUnsetTestsData(),

		FeatureFlagVersionTestData(),
	}

	var r1 []Test
	for _, r := range routes {
		r1 = append(r1, r...)
	}
	return r1
}
