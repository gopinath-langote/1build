from onebuild.commands.command import Command
from onebuild.utils import print_help


class HelpCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        print_help(arg_parser)
