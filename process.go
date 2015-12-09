package ps

import (
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	Pid     int64
	User    string
	Cpu     float64
	Mem     float64
	Time    string
	Command string
}

func compact(collection []string) []string {
	var result []string
	for i := 0; i < len(collection); i++ {
		if collection[i] != "" {
			result = append(result, collection[i])
		}
	}
	return result
}

// USER PID %CPU %MEM VSZ RSS TT STAT STARTED TIME COMMAND
func Processes() []Process {
	var result []Process
	output, _ := exec.Command("ps", "aux").CombinedOutput()
	process_list := strings.Split(string(output), "\n")

	for i := 1; i < len(process_list); i++ {
		parts := compact(strings.Split(process_list[i], " "))
		if len(parts) >= 10 {
			var process Process
			process.User = parts[0]
			process.Pid, _ = strconv.ParseInt(parts[1], 10, 64)
			process.Cpu, _ = strconv.ParseFloat(parts[2], 64)
			process.Mem, _ = strconv.ParseFloat(parts[3], 64)
			process.Time = parts[9]
			for j := 10; j < len(parts); j++ {
				process.Command = process.Command + " " + parts[j]
			}
			result = append(result, process)
		}
	}
	return result
}

func main() {
	Processes()
}
