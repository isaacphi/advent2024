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

	setsOfSize := make(map[int]map[string]set)
	setsOfSize[2] = make(map[string]set)

	for _, conn := range strings.Fields(input) {
		computerNames := strings.Split(conn, "-")
		for _, c := range computerNames {
			if _, exists := nodes[c]; !exists {
				nodes[c] = &node{
					name:        c,
					connections: make(set),
					setsOfSize:  make(map[int]map[string]set),
				}
				nodes[c].setsOfSize[2] = make(map[string]set)
			}
		}

		for i, c := range computerNames {
			var c2 string
			if i == 0 {
				c2 = computerNames[1]
			} else {
				c2 = computerNames[0]
			}
			nodes[c].connections[c2] = nodes[c2]
		}

		s := newSet(nodes[computerNames[0]], nodes[computerNames[1]])
		nodes[computerNames[0]].setsOfSize[2][s.getPassword()] = s
		nodes[computerNames[1]].setsOfSize[2][s.getPassword()] = s
		setsOfSize[2][s.getPassword()] = s
		setsOfSize[2][s.getPassword()] = s
	}

	// count := 0
	// for _, n := range nodes {
	// 	for _, c := range n.connections {
	// 		for _, c2 := range n.connections {
	// 			if n.name[0] != 't' && c.name[0] != 't' && c2.name[0] != 't' {
	// 				continue
	// 			}
	// 			if c.hasConnection(c2) {
	// 				count++
	// 			}
	// 		}
	// 	}
	// }
	//
	// fmt.Println("part 1", count/6)

	for level := 3; ; level++ {
		// For each size of cluster, starting a 3 and increasing until there are no more clusters
		setsOfSize[level] = make(map[string]set)

		for _, smallerSet := range setsOfSize[level-1] {
			// For each cluster of the previous size
			otherNodes := newSet() // Nodes not in this cluster
			for _, n := range nodes {
				otherNodes.add(n)
			}

			for _, n := range smallerSet {
				// For each other node
				for _, otherNode := range otherNodes {
					// If not connected to n, remove from otherNodes
					if !n.isConnected(otherNode) {
						otherNodes.remove(otherNode)
					}
				}
			}

			// Anything remaining in otherNodes is connected to all nodes in smallerSet
			for _, n := range otherNodes {
				s := newSet()
				for _, oldNode := range smallerSet {
					s.add(oldNode)
				}
				s.add(n)
				setsOfSize[level][s.getPassword()] = s
			}
		}

		fmt.Println("part 2, level", level)
		for s := range setsOfSize[level] {
			fmt.Println(s)
		}
		if len(setsOfSize[level]) == 0 {
			break
		}
	}
}

type set map[string]*node

func (s set) isEqual(s2 set) bool {
	return s.getPassword() == s2.getPassword()
}

func newSet(nodes ...*node) set {
	s := make(set)
	for _, n := range nodes {
		s[n.name] = n
	}
	return s
}

func (s set) add(n *node) {
	s[n.name] = n
}

func (s set) remove(n *node) {
	delete(s, n.name)
}

func (s set) contains(n *node) bool {
	_, exists := s[n.name]
	return exists
}

func (s set) getPassword() string {
	l := make([]string, 0)
	for _, n := range s {
		l = append(l, n.name)
	}
	sort.Strings(l)
	var sb strings.Builder
	for i, n := range l {
		sb.WriteString(n)
		if i != len(l)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

type node struct {
	connections set
	name        string
	setsOfSize  map[int]map[string]set
}

func (n *node) hasConnection(otherNode *node) bool {
	return n.connections.contains(otherNode)
}

func (n *node) printNode() {
	fmt.Println(n.name)
	for size, sets := range n.setsOfSize {
		fmt.Println(size)
		for password := range sets {
			fmt.Print(password, " ")
		}
		fmt.Println()
	}
}

func (n *node) isConnected(node *node) bool {
	return n.connections.contains(node)
}
