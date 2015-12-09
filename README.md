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
  processes := ps.Processes()
  for i := 0; i < len(processes); i++ {
    fmt.Println("user:", processes[i].User, "pid:", processes[i].Pid, "command:", processes[i].Command)
  }
}
```