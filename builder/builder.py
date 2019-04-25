#!/usr/bin/env python

import yaml
import io as io
import utils
from project import Project
from command import Command


def execute(command):
    max_length = utils.max_length_from(command.name, command.cmd, command.description)
    io.print_separator(max_length)
    io.pretty_print("Name: " + command.name)
    io.pretty_print("Command: " + command.cmd)
    io.pretty_print("Description: " + command.description)
    io.print_separator(max_length)
    io.run(command.cmd)


def parse_command(raw_string):
    command_name = next(iter(raw_string), None)
    cmd = raw_string.get(command_name)["cmd"]
    description = raw_string.get(command_name)["description"]
    return Command(
        name=command_name,
        cmd=cmd,
        description=description
    )


def get_command_list_from_config(raw_string):
    commands = []
    for cmd in raw_string:
        commands.append(parse_command(cmd))
    return commands


def parse_project_config():
    with open("1build.yaml", 'r') as stream:
        try:
            content = yaml.safe_load(stream)
        except yaml.YAMLError, exc:
            raise ValueError(
                "Error in parsing `1build.yaml` config file. Make sure file is in correct format." +
                "\n\n" +
                exc.__str__()
            )
        name = content["project"]
        commands = get_command_list_from_config(content["commands"])
        return Project(
            name=name,
            commands=commands
        )


def run(arguments):
    try:
        project = parse_project_config()
        command = project.get_command(io.command_to_run(arguments))
        execute(command)
    except ValueError as error:
        print error
