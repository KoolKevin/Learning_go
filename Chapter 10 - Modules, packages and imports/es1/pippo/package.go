// Questo è il commento che apparirà su pkg.go.dev per il package pippo.
// Può essere multilinea e spiegare lo scopo del package.
package pippo

import "fmt"

func Pippo() {
	// riesco a vedere a anche se in un altro file e non esportato, dato che sono nello stesso package
	fmt.Println(a)
}
