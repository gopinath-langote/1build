from onebuild.predefined_actions import PredefinedActions
from onebuild.help_command import HelpCommand
from onebuild.init_command import InitCommand
from onebuild.list_command import ListCommand
from onebuild.perform_command import PerformCommand
from onebuild.version_command import VersionCommand


class ActionToCommandLookup:

    def __init__(self):
        self._command_to_action_dictionary = dict({
            PredefinedActions.HELP: HelpCommand(),
            PredefinedActions.INIT: InitCommand(),
            PredefinedActions.LIST: ListCommand(),
            PredefinedActions.VERSION: VersionCommand(),
            PredefinedActions.PERFORM: PerformCommand()
        })

    def get_command_for_action(self, action):
        return self._command_to_action_dictionary.get(action, PerformCommand())
