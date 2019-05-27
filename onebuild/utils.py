# !/usr/bin/env python


DASH = '-' * 50


def sample_yaml_file():
    return "project: Sample Project" + "\n" + \
           "commands:" + "\n" + \
           "  - build: ./gradlew clean build" + "\n" + \
           "  - lint: ./gradlew spotlessApply"


def help_message(project):
    return "Usage: 1build <command_name> \n\n" + project.__str__()
