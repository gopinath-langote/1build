package fixtures

// ExecFixtures returns test fixtures for command execution tests.
func ExecFixtures() []struct {
	Name                string
	InitialYAML         string
	Args                []string
	ExpectedOutContains []string
} {
	return []struct {
		Name                string
		InitialYAML         string
		Args                []string
		ExpectedOutContains []string
	}{
		{
			Name: "exec with all hooks",
			InitialYAML: `
project: test
before: echo beforeAll
after: echo afterAll
commands:
  - build:
      before: echo before
      command: echo main
      after: echo after
`,
			Args: []string{"exec", "build"},
			ExpectedOutContains: []string{
				"beforeAll: echo running pre-command",
				"-----------------------------[ beforeAll ]------------------------------",
				"Executing command: build",
				"-------------------------------[ build ]--------------------------------",
				"FAILURE - Total Time:",
			},
		},
	}
}
