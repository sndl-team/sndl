
# 2. tooling
======
Date: 21-08-2020 20:58:47

## Status
======
Proposed

## Context
======

SNDL needs propert tools both static analysis tools, and linting 
utilities. Beyond this there are requirements around publishing, 
code quality, testing and design that must be addressed.

## Decision
======

The decisions are listed as follows:

- formatter
    - gofmt

This tool also does some static analysis and code quality

- testing
    - go test

## Consequences
======

We are able to do testing and formatting to eliminate 
inconsistencies across developers.

