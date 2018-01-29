# go_history [![Build Status](https://travis-ci.org/adamryman/go_history.svg?branch=master)](https://travis-ci.org/adamryman/go_history)

Basically does a `cat $HOME/.bash_history | sort | uniq > .bash_history.go_history`, but preserves order. If you know what I mean.

Fun little thing, might try to optimize for fun in the future.

# Install

```
go get github.com/adamryman/go_history
```

# Usage

```
go_history -f $HOME/.bash_history && mv .bash_history.go_history .bash_history
```
