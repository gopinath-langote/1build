package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func featureFlagVersionTestData() []Test {
	feature := "version"
	return []Test{
		shouldPrintCurrentVersion(feature),
		shouldPrintCurrentVersionWithShortFlag(feature),
	}
}

func shouldPrintCurrentVersion(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintCurrentVersion",
		CmdArgs: Args("--version"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "1build version test-version")
		},
	}
}

func shouldPrintCurrentVersionWithShortFlag(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintCurrentVersionWithShortFlag",
		CmdArgs: Args("-v"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "1build version test-version")
		},
	}
}
