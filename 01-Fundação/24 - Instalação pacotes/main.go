package main

import (
	"fmt"
	"github.com/google/uuid"
)
// para instalar o pacote uuid
// go get github.com/google/uuid

// caso rode go mod tidy, o pacote uuid será instalado automaticamente
// caso não rode go mod tidy, o pacote uuid não será instalado
// caso não esteja instalado o go mod tidy, remove a linha que importa o pacote uuid do arquivo main.go e remove o arquivo go.mod
func main() {
	fmt.Println("UUID:", uuid.NewString())
}
