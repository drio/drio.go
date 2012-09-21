package main

import "fmt"

// Problem 17
//    If the numbers 1 to 5 are written out in words: one, two, three, four,
//    five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
//    If all the numbers from 1 to 1000 (one thousand) inclusive were written
//    out in words, how many letters would be used?
//
//    NOTE: Do not count spaces or hyphens. For example, 342 (three hundred
//    and forty-two) contains 23 letters and 115 (one hundred and fifteen)
//    contains 20 letters. The use of "and" when writing out numbers is in
//    compliance with British usage.

var one = map[int]string{
  0: "",
  1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
  6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
  11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
  15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen",
  19: "nineteen",
}

var two = map[int]string{
  0: "",
  1: "ten",
  2: "twenty", 3: "thirty", 4: "forty", 5: "fifty",
  6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety",
}

func main() {
  result := ""
  for i := 1; i < 1000; i++ {
    if i < 20 {
      result += one[i]
    }
    if i >= 20 && i < 100 {
      f, s := i/10, i%10 // first and second digits
      result = result + two[f] + one[s]
    }
    if i >= 100 && i < 1000 {
      f, s, t := i/100, (i%100)/10, (i%100)%10
      and := ""
      twoDigits := i % 100
      if twoDigits != 0 {
        and = "and"
      }
      if twoDigits > 10 && twoDigits < 20 {
        result = result + one[f] + "hundred" + and + one[twoDigits]
      } else {
        result = result + one[f] + "hundred" + and + two[s] + one[t]
      }
    }
  }
  fmt.Println(len(result + "onethousand"))
}
