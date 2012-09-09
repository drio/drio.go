#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "cd $DIR"
for g in `find . -name "*.go"`
do
  echo "gofmt -tabs=false -tabwidth=2 $g > $g.fix"
  echo "mv $g.fix $g"
done
echo "cd - > /dev/null"
