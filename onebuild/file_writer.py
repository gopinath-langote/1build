def write(file_name, mode, content):
    with open(file_name, mode) as h:
        h.write(content)
        h.close()
