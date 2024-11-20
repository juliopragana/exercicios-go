package main

import (
	"exe1/models"
	"fmt"
	"time"
)

func main() {

	var itensCompras []models.ItensCompras
	itensCompras = append(itensCompras, models.ItensCompras{
		Nome:       "Bola",
		Valor:      200.50,
		Quantidade: 3,
	})
	itensCompras = append(itensCompras, models.ItensCompras{
		Nome:       "Chuteira",
		Valor:      300.50,
		Quantidade: 2,
	})

	compras := models.Compras{
		DataCompra:   time.Now(),
		NomeMercado:  "Mercadinho bom de Preço",
		ItensCompras: itensCompras,
	}

	compras.CalculaValorTotal()

	//formato da Data
	formatoData := "02/01/2006 15:04:05"
	dataFormatada := compras.DataCompra.Format(formatoData)

	fmt.Println("Data da Compra: ", dataFormatada)
	fmt.Println("NomeMercado: ", compras.NomeMercado)
	fmt.Println("---- Itens da compras --- ")
	fmt.Println("Nome do item:       Quantidade         Valor do Item")
	for i := 0; i < len(compras.ItensCompras); i++ {
		fmt.Printf("%s               	%d                   R$%g \n",
			compras.ItensCompras[i].Nome,
			compras.ItensCompras[i].Quantidade,
			compras.ItensCompras[i].Valor)
	}

	fmt.Println("---- Itens da compras --- ")
	fmt.Println("Valor total da compra:  R$", compras.ValorTotal)

}
