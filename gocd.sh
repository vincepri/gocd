#!/bin/bash

gocd() {
    cd `tree --noreport -f -L 3 -d -i $GOPATH/src/ | grep $1`
}

_gocd() {
    cur=${COMP_WORDS[COMP_CWORD]}
    use=`tree --noreport -f -L 3 -d -i $GOPATH/src/ | sed -e "s#$GOPATH/src/##g" | tr / '\t' | awk '{print $(NF-1),$NF; print $(NF)}' | tr ' ' /`
    COMPREPLY=(`compgen -W "$use" -- $cur`)
}

complete -o default -o nospace -F _gocd gocd
