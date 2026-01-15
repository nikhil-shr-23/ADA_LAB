package main

import (
	"fmt"
	"sort"
)

type Project struct {
	Name  string
	Start int
	End   int
}

func AllocateResources(projects []Project) []Project {
	// Sort projects by finish time
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].End < projects[j].End
	})

	var selected []Project
	if len(projects) == 0 {
		return selected
	}

	// The first project is always selected
	selected = append(selected, projects[0])
	lastEnd := projects[0].End

	for i := 1; i < len(projects); i++ {
		if projects[i].Start >= lastEnd {
			selected = append(selected, projects[i])
			lastEnd = projects[i].End
		}
	}

	return selected
}

func main() {
	projects := []Project{
		{"Project A", 1, 3},
		{"Project B", 2, 5},
		{"Project C", 4, 7},
		{"Project D", 1, 8},
		{"Project E", 5, 9},
		{"Project F", 8, 10},
		{"Project G", 9, 11},
		{"Project H", 11, 14},
		{"Project I", 13, 16},
	}

	fmt.Println("Available Projects:")
	fmt.Printf("%-10s | %-5s | %-5s\n", "Name", "Start", "End")
	fmt.Println("------------------------------")
	for _, p := range projects {
		fmt.Printf("%-10s | %-5d | %-5d\n", p.Name, p.Start, p.End)
	}
	fmt.Println("------------------------------")

	selected := AllocateResources(projects)

	fmt.Println("\nSelected Projects (Optimal Schedule):")
	fmt.Printf("%-10s | %-5s | %-5s\n", "Name", "Start", "End")
	fmt.Println("------------------------------")
	for _, p := range selected {
		fmt.Printf("%-10s | %-5d | %-5d\n", p.Name, p.Start, p.End)
	}
	fmt.Println("------------------------------")
	fmt.Printf("Total Projects Selected: %d\n", len(selected))
}
