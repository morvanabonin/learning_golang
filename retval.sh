#! /bin/bash

# Create and define de var retval
retval=$?

function F1() {
    retval='yellow'
}

retval='blue'
echo $retval
F1
echo $retval
