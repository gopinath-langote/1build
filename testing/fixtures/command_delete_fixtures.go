package fixtures

import (
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureDeleteTestData() []Test {
	feature := "delete"
	return []Test{
		shouldDeleteConfigFile(feature),
		shouldFailIfFileDoesntExists(feature, ""),
		shouldFailIfFileDoesntExists(feature, "--force"),
	}
}

func shouldDeleteConfigFile(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldDeleteConfigFile",
		CmdArgs: []string{"delete", "--force"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assertFileNotExists(t, dir+"/"+def.ConfigFileName)
		},
	}
}

func shouldFailIfFileDoesntExists(feature string, arg string) Test {
	expectedOutput := "No configuration file found!"
	return Test{
		Feature: feature,
		Name:    "shouldFailIfFileDoesntExists",
		CmdArgs: []string{"delete", arg},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func assertFileNotExists(t *testing.T, path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		assert.Fail(t, "Delete command did not delete config file!")
		return false
	} else if !os.IsNotExist(err) {
		assert.Fail(t, "error running os.Stat(%q): %s", path, err)
		return false
	}
	return true
}
