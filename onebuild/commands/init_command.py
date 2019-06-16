from onebuild.commands.command import Command
from onebuild.file_writer import write_default_config_file


class InitCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        write_default_config_file(arguments[1])
