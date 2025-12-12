package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
	"sort"
	"time"
)

func sortContent(contents [][]string) {
	sort.Slice(contents, func(i, j int) bool {
		v1, _ := strconv.Atoi(contents[i][0])
		v2, _ := strconv.Atoi(contents[j][0])
		return v1 < v2
	})
}

func getContents() (*[][]string, *os.File) {
	file, err := os.OpenFile("tasks.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	contents := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := strings.Split(scanner.Text(), ",")
		contents = append(contents, val)
	}

	sortContent(contents)
	contents = append(contents, []string{"ID", "TASK", "CREATED", "DONE"})

	x, y := len(contents) - 1, len(contents) - 2
	for {
		if x < 1 || y < 0 {
			break
		}

		temp := contents[x]
		contents[x] = contents[y]
		contents[y] = temp

		x--
		y--
	}

	return &contents, file
}

func printTasks(complete bool) {
	contents, _ := getContents()
	
	maxWords := make([]int, 4)
	for _, content := range *contents {
		for i, v := range content {
			if i == 2 {
				x := time.Now().Format(time.RFC850)
				maxWords[i] = max(maxWords[i], len(x))
			} else {
				maxWords[i] = max(maxWords[i], len(v))
			}
		}
	}

	for k, content := range *contents {
		val := content
		for j, v := range val {
			if !complete && j == 3 {
				continue
			}

			if j == 2 && k != 0 {
				timestamp, _ := strconv.ParseInt(v, 10, 64)
				fmt.Printf("%s", time.Unix(0, timestamp).Format(time.RFC850))
			} else {
				fmt.Printf("%s", v)
				for i := 0; i < maxWords[j]-len(v); i++ {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\t")
		}
		fmt.Printf("\n")
	}
}

func addTask(taskString string, creationTime int64, completed bool) {
	contents, file := getContents()

	newTaskId := len(*contents)
	for i, content := range *contents {
		// skip the HEADERS
		if i == 0 {
			continue
		}

		val, _ := strconv.Atoi(content[0])
		if val != i {
			newTaskId = i
			break
		}
	}

	// fmt.Printf("The new task id is %d\n", newTaskId)
	_, err := file.Write([]byte(fmt.Sprintf("%d,%s,%d,%t\n", newTaskId, taskString, creationTime, completed)))
	if err != nil {
		panic(err)
	}
}

func deleteTask(id int) {
	contents, file := getContents()

	// for _, content := range *contents {
	// 	fmt.Println(content)
	// }

	n := len(*contents)
	for i, content := range *contents {
		// skip the headers
		if i == 0 {
			continue
		}

		val, _ := strconv.Atoi(content[0])
		if val == id {
			(*contents)[i] = (*contents)[n - 1]
			*contents = (*contents)[:n-1]
		}
	}

	// for _, content := range *contents {
	// 	fmt.Println(content)
	// }

	file.Truncate(0)
	file.Seek(0,0)

	for i, content := range *contents {
		if i == 0 {
			continue
		}
		val := strings.Join(content, ",")
		file.Write([]byte(fmt.Sprintf("%s\n", val)))
	}
}

func completeTask(id int) {
	contents, _ := getContents()

	for _, val := range *contents {
		taskId, _ := strconv.Atoi(val[0])
		if id == taskId {
			val[3] = "true"
			deleteTask(id)
			t, _ := strconv.ParseInt(val[2], 10, 64)
			addTask(val[1], t, true)
			return
		}
	}

	fmt.Printf("Task with id: %d not found\n", id)
	return
}

func main() {
	fmt.Println("original tasks:")
	printTasks(true)

	fmt.Println()
	fmt.Printf("adding a new task: %s\n", "compiler??")
	addTask("comiler??", time.Now().UnixNano(), false)
	fmt.Println("task added:")
	printTasks(true)

	fmt.Println()
	fmt.Println("deleting task 5:")
	deleteTask(5)
	printTasks(true)

	fmt.Println()
	fmt.Println("marking task 2 complete:")
	completeTask(2)
	printTasks(true)
}

