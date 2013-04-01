package urlness

// findOptimalSet Given a set of samples, reduce it to the point it contains
// only the maximum number of unrelated samples.
// We also need the phi scores per each samle (m) and
// the phi score we want to use as a cut off
func findOptimalSet(set map[string]bool,
  m Samples,
  phi float64,
  forceList map[string]bool) map[string]bool {
  // Per all samples in set, group them by the number of samples they relate to
  // 0 -> [ "sample1", "sample7" ...] ; sample1 sample7 ... related to 0 animals
  // 1 -> [ "sample2" ]
  // ..
  histRelated := make(map[int][]string)
  for e, _ := range set {
    nOfRelated := 0
    for o, _ := range set {
      // Check if e and o are related only if not in the forcelist
      if !forceList[o] && m[e].Phis[o] > phi {
        nOfRelated++
      }
    }
    histRelated[nOfRelated] = append(histRelated[nOfRelated], e)
  }

  if len(histRelated[0]) == len(set) { // All elements in set are unrelated
    for e, _ := range forceList { // Add back the samples the user wants in the list for sure
      set[e] = true
    }
    return set
  }

  max := 0 // what's the group with more related animals?
  for num, _ := range histRelated {
    if num > max {
      max = num
    }
  }
  // remove the first one in the slice.
  delete(set, histRelated[max][0])

  // .. and make a recursive call with the reduced set
  return findOptimalSet(set, m, phi, forceList)
}
