package gocodegen

//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)
