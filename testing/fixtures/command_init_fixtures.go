package fixtures

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureInitTestsData() []Test {
	feature := "init"
	return []Test{
		shouldInitialiseNewProject(feature),
		shouldFailIfFileAlreadyExists(feature),
	}
}

func shouldInitialiseNewProject(feature string) Test {
	expectedOutput := `project: trial
commands:
  - build: echo 'Running build'
`
	return Test{
		Feature: feature,
		Name:    "shouldInitialiseNewProject",
		CmdArgs: []string{"init", "--name", "trial"},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldFailIfFileAlreadyExists(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`
	expectedOutput := "'" + def.ConfigFileName + "' configuration file already exists."
	return Test{
		Feature: feature,
		Name:    "shouldFailIfFileAlreadyExists",
		CmdArgs: []string{"init", "--name", "trial"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
