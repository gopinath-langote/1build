# !/usr/bin/env python


def max_length_from(*args):
    max_length = 0  # type: int
    for string in args:
        if len(string) > max_length:
            max_length = len(string)
    return max_length + 20
