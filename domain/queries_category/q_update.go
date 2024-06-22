package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"

//	"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func UpdateProductQuery(body string, User string, pathParams int) (int, string) {
	//creamos la variable de la estructura que almacenará todo lo del producto relacionado
	var t entities.Product

	//Decodificamos el json que nos mandan en el endpoint en la estructura del producto para poder guardarla.
	err := json.Unmarshal([]byte(body), &t)

	//Verificamos que no tengamos un error al decodificar la información en la estructura.
	if err != nil {
		return 400, "Error en los datos recibidos con el error: " + err.Error()
	}

	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// el id del producto lo asignamos que nos pasan como parametro
	t.ProdId = pathParams

	//Actualizamos el Producto.
	err2 := database.UpdateProductDatabase(t)

	//Verificamos no exista un error al momento en que actualizamos el producto.
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del producto" + strconv.Itoa(pathParams) + " > " + err2.Error()
	}

	return 200, "Update OK Producto"

}
