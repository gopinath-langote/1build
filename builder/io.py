#!/usr/bin/env python

import os


def print_separator(length):
    print('-' * length)


def pretty_print(message, separator=False, separator_length=0):
    if separator:
        print_separator(separator_length)
    print message


def command_to_run(arguments):
    if len(arguments) is 1:
        return "build"
    return arguments[1]


def run(command):
    os.system(command)
