#!/usr/bin/env python

import sys

from .config_parser import parse_project_config
from .executor import execute
from .input_parser import command_to_run, argument_parser
from .utils import print_help, config_string, PredefinedActions

BUILD_FILE_NAME = "1build.yaml"


def run(build_file_name, arguments):
    try:
        arg_parser = argument_parser()
        command_name = command_to_run(arg_parser, arguments)
        if command_name is PredefinedActions.HELP:
            print_help(arg_parser)
        else:
            project = parse_project_config(build_file_name)
            if command_name is PredefinedActions.LIST:
                print(config_string(project))
            else:
                command = project.get_command(command_name)
                execute(command, project.before, project.after)
    except ValueError as error:
        print(error)


def main():
    run(BUILD_FILE_NAME, sys.argv[1:])
