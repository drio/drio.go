#!/usr/bin/env ruby

# TODO: the problem may not exist.
def getDesc(n)
  start = false
  desc = []
  `lynx -dump http://projecteuler.net/problem=#{n}`.each_line do |l|
    start = true  if l =~ /^Problem\s/
    start = false if l =~ /Copyright Information$/
    desc << "// " + l.chomp if start
  end
  desc.join("\n")
end

# Main
#
n = []
Dir["p*"].each {|d| n << d.match(/(\d+)$/)[1]}
nextNum = n.sort[-1].to_i + 1
dir = "./p#{nextNum}"

`mkdir #{dir}`
File.open(dir + "/euler-p#{nextNum}.go", "w") do |f|
  f.puts DATA.read.gsub(/__DESC__/, getDesc(nextNum))
end

__END__
package main

import "fmt"

__DESC__
func main() {
}
