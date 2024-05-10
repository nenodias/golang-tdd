package mock

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyContagemOperacoes struct {
	Chamadas []string
}

const escrita = "escrita"
const pausa = "pausa"

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return 1, nil
}

type TimeSpy struct {
	duracaoSleep time.Duration
}

func (t *TimeSpy) Sleep(duracao time.Duration) {
	t.duracaoSleep = duracao
}

func TestContagem(t *testing.T) {
	t.Run("imprime 3 até Vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SpyContagemOperacoes{}
		Contagem(buffer, sleeperSpy)

		got := buffer.String()
		want := `3
2
1
Vai!`
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("pausa antes de cada impressao", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		want := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}
		got := spyImpressoraSleep.Chamadas
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestSleepConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second
	timeSpy := &TimeSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, timeSpy.Sleep}
	sleeper.Sleep()
	if timeSpy.duracaoSleep != tempoPausa {
		t.Errorf("should sleep for %v, but sleep for %v", tempoPausa, timeSpy.duracaoSleep)
	}
}
