#!/usr/bin/env python

import pytest
import imp

build = imp.load_source('1build', '1build')


def test_build_successful_command(capsys):
    build.run("tests/build_file.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_message = "---------------------------------------------------\n" \
                       "Name: build\n" \
                       "Command: ls\n" \
                       "---------------------------------------------------\n"
    assert expected_message in captured.out
