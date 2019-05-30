# !/usr/bin/env python


DASH = '-' * 50
NEWLINE = "\n"


def sample_yaml_file():
    return "project: Sample Project" + "\n" + \
           "commands:" + "\n" + \
           "  - build: ./gradlew clean build" + "\n" + \
           "  - lint: ./gradlew spotlessApply"


def help_message(project):
    return "Usage: 1build <command_name> \n\n" + project.__str__()


def print_help(parser):
    parser.print_help()


def config_string(project):
    return "" + DASH + NEWLINE + project.__str__() + NEWLINE + DASH
