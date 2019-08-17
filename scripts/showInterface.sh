#!/usr/bin/sh

ifconfig -a | sed 's/[ \t].*//;/^$/d'
