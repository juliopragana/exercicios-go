package models

import "time"

type Compras struct {
	DataCompra   time.Time
	NomeMercado  string
	ItensCompras []ItensCompras
	ValorTotal   float64
}

type ItensCompras struct {
	Nome       string
	Valor      float64
	Quantidade int
}

func (c *Compras) CalculaValorTotal() {

	var valorTotal float64 = 0
	// var quantidade int = 0

	for i := 0; i < len(c.ItensCompras); i++ {
		valorTotal = valorTotal + c.ItensCompras[i].Valor*float64(c.ItensCompras[i].Quantidade)
		// quantidade = quantidade + c.ItensCompras[i].Quantidade
	}

	c.ValorTotal = valorTotal
}
