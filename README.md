# Go Processlist
This is a library to get the process list in a format compatible with golang using the system 'ps' command.


### Installation

```sh
$ go get github.com/jagdeep/gops
```

### Example

```sh
package main

import (
  "fmt"
  "github.com/jagdeep/gops"
)

main(){
  // find and print all running processes
  processes := ps.Processes()
  for i := 0; i < len(processes); i++ {
    fmt.Println("user:", processes[i].User, "pid:", processes[i].Pid, "command:", processes[i].Command)
  }

  // find a process with PID
  process := ps.FindByPid(processes[0].Pid)
  fmt.Println("process found with PID:", process)

  // find a process with PID file path
  pid_file_path := "REPLACE WITH A VALID PID FILE PATH"
  process := ps.FindByPidFilePath(pid_file_path)
  fmt.Println("process found with PID file path:", process)
}
```