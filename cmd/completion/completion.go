package completion

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// profileFiles maps each supported shell to its default user profile file.
var profileFiles = map[string]string{
	"bash": ".bash_profile",
	"zsh":  ".zshrc",
	"fish": ".config/fish/config.fish",
}

// Cmd is the cobra command for generating and optionally installing shell completion scripts.
var Cmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion script",
	Long: `Generate a shell completion script for 1build.

Without --install, the script is printed to stdout so you can source it manually:

  source <(1build completion bash)
  source <(1build completion zsh)

With --install, the script is appended to your shell profile automatically:

  1build completion bash --install    # appends to ~/.bash_profile
  1build completion zsh  --install    # appends to ~/.zshrc
  1build completion fish --install    # appends to ~/.config/fish/config.fish

Note: powershell does not support --install.`,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		install, _ := cmd.Flags().GetBool("install")
		shell := args[0]

		if install {
			return installCompletion(cmd, shell)
		}
		return printCompletion(cmd, shell)
	},
}

func init() {
	Cmd.Flags().Bool("install", false, "Append the completion script to your shell profile file")
}

// printCompletion writes the completion script for the given shell to stdout.
func printCompletion(cmd *cobra.Command, shell string) error {
	root := cmd.Root()
	switch shell {
	case "bash":
		return root.GenBashCompletion(os.Stdout)
	case "zsh":
		return root.GenZshCompletion(os.Stdout)
	case "fish":
		return root.GenFishCompletion(os.Stdout, true)
	case "powershell":
		return root.GenPowerShellCompletionWithDesc(os.Stdout)
	default:
		return fmt.Errorf("unsupported shell %q — supported: bash, zsh, fish, powershell", shell)
	}
}

// installCompletion appends the completion script to the user's shell profile.
func installCompletion(cmd *cobra.Command, shell string) error {
	if shell == "powershell" {
		return fmt.Errorf("--install is not supported for powershell; add the script manually to your PowerShell profile")
	}

	profileName, ok := profileFiles[shell]
	if !ok {
		return fmt.Errorf("unsupported shell %q — supported: bash, zsh, fish, powershell", shell)
	}

	if runtime.GOOS == "windows" {
		return fmt.Errorf("--install is not supported on Windows; source the script manually")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	profilePath := filepath.Join(home, profileName)

	// Ensure parent directory exists (needed for fish).
	if err := os.MkdirAll(filepath.Dir(profilePath), 0750); err != nil {
		return fmt.Errorf("could not create directory for profile %s: %w", profilePath, err)
	}

	// Generate the completion script into a string builder.
	var sb strings.Builder
	root := cmd.Root()
	switch shell {
	case "bash":
		err = root.GenBashCompletion(&sb)
	case "zsh":
		err = root.GenZshCompletion(&sb)
	case "fish":
		err = root.GenFishCompletion(&sb, true)
	}
	if err != nil {
		return fmt.Errorf("failed to generate completion script: %w", err)
	}

	script := sb.String()

	// Check if already installed to avoid duplicates.
	existing, readErr := os.ReadFile(profilePath)
	if readErr == nil && strings.Contains(string(existing), script) {
		fmt.Printf("Completion for %s is already installed in %s\n", shell, profilePath)
		return nil
	}

	f, err := os.OpenFile(profilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open %s: %w", profilePath, err)
	}
	defer func() { _ = f.Close() }()

	_, err = fmt.Fprintf(f, "\n# 1build shell completion (%s)\n%s\n", shell, script)
	if err != nil {
		return fmt.Errorf("could not write to %s: %w", profilePath, err)
	}

	fmt.Printf("Installed %s completion to %s\n", shell, profilePath)
	fmt.Printf("Restart your shell or run: source %s\n", profilePath)
	return nil
}
