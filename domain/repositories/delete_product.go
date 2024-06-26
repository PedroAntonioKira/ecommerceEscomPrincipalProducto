package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/queries_category"
)

func DeleteProductRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Delete Category Repositories"

	fmt.Println("Entramos a DeleteCategoryRepositories")
	status, response = queries_category.DeleteProductQuery(body, user, pathParams)

	return status, response
}