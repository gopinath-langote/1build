#!/usr/bin/env python

try:
    from mock import patch, mock_open
except ImportError:
    from unittest.mock import patch, mock_open

from onebuild.utils import default_yaml_file
import onebuild.file_writer as file_writer


def test_create_default_yaml_file():
    open_mock = mock_open()

    with patch("onebuild.file_writer.open", open_mock, create=True):
        file_writer.write_default_config_file("Some Project")

    open_mock.assert_called_with("1build.yaml", "w")
    open_mock.return_value.write.assert_called_once_with(
        default_yaml_file("Some Project"))
