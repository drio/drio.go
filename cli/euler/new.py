#!/usr/bin/env python
#
import os, glob, re, sys

go = '''package main

import "fmt"

__DESC__
func main() {
}'''

def getDesc(n):
  start = False
  desc = []
  cmd = "lynx -dump http://projecteuler.net/problem=" + str(n)
  for l in os.popen(cmd).readlines():
    if re.search('^Problem\s', l): start = True
    if re.search('Copyright Information', l): start = False
    if start: desc.append("// " + l.rstrip())
  return '\n'.join(desc)

n = []
for f in glob.glob('p*'):
  n.append(int(re.search("\d+", f).group(0)))
nextNum = str(sorted(n)[-1] + 1)
newDir = "./p" + nextNum
newFile = newDir + "/euler-p" + nextNum + ".go"

os.mkdir(newDir)
f = open(newFile, "w")
f.write(go.replace('__DESC__', getDesc(nextNum)))
f.close()
