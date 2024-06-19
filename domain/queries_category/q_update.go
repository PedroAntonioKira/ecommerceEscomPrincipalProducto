package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func UpdateCategoryQuery(body string, User string, pathParams int) (int, string) {
	//creamos la variable de la estructura que almacenará todo lo de la categoria relacionada
	var t entities.Category

	//Decodificamos el json que nos mandan en el endpoint en la estructura del producto para poder guardarla.
	err := json.Unmarshal([]byte(body), &t)

	//Verificamos que no tengamos un error al decodificar la información en la estructura.
	if err != nil {
		return 400, "Error en los datos recibidos con el error: " + err.Error()
	}

	//Verificamos que nos mande la información
	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar CategName y/o CategPath para actualizar"
	}

	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// el id de la categoria lo asignamos que nos pasan como parametro
	t.CategID = pathParams

	err2 := database.UpdateCategoryQuery(t)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realziar el UPDATE de la categoria" + strconv.Itoa(pathParams) + " > " + err2.Error()
	}

	return 200, "Update OK Very Good"

}
