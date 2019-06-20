from onebuild.command import Command
from onebuild.file_writer import write


def default_yaml_file(project_name):
    return "project: " + project_name + "\n" + \
           "commands:" + "\n" + \
           "  - build: echo 'Running build'"


class InitCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        write("1build.yaml", "w", default_yaml_file(arguments[1]))
