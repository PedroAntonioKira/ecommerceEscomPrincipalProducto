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
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func DeleteCategoryQuery(body string, User string, pathParams int) (int, string) {
	
	// Validamos que nos hayan pasado un Id valido
	if pathParams == 0 {
		return 400, "Debe de especificar el ID de la categoria a borrar."
	}

	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Eliminamos la categoria que corresponde al id
	err := database.DeleteCategoryQuery(pathParams)

	// Validamos que no haya surgido un error al eliminar la categoria
	if err != nil {
		return 400, "Ocurrió un error al intentar realziar el DELETE de la categoria" + strconv.Itoa(pathParams) + " > " + err.Error()
	}

	return 200, "Delete OK"

}