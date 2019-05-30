#!/usr/bin/env python

import argparse


def parser():
    # , formatter_class = argparse.ArgumentDefaultsHelpFormatter
    parser = argparse.ArgumentParser(prog='1build')
    parser.add_argument('command', nargs='?', default='help',
                        help='Command to run from the `1build.yaml` config file',
                        choices=['a', 'help']
                        )
    parser.add_argument('--set', nargs='*', help='set new command shortcut to config')
    # return parser.parse_args()
    return parser


def parser1(project):
    cmds = list(map(lambda x: x.name, project.commands))
    cmds.append("help")
    parser = argparse.ArgumentParser(prog='1build')
    parser.add_argument('command', nargs='?', default='help',
                        help='Command to run from the `1build.yaml` config file',
                        choices=cmds
                        )
    parser.add_argument('--set', nargs='*', help='set new command shortcut to config')
    # return parser.parse_args()
    return parser


def command_to_run(args):
    print(args)
    return args.command
