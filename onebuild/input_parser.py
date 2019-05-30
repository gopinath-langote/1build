#!/usr/bin/env python

import argparse


def argument_parser():
    parser = argparse.ArgumentParser(prog='1build', add_help=False)
    parser.add_argument(
        'command',
        nargs='?',
        default="help",
        help='Command to run from the `1build.yaml` config file',
    )
    parser.add_argument(
        '-h', '--help',
        action='store_true',
        default=False,
        help="Print this help message"
    )
    return parser


def command_to_run(argument_parser, arguments):
    args = argument_parser.parse_args(args=arguments)
    if args.help is True:
        return "help"
    return args.command
