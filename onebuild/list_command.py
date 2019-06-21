from onebuild.command import Command
from onebuild.config_parser import parse_project_config
from onebuild.utils import config_string


class ListCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        project = parse_project_config(build_file_name)
        print(config_string(project))
