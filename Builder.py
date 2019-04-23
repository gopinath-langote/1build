import yaml
import sys
import CommandRunner as cmdRunner


class Command:
    def __init__(self, cmd, description):
        self.cmd = cmd
        self.description = description


if __name__ == '__main__':
    commandToExecute = sys.argv[1]
    with open("1build.yaml", 'r') as stream:
        try:
            content = yaml.safe_load(stream)
            project = content["project"]
            commands = []
            for cmd in content["commands"]:
                command = cmd.get(commandToExecute)
                if command is not None:
                    print "executing " + command["description"]
                    print command["cmd"]
                    cmdRunner.run(command["cmd"])
        except yaml.YAMLError as exc:
            print(exc)
