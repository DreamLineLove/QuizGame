# Welcome to QuizGame!

<img src="https://camo.githubusercontent.com/94761affed6454156a526a0fcab454ed4a432d9472087a9d330598a38ffe56cd/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67" width="175px" />

A program written in Go for running timed quizes via the command line.

## Usage

- First, clone the repository to your local machine:

```
    $ git clone https://github.com/DreamLineLove/QuizGame.git
```

- Run the binary or build it yourself (<a href="https://go.dev/learn/" target="_blank">Go toolchain required</a>):

```go
    // This will run the binary with default flags
    $ ./QuizGame

    // Or build the binary yourself
    $ go build .
```

## Flags
- By default -timelimit is 30 (seconds) and -filename is "problems.csv".
- However, you can set the flags anyway you like:
```go
    // The time limit is 60 seconds and the data source is
    // set as "mydata.csv" from the same directory
    $ ./QuizGame -timelimit 60 -filename mydata.csv
```

## To Add
- Tracking time for each question
- Turn-based 2 player mode
- Simultaneous 2 player mode
