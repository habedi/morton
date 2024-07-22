#!/bin/bash

# You need to have graphviz installed to run this script
# sudo apt-get install graphviz # for Debian-based systems

# Make figures from .dot files
for f in *.dot; do
    dot -Tsvg $f -o ${f%.dot}.svg
done
