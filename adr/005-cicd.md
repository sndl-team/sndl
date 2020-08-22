
# 5. cicd

Date: 21-08-2020 22:06:18

## Status

Accepted

## Context

We need to utilize a CI process for our code.

Travis requires a open source project.

Github actions is free and we get 2000 build minutes a month.

- linux builds are 1x
- windows builds are 2x
- mac builds are 10x

## Decision

We will use github actions because it is free.

## Consequences

We need to keep build times low and minimize unnecessary builds.

We will try to have a keyword in our commit messages that will not fire a build. [ci skip]