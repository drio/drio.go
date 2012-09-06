package urlness

import (
  "fmt"
  "strconv"
  "strings"
)

type DsError struct {
  What string
}

func (e *DsError) Error() string {
  return fmt.Sprintf("Error: %s", e.What)
}

type Samples map[string]*Relations

type Relations struct {
  Phis      map[string]float64
  Sex       string
  PhenoType map[string]float64
}

func (s *Samples) Init() {
  *s = make(map[string]*Relations)
}

func (s *Samples) add(id string) {
  r := new(Relations)
  r.Phis = make(map[string]float64)
  r.PhenoType = make(map[string]float64)
  (*s)[id] = r
}

func (s *Samples) AddRelation(s_line []string) error {
  one, two, phi := s_line[0], s_line[1], s_line[2]

  if _, present := (*s)[one]; present == false {
    s.add(one)
  }
  if _, present := (*s)[two]; present == false {
    s.add(two)
  }

  if f, err := strconv.ParseFloat(phi, 32); err != nil {
    return err
  } else {
    (*s)[one].Phis[two] = f
    (*s)[two].Phis[one] = f
  }

  return nil
}

func (s *Samples) Ids() []string {
  var values []string
  for k, _ := range *s {
    values = append(values, k)
  }
  return values
}

func (s *Samples) AddSex(s_line []string) error {
  id, sex := strings.Trim(s_line[0], " "), strings.Trim(s_line[1], " ")

  if _, present := (*s)[id]; present == false {
    s.add(id)
  }

  if sex != "M" && sex != "F" {
    return &DsError{fmt.Sprintf("(%s) is not valid sex type (M|F) only.", sex)}
  }

  (*s)[id].Sex = sex

  return nil
}

func (s *Samples) AddPheno(s_line, s_header []string) error {
  id := strings.Trim(s_line[0], " ")

  if len(s_line) != len(s_header) {
    return &DsError{fmt.Sprintf("Line doesn't match header size for phenotype file.")}
  }

  if _, present := (*s)[id]; present == false {
    s.add(id)
  }

  for i, v := range s_header {
    if i != 0 {
      phe := strings.Trim(v, " ")
      if fVal, err := strconv.ParseFloat(strings.Trim(s_line[i], " "), 32); err != nil {
        return err
      } else {
        (*s)[id].PhenoType[phe] = fVal
      }
    }
  }

  return nil
}

func (s *Samples) ListPhenoTypes() []string {
  if len(*s) == 0 {
    return []string{}
  }

  var slicePhenoTypes []string
  for id, _ := range *s {
    for phe, _ := range (*s)[id].PhenoType {
      slicePhenoTypes = append(slicePhenoTypes, phe)
    }
    break
  }

  return slicePhenoTypes
}
