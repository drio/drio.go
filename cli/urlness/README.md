urlness
=======

This is a tool for finding unrelateness giving a phi score for a set of animals.
Currently we have two methods. The first one is very naive: We go over
the phi matrix and drop animals that are related (based on the phi score).
But we can do better than that. That's with the optimal approach does. Giving
two animals, it computes what would happen if we drop each of the animals and
then picks the one that would leads us to the biggest final set of unrelated
animals.
