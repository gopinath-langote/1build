package fixtures

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureInitTestsData() []Test {
	feature := "init"
	return []Test{
		shouldInitialiseNewProject(feature),
		shouldInitialiseNewProjectInSpecifiedFile(feature),
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
		CmdArgs: Args("init", "--name", "trial"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldInitialiseNewProjectInSpecifiedFile(feature string) Test {
	expectedOutput := `project: trial
commands:
  - build: echo 'Running build'
`
	return Test{
		Feature: feature,
		Name:    "shouldInitialiseNewProjectInSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"init", "--name", "trial", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			return os.MkdirAll(dir+"/some-dir/", 0750)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/some-dir/some-config.yaml"
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
		CmdArgs: Args("init", "--name", "trial"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
