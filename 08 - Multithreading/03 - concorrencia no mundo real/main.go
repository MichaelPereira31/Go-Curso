package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// Para evitar condições de corrida, podemos usar um mutex para proteger o acesso à variável compartilhada
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()

		// Ou podemos usar a função atomic para incrementar a variável de forma atômica, garantindo que não haja condições de corrida
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("voce e o visitante numero %d \n", number)))
	})

	http.ListenAndServe(":3000", nil)
}
