#!/usr/bin/env python

from onebuild.main import run

USAGE_HELP_MESSAGE = \
    "usage: 1build [-h] [-l] [command]\n\n" \
    "positional arguments:\n  " \
    "command     Command to run – from `1build.yaml` file\n\n" \
    "optional arguments:\n" \
    "  -h, --help  Print this help message\n"\
    "  -l, --list  Show all available commands – from `1build.yaml` file\n"


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
