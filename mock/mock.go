package mock

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperConfiguravel struct {
	Duracao time.Duration
	Pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	s.Pausa(s.Duracao)
}

var inicioContagem = 3
var ultimaPalavra = "Vai!"

func Contagem(w io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, ultimaPalavra)
}
