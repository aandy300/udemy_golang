//常數的索引?
//參考 BLOG經驗分享 https://blog.wolfogre.com/posts/golang-iota/
package main

import (
	"fmt"
)

const a = iota

const (
	b = iota
)

const (
	c = 0
	d = iota
	e
	f = "hello"
	// nothing
	g
	h = iota
	i
	j = 0
	k
	l, m = iota, iota
	n, o

	p = iota + 1
	q
	_
	r = iota * iota
	s
	t = r
	u
	v = 1 << iota
	w
	x         = iota * 0.01
	y float32 = iota * 0.01
	z
)

func main() {
	fmt.Printf("a : %T = %v\n", a, a)
	fmt.Printf("b : %T = %v\n", b, b)
	fmt.Printf("c : %T = %v\n", c, c)
	fmt.Printf("d : %T = %v\n", d, d)
	fmt.Printf("e : %T = %v\n", e, e)
	fmt.Printf("f : %T = %v\n", f, f)
	fmt.Printf("g : %T = %v\n", g, g)
	fmt.Printf("h : %T = %v\n", h, h)
	fmt.Printf("i : %T = %v\n", i, i)
	fmt.Printf("j : %T = %v\n", j, j)
	fmt.Printf("k : %T = %v\n", k, k)
	fmt.Printf("l : %T = %v\n", l, l)
	fmt.Printf("m : %T = %v\n", m, m)
	fmt.Printf("n : %T = %v\n", n, n)
	fmt.Printf("o : %T = %v\n", o, o)
	fmt.Printf("p : %T = %v\n", p, p)
	fmt.Printf("q : %T = %v\n", q, q)
	fmt.Printf("r : %T = %v\n", r, r)
	fmt.Printf("s : %T = %v\n", s, s)
	fmt.Printf("t : %T = %v\n", t, t)
	fmt.Printf("u : %T = %v\n", u, u)
	fmt.Printf("v : %T = %v\n", v, v)
	fmt.Printf("w : %T = %v\n", w, w)
	fmt.Printf("x : %T = %v\n", x, x)
	fmt.Printf("y : %T = %v\n", y, y)
	fmt.Printf("z : %T = %v\n", z, z)
}
