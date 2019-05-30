# !/usr/bin/env python

from .utils import config_string


class Command:
    def __init__(self, name, cmd):
        self.name = name
        self.cmd = cmd

    def __str__(self):
        return self.name + " | " + self.cmd


class Project:
    def __init__(self, name, before, after, commands):
        self.name = name
        self.before = before
        self.after = after
        self.commands = commands

    def get_command(self, command_name):
        if self.__has_command__(command_name):
            for cmd in self.commands:
                if cmd.name == command_name:
                    return cmd
        else:
            raise ValueError(
                "No command '" + command_name + "' found in config file" +
                "\n\n" + config_string(self)
            )

    def __has_command__(self, command_name):
        for cmd in self.commands:
            if cmd.name == command_name:
                return True
        return False

    def available_commands(self):
        return "\n".join(map(str, self.commands))

    def __str__(self):
        return "project: " + self.name + "\ncommands:\n" + \
               self.available_commands()
