#!/usr/bin/env python

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
