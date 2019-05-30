# !/usr/bin/env python
from enum import Enum


DASH = '-' * 50
NEWLINE = "\n"


class PredefinedActions(Enum):
    ONEBUILD_HELP = 1
    ONEBUILD_LIST = 2


def sample_yaml_file():
    return "project: Sample Project" + "\n" + \
           "commands:" + "\n" + \
           "  - build: ./gradlew clean build" + "\n" + \
           "  - lint: ./gradlew spotlessApply"


def print_help(parser):
    parser.print_help()


def config_string(project):
    return "" + DASH + NEWLINE + project.__str__() + NEWLINE + DASH
