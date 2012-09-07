package urlness

import ("sort")

// Tuple that holds a sample name and the number of samples that
// are related to it
type pair struct {
	name string
	num int
}

type pairList []pair // This is really a set of pairs
func (p pairList) Swap (i, j int) { p[i], p[j] = p[j], p[i]}
func (p pairList) Len() int { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].num < p[j].num }

// Given a set of samples, reduce it to the point it contains
// only the maximum number of unrelated samples.
// We obviously also need the phi scores per each samle (m) and
// the phi score we want to use as a cut off
func findOptimalSet(set map[string]bool, m Samples, phi float64) map[string]bool {

	// Find the number of related samples for each of the samples in set
	nOfRelated := make(map[string]int, len(set))
	for e, _ := range set {
		for o, _ := range set {
			if m[e].Phis[o] > phi { // sample e is related to sample o
				nOfRelated[e]++
			}
		}
	}

	// convert the nOfRelated to a list of pairs and sort the
	// list of pairs by value
	pl := make(pairList, len(set))
	i, sum := 0, 0
	for e, n := range nOfRelated {
		pl[i].name, pl[i].num = e, n
		i++
		sum+=n
	}
	allZero := sum == 0 // All the elements in set are unrelated
	sort.Sort(pl)

	if allZero { // We found a subset where all the samples are unrelated
		return set
	}

	// We have samples in set that are related
	// Delete from the set the sample with more related individuals
	// If there is more than one, pick the one in the end of the sorted pl
	delete(set, pl[len(set)-1].name)
	// .. and make a recursive call with the reduced set
	return findOptimalSet(set, m, phi)
}

