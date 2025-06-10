package main


type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

// Cria uma nova instância de Conta e retorna um ponteiro para ela (&Conta)
func NewConta() *Conta {
	return &Conta{
		saldo: 0,
	}
}
func main() {
	conta := Conta{
		saldo: 100,
	}

	conta.simular(200) // Simula um depósito de 200
	println("Saldo após simulação:", conta.saldo) // Exibe o saldo após a simulação
	
}
