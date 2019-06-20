from onebuild import __version__
from onebuild.command import Command


class VersionCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        version = '1build {version} '.format(version=__version__)
        print(version)
