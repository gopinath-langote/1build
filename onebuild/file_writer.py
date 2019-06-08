from onebuild.utils import default_yaml_file


def write_default_config_file(project_name):
    with open("1build.yaml", "w") as h:
        h.write(default_yaml_file(project_name))
        h.close()
