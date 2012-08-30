#!/bin/bash

for g in `find . -name "*.go"`
do
  echo "gofmt -tabs=false -tabwidth=2 $g > $g.fix"
  echo "mv $g.fix $g"
done
