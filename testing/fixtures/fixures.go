package fixtures

import "testing"

type cmdArgs func(dir string) []string
type setup func(dir string) error
type assertion func(dir string, actualOutput string, t *testing.T) bool
type teardown func(dir string) error

//Test represents all the necessary information to run the test case
type Test struct {
	Feature   string
	Name      string
	CmdArgs   cmdArgs
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
		featureFlagTestData(),
		featureDeleteTestData(),
	}

	var r1 []Test
	for _, r := range routes {
		r1 = append(r1, r...)
	}
	return r1
}

func Args(args ...string) func(dir string) []string {
	return func(dir string) []string {
		return args
	}
}
