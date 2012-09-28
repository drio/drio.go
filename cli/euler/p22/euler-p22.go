package main

import (
  "fmt"
  "github.com/drio/drio.go/common/files"
  "io"
  "log"
  "sort"
  "strings"
)

// Problem 22
//
//    19 July 2002
//
//    Using [7]names.txt (right click and 'Save Link/Target As...'), a 46K
//    text file containing over five-thousand first names, begin by sorting
//    it into alphabetical order. Then working out the alphabetical value for
//    each name, multiply this value by its alphabetical position in the list
//    to obtain a name score.
//
//    For example, when the list is sorted into alphabetical order, COLIN,
//    which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the
//    list. So, COLIN would obtain a score of 938x53=49714.
//
//    What is the total of all the name scores in the file?
func main() {
  names := sliceData(loadData("names.txt"))
  fmt.Println(processData(names))
}

func processData(names []string) int {
  total := 0
  for i, n := range names {
    sum := 0
    for _, c := range n {
      if c >= 'A' && c <= 'Z' {
        sum += int(c) - 'A' + 1
      }
    }
    total += sum * (i + 1)
  }
  return total
}

func sliceData(lineData []byte) sort.StringSlice {
  var names sort.StringSlice
  for _, n := range strings.Split(string(lineData), ",") {
    names = append(names, n)
  }
  names.Sort()
  return names
}

func loadData(fn string) []byte {
  f, r := files.Xopen(fn)
  defer f.Close()

  var lineData []byte
  buff := make([]byte, 100000)
  var err error
  n := 1
  for n != 0 {
    if n, err = r.Read(buff); err == nil {
      lineData = append(lineData, buff...)
    } else {
      if err == io.EOF {
        break
      } else {
        log.Fatal("Problems reading file", err)
      }
    }
  }
  return lineData
}
