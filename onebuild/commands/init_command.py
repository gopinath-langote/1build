from onebuild.commands.command import Command
from onebuild.file_writer import write
from onebuild.utils import default_yaml_file


class InitCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        write("1build.yaml", "w", default_yaml_file(arguments[1]))
