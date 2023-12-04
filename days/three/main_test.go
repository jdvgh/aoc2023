package main
import (
  "fmt"
  "testing"
  "reflect"
)
func TestDaythree(t *testing.T) {
  var tests = []struct {
  in  DaythreeInput
  res int
  }{
    { DaythreeInput{0}, 2},
  }
  for _, tt := range tests {
    testname := fmt.Sprintf("Daythree(%v) - wanted: %v", tt.in, tt.res)
      t.Run(testname, func(t *testing.T) {
        ans := Daythree(tt.in)
        if !reflect.DeepEqual(ans, tt.res) {
          t.Errorf("%v got: %v \n", testname, ans)
        }
      })
    }
}
