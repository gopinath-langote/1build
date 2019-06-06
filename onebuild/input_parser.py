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
    parser.add_argument(
        '-v', '--version',
        dest='version',
        action='store_true',
        default=False,
        help="Show version of 1build and exit"
    )
    return parser


def command_to_run(arg_parser, arguments):
    args = arg_parser.parse_args(args=arguments)
    if args.help:
        return PredefinedActions.HELP
    if args.version:
        return PredefinedActions.VERSION
    if args.list:
        return PredefinedActions.LIST
    return args.command
