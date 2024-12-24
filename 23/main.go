package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)

	nodes := make(map[string]*node)
	MAX_GROUP_SIZE := 40
	setsOfSize := make([]map[string]set, MAX_GROUP_SIZE+1)
	invalidSetsOfSize := make([]map[string]bool, MAX_GROUP_SIZE+1)
	setsOfSize[1] = make(map[string]set)
	setsOfSize[2] = make(map[string]set)

	for _, conn := range strings.Fields(input) {
		// fmt.Println(conn)
		computers := strings.Split(conn, "-")
		for _, c := range computers {
			// Populate sets of 1
			s := newSet(c)
			setsOfSize[1][s.getPassword()] = s

			if _, exists := nodes[c]; !exists {
				nodes[c] = &node{
					name: c,
				}
			}
		}
		for i, c := range computers {
			var c2 string
			if i == 0 {
				c2 = computers[1]
			} else {
				c2 = computers[0]
			}
			nodes[c].connections = append(nodes[c].connections, nodes[c2])
		}
		// Populate sets of 2
		s := newSet(computers...)
		setsOfSize[2][s.getPassword()] = s
	}

	// fmt.Println(nodes)
	count := 0
	for _, n := range nodes {
		for _, c := range n.connections {
			for _, c2 := range n.connections {
				if n.name[0] != 't' && c.name[0] != 't' && c2.name[0] != 't' {
					continue
				}
				if c.hasConnection(c2) {
					// fmt.Println(n.name, c.name, c2.name)
					count++
				}
			}
		}
	}
	// fmt.Println("part 1", count/6)
	biggestN := 0
	for n := 3; n <= MAX_GROUP_SIZE; n++ {
		fmt.Println(n)
		setsOfSize[n] = make(map[string]set)
		invalidSetsOfSize[n] = make(map[string]bool)
		// fmt.Println(n)

		for groupName, group := range setsOfSize[n-1] {
			fmt.Println(n, groupName)

			largerGroup := newSet(group.contents()...)
			for nodeName := range setsOfSize[1] {
				// fmt.Print(nodeName)
				if group.contains(nodeName) {
					continue
				}
				matches := 0
				largerGroup.add(nodeName)
				pw := largerGroup.getPassword()
				if _, exists := setsOfSize[n][pw]; exists {
					largerGroup.remove(nodeName)
					continue
				}
				if _, exists := invalidSetsOfSize[n][pw]; exists {
					largerGroup.remove(nodeName)
					continue
				}

				for otherGroupName, otherGroup := range setsOfSize[n-1] {
					if groupName == otherGroupName {
						continue
					}
					if largerGroup.isSupersetOf(otherGroup) {
						matches++
					}
					// fmt.Println(largerGroup, otherGroup, matches)
				}
				if matches == n-1 {
					setsOfSize[n][pw] = largerGroup
				} else {
					invalidSetsOfSize[n][pw] = true
				}
				largerGroup.remove(nodeName)
			}
		}
		if len(setsOfSize[n]) == 0 {
			biggestN = n - 1
			fmt.Println("done", n)
			break
		}
	}
	for s := range setsOfSize[biggestN] {
		fmt.Println(s)
	}
}

type set map[string]struct{}

func (s set) isSupersetOf(s2 set) bool {
	count := 0
	for val := range s2 {
		if s.contains(val) {
			count++
		}
	}
	return count == len(s2)
}

func (s set) add(nodeName string) {
	s[nodeName] = struct{}{}
}

func (s set) remove(nodeName string) {
	delete(s, nodeName)
}

func (s set) contains(nodeName string) bool {
	_, exists := s[nodeName]
	return exists
}

func (s set) contents() []string {
	c := make([]string, 0)
	for val := range s {
		c = append(c, val)
	}
	return c
}

func (s set) getPassword() string {
	l := make([]string, 0)
	for node := range s {
		l = append(l, node)
	}
	sort.Strings(l)
	var sb strings.Builder
	for i, s := range l {
		sb.WriteString(s)
		if i < len(l)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func newSet(vals ...string) set {
	s := make(set)
	for _, val := range vals {
		s.add(val)
	}
	return s
}

type node struct {
	connections []*node
	name        string
}

func (n *node) hasConnection(n2 *node) bool {
	for _, c := range n.connections {
		if c == n2 {
			return true
		}
	}
	return false
}
