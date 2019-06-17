#!/usr/bin/env python

from unittest.mock import patch, mock_open

from onebuild.utils import default_yaml_file
import onebuild.file_writer as file_writer


def test_write():
    open_mock = mock_open()

    with patch("onebuild.file_writer.open", open_mock, create=True):
        some_file_name = "Some file"
        some_mode = "write"
        some_content = "Some content"

        file_writer.write(some_file_name, some_mode, some_content)

    open_mock.assert_called_with(some_file_name, some_mode)
    open_mock.return_value.write.assert_called_once_with(some_content)
