
# 7. requirements

Date: 22-08-2020 10:40:15

## Status

Proposed

## Context

We need to determine how to do logging and whether we should utilitze a utility like sentry in our release builds.

We have, or will have a mechanism to log based off of a verbosity flag. What we must decide is whether to use a tool like
[sentry](https://sentry.io/for/go/) to track errors of users. Or if we should just let users report errors as they discover them.

The benefits of sentry are that bugs can be discovered more immediately and we are less reliant on users reporting problems.

## Decision

Proposed for decision

## Consequences

None until a decision is made
