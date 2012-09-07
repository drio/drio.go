package main

import (
  "flag"
  "fmt"
  "github.com/drio/drio.go/urlness"
  "io"
  "log"
  "os"
  "strings"
)

var usage = []string{
  //"urlness, a tool to find unrelated individuals.",
  "urlness (v0.1)",
  "Usage: urlness -ks <kinship_file> [optional params]",
  "",
  "Madatory params: ",
  "  -ks   : csv file containing phi coefficients per each relation.",
  "",
  "Optional params: ",
  "  -sex     : csv file containing samples' sex",
  "  -phe     : csv file containing the gender of the samples",
  "  -only    : only use this sample when filtering by phi score",
  "  -phi     : maximum phi score allowed between relations",
  "  -optimal : find the biggest subset of unrelated samples ",
  "",
  "Output:",
  "  stdin : csv matrix of phi coefficients for ALL the samples.",
  "  stdout: If -phi used: csv list of samples that pass the phi score",
  "",
  "Examples:",
  "  $ urlness -ks data/kinship.csv > matrix.csv",
  "  $ urlness -ks data/kinship.csv -phi 0.5 > matrix.csv 2> l.csv",
  "  $ urlness -ks data/kinship.csv -phi 0.5 -sex data/441-Gender.csv -phe data/441-Pheno.csv > m.csv 2> l.csv",
  "  $ urlness -ks data/kinship.csv -phi 0.5 -optimal > list.optimal.csv",
}

// For processing the input parameters from the user.
type options struct {
  ksFname, sexFname, pheFname, onlyFname *string
  phiFilter                              *float64
  optimal                                *bool
}

func parseArgs() *options {
  o := new(options)
  o.ksFname = flag.String("ks", "", "Kinship csv file.")
  o.sexFname = flag.String("sex", "", "Gender csv file.")
  o.pheFname = flag.String("phe", "", "Gender csv file.")
  o.onlyFname = flag.String("only", "", "Gender csv file.")
  o.phiFilter = flag.Float64("phi", 0, "Maximum phi coefficient allowed.")
  o.optimal = flag.Bool("optimal", false, "Enable optimal.")
  flag.Parse()

  // TODO: check that the file exists!!!!!!!!!!!!!!!!!!!!
  err := false
  if len(flag.Args()) != 0 {
    fmt.Fprintln(os.Stderr, "Invalid parameter.")
    err = true
  }

  if *o.ksFname == "" {
    fmt.Fprintln(os.Stderr, "ERROR: Need kinship file")
    err = true
  }

  if *o.onlyFname != "" && *o.phiFilter == 0 {
    fmt.Fprintln(os.Stderr, "ERROR: -only requires -phi param")
    err = true
  }

  if *o.optimal && *o.phiFilter == 0 {
    fmt.Fprintln(os.Stderr, "ERROR: -optimal requires -phi param")
    err = true
  }

  if err {
    fmt.Println(strings.Join(usage, "\n"))
    os.Exit(0)
  }

  return o
}

func main() {
  o := parseArgs()
  inputData := new(urlness.Options)

  fNamesToFiles := map[*string]*io.Reader{
    o.ksFname:   &inputData.KS,
    o.sexFname:  &inputData.Sex,
    o.pheFname:  &inputData.Phe,
    o.onlyFname: &inputData.Only,
  }

  for path, reader := range fNamesToFiles {
    if *path != "" {
      if file, err := os.Open(*path); err != nil {
        log.Fatal("Error opening csv file: ", path, " err: ", err)
      } else {
        *reader = file
        defer file.Close()
      }
    }
  }

  // We know the Phi param has to be there for sure
  inputData.PhiFilter = *o.phiFilter

  if *o.optimal {
    fmt.Println(urlness.ComputeOptimal(*inputData))
  } else {
    m, l := urlness.Compute(*inputData) // matrix and list of unrelated samples
    fmt.Println(m)
    if *o.phiFilter != 0 {
      fmt.Fprintln(os.Stderr, l)
    }
  }
}
