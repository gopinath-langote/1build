# !/usr/bin/env python

from onebuild import __version__

DASH = '-' * 50
NEWLINE = "\n"


def sample_yaml_file():
    return "project: Sample Project" + "\n" + \
           "commands:" + "\n" + \
           "  - build: ./gradlew clean build" + "\n" + \
           "  - lint: ./gradlew spotlessApply"


def default_yaml_file(project_name):
    return "project: " + project_name + "\n" + \
           "commands:" + "\n" + \
           "  - build: echo 'Running build'"


def print_help(parser):
    parser.print_help()


def config_string(project):
    return "" + DASH + NEWLINE + project.__str__() + NEWLINE + DASH


def version_string():
    return '1build {version} '.format(version=__version__)
