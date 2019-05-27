#!/usr/bin/env python

import imp

from .test_utils import DASH
from onebuild.main import run
build = imp.load_source('1build', '1build')


def test_build_successful_command(capsys):
    run("tests/data/build_file.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_message = "" + DASH + "\n" \
                                   "Name: build\n" \
                                   "Command: echo 'Running build'\n" \
                       + DASH + "\n"
    expected_command_output = "Running build"
    assert expected_message in captured.out
    assert expected_command_output in captured.out


def test_should_fail_with_invalid_file_message_if_file_is_not_in_correct_yaml_format(capsys):
    run("tests/data/invalid_yaml_file.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()

    invalid_file_error_message = "Error in parsing 'tests/data/invalid_yaml_file.yaml' config file. Make sure file is in correct format.\n" \
                                 "Sample format is:\n\n" \
                                 + DASH + "\n" \
                                          "project: Sample Project\n" \
                                          "commands:\n" \
                                          "  - build: ./gradlew clean build\n" \
                                          "  - lint: ./gradlew spotlessApply\n" \
                                 + DASH

    assert invalid_file_error_message in captured.out


def test_should_print_help_on_help_command(capsys):
    run("tests/data/build_file.yaml", ['file_name', 'help'])
    captured = capsys.readouterr()

    invalid_file_error_message = "Usage: 1build <command_name> \n\n" \
                                 "project: Sample Project\n" + \
                                 "commands:\n" + \
                                 "build | echo 'Running build'\n" + \
                                 "lint | echo 'Running lint'"

    assert invalid_file_error_message in captured.out


def test_should_print_help_if_no_command_specified(capsys):
    run("tests/data/build_file.yaml", ['file_name'])
    captured = capsys.readouterr()

    invalid_file_error_message = "Usage: 1build <command_name> \n\n" \
                                 "project: Sample Project\n" + \
                                 "commands:\n" + \
                                 "build | echo 'Running build'\n" + \
                                 "lint | echo 'Running lint'"

    assert invalid_file_error_message in captured.out


def test_should_print_command_not_found_if_no_command_found_with_given_name(capsys):
    run("tests/data/build_file.yaml", ['file_name', 'random'])
    captured = capsys.readouterr()

    invalid_file_error_message = "No command 'random' found in config file\n\n" \
                                 "Usage: 1build <command_name> \n\n" \
                                 "project: Sample Project\n" + \
                                 "commands:\n" + \
                                 "build | echo 'Running build'\n" + \
                                 "lint | echo 'Running lint'"

    assert invalid_file_error_message in captured.out
