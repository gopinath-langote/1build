#!/usr/bin/env python

from onebuild.main import run

USAGE_HELP_MESSAGE \
    = """usage: 1build [-h] [-l] [-v] [-i] [command [command ...]]

positional arguments:
  command        Command(s) to run - from `1build.yaml` file

optional arguments:
  -h, --help     Print this help message
  -l, --list     Show all available commands - from `1build.yaml` file
  -v, --version  Show version of 1build and exit
  -i, --init     Create default `1build.yaml` configuration file
"""


def test_show_help(capsys):
    run("tests/data/build_file.yaml", ['--help'])

    assert_usage_help(capsys)


def test_show_help_on_short_input(capsys):
    run("tests/data/build_file.yaml", ['-h'])

    assert_usage_help(capsys)


def test_show_help_on_not_input(capsys):
    run("tests/data/build_file.yaml", [])
    assert_usage_help(capsys)


def assert_usage_help(capsys):
    captured = capsys.readouterr()

    assert USAGE_HELP_MESSAGE in captured.out
