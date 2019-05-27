#!/usr/bin/env python

import os

from .utils import dash


def execute(command, before=None, after=None):
    cmd = command.cmd
    print(dash + "\nName: " + command.name + "\nCommand: " + command.cmd)
    if before:
        print("Before: " + before)
        cmd = before + " && " + cmd
    if after:
        print("After: " + after)
        cmd = cmd + " && " + after
    print(dash)
    os.system(cmd)
