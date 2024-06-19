package use_cases

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/repositories"
)

func AddProductUC(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Product Use Case"

	fmt.Println("Entramos a AddProductUC")
	status, response = repositories.AddProductRepositories(body, user)

	return status, response
}