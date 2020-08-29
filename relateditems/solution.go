package main

import (
	"fmt"
	"sort"
	"strings"
)

// Pair
type pair struct {
	first  string
	second string
}

// group
type group struct {
	items     []string
	itemCount int32
}

// check a group for an item
func (g group) contains(item string) bool {

	for _, itm := range g.items {
		if itm == item {
			return true
		}
	}

	return false
}

// Add an item to a group
func add(g group, item string) group {
	if !g.contains(item) {
		g.items = append(g.items, item)
		g.itemCount++
		sort.Strings(g.items)
		fmt.Println("Added ", item, "to group", g)
	}
	return g
}

//Pring group
func print(g group) {
	fmt.Println(g)
}

// Create a new group, for a pair
func createGroup(p pair) group {

	group := group{[]string{p.first, p.second}, 2}
	fmt.Println("Created new group:", group)
	return group
}

func main() {

	pairs := []pair{{"item1", "item2"}, {"item2", "item3"}, {"item7", "item97"}, {"item97", "item93"}}

	groups := []group{}
	for _, pair := range pairs {

		fmt.Println("Processing pair :", pair)
		if len(groups) == 0 {
			group := createGroup(pair)
			groups = append(groups, group)
			continue
		}

		foundInGroup := false
		for i := 0; i < len(groups); i++ {
			if groups[i].contains(pair.first) {
				groups[i] = add(groups[i], pair.second)
				print(groups[i])
				foundInGroup = true
				break
			} else if groups[i].contains(pair.second) {
				groups[i] = add(groups[i], pair.first)
				print(groups[i])
				foundInGroup = true
				break
			}
		}

		if !foundInGroup {
			group := createGroup(pair)
			groups = append(groups, group)
		}

	}

	// Derive largest association group
	var bg group = group{}
	for _, g := range groups {
		if g.itemCount > bg.itemCount {
			bg = g
		} else if g.itemCount == bg.itemCount {
			if strings.Compare(g.items[0], bg.items[0]) < 0 {
				bg = g
			}
		}
	}

	fmt.Println(bg)
}
