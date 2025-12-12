package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	distSq int
	i, j   int
}

var uf []int

func find(x int) int {
	if uf[x] == x {
		return x
	}
	uf[x] = find(uf[x])
	return uf[x]
}

func mix(x, y int) {
	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		uf[rootX] = rootY
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var P []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			P = append(P, Point{x, y, z})
		}
	}

	var D []Edge
	n := len(P)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > j {
				p1, p2 := P[i], P[j]
				dist := (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)
				D = append(D, Edge{distSq: dist, i: i, j: j})

			}

		}
	}

	sort.Slice(D, func(k, l int) bool {
		if D[k].distSq != D[l].distSq {
			return D[k].distSq < D[l].distSq
		}
		if D[k].i != D[l].i {
			return D[k].i < D[l].i
		}
		return D[k].j < D[l].j
	})

	uf = make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}

	connections := 0
	for t, edge := range D {

		if t == 1000 {
			SZ := make(map[int]int)
			for x := 0; x < n; x++ {
				root := find(x)
				SZ[root]++
			}

			var sValues []int
			for _, v := range SZ {
				sValues = append(sValues, v)
			}
			sort.Ints(sValues)

			l := len(sValues)
			if l >= 3 {
				fmt.Println(sValues[l-1] * sValues[l-2] * sValues[l-3])
			}
		}
		if find(edge.i) != find(edge.j) {
			connections++
			if connections == n-1 {
				fmt.Println(P[edge.i].x * P[edge.j].x)
			}
			mix(edge.i, edge.j)
		}
	}
}
