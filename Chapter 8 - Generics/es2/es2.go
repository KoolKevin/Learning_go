package main

import (
	"fmt"
)

type printable interface {
	~int | ~float64
	fmt.Stringer
}

// questo non funziona, l'underlying type deve essere diretto
// type PrintableInt struct {
// 	intero int
// }

// func (pi PrintableInt) String() string {
// 	return strconv.Itoa(pi.intero)
// }

type PrintableFloat float64

func (pf PrintableFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintPrintable[T printable](parameter T) {
	fmt.Println(parameter)
}

func main() {
	var pf PrintableFloat = 10.1
	PrintPrintable(pf)

}
