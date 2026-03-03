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
		shouldDeleteConfigSpecifiedFile(feature),
		shouldFailIfFileDoesntExists(feature),
		shouldFailIfFileDoesntExistsWithForceFlag(feature),
		shouldDryRunDelete(feature),
	}
}

func shouldDeleteConfigFile(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldDeleteConfigFile",
		CmdArgs: Args("delete", "--force"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assertFileNotExists(t, dir+"/"+def.ConfigFileName)
		},
	}
}

func shouldDeleteConfigSpecifiedFile(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldDeleteConfigSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"delete", "-f", dir + "/custom-directory/some-file.yaml", "--force"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/custom-directory", 0750)
			return utils.CreateConfigFileWithName(
				dir+"/custom-directory", "some-file.yaml", "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assertFileNotExists(t, dir+"/custom-directory/some-file.yaml")
		},
	}
}

func shouldFailIfFileDoesntExists(feature string) Test {
	expectedOutput := "No configuration file found!"
	return Test{
		Feature:          feature,
		Name:             "shouldFailIfFileDoesntExists",
		CmdArgs:          Args("delete"),
		ExpectedExitCode: 1,
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldFailIfFileDoesntExistsWithForceFlag(feature string) Test {
	expectedOutput := "No configuration file found!"
	return Test{
		Feature:          feature,
		Name:             "shouldFailIfFileDoesntExistsWithForceFlag",
		CmdArgs:          Args("delete", "--force"),
		ExpectedExitCode: 1,
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

func shouldDryRunDelete(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldDryRunDelete",
		CmdArgs: Args("delete", "--dry-run"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			// Output should mention dry-run
			if !assert.Contains(t, actualOutput, "[dry-run]") {
				return false
			}
			// Config file must still exist
			_, err := os.Stat(dir + "/" + def.ConfigFileName)
			return assert.NoError(t, err, "config file should still exist after dry-run")
		},
	}
}
