from onebuild.commands.help_command import HelpCommand
from onebuild.commands.init_command import InitCommand
from onebuild.commands.list_command import ListCommand
from onebuild.commands.perform_command import PerformCommand
from onebuild.commands.version_command import VersionCommand
from onebuild.utils import PredefinedActions


class ActionToCommandLookup:

    def __init__(self):
        self._command_to_action_dictionary = dict({
            PredefinedActions.HELP: HelpCommand(),
            PredefinedActions.INIT: InitCommand(),
            PredefinedActions.LIST: ListCommand(),
            PredefinedActions.VERSION: VersionCommand()})

    def get_command_for_action(self, action):
        return self._command_to_action_dictionary.get(action, PerformCommand())
