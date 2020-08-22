![logo](assets/twitter_header_photo_2.png)

<div style="text-align: center">
    <i>
        A cli tool to automate tasks on the senturion site.
        <img src="https://travis-ci.com/sndl-team/sndl.svg?branch=master" />
        <a href="https://www.codefactor.io/repository/github/sndl-team/sndl"><img src="https://www.codefactor.io/repository/github/sndl-team/sndl/badge" alt="CodeFactor" /></a>
    </i>
</div>


# Start

To build the project from source just clone the repo and run

```bash
$ go build -o sndl main.go 
```

If you would like to utilitze [go-task](https://taskfile.dev) as a
makefile replacement then run the following

```bash
$ task build
```

**We are working on a devcontainer for easier onboarding.**

# ADR

The adr folder contains all the accepted, proposed, rejected, and 
deprecated adr's (architectural design records) for this project.
Any person may refer to this directory for information regarding
decisions and may use this folder to propose new changes.

Proposed adr's must follow the convention of all other adr's.

# Assets

The asset folder contains any asset that provides no code or 
architectural benefit but does influence or otherwise improve the
project.

# Config

This project doesnt include a configuration file to read from
but we do utilize a file in our mirror subcommand. The format of that
file is:

```text
SEARCH TERM:ID
SEARCH TERM:ID
```

This allows us to easily parse and download the specified files. We use a specific 
naming scheme for the files to ensure we dont repeat downloads. Unfortunately we don't
have plans right now for implementing a mechanism to remember files other than to check
that they exist.

There is an [adr](adr/009-avoid-repeat-downloads.md) to define methods to avoid this problem
but this wont be implemented until a v2 release.

# Communication

We utilitze discord to communicate, if you would like to join our discord server then please reach out
to @Woodenikki.

---

Future adrs:

- usage docs


