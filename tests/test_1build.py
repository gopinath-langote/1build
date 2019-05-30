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
    expected_command_output = "Running build"
    assert expected_message in captured.out
    assert expected_command_output in captured.out


def test_should_fail_if_file_is_not_in_correct_yaml_format(capsys):
    run("tests/data/invalid_yaml_file.yaml", ['build'])
    captured = capsys.readouterr()

    invalid_file_error_message = \
        "Error in parsing 'tests/data/invalid_yaml_file.yaml' config file." \
        " Make sure file is in correct format.\n" \
        "Sample format is:\n\n" \
        + DASH + "\n" \
                 "project: Sample Project\n" \
                 "commands:\n" \
                 "  - build: ./gradlew clean build\n" \
                 "  - lint: ./gradlew spotlessApply\n" \
        + DASH

    assert invalid_file_error_message in captured.out


def test_should_fails_if_no_command_found_with_given_name(capsys):
    run("tests/data/build_file.yaml", ['random'])
    captured = capsys.readouterr()

    invalid_file_error = "No command 'random' found in config file\n\n" \
                         "" + DASH + "\n" \
                                     "project: Sample Project\n" \
                                     "commands:\n" \
                                     "build | echo 'Running build'\n" \
                                     "lint | echo 'Running lint'\n" \
                                     "" + DASH

    assert invalid_file_error in captured.out
