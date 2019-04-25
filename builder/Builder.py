#!/usr/bin/env python
# !/usr/bin/env python

import yaml
import sys
import CommandRunner as cmdRunner


class Command:
    def __init__(self, cmd, description):
        self.cmd = cmd
        self.description = description


def getCommandName(arguments):
    if len(arguments) is 1:
        return "build"
    return arguments[1]


def execute(cmdObject):
    print "executing " + cmdObject["description"]
    print cmdObject["cmd"]
    cmdRunner.run(cmdObject["cmd"])


def run(arguments):
    command_to_execute = getCommandName(arguments)
    with open("1build.yaml", 'r') as stream:
        try:
            content = yaml.safe_load(stream)
            project = content["project"]
            for cmd in content["commands"]:
                command = cmd.get(command_to_execute)
                if command is not None:
                    execute(command)
        except yaml.YAMLError as exc:
            print(exc)
