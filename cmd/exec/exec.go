package exec

import (
	"errors"
	"github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
)

func ExecuteCommandAndExit(commands ...string) {
	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		utils.PrintErr(err)
		return
	}
	beforeCmd := configuration.Before
	afterCmd := configuration.After

	cmdMap, err := createMapOfCommandsToExecute(configuration, commands...)
	if err != nil {
		return
	}

	printMessage(cmdMap, beforeCmd, afterCmd)

	if beforeCmd != "" {
		err := executeAndStopIfFailed(beforeCmd)
		if err != nil {
			utils.Println("\nFailed to execute '" + beforeCmd + "'\n")
			return
		}
	}
	for _, v := range cmdMap {
		err = executeAndStopIfFailed(v)
		if err != nil {
			utils.Println("\nFailed to execute '" + v + "'\n")
			return
		}
	}

	if afterCmd != "" {
		err = executeAndStopIfFailed(afterCmd)
		if err != nil {
			utils.Println("\nFailed to execute '" + afterCmd + "'\n")
			return
		}
	}
}

func printMessage(m map[string]string, beforeCmd string, afterCmd string) {
	var message = utils.DASH() + "\n"

	if beforeCmd != "" {
		message = message + "Before: " + beforeCmd + "\n\n"
	}
	for k, v := range m {
		message = message + k + " : " + v + "\n"
	}

	if afterCmd != "" {
		message = message + "\nAfter: " + afterCmd + "\n"
	}
	message = message + utils.DASH()
	utils.Println(message)
}

func getCommandFromName(config config.OneBuildConfiguration, cmd string) string {
	for _, command := range config.Commands {
		for k, v := range command {
			if k == cmd {
				return v
			}
		}
	}
	return ""
}

func createMapOfCommandsToExecute(oneBuildConfiguration config.OneBuildConfiguration, commands ...string) (map[string]string, error) {
	var cmdMap map[string]string
	cmdMap = make(map[string]string)

	for _, name := range commands {
		value := getCommandFromName(oneBuildConfiguration, name)
		if value == "" {
			utils.Println("No command '" + name + "' found in config file\n")
			config.PrintConfiguration(oneBuildConfiguration)
			return nil, errors.New("")
		}
		cmdMap[name] = value
	}
	return cmdMap, nil
}

func executeAndStopIfFailed(command string) error {
	return sh.NewSession().Command("bash", "-c", command).Run()
}
