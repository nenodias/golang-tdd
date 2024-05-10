package main

import (
	"fmt"
	"net/http"
	"os"
	"tdd/injecao"
	"tdd/mock"
	"time"
)

func main() {
	mainMock()
}

func mainMock() {
	sleeper := &mock.SleeperConfiguravel{
		Duracao: 500 * time.Millisecond,
		Pausa:   time.Sleep,
	}
	mock.Contagem(os.Stdout, sleeper)
}

func mainInjecao() {
	//injecao.Cumprimenta(os.Stdout, "Elodie")
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))
	if err != nil {
		fmt.Println(err)
	}
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	injecao.Cumprimenta(w, "mundo")
}
