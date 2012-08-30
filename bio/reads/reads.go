package reads

import "strings"

type Read struct {
	Id, Seq string
}

func NewRead(line string) *Read {
	s := strings.Split(line, "\t")
	id, seq := s[0], s[9]
	return &Read{
		id,
		seq,
	}
}
