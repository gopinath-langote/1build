#!/usr/bin/env python
import yaml

import os
import sys


def run(arguments):
    try:
        project = parse_project_config()
        command = project.get_command(command_to_run(arguments))
        execute(command)
    except ValueError as error:
        print error


class Project:
    def __init__(self, name, commands):
        self.name = name
        self.commands = commands

    def get_command(self, command_name):
        if self.__has_command__(command_name):
            for cmd in self.commands:
                if cmd.name == command_name:
                    return cmd
        else:
            raise ValueError("No command " + command_name + " found in config file `1build.yaml` \n" +
                             "Available commands:" + "\n" +
                             self.available_commands()
                             )

    def __has_command__(self, command_name):
        for cmd in self.commands:
            if cmd.name == command_name:
                return True
        return False

    def available_commands(self):
        return "\n".join(map(str, self.commands))

    def __str__(self):
        return "Project Name: " + self.name + " \nAvailable commands:\n" + self.available_commands()


class Command:
    def __init__(self, name, cmd, description):
        self.name = name
        self.cmd = cmd
        self.description = description

    def __str__(self):
        return "Name: " + self.name + " | command: " + self.cmd + " | description: " + self.description


def execute(command):
    print("---------------------------------------------------")
    print ("Name: " + command.name)
    print ("Command: " + command.cmd)
    print ("Description: " + command.description)
    print("---------------------------------------------------")
    os.system(command.cmd)


def parse_command(raw_string):
    command_name = next(iter(raw_string), None)
    return Command(name=command_name, cmd=raw_string.get(command_name)["cmd"],
                   description=raw_string.get(command_name)["description"])


def get_command_list_from_config(raw_string):
    commands = []
    for cmd in raw_string: commands.append(parse_command(cmd))
    return commands


def parse_project_config():
    with open("1build.yaml", 'r') as stream:
        try:
            content = yaml.safe_load(stream)
        except yaml.YAMLError, exc:
            raise ValueError(
                "Error in parsing `1build.yaml` config file. Make sure file is in correct format. \n\n" + exc.__str__()
            )
        return Project(name=(content["project"]), commands=(get_command_list_from_config(content["commands"])))


def command_to_run(arguments):
    if len(arguments) is 1:
        return "build"
    else:
        return arguments[1]


run(sys.argv)
