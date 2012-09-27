package urlness

import (
//"sort"
//"fmt"
//"math/rand"
)

// findOptimalSet Given a set of samples, reduce it to the point it contains
// only the maximum number of unrelated samples.
// We also need the phi scores per each samle (m) and
// the phi score we want to use as a cut off
func findOptimalSet(set map[string]bool, m Samples, phi float64) map[string]bool {
  // Per all samples in set, group them by the number of samples they relate to
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

  max := 0 // what's the group with more related animals?
  for num, _ := range histRelated {
    if num > max {
      max = num
    }
  }

  /*
       // What's the sample I have to remove to get to the bigger set ?
     	bestName, bestLen := "", 0
     	for s, _ := range set {
     		delete(set, s)
     		if willGetMe := len(findOptimalSet(set, m, phi)); willGetMe > bestLen {
     			bestName, bestLen = s, willGetMe
     			//fmt.Println("  ", set, " -> (", bestLen, "," , bestName, ")")
     		}
     		set[s] = true // Put it back, and try next
     	}
     	delete(set, bestName)
  */

  delete(set, histRelated[max][0])
  //delete(set, histRelated[max][rand.Intn(len(histRelated[max]))])

  // .. and make a recursive call with the reduced set
  return findOptimalSet(set, m, phi)
}
