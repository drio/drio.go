package urlness

import (
	"testing"
	"strings"
)

func TestAddRelation(t *testing.T) {
	var s urlness_ds.Samples
	s.Init()

	s_line := strings.Split("id1,id2,0.5", ",")
	if s.AddRelation(s_line); s["id1"].Phis["id2"] != 0.5 {
		v := s["id1"].Phis["id2"]
		t.Errorf("Failed adding a relation for sample id1 against id2 %d", v)
	}
	// We have to add the relationship the other way around !!
	if s.AddRelation(s_line); s["id2"].Phis["id1"] != 0.5 {
		v := s["id2"].Phis["id1"]
		t.Errorf("Failed adding a relation for sample id2 against id1 %d", v)
	}

	s_line = strings.Split("id1,id3,2", ",")
	if s.AddRelation(s_line); s["id1"].Phis["id3"] != 2 {
		v := s["id1"].Phis["id3"]
		t.Errorf("Failed adding a relation for sample id1 against id3 %d", v)
	}

	s_line = strings.Split("id1,id3,20", ",")
	if s.AddRelation(s_line); s["id1"].Phis["id3"] != 20 {
		v := s["id1"].Phis["id3"]
		t.Errorf("Failed re-adding relation for sample id1 against id3 %d", v)
	}

	s_line = strings.Split("id1,id3,XXX", ",")
	if e := s.AddRelation(s_line); e == nil {
		t.Errorf("I should have returned an error when processing line with invalid phi")
	}

}

func TestIds(t *testing.T) {
	var s urlness_ds.Samples
	s.Init()

	s.AddRelation(strings.Split("id1,id2,0.1", ","))
	s.AddRelation(strings.Split("id1,id3,0.2", ","))
	s.AddRelation(strings.Split("id2,id1,0.3", ","))
	s.AddRelation(strings.Split("id2,id3,0.4", ","))
	s.AddRelation(strings.Split("id3,id1,0.5", ","))
	s.AddRelation(strings.Split("id3,id2,0.6", ","))

	expectedIds := map[string]int { "id1":0, "id2":0, "id3":0 }
	ids := s.Ids()
	if len(ids) != len(expectedIds) {
		t.Errorf("Wrong number of ids. %d != %d", len(expectedIds), len(ids))
 	}

	for _, v := range ids {
		if _, present := expectedIds[v]; present == false {
			t.Errorf("I cannot find %s in expectedIds", v)
		}
 	}

}

func TestAddSex(t *testing.T) {
	var s urlness_ds.Samples
	s.Init()

	s_line := strings.Split("id1, M", ",")
	if err := s.AddSex(s_line); err != nil {
		t.Errorf(err.Error())
	}
	if s["id1"].Sex != "M" {
		t.Errorf("Sex not properly set. Expecting M, go (%s)", s["id1"].Sex)
	}

	s_line = strings.Split("id1, X", ",")
	if err := s.AddSex(s_line); err == nil {
		t.Errorf("I am not properly detecting when I use a wrong character for the sex.")
	}
}

func TestAddPheno(t *testing.T) {
	var s urlness_ds.Samples
	s.Init()

	s_header := strings.Split("sample, anxiety, obesity", ",")

	s_line := strings.Split("id1, 12, 5", ",")
	if err := s.AddPheno(s_line, s_header); err != nil {
		t.Errorf(err.Error())
	}
	if r := s["id1"].PhenoType["anxiety"]; r != 12 {
		t.Errorf("Phenot. anxiety not properly set. Expecting ", r)
	}
	if r := s["id1"].PhenoType["obesity"]; r != 5 {
		t.Errorf("Phenot. anxiety not properly set. Expecting 5, got ", r)
	}
}

func TestListPhenos(t *testing.T) {
	var s urlness_ds.Samples
	s.Init()

	if s := s.ListPhenoTypes(); len(s) != 0 {
		t.Errorf("not returning an empty slice when we haven't loaded any sample")
	}

	phenos := map[string]bool {
		"anxiety" : true,
		"obesity" : true,
	}
	s_header := strings.Split("sample, anxiety, obesity", ",")
	s_line   := strings.Split("id1, 12, 5", ",")
	s.AddPheno(s_line, s_header)

	slist := s.ListPhenoTypes()

	if len(slist) != 2 {
		t.Errorf("Wrong number of phenotypes returned")
	}

	for _, v := range slist {
		if _, present := phenos[v]; present == false {
			t.Errorf("Phenotype ", v, " should not be there")
		}
	}

}

