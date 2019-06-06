#!/usr/bin/env python

from onebuild.main import run

VERSION_MESSAGE = "1build 1.0.0"


def test_show_version(capsys):
    run("tests/data/build_file.yaml", ['--version'])
    assert_version(capsys)


def test_show_version_on_short_input(capsys):
    run("tests/data/build_file.yaml", ['-v'])

    assert_version(capsys)


def test_show_version_mismatch(capsys):
    version_message = "1build 1.0.1"
    run("tests/data/build_file.yaml", ['-v'])
    captured = capsys.readouterr()
    assert version_message != captured.out


def assert_version(capsys):
    captured = capsys.readouterr()
    assert VERSION_MESSAGE in captured.out
