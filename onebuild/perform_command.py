import os

from onebuild.command import Command
from onebuild.config_parser import parse_project_config
from onebuild.utils import DASH


class PerformCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        project = parse_project_config(build_file_name)
        args = arg_parser.parse_args(args=arguments)
        for cmd in args.command:
            self.__execute_command__(project, cmd)

    @staticmethod
    def __execute_command__(project, command_name):
        command = project.get_command(command_name)

        cmd = command.cmd
        print(DASH + "\nName: " + command.name +
              "\nCommand: " + command.cmd)

        before = project.before
        after = project.after

        if before:
            print("Before: " + before)
            cmd = before + " && " + cmd
        if after:
            print("After: " + after)
            cmd = cmd + " && " + after

        print(DASH)
        os.system(cmd)
