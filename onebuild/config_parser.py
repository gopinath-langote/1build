#!/usr/bin/env python

import os

from ruamel.yaml import YAML

from .project import Command, Project
from .utils import DASH, sample_yaml_file


def parse_project_config(build_file_name):
    if os.path.exists(build_file_name):
        with open(build_file_name, 'r') as stream:
            try:
                yaml = YAML(typ="safe")
                content = yaml.load(stream)
                before = content.get("before", None)
                after = content.get("after", None)
                return Project(
                    name=(content["project"]),
                    before=before,
                    after=after,
                    commands=__command_list__(content["commands"])
                )
            except Exception:
                raise ValueError(__parsing_error_message__(build_file_name))
    else:
        raise ValueError(__file_not_found_error_message__(build_file_name))


def __command_list__(raw_string):
    commands = []
    for cmd in raw_string:
        for key, val in cmd.items():
            commands.append(Command(name=key, cmd=val))
    return commands


def __file_not_found_error_message__(build_file_name):
    return "No '" + build_file_name + "' file found in current directory."


def __parsing_error_message__(build_file_name):
    return "Error in parsing '" + build_file_name + "' config file." + \
           " Make sure file is in correct format." + \
           "\nSample format is:\n\n" + DASH + "\n" + sample_yaml_file() + \
           "\n" + DASH + "\n"
