#!/usr/bin/env python
import sys
from io import StringIO
from unittest.mock import patch, MagicMock

from onebuild.main import run


@patch('os.path.isfile')
def test_create_default_yaml_file(mock_isfile):
    mock_file_writer = MagicMock()

    mock_isfile.return_value = False

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


@patch('os.path.isfile')
def test_error_message_if_file_already_exists(mock_isfile):
    mock_file_writer = MagicMock()

    mock_isfile.return_value = True

    captured_output = StringIO()
    sys.stdout = captured_output

    with patch("onebuild.init_command.write",
               mock_file_writer,
               create=True):
        run("", ['--init', 'some project'])

    sys.stdout = sys.__stdout__
    assert "1build.yaml configuration file already exists." \
           in captured_output.getvalue()

    mock_file_writer.assert_not_called()
