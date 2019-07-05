#!/usr/bin/env python

from onebuild.main import run
from .test_utils import DASH


def test_build_successful_command(capsys):
    run("tests/data/build_file.yaml", ['build'])
    captured = capsys.readouterr()
    expected_message = "" + DASH + "\n" \
                                   "Name: build\n" \
                                   "Command: echo 'Running build'\n" \
                       + DASH + "\n"
    assert expected_message in captured.out


def test_multiple_commands_successful(capsys):
    run("tests/data/build_file.yaml", ['build', 'lint'])
    captured = capsys.readouterr()
    expected_build_cmd_message = \
        "" + DASH + "\n" \
                    "Name: build\n" \
                    "Command: echo 'Running build'\n" \
        + DASH + "\n"
    expected_lint_cmd_message = \
        "" + DASH + "\n" \
                    "Name: lint\n" \
                    "Command: echo 'Running lint'\n" \
        + DASH + "\n"
    assert expected_build_cmd_message in captured.out
    assert expected_lint_cmd_message in captured.out
