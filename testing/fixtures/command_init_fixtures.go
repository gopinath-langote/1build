package fixtures

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func featureInitTestsData() []test {
	feature := "init"
	return []test{
		shouldInitialiseNewProject(feature),
		shouldFailIfFileAlreadyExists(feature),
	}
}

func shouldInitialiseNewProject(feature string) test {
	expectedOutput := `project: trial
commands:
  - build: echo 'Running build'
`
	return test{
		Feature: feature,
		Name:    "shouldInitialiseNewProject",
		CmdArgs: []string{"init", "--name", "trial"},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldFailIfFileAlreadyExists(feature string) test {
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`
	expectedOutput := "'" + def.ConfigFileName + "' configuration file already exists."
	return test{
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
