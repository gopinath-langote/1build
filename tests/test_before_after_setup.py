import imp

build = imp.load_source('1build', '1build')
from .test_utils import DASH
from onebuild.main import run


def test_build_successful_with_before_and_after_command(capsys):
    run("tests/data/build_file_with_before_and_after.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_cmd_message = "" + DASH + "\n" \
                                       "Name: build\n" \
                                       "Command: echo 'Running build'\n" \
                                       "Before: echo 'Running some setup script'\n" \
                                       "After: echo 'Running some cleanup script'\n" \
                           + DASH + "\n"
    expected_before_command_output = "Running some setup script"
    expected_command_output = "Running build"
    expected_after_command_output = "Running some cleanup script"
    assert expected_cmd_message in captured.out
    assert expected_before_command_output in captured.out
    assert expected_command_output in captured.out
    assert expected_after_command_output in captured.out


def test_should_work_with_only_before_command(capsys):
    run("tests/data/build_file_with_before_only.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_cmd_message = "" + DASH + "\n" \
                                       "Name: build\n" \
                                       "Command: echo 'Running build'\n" \
                                       "Before: echo 'Running some setup script'\n" \
                           + DASH + "\n"
    expected_before_command_output = "Running some setup script"
    expected_command_output = "Running build"
    assert expected_cmd_message in captured.out
    assert expected_before_command_output in captured.out
    assert expected_command_output in captured.out


def test_should_work_with_only_after_command(capsys):
    run("tests/data/build_file_with_after_only.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_cmd_message = "" + DASH + "\n" \
                                       "Name: build\n" \
                                       "Command: echo 'Running build'\n" \
                                       "After: echo 'Running some cleanup script'\n" \
                           + DASH + "\n"
    expected_command_output = "Running build"
    expected_after_command_output = "Running some cleanup script"
    assert expected_cmd_message in captured.out
    assert expected_command_output in captured.out
    assert expected_after_command_output in captured.out
