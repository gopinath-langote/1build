package fixtures

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func featureCompletionTestData() []Test {
	feature := "completion"
	return []Test{
		shouldPrintBashCompletion(feature),
		shouldPrintZshCompletion(feature),
		shouldPrintFishCompletion(feature),
		shouldPrintPowershellCompletion(feature),
		shouldInstallBashCompletion(feature),
		shouldFailCompletionWithUnsupportedShell(feature),
		shouldFailInstallForPowershell(feature),
	}
}

func shouldPrintBashCompletion(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintBashCompletion",
		CmdArgs: Args("completion", "bash"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "bash completion for 1build")
		},
	}
}

func shouldPrintZshCompletion(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintZshCompletion",
		CmdArgs: Args("completion", "zsh"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "zsh")
		},
	}
}

func shouldPrintFishCompletion(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintFishCompletion",
		CmdArgs: Args("completion", "fish"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "fish")
		},
	}
}

func shouldPrintPowershellCompletion(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldPrintPowershellCompletion",
		CmdArgs: Args("completion", "powershell"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "powershell")
		},
	}
}

func shouldInstallBashCompletion(feature string) Test {
	// Use a temp profile file via HOME override so we don't touch the real ~/.bash_profile.
	return Test{
		Feature: feature,
		Name:    "shouldInstallBashCompletion",
		CmdArgs: func(dir string) []string {
			return []string{"completion", "bash", "--install"}
		},
		Setup: func(dir string) error {
			// Point HOME to dir so the profile is written inside the temp dir.
			return os.Setenv("HOME", dir)
		},
		Teardown: func(dir string) error {
			return os.Unsetenv("HOME")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			profilePath := dir + "/.bash_profile"
			assert.FileExists(t, profilePath)
			content, _ := os.ReadFile(profilePath)
			return assert.Contains(t, string(content), "1build shell completion (bash)") &&
				assert.Contains(t, actualOutput, "Installed bash completion to") &&
				assert.Contains(t, actualOutput, "Restart your shell")
		},
	}
}

func shouldFailCompletionWithUnsupportedShell(feature string) Test {
	return Test{
		Feature:          feature,
		Name:             "shouldFailCompletionWithUnsupportedShell",
		CmdArgs:          Args("completion", "nushell"),
		ExpectedExitCode: 1,
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "unsupported shell")
		},
	}
}

func shouldFailInstallForPowershell(feature string) Test {
	return Test{
		Feature:          feature,
		Name:             "shouldFailInstallForPowershell",
		CmdArgs:          Args("completion", "powershell", "--install"),
		ExpectedExitCode: 1,
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "--install is not supported for powershell")
		},
	}
}
