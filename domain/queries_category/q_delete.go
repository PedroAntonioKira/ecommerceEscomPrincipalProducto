package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func DeleteProductQuery(body string, User string, pathParams int) (int, string) {
	
	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Actualizamos el Producto.
	err2 := database.DeleteProductQuery(pathParams)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el DELETE del producto" + strconv.Itoa(pathParams) + " > " + err2.Error()
	}

	return 200, "Delete Product OK"

}