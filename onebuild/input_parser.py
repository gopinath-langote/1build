#!/usr/bin/env python


def command_to_run(arguments):
    return "help" if len(arguments) == 1 else arguments[1]