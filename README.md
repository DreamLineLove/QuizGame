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
    // This will run the binary with default flags
    $ ./bin/mac-apple-silicon
    $ ./bin/mac-intel
    $ ./bin/windows
    $ ./bin/linux

    // Or build the binary for your operating system
    $ go build .
```

## Flags
- By default -timelimit is 30 (seconds) and -path is "./problems.csv".
- However, you can set the flags anyway you like:
```go
    // The time limit is 60 seconds and the data source is set as
    // "mydata.csv" file from the directory at which the command was run
    $ ./QuizGame -timelimit 60 -path mydata.csv
```
