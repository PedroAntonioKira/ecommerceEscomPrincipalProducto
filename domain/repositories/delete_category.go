package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/queries_category"
)

func DeleteCategoryRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Delete Category Repositories"

	fmt.Println("Entramos a DeleteCategoryRepositories")
	status, response = queries_category.DeleteCategoryQuery(body, user, pathParams)

	return status, response
}