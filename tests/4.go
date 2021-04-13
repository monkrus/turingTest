package main

type S struct {
	m string
}

func f() *S {
	return &S{"foo"}
}

func main() {
	p := f()
	print(p.m) //print"foo"

}

// fill the blanks for lines marked with A and B, to ensure that the printed output is "foo"
// we return a memory adress first, then print "foo"
// https://play.golang.org/p/guR0ASOtGUu
