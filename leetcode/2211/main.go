package main

import "fmt"

type P struct {
	dir   byte
	count int
}

func countCollisions(directions string) int {
	res := 0
	n := len(directions)
	p := []*P{}
	p = append(p, &P{
		dir:   directions[0],
		count: 1,
	})
	for i := 1; i < n; i++ {
		top := len(p) - 1
		cur := directions[i]
		if p[top].dir == cur {
			p[top].count++
		} else {
			p = append(p, &P{
				dir:   cur,
				count: 1,
			})
		}
	}

	for i := 0; i < len(p)-1; i++ {
		p1 := p[i]
		p2 := p[i+1]
		if p1.dir == 'R' && p2.dir == 'L' {
			res += (p1.count + p2.count)
		}
		if p1.dir == 'R' && p2.dir == 'S' {
			res += p1.count
		}
		if p1.dir == 'S' && p2.dir == 'L' {
			res += p2.count
		}
	}
	return res
}

func main() {
	fmt.Println(countCollisions("RLRSLL"))
}
