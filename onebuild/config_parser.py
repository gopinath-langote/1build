#!/usr/bin/env python

import os

from ruamel.yaml import YAML

from .project import Command, Project
from .utils import DASH, sample_yaml_file


def parse_project_config(build_file_name):
    """
    :param build_file_name: 1build config file name, default is `1build.yaml`
    :return: configuration from file as `Project` class.
    """
    if os.path.exists(build_file_name):
        with open(build_file_name, 'r') as stream:
            try:
                yaml = YAML(typ="safe")
                content = yaml.load(stream)
                before = content.get("before", None)
                after = content.get("after", None)
                return Project(name=(content["project"]),
                               before=before,
                               after=after,
                               commands=__get_command_list_from_config__(content["commands"]))
            except:
                raise ValueError(
                    "Error in parsing '" + build_file_name + "' config file."
                    + " Make sure file is in correct format.\nSample format is:\n\n" +
                    DASH + "\n" + sample_yaml_file() + "\n" + DASH + "\n"
                )
    else:
        raise ValueError("No '" + build_file_name + "' file found in current directory.")


def __get_command_list_from_config__(raw_string):
    commands = []
    for cmd in raw_string:
        for key, val in cmd.items():
            commands.append(Command(name=key, cmd=val))
    return commands
