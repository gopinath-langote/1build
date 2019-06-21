import abc as abstract_class


class Command(metaclass=abstract_class.ABCMeta):

    @abstract_class.abstractmethod
    def execute(self, arg_parser, arguments, build_file_name, command_name):
        pass
