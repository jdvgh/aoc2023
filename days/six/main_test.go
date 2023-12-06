package main
import (
  "fmt"
  "testing"
  "reflect"
)
func TestDaysix(t *testing.T) {
  var tests = []struct {
  in  DaysixInput
  res int
  }{
    { DaysixInput{
[]string {
"Time:      7  15   30",
"Distance:  9  40  200",
},
    }, 288},
    { DaysixInput{
[]string {
"Time:        40     81     77     72",
"Distance:   219   1012   1365   1089",
},
    }, 861300},
  }
  for _, tt := range tests {
    testname := fmt.Sprintf("Daysix(%v) - wanted: %v", tt.in, tt.res)
      t.Run(testname, func(t *testing.T) {
        ans := Daysix(tt.in)
        if !reflect.DeepEqual(ans, tt.res) {
          t.Errorf("%v got: %v \n", testname, ans)
        }
      })
    }
}
func TestDaysixPartTwo(t *testing.T) {
  var tests = []struct {
  in  DaysixInput
  res int
  }{
    { DaysixInput{
[]string {
"Time:      7  15   30",
"Distance:  9  40  200",
},
    }, 71503},
    { DaysixInput{
[]string {
"Time:        40     81     77     72",
"Distance:   219   1012   1365   1089",
},
    }, 28101347},
  }
  for _, tt := range tests {
    testname := fmt.Sprintf("DaysixPartTwo(%v) - wanted: %v", tt.in, tt.res)
      t.Run(testname, func(t *testing.T) {
        ans := DaysixPartTwo(tt.in)
        if !reflect.DeepEqual(ans, tt.res) {
          t.Errorf("%v got: %v \n", testname, ans)
        }
      })
    }
}
