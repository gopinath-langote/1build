#!/usr/bin/env python

import pytest
import imp

build = imp.load_source('1build', '1build')

dash = '-' * 50


def test_build_successful_command(capsys):
    build.run("tests/build_file.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_message = "" + dash + "\n" \
                                   "Name: build\n" \
                                   "Command: ls\n" \
                       + dash + "\n"
    assert expected_message in captured.out


def test_should_fail_with_invalid_file_message_if_file_is_not_in_correct_yaml_format(capsys):
    build.run("tests/invalid_yaml_file.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()

    invalid_file_error_message = "Error in parsing 'tests/invalid_yaml_file.yaml' config file. Make sure file is in correct format.\n" \
                                 "Sample format is:\n\n" \
                                 + dash + "\n" \
                                          "project: Sample Project\n" \
                                          "commands:\n" \
                                          "  - build: ./gradlew clean build\n" \
                                          "  - lint: ./gradlew spotlessApply\n" \
                                 + dash

    assert invalid_file_error_message in captured.out


def test_should_print_help_on_help_command(capsys):
    build.run("tests/build_file.yaml", ['file_name', 'help'])
    captured = capsys.readouterr()

    invalid_file_error_message = "Usage: 1build <command_name> \n\n" \
                                 "project: Sample Project\n" + \
                                 "commands:\n" + \
                                 "build | ls\n" + \
                                 "lint | ls"

    assert invalid_file_error_message in captured.out


def test_should_print_help_if_no_command_specified(capsys):
    build.run("tests/build_file.yaml", ['file_name'])
    captured = capsys.readouterr()

    invalid_file_error_message = "Usage: 1build <command_name> \n\n" \
                                 "project: Sample Project\n" + \
                                 "commands:\n" + \
                                 "build | ls\n" + \
                                 "lint | ls"

    assert invalid_file_error_message in captured.out


def test_should_print_command_not_found_if_no_command_found_with_given_name(capsys):
    build.run("tests/build_file.yaml", ['file_name', 'random'])
    captured = capsys.readouterr()

    invalid_file_error_message = "No command 'random' found in config file 'tests/build_file.yaml'\n\n" \
                                 "Usage: 1build <command_name> \n\n" \
                                 "project: Sample Project\n" + \
                                 "commands:\n" + \
                                 "build | ls\n" + \
                                 "lint | ls"

    assert invalid_file_error_message in captured.out
