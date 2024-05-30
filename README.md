# Welcome to QuizGame!

<img src="https://cdn.worldvectorlogo.com/logos/gopher.svg" width="175px" />

A program written in Go for running timed quizes via the command line.

## Usage

- First, clone the repository to your local machine:

```
    $ git clone https://github.com/DreamLineLove/QuizGame.git
```

- Run the binary or build it yourself (<a href="https://go.dev/learn/" target="_blank">Go toolchain required</a>):

```go
    // The following each runs the binary with default flags
    // Run the one for your operating system
    $ ./bin/apple-silicon-mac
    $ ./bin/intel-mac
    $ ./bin/windows
    $ ./bin/linux

    // Or build the binary yourself
    $ go build .
```

## Flags
- By default -timelimit is 30 (seconds) and -path is "./problems.csv".
- However, you can set the flags to anything you like:
```go
    // Here, the time limit is 60 seconds and the file path
    // is set as "another.csv" from the same directory.
    // It will be run on an M-series CPU mac (M1, M2, chips)
    $ ./bin/apple-silicon-mac -timelimit 60 -path another.csv
```
