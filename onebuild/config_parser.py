#!/usr/bin/env python

import os

from ruamel.yaml import YAML

from .project import Command, Project
from .utils import DASH, sample_yaml_file


def __get_project_config(working_directory, build_file_name):
    potential_path = os.path.join(working_directory, build_file_name)
    if os.path.exists(potential_path) and os.path.isfile(potential_path):
        return os.path.abspath(potential_path)
    else:
        parent = os.path.abspath(os.path.join(working_directory,
         os.path.pardir))
        if(parent == working_directory):
            return None
        return __get_project_config(parent, build_file_name)

def parse_project_config(build_file_name):
    project_config_path = __get_project_config(os.getcwd(), build_file_name)
    if project_config_path is not None:
        with open(project_config_path, 'r') as stream:
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
    return "No '" + build_file_name + "' file found in working tree."


def __parsing_error_message__(build_file_name):
    return "Error in parsing '" + build_file_name + "' config file." + \
           " Make sure file is in correct format." + \
           "\nSample format is:\n\n" + DASH + "\n" + sample_yaml_file() + \
           "\n" + DASH + "\n"
