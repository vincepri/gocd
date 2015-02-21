Gocd
====
Quick change directory in GOPATH


Build & Install
---------------
```sh
go get -u github.com/vinceprignano/gocd
echo "source $(gocd -shellinit)" >> ~/.bash_profile
```

Usage
-----
`gocd [directory]`
Where directory can be any directory inside in your GOPATH.

Using `gocd` with no arguments will change the current working directory in `$GOPATH/src`.

## Example
If you have gocd in your `$GOPATH/src` folder. You will be able to cd into the project directory by doing:
```sh
$ gocd gocd
-> Will change directory in $GOPATH/src/github.com/vinceprignano/gocd
```
In case of multiple directories with the same name you will be required to choose one from a list using `gocd directory:index`.
```sh
$ gocd example
Found multiple folders with this name:
	0: /Users/vincenzo/src/github.com/b2aio/typhon/example
	1: /Users/vincenzo/src/github.com/kylelemons/go-gypsy/example
	2: /Users/vincenzo/src/github.com/qur/withmock/example
	5: /Users/vincenzo/src/github.com/wsxiaoys/terminal/example
```