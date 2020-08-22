
# 7. requirements

Date: 22-08-2020 10:40:15

## Status

Proposed

## Context

We need a mechanism to avoid repeat downloads when using the mirror or favorites command.

## Decision

For our V1 release we will simply utilitize the filename and do a file system check
to see if the file exists and skip if it does.

In a future release we may use a sqlite3 database to keep track of what and where the files are.
The complexity there is we may require the user to install sqlite3 which is not ideal.

We may also use a simple text file, however the file size that tracks downloads may get large.

More research is required.

## Consequences

We dont have to repeat every download on every run.
