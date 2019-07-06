#!/usr/bin/bash
_1build_completions()
{
    COMPREPLY=($(compgen -W "$(1build -lr | sed 's/\\t//')" -- "${COMP_WORDS[COMP_CWORD]}"))
}

complete -F _1build_completions 1build
