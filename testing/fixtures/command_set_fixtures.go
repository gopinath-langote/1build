package fixtures

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func featureSetTestsData() []Test {
	feature := "set"

	return []Test{
		shouldSetNewCommand(feature),
		shouldSetNewCommandInSpecifiedFile(feature),
		shouldUpdateExistingCommand(feature),
		shouldFailWhenConfigurationFileIsNotFound(feature),
		shouldFailWhenConfigurationFileIsInInvalidFormat(feature),
		shouldSetBeforeCommand(feature),
		shouldSetAfterCommand(feature),
	}
}

func shouldSetNewCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build: go build
  - Test: go Test
`

	return Test{
		Feature: feature,
		Name:    "shouldSetNewCommand",
		CmdArgs: Args("set", "Test", "go Test"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldSetNewCommandInSpecifiedFile(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build: go build
  - Test: go Test
`

	return Test{
		Feature: feature,
		Name:    "shouldSetNewCommandInSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"set", "Test", "go Test", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/some-dir", 0777)
			return utils.CreateConfigFileWithName(dir+"/some-dir", "some-config.yaml", defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/some-dir/some-config.yaml"
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldUpdateExistingCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build: go build -o
`

	return Test{
		Feature: feature,
		Name:    "shouldUpdateExistingCommand",
		CmdArgs: Args("set", "build", "go build -o"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldFailWhenConfigurationFileIsNotFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailWhenConfigurationFileIsNotFound",
		CmdArgs: Args("set", "build", "go build -o"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found")
		},
	}
}

func shouldFailWhenConfigurationFileIsInInvalidFormat(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailWhenConfigurationFileIsInInvalidFormat",
		CmdArgs: Args("set", "build", "go build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid config content")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Unable to parse '"+def.ConfigFileName+"' config file. Make sure file is in correct format.")
		},
	}
}

func shouldSetBeforeCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
after: yo
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "shouldSetBeforeCommand",
		CmdArgs: Args("set", "after", "yo"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldSetAfterCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
after: yo
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "shouldSetBeforeCommand",
		CmdArgs: Args("set", "after", "yo"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}
