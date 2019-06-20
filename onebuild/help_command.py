from onebuild.command import Command


class HelpCommand(Command):

    def execute(self, arg_parser, arguments, build_file_name, command_name):
        arg_parser.print_help()
