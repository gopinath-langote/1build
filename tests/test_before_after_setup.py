import imp

build = imp.load_source('1build', '1build')
dash = '-' * 50


def test_build_successful_with_before_and_after_command(capsys):
    build.run("tests/data/build_file_with_before_and_after.yaml", ['file_name', 'build'])
    captured = capsys.readouterr()
    expected_cmd_message = "" + dash + "\n" \
                                       "Name: build\n" \
                                       "Command: echo 'Running build'\n" \
                                       "Before: echo 'Running some setup script'\n" \
                                       "After: echo 'Running some cleanup script'\n" \
                           + dash + "\n"
    expected_before_command_output = "Running some setup script"
    expected_after_command_output = "Running some cleanup script"
    expected_command_output = "Running build"
    assert expected_cmd_message in captured.out
    assert expected_before_command_output in captured.out
    assert expected_after_command_output in captured.out
    assert expected_command_output in captured.out
