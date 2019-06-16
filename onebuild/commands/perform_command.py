from onebuild.commands.command import Command
from onebuild.config_parser import parse_project_config
from onebuild.executor import execute


class PerformCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        project = parse_project_config(build_file_name)
        command = project.get_command(command_name)
        execute(command, project.before, project.after)
