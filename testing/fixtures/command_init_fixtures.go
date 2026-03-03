package fixtures

import (
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
		shouldInitialiseNewProjectWithName(feature),
	}
}

func shouldInitialiseNewProject(feature string) Test {
	expectedOutput := `project: Sample Project
commands:
  - setup:
        command: go get -u golang.org/x/lint/golint
  - test:
        command: go test ./...
  - lint:
        command: golint ./...
  - build:
        before: echo "before build"
        command: go build
        after: echo "after build"
`
	return Test{
		Feature: feature,
		Name:    "shouldInitialiseNewProject",
		CmdArgs: Args("init"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldInitialiseNewProjectWithName(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldInitialiseNewProjectWithName",
		CmdArgs: Args("init", "--name", "MyAwesomeProject"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, filePath)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), "project: MyAwesomeProject")
		},
	}
}

func shouldInitialiseNewProjectInSpecifiedFile(feature string) Test {
	expectedOutput := `project: Sample Project
commands:
  - setup:
        command: go get -u golang.org/x/lint/golint
  - test:
        command: go test ./...
  - lint:
        command: golint ./...
  - build:
        before: echo "before build"
        command: go build
        after: echo "after build"
`
	return Test{
		Feature: feature,
		Name:    "shouldInitialiseNewProjectInSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"init", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			return os.MkdirAll(dir+"/some-dir/", 0750)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/some-dir/some-config.yaml"
			assert.FileExists(t, filePath)
			content, _ := os.ReadFile(filePath)
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
	expectedOutput := "'" + def.ConfigFileName + "' already exists in the current directory."
	return Test{
		Feature: feature,
		Name:    "shouldFailIfFileAlreadyExists",
		CmdArgs: Args("init"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
