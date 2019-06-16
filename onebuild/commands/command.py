import abc


class Command(metaclass=abc.ABCMeta):

    @abc.abstractmethod
    def execute(self, arg_parser, arguments, build_file_name, command_name):
        pass
