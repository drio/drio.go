package urlness

import (
//"sort"
//"fmt"
//"math/rand"
)

// pair Tuple that holds a sample name and the number of samples that
// are related to it
type pair struct {
  name string
  num  int
}

// pairList holds a list of pairs. And implements sort.Interface so
// we can sort the list.
type pairList []pair                  // This is really a set of pairs
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].num < p[j].num }

// findOptimalSet Given a set of samples, reduce it to the point it contains
// only the maximum number of unrelated samples.
// We also need the phi scores per each samle (m) and
// the phi score we want to use as a cut off
func findOptimalSet(set map[string]bool, m Samples, phi float64) map[string]bool {
  // Per all samples in set, I want to group animals together that have the
  // same number of related animals
  histRelated := make(map[int][]string)
  for e, _ := range set {
    nOfRelated := 0
    for o, _ := range set {
      if m[e].Phis[o] > phi { // sample e is related to sample o
        nOfRelated++
      }
    }
    histRelated[nOfRelated] = append(histRelated[nOfRelated], e)
  }

  if len(histRelated[0]) == len(set) { // All elements in set are unrelated
    return set
  }

  max := 0 // key in histRelated for the animals with the highest number of relateness
  for num, _ := range histRelated {
    if num > max {
      max = num
    }
  }

  /*
  	worseName, worseLen := "", len(set)
  	for _, s := range histRelated[max] {
  		delete(set, s)
  		if lenCurrent := len(findOptimalSet(set, m, phi)); lenCurrent < worseLen {
  			worseName, worseLen = s, lenCurrent
  		}
  		set[s] = true // Put it back, and try next
  	}
  	delete(set, worseName)
  */

  delete(set, histRelated[max][0])
  //delete(set, histRelated[max][rand.Intn(len(histRelated[max]))])

  // .. and make a recursive call with the reduced set
  return findOptimalSet(set, m, phi)
}
