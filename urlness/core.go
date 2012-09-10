// Package urlness help us to find unrelated individuals in a population
// ds.go holds all the datastructure
// core.go holds the main algorithms
// optimal_sets.go contains code to find the optimal subset of individuals
// that are unrelated
package urlness

import (
  "encoding/csv"
  "fmt"
  "io"
  "log"
  "math/rand"
  "strings"
  "time"
)

// action is a signature for an action
// what to do when we have a csv line ready to process
type action func(m Samples, s_line, s_header []string) error

// processFile iterates over a csv file and adds the data
// to the main datastructure
func processFile(m Samples, rf io.Reader, a action) {
  csv := csv.NewReader(rf)
  header := true
  var s_header []string
  for {
    s_line, err := csv.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal("Error reading from csv file in ulrness.processFile() ", " err: ", err)
    }
    if header {
      s_header = s_line
      header = false
    } else {
      err = a(m, s_line, s_header)
      if err != nil {
        log.Fatal("Error processing line in urlness.processFile()", "error: ", err)
      }
    }
  }
}

// processData iterate over the list of samples in the main data structure (Samples) and
// create the matrix of coefficients and list of unrelated samples.
// This last one only if we are providing a Phi score.
func processData(m Samples, o Options, listOnly map[string]bool) (string, string) {
  ids := m.Ids()
  unrelatedList := []string{} // List of animals that are unrelated based on phi coef.
  phiEnabled := o.PhiFilter != 0
  onlyEnabled := false
  if len(listOnly) > 0 {
    onlyEnabled = true
  }

  // Let's prepare the header for the matrix
  sHeader := append([]string{"sample"}, ids...)
  matrix := append([]string{}, strings.Join(sHeader, ","))

  for _, i := range ids {
    line := append([]string{}, i)
    isUnrelated := true
    for _, j := range ids {
      line = append(line, fmt.Sprintf("%f", m[i].Phis[j]))
      if phiEnabled && i != j { // phi coef enabled
        if onlyEnabled {
          if listOnly[j] { // only check for phi threshold for the ids in unrelatedList
            isUnrelated = isUnrelated && (m[i].Phis[j] <= o.PhiFilter)
          }
        } else { // Check phi coef against all ids
          isUnrelated = isUnrelated && (m[i].Phis[j] <= o.PhiFilter)
        }
      }
    }
    matrix = append(matrix, strings.Join(line, ","))
    if phiEnabled && isUnrelated { // Save as unrelated if passes phi coeff against the others
      unrelatedList = append(unrelatedList, i)
    }
  }

  //return toBytes(matrix, unrelatedList, o, m)
  return strings.Join(matrix, "\n"), strings.Join(unrelatedList, "\n")
}

// Options holds the info we use as entry point
// to this package
type Options struct {
  KS, Sex, Phe, Only io.Reader // Data for all the files
  PhiFilter          float64   // What's the filter phi value
}

// ComputeOptimal is the entry point to use the optimal approach
// for finding the best subset of animals in a population set
func ComputeOptimal(o Options) string {
  var m Samples // Map of samples and its relation ships
  m.Init()      // Prepare data structure for data

  // Load Kindship data
  processFile(m, o.KS, func(m Samples, s_line, s_header []string) error {
    return m.AddRelation(s_line)
  })

  // Prepare a set with all the samples
  set := make(map[string]bool)
  for _, e := range m.Ids() {
    set[e] = true
  }

  // Prepare the seed for random
  rand.Seed(time.Now().UTC().UnixNano())

  // Call the optimal routine and iterate over the elements in the results
  final := []string{}
  for e, _ := range findOptimalSet(set, m, o.PhiFilter) {
    final = append(final, e)
  }

  return strings.Join(final, "\n")
}

// Compute is one of the two entry points to the package.
// It computes the urlness and returns the matrix and the list (if possible)
// It retuns the data in matrix and the list as a slice of bytes
// This is the basic approach, per each individual, make sure the relateness
// against all the other individuals is below or equal the phi score provided
// by the user
func Compute(o Options) (string, string) {
  var m Samples // Map of samples and its relation ships
  m.Init()      // Prepare data structure for data

  if o.Sex != nil { // Load sex data
    processFile(m, o.Sex, func(m Samples, s_line, s_header []string) error {
      return m.AddSex(s_line)
    })
  }

  if o.Phe != nil { // Load phenotype data
    processFile(m, o.Phe, func(m Samples, s_line, s_header []string) error {
      return m.AddPheno(s_line, s_header)
    })
  }

  // Load list of ids to use for phi comparison
  listOnly := make(map[string]bool)
  if o.Only != nil {
    processFile(m, o.Only, func(m Samples, s_line, s_header []string) error {
      listOnly[strings.Trim(s_line[0], " ")] = true
      return nil
    })
  }

  // Load Kindship data
  processFile(m, o.KS, func(m Samples, s_line, s_header []string) error {
    return m.AddRelation(s_line)
  })
  //fmt.Println(m["sample1"].Phis["sample2"])

  matrix, list := processData(m, o, listOnly)
  return matrix, list
}

// GenRandomKindShip will create a random kinship file (string)
func GenRandomKindShip(nSamples int) string {
  // Prepare the seed for random
  rand.Seed(time.Now().UTC().UnixNano())

  ks := "ego1,ego2,phi\n" // header
  for i := 0; i < nSamples; i++ {
    for j := i; j < nSamples; j++ {
      if i == j {
        ks += fmt.Sprintf("s_%d,s_%d,%f\n", i, j, 0.0)
      } else {
        ks += fmt.Sprintf("s_%d,s_%d,%f\n", i, j, float32(rand.Int31n(200))/float32(1000))
      }
    }
  }

  return ks
}
