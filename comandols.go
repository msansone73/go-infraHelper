package main

import "fmt"
import "exec"

func main() {
  cmd, err := exec.Run("/bin/ls", []string{"/bin/ls"}, []string{}, "", exec.DevNull, exec.PassThrough, exec.PassThrough)
  if (err != nil) {
    fmt.Println(err)
    return
  }
  cmd.Close()

}
