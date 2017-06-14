package kata

import (
    "bytes"
)

//alternatives: strings.Join and + operator

func RepeatStr(repetitions int, value string) string {
  var buffer bytes.Buffer

  for i := 0; i < repetitions; i++ {
        buffer.WriteString(value)
    }
  
  return buffer.String()
}