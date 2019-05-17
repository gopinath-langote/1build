#!/usr/bin/env python

import imp

from test_utils import dash

build = imp.load_source('1build', '1build')


def test_build_successful_run_file_command(capsys):
    build.run("tests/data/build_file_with_run_shell_script_cmd.yaml", ['file_name', 'run_file'])
    captured = capsys.readouterr()
    expected_message = "" + dash + "\n" \
                                   "Name: run_file\n" \
                                   "Command: sh tests/data/test_shell_script.sh\n" \
                       + dash + "\n"
    expected_command_output = "Shell Script Execution"
    assert expected_message in captured.out
    assert expected_command_output in captured.out
