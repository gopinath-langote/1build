#!/usr/bin/env python

class Command:
    def __init__(self, name, cmd, description):
        self.name = name
        self.cmd = cmd
        self.description = description

    def __str__(self):
        return "Name: " + self.name + " | command: " + self.cmd + " | description: " + self.description
