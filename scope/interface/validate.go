package main

type email string

func (e email) validate() bool {
	return true
}

type strtofloat string

func (s strtofloat) validate() bool {
	return true
}

func (s strtofloat) convert() float64 {
	return 0.0
}