from onebuild.utils import default_yaml_file


def write_default_config_file():
    with open("output.txt", "w") as h:
        h.write(default_yaml_file())
        h.close()
