from onebuild.commands.command import Command
from onebuild.utils import version_string


class VersionCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        print(version_string())
