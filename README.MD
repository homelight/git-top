# git top

`git top` is an analytics tool for git log.

# Prerequisites

1. [Go](https://golang.org/doc/install)
1. [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

# Usage

```
> git top --help

NAME:
   git top - view the top analytics for a git log.

USAGE:
   git top [global options] command [command options] [arguments...]

COMMANDS:
   files  view the most changed files for a time period.

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## files

The `files` command gives you visibility into the most changed files for a time period.

```
> git top files --help

NAME:
   git top files - view the most changed files for a time period.

USAGE:
   git top files [command options] [arguments...]

OPTIONS:
   --after value   date to start lookback. Format: YYYY-MM-DD (default: two weeks ago)
   --before value  date to end lookback. Format: YYYY-MM-DD (default: now)
   --places value  number of files returned (default: 10)
   --help, -h      show help (default: false)
```

### Example

```
> git top files --places 5 --after="2020-01-01"

 Place |  Count | File
     1 |    147 | db/db_test.go
     2 |    119 | handlers/routes.go
     3 |    119 | handlers/loan_applications.go
     4 |     91 | underwriting/loan_application.go
     5 |     87 | handlers/tasking.go
```

# Installation

```
git clone git@github.com:homelight/git-top.git
go install
```

## Troubleshooting

1. If the binary is not found when you run `git top` make sure your $GOPATH/bin is part of your path.
