package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Read and Parse Input
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	G := make(map[string][]string)
	I := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		x := strings.TrimSuffix(parts[0], ":")

		// Ensure x exists in G
		if _, exists := G[x]; !exists {
			G[x] = []string{}
		}

		// Process neighbors
		neighbors := parts[1:]
		G[x] = neighbors

		// Ensure neighbors exist in G and update In-degrees
		for _, n := range neighbors {
			if _, exists := G[n]; !exists {
				G[n] = []string{}
			}
			I[n]++
		}
	}

	// 2. Topological Sort (Kahn's Algorithm)
	// Q starts with nodes having 0 in-degree
	var Q []string
	for k := range G {
		if I[k] == 0 {
			Q = append(Q, k)
		}
	}

	// Iterate Q. Note: In Go, we use an index loop because we append to Q
	// inside the loop, similar to the Python code behavior.
	for idx := 0; idx < len(Q); idx++ {
		u := Q[idx]
		for _, v := range G[u] {
			I[v]--
			if I[v] == 0 {
				Q = append(Q, v)
			}
		}
	}

	// 3. Dynamic Programming (Reverse Topological Order)
	// Z maps node -> [4]int state
	Z := make(map[string][4]int)

	// Initialize sinks (nodes with no outgoing edges)
	// In Python: Z = {x:[1,0,0,0] for x in Q if not G[x]}
	for _, x := range Q {
		if len(G[x]) == 0 {
			Z[x] = [4]int{1, 0, 0, 0}
		}
	}

	// Process Q in reverse (pop())
	for j := len(Q) - 1; j >= 0; j-- {
		u := Q[j]

		// Determine the bit shift flags based on node name
		// Python: (u=='dac') | 2*(u=='fft')
		shift := 0
		if u == "dac" {
			shift |= 1
		}
		if u == "fft" {
			shift |= 2
		}

		// Accumulate values from neighbors (v) into u
		// Note: Z[u] might already have values if it was a sink,
		// but for non-sinks it starts as [0,0,0,0] (default map value).
		// We need to copy the current state to modify it or access map directly.
		currentZ := Z[u]

		for _, v := range G[u] {
			childZ := Z[v]
			for i := 0; i < 4; i++ {
				// Python: Z[u][i | shift] += Z[v][i]
				targetIdx := i | shift
				if targetIdx < 4 {
					currentZ[targetIdx] += childZ[i]
				}
			}
		}
		Z[u] = currentZ
	}

	// 4. Output Results
	// Part 1: Sum of Z['you']
	youSum := 0
	if val, ok := Z["you"]; ok {
		for _, v := range val {
			youSum += v
		}
	}
	fmt.Printf("Part 1: %d\n", youSum)

	// Part 2: Z['svr'][3]
	svrVal := 0
	if val, ok := Z["svr"]; ok {
		svrVal = val[3]
	}
	fmt.Printf("Part 2: %d\n", svrVal)
}
