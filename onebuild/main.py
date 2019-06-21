#!/usr/bin/env python

import sys

from onebuild.action_to_command_lookup import ActionToCommandLookup
from .input_parser import command_to_run, argument_parser

BUILD_FILE_NAME = "1build.yaml"


def run(build_file_name, arguments):
    try:
        arg_parser = argument_parser()
        command_name = command_to_run(arg_parser, arguments)

        ActionToCommandLookup() \
            .get_command_for_action(command_name) \
            .execute(arg_parser, arguments, build_file_name, command_name)

    except ValueError as error:
        print(error)


def main():
    run(BUILD_FILE_NAME, sys.argv[1:])
