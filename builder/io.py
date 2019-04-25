#!/usr/bin/env python

import os


def pretty_print(message):
    print (message)


def command_to_run(arguments):
    if len(arguments) is 1:
        return "build"
    return arguments[1]


def run(command):
    os.system(command)
