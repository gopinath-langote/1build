#!/usr/bin/env python
try:
    from mock import patch, MagicMock
except ImportError:
    from unittest.mock import patch, MagicMock

from onebuild.main import run


def test_create_default_yaml_file():
    mock_file_writer = MagicMock()

    with patch("onebuild.main.write_default_config_file", mock_file_writer,
               create=True):
        run("", ['--init', 'some project'])

    mock_file_writer.assert_called_with('some project')
