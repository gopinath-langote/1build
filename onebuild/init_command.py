import os

from onebuild.command import Command
from onebuild.file_writer import write


def default_yaml_file(project_name):
    return "project: " + project_name + "\n" + \
           "commands:" + "\n" + \
           "  - build: echo 'Running build'"


class InitCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        if len(arguments) < 2:
            print(__project_name_not_found_error_message__())
        elif os.path.isfile("1build.yaml"):
            print(__file_already_exists_message__())
        else:
            write("1build.yaml", "w", default_yaml_file(arguments[1]))


def __project_name_not_found_error_message__():
    return "The 'project name' parameter is missing with --init" \
           "\n\nusage: 1build --init project_name"


def __file_already_exists_message__():
    return "1build.yaml configuration file already exists."
