package injecao

import (
	"fmt"
	"io"
)

func Cumprimenta(writer io.Writer, nome string) {
	fmt.Fprintf(writer, "Olá, %s", nome)
}
