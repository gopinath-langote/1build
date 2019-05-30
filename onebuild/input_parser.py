#!/usr/bin/env python

import argparse

from .utils import PredefinedActions


def argument_parser():
    parser = argparse.ArgumentParser(prog='1build', add_help=False)
    parser.add_argument(
        'command',
        nargs='?',
        default=PredefinedActions.HELP,
        help='Command to run - from `1build.yaml` file',
    )
    parser.add_argument(
        '-h', '--help',
        action='store_true',
        default=False,
        help="Print this help message"
    )
    parser.add_argument(
        '-l', '--list',
        action='store_true',
        default=False,
        help="Show all available commands - from `1build.yaml` file"
    )
    return parser


def command_to_run(arg_parser, arguments):
    args = arg_parser.parse_args(args=arguments)
    if args.help is True:
        return PredefinedActions.HELP
    if args.list is True:
        return PredefinedActions.LIST
    return args.command
