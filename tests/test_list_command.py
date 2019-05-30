#!/usr/bin/env python

from onebuild.main import run
from .test_utils import DASH

LIST_HELP_MESSAGE = "" + DASH + "\n" + \
                    "project: Sample Project" + "\n" + \
                    "commands:" + "\n" + \
                    "build | echo 'Running build'" + "\n" + \
                    "lint | echo 'Running lint'" + "\n" + \
                    "" + DASH


def test_show_list_of_commands(capsys):
    run("tests/data/build_file.yaml", ['--list'])

    assert_list_of_command_help(capsys)


def test_show_list_of_commands_on_short_input(capsys):
    run("tests/data/build_file.yaml", ['-l'])

    assert_list_of_command_help(capsys)


def assert_list_of_command_help(capsys):
    captured = capsys.readouterr()

    assert LIST_HELP_MESSAGE in captured.out
