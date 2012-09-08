//Package files contains some helpers to work with files.
package files

import (
  "bufio"
  "bytes"
  "compress/bzip2"
  "compress/gzip"
  "fmt"
  "io"
  "os"
  "strings"
)

// IterLines returns an iterator (channel) so we can
// iterate over the lines in a file
func IterLines(r *bufio.Reader) chan string {
  ch := make(chan string, 10000)
  go func() {
    for {
      if b_line, _, err := r.ReadLine(); err != nil {
        if err == io.EOF {
          close(ch)
          break
        } else {
          fmt.Println("Problems iterating over lines:")
          fmt.Println(err)
          os.Exit(1)
        }
      } else {
        line := bytes.NewBuffer(b_line).String()
        ch <- line
      }
    }
  }()

  return ch
}

// Xopen opens a file and retuns a Reader.
// It works with compressed files by inspecting the
// file extension.
// Note I originally read this code in a lh3 snippet:
// https://github.com/lh3/misc/blob/master/klib.lua#L143
func Xopen(fName string) (*os.File, *bufio.Reader) {
  var (
    fos *os.File
    err error
  )

  // Open the file (unless we are reading from stdin
  if fName != "-" {
    fos, err = os.Open(fName)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Can't open %s: error: %s\n", fName, err)
      os.Exit(1)
    }
  } else {
    fos = os.Stdin
  }

  // Deal with the compression of the file
  var newReader io.Reader
  err = nil
  if strings.HasSuffix(fName, ".gz") {
    newReader, err = gzip.NewReader(fos)
  } else if strings.HasSuffix(fName, ".bz2") {
    newReader = bzip2.NewReader(fos)
  } else {
    newReader = fos // No compression, stdin
  }

  if err != nil {
    fmt.Fprintf(os.Stderr,
      "%s file is not in the correct format: error: %s\n", fName, err)
    os.Exit(1)
  }

  // We want a buffered stream to increase performance
  return fos, bufio.NewReader(newReader)
}
