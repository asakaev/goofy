package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tuple[A, B any] struct {
	_1 A
	_2 B
}

type State struct {
	t  int
	mx int
	n  int
}

func read() (string, error) {
	bytes, err := io.ReadAll(os.Stdin)
	return string(bytes), err
}

func filter[A any](xs []A, p func(a A) bool) []A {
	var ys []A
	for i := range xs {
		if p(xs[i]) {
			ys = append(ys, xs[i])
		}
	}
	return ys
}

func tail[A any](xs []A) []A {
	return xs[1:]
}

func parse(s string) Tuple[int, int] {
	xs := strings.Split(s, " ")
	a, _ := strconv.Atoi(xs[0])
	b, _ := strconv.Atoi(xs[1])
	return Tuple[int, int]{a, b}
}

func fmap[A, B any](xs []A, f func(a A) B) []B {
	var ys []B
	for i := range xs {
		ys = append(ys, f(xs[i]))
	}
	return ys
}

func sortBy[A any](xs []A, f func(a A) int) []A {
	ys := make([]A, len(xs))
	copy(ys, xs)
	sort.Slice(ys, func(i, j int) bool { return f(ys[i]) < f(ys[j]) })
	return ys
}

func flatten[A any](xs [][]A) []A {
	var ys []A
	for i := range xs {
		for j := range xs[i] {
			ys = append(ys, xs[i][j])
		}
	}
	return ys
}

func fold[A, B any](xs []A, zero B, f func(acc B, a A) B) B {
	b := zero
	for i := range xs {
		b = f(b, xs[i])
	}
	return b
}

func step(s State, x Tuple[int, bool]) State {
	if !x._2 {
		return State{s.t, s.mx, s.n - 1}
	} else if s.n+1 > s.mx {
		return State{x._1, s.n + 1, s.n + 1}
	} else {
		return State{s.t, s.mx, s.n + 1}
	}
}

func Overlapping(xs []Tuple[int, int]) int {
	xs1 := fmap(xs, func(t Tuple[int, int]) []Tuple[int, bool] {
		return []Tuple[int, bool]{{t._1, true}, {t._2, false}}
	})
	xs2 := flatten(xs1)
	xs3 := sortBy(xs2, func(t Tuple[int, bool]) int { return t._1 })
	return fold(xs3, State{-1, 0, 0}, step).t
}

func main() {
	s, _ := read()
	xs1 := strings.Split(s, "\n")
	xs2 := filter(xs1, func(s string) bool { return s != "" })
	xs3 := tail(xs2)
	xs4 := fmap(xs3, parse)
	n := Overlapping(xs4)
	fmt.Println(n)
}
