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

// Return an iterator to iterate over lines
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

// Smart way (I think) to open a file. It can be compressed. Use "-" to read from stdin
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
