#!/usr/bin/env python

import yaml
import io as io


class Command:
    def __init__(self, cmd, description):
        self.cmd = cmd
        self.description = description


def execute(project, cmd_object):
    cmd = cmd_object["cmd"]
    cmd_description = cmd_object["description"]
    io.pretty_print("-----------------------")
    io.pretty_print("Project: " + project)
    io.pretty_print("Command: " + cmd)
    io.pretty_print("Description: " + cmd_description)
    io.pretty_print("-----------------------")
    io.run(cmd)


def run(arguments):
    command_to_execute = io.command_to_run(arguments)
    with open("1build.yaml", 'r') as stream:
        try:
            content = yaml.safe_load(stream)
            project = content["project"]
            for cmd in content["commands"]:
                command = cmd.get(command_to_execute)
                if command is not None:
                    execute(project, command)
        except yaml.YAMLError as exc:
            print(exc)
