package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/queries_product"
)

func AddProductRepositories(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Repositories"

	fmt.Println("Entramos a AddProductRepositories")
	status, response = queries_product.AddProductQuery(body, user)

	return status, response
}