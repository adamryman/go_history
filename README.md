# go_history [![Build Status](https://travis-ci.org/adamryman/go_history.svg?branch=master)](https://travis-ci.org/adamryman/go_history)

Basically does a `cat $HOME/.bash_history | sort | uniq > .bash_history.go_history`, but preserves order. If you know what I mean.

# Install

```
go get github.com/adamryman/go_history
```

# Usage

```
# To replace $HOME/.bash_history with all leading dupicates removed
go_history -r
```

## .bashrc

```
# Avoid duplicates
export HISTCONTROL=erasedups
export PROMPT_COMMAND="history -a"
if which go_history >/dev/null; then
	export PROMPT_COMMAND="history -a;go_history -r"
fi
```
