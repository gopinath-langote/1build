#!/usr/bin/env python

from unittest.mock import patch, MagicMock

from onebuild.main import run


def test_create_default_yaml_file():
    mock_file_writer = MagicMock()

    with patch("onebuild.init_command.write",
               mock_file_writer,
               create=True):
        run("", ['--init', 'some project'])

    mock_file_writer.assert_called_with(
        '1build.yaml', 'w',
        "project: some project\ncommands:\n  - build: echo 'Running build'")
