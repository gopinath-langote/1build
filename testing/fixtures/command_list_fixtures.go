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
	}
}

func shouldShowListOfCommands(feature string) Test {
	commandListMessage := `Project: Sample Project
commands:
  build: npm run build
  lint: eslint
`
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
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldShowListOfCommandsFromSpecifiedConfigFile(feature string) Test {
	commandListMessage := `Project: Sample Project
commands:
  build: npm run build
`
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
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName(feature string) Test {
	commandListMessage := `Project: Sample Project
commands:
  build: npm run build
`
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
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldNotShowAnyCommandsIfNoCommandsFound(feature string) Test {
	emptyCommandListMessage := `Project: Sample Project
commands:
`
	return Test{
		Feature: feature,
		Name:    "shouldNotShowAnyCommandsIfNoCommandsFound",
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, emptyCommandListMessage)
		},
	}
}

func shouldShowCommandsWithBeforeAndAfterIfPresent(feature string) Test {
	expectedOutput := `Project: Sample Project
beforeAll: pre_command
afterAll: post_command
commands:
  build: npm run build
`
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
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldShowListWithNestedAndInline(feature string) Test {
	expectedOutput := `Project: test
commands:
  build:
    before: echo before
    command: go build
    after: echo after
  test: go test
`
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
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

// No extra closing brace here!
