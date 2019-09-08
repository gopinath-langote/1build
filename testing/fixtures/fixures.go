package fixtures

import "testing"

type setup func(dir string) error
type assertion func(dir string, actualOutput string, t *testing.T) bool
type teardown func(dir string) error

type test struct {
	Feature   string
	Name      string
	CmdArgs   []string
	Setup     setup
	Assertion assertion
	Teardown  teardown
}

// GetFixtures returns all the fixtures to be tested
func GetFixtures() []test {

	routes := [][]test{
		featureRootTestData(),
		featureExecuteCmdTestData(),

		featureInitTestsData(),
		featureListTestData(),
		featureSetTestsData(),
		featureUnsetTestsData(),

		featureFlagVersionTestData(),
	}

	var r1 []test
	for _, r := range routes {
		r1 = append(r1, r...)
	}
	return r1
}
