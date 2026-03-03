package fixtures

import (
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureListTestData() []Test {
	var feature = "list"
	return []Test{
		shouldShowListOfCommands(feature),
		shouldShowListOfCommandsFromSpecifiedConfigFile(feature),
		shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName(feature),
		shouldNotShowAnyCommandsIfNoCommandsFound(feature),
		shouldShowCommandsWithBeforeAndAfterIfPresent(feature),
		shouldShowListWithNestedAndInline(feature),
		shouldShowListAsJSON(feature),
	}
}

func shouldShowListOfCommands(feature string) Test {
	defaultFileContent := `project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`

	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommands",
		CmdArgs: Args("list"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: Sample Project") &&
				assert.Contains(t, actualOutput, "build | npm run build") &&
				assert.Contains(t, actualOutput, "lint | eslint")
		},
	}
}

func shouldShowListOfCommandsFromSpecifiedConfigFile(feature string) Test {
	defaultFileContent := `project: Sample Project
commands:
  - build: npm run build
`
	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommandsFromSpecifiedConfigFile",
		CmdArgs: func(dir string) []string {
			return []string{"list", "-f", dir + "/custom-directory/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/custom-directory", 0750)
			return utils.CreateConfigFileWithName(dir+"/custom-directory", "some-config.yaml", defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: Sample Project") &&
				assert.Contains(t, actualOutput, "build | npm run build")
		},
	}
}

func shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName(feature string) Test {
	defaultFileContent := `project: Sample Project
commands:
  - build: npm run build
`
	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName",
		CmdArgs: func(dir string) []string {
			return []string{"list", "--file", dir + "/custom-directory/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/custom-directory", 0750)
			return utils.CreateConfigFileWithName(dir+"/custom-directory", "some-config.yaml", defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: Sample Project") &&
				assert.Contains(t, actualOutput, "build | npm run build")
		},
	}
}

func shouldNotShowAnyCommandsIfNoCommandsFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldNotShowAnyCommandsIfNoCommandsFound",
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: Sample Project") &&
				assert.Contains(t, actualOutput, "commands:")
		},
	}
}

func shouldShowCommandsWithBeforeAndAfterIfPresent(feature string) Test {
	fileContent := `project: Sample Project
beforeAll: pre_command
afterAll: post_command
commands:
  - build: npm run build
`

	return Test{
		Feature: feature,
		Name:    "shouldShowCommandsWithBeforeAndAfterIfPresent",
		CmdArgs: Args("list"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: Sample Project") &&
				assert.Contains(t, actualOutput, "beforeAll: pre_command") &&
				assert.Contains(t, actualOutput, "afterAll: post_command") &&
				assert.Contains(t, actualOutput, "build | npm run build")
		},
	}
}

func shouldShowListWithNestedAndInline(feature string) Test {
	fileContent := `project: test
commands:
  - build:
      before: echo before
      command: go build
      after: echo after
  - test: go test
`
	return Test{
		Feature: feature,
		Name:    "list with nested and inline",
		CmdArgs: Args("list"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: test") &&
				assert.Contains(t, actualOutput, "build | go build") &&
				assert.Contains(t, actualOutput, "test | go test")
		},
	}
}

// No extra closing brace here!

func shouldShowListAsJSON(feature string) Test {
	defaultFileContent := `project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`

	return Test{
		Feature: feature,
		Name:    "shouldShowListAsJSON",
		CmdArgs: Args("list", "--output", "json"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, `"project": "Sample Project"`) &&
				assert.Contains(t, actualOutput, `"name": "build"`) &&
				assert.Contains(t, actualOutput, `"name": "lint"`)
		},
	}
}
