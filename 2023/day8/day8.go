package main

// pomnozi gi

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph struct {
	vertices []*Vertex
}

// Adjacent Vertex
type Vertex struct {
	key      string
	adjacent []*Vertex
}

// AddVertext will add a vertex to a graph
func (g *Graph) AddVertex(vertex string) error {
	if contains(g.vertices, vertex) {
		err := fmt.Errorf("Vertex %s already exists", vertex)
		return err
	} else {
		v := &Vertex{
			key: vertex,
		}
		g.vertices = append(g.vertices, v)
	}
	return nil
}

func (g *Graph) AddEdge(to, from string) error {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("Not a valid edge from %s ---> %s", from, to)
	} else if contains(fromVertex.adjacent, toVertex.key) {
		return fmt.Errorf("Edge from vertex %s ---> %s already exists", fromVertex.key, toVertex.key)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		return nil
	}
}

// getVertex will return a vertex point if exists or return nil
func (g *Graph) getVertex(vertex string) *Vertex {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(v []*Vertex, key string) bool {
	for _, v := range v {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%s : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%s ", v.key)
		}
		fmt.Println()
	}
}

func findZZZ(LR string, root Vertex) int {
	var sum int = 0
	for i := 0; i < len(LR); i = (i + 1) % len(LR) {
		if root.key == "ZZZ" {
			break
		}
		if len(root.adjacent) == 1 {
			root = *root.adjacent[0]
			sum++
			continue
		}
		if LR[i] == 'L' {
			root = *root.adjacent[0]
			sum++
		} else {
			root = *root.adjacent[1]
			sum++
		}
	}
	return sum
}

func findZ(LR string, root Vertex) int {
	var sum int = 0
	for i := 0; i < len(LR); i = (i + 1) % len(LR) {
		if root.key[2] == 'Z' {
			break
		}
		if len(root.adjacent) == 1 {
			root = *root.adjacent[0]
			sum++
			continue
		}
		if LR[i] == 'L' {
			root = *root.adjacent[0]
			sum++
		} else {
			root = *root.adjacent[1]
			sum++
		}
	}
	return sum
}

func reachedEnd(roots []Vertex) bool {
	for _, el := range roots {
		if el.key[2] != 'Z' {
			return false
		}
	}
	return true
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findlcm(arr []int) int {
	var ans int = arr[0]
	for i := 1; i < len(arr); i++ {
		ans = ((arr[i] * ans) / (gcd(arr[i], ans)))
	}
	return ans
}

func main() {
	var graph Graph
	// file, err := os.Open("shortInput.txt")
	file, err := os.Open("longInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var LR string = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		var root string = strings.Split(scanner.Text(), " = ")[0]
		var left string = strings.Split(strings.Split(scanner.Text(), " = ")[1], ", ")[0][1:]
		var right string = strings.Split(strings.Split(scanner.Text(), " = ")[1], ", ")[1]
		right = strings.TrimSuffix(right, ")")
		graph.AddVertex(root)
		graph.AddVertex(left)
		graph.AddVertex(right)
		graph.AddEdge(left, root)
		graph.AddEdge(right, root)
	}
	var nums []int
	for _, el := range graph.vertices {
		if el.key[2] == 'A' {
			nums = append(nums, findZ(LR, *el))
		}
	}
	fmt.Printf("==== Part 1 ====\n")
	var root Vertex = *graph.getVertex("AAA")
	fmt.Println(findZZZ(LR, root))
	fmt.Printf("==== Part 2 ====\n")
	fmt.Println(findlcm(nums))
}
