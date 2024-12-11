package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func sortBlocks1(diskmap []string) {
	front, back := 0, len(diskmap)-1
	for front < back {
		for diskmap[front] != "." {
			front++
		}
		for diskmap[back] == "." {
			back--
		}
		if front < back {
			diskmap[front] = diskmap[back]
			diskmap[back] = "."
			// fmt.Println(diskmap)
		}
	}
}

func sortBlocks2(diskmap []string) {
	front, back := 0, len(diskmap)-1
	for front < back {
		for diskmap[front] != "." {
			front++
		}
		for diskmap[back] == "." {
			back--
		}
		if front < back {
			fileId := diskmap[back]
			fileSize := 0
			for diskmap[back-fileSize] == fileId {
				fileSize++
			}

			spaceLocation := front
			spaceSize := 0
			for spaceSize < fileSize && spaceLocation+spaceSize < back-fileSize {
				if diskmap[spaceLocation+spaceSize] != "." {
					spaceLocation++
					spaceLocation += spaceSize
					spaceSize = 0
					continue
				}
				for diskmap[spaceLocation+spaceSize] == "." {
					spaceSize++
				}
			}

			if spaceSize >= fileSize {
				for i := 0; i < fileSize; i++ {
					diskmap[spaceLocation+i] = fileId
					diskmap[back-i] = "."
				}
				// front += fileSize
			}
			// fmt.Println(diskmap)
			back -= fileSize
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func checksum(diskmap []string) int {
	result := 0
	for i, val := range diskmap {
		intVal, _ := strconv.Atoi(val)
		result += i * intVal
	}
	return result
}

func main() {
	// data, _ := os.ReadFile("test_data")
	data, _ := os.ReadFile("data")
	input := string(data)

	// fmt.Println(input)

	diskmap := make([]string, 0)
	isFile := true
	for i, digit := range input {
		num, _ := strconv.Atoi(string(digit))
		for range num {
			if isFile {
				diskmap = append(diskmap, strconv.Itoa(i/2))
			} else {
				diskmap = append(diskmap, ".")
			}
		}
		isFile = !isFile
	}

	// fmt.Println(diskmap)
	// sortBlocks1(diskmap)
	sortBlocks2(diskmap)
	result := checksum(diskmap)
	fmt.Println(result)
}
