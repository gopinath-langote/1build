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


def test_error_message_if_file_name_is_not_provided_with_init(capsys):
    run("", ['--init'])
    captured = capsys.readouterr()
    expected_message = "The 'project name' parameter is missing with --init" \
                       "\n\nusage: 1build --init project_name"
    assert expected_message in captured.out
