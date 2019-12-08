# east-to-west
A tool to convert between Indian music notation and Western notation

### Background

In indian notation, music notes are as below:

```
sa re ga ma pa dha ni sa
```

Where note `sa` is not fixed i.e., it can begin with any note (like `a` or `c`). Difference between the notes is similar to major scale. 

To convert, from indian notation to western notation, you would need to know
what your `sa` is. And then it's simple major scale from that note.

### Why

There are awesome tools like [alda](https://github.com/alda-lang/alda). That can
help you compose, but they expect western notation.

### Setup

Clone the repo and build it using `go build`:

```
$ git clone https://github.com/ms-choudhary/east-to-west.git
$ cd east-to-west && go build
```

### Usage

```
$ ./east-to-west --help
Usage: ./east-to-west [options] filename

Options:
  -from string
        indian or western (default "indian")
  -sa string
        which western note is sa (default "a")
```
