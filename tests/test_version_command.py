#!/usr/bin/env python

from onebuild.main import run
from onebuild import __version__

VERSION_MESSAGE = "1build {version}".format(version=__version__)


def test_show_version(capsys):
    run("tests/data/build_file.yaml", ['--version'])
    assert_version(capsys)


def test_show_version_on_short_input(capsys):
    run("tests/data/build_file.yaml", ['-v'])

    assert_version(capsys)


def assert_version(capsys):
    captured = capsys.readouterr()
    assert VERSION_MESSAGE in captured.out
