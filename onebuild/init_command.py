from onebuild.command import Command
from onebuild.file_writer import write


def default_yaml_file(project_name):
    return "project: " + project_name + "\n" + \
           "commands:" + "\n" + \
           "  - build: echo 'Running build'"


class InitCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        try:
            write("1build.yaml", "w", default_yaml_file(arguments[1]))
        except IndexError:
            raise ValueError(
                __project_name_not_found_error_message__()
            )


def __project_name_not_found_error_message__():
    return "Please add project name after init command as " \
           "follows - 1build --init myproject"
