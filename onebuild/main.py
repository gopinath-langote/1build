#!/usr/bin/env python

import sys

from .config_parser import parse_project_config
from .executor import execute
from .input_parser import command_to_run, argument_parser
from .utils import print_help

BUILD_FILE_NAME = "1build.yaml"


def run(build_file_name, arguments):
    try:
        arg_parser = argument_parser()
        command_name = command_to_run(arg_parser, arguments)
        project = parse_project_config(build_file_name)
        if command_name is "help":
            print_help(arg_parser, project)
        else:
            command = project.get_command(command_name)
            execute(command, project.before, project.after)
    except ValueError as error:
        print(error)


def main():
    run(BUILD_FILE_NAME, sys.argv[1:])
