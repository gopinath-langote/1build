#!/usr/bin/env python

import sys

from .config_parser import parse_project_config
from .executor import execute
from .input_parser import command_to_run
from .utils import help_message

BUILD_FILE_NAME = "1build.yaml"


def run(build_file_name, arguments):
    global BUILD_FILE_NAME
    BUILD_FILE_NAME = build_file_name
    try:
        project = parse_project_config(build_file_name)
        command_name = command_to_run(arguments)
        if command_name == "help":
            print(help_message(project))
        else:
            command = project.get_command(command_name)
            execute(command, project.before, project.after)
    except ValueError as error:
        print(error)


def main():
    run(BUILD_FILE_NAME, sys.argv)
