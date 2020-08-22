
# 7. requirements

Date: 21-08-2020 22:31:45

## Status

Accepted

## Context

We need to have an understanding of what sndl is trying to accomplish
and how we expect to do it.

## Decision

V1
- CLI
    - help menus
    - verbosity
    - search subcommand
        - search term argument (required)
        - series flag (default: any)
        - movie flag (default: any)
    - download subcommand
        - download term argument (required)
        - series or movie (default: any)
        - resolution (default: 720p) (choices: 720p || 10800p)
        - server (default: recommended) (range: 1-6 + names mapping) (fallback: on failure)
    - mirror subcommand
        - path argument (default: current dir/sndl.txt) to sync file
        - delta flag -- tell me what has been updated dont download
    - favorites subcommand -- download list from favorites menu
        - list (default) -- show me
        - mirror flag -- download
    - popular subcommand
        - series flag
        - movies flag (default)
        - new flag -- new episodes

V2
- UI
    - to be determined

## Consequences

We have a defined spec to build from.
