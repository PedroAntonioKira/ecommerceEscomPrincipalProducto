package queries_product

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"
	"fmt"

	//importaciones externas (descargadas)
	//"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

// Declaramos función para insertar los productos
func AddProductQuery(body string, User string) (int, string) {


	fmt.Println("Entramos a AddProductQuery")

	// Creamos la variable que tiene la estructura de todo lo que conforma a un producto.
	var t entities.Product

	// Decodificamos lo que nos viene en el endpoint (json) para guardarlo en la estructura.
	err := json.Unmarshal([]byte(body), &t)

	// Verificamos que no haya existido algun error al decodificar el json a la estructura
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	// Verificamos que nos hayan pasado el titulo del Producto
	if len(t.ProdTitle) == 0 {
		return 400, "Se debe especificar el Nombre (Title) del Producto"
	}

	// Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	// Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	// Insertamos el producto
	result, err2 := database.AddProductDatabase(t)
	fmt.Println(result)

	// Verificamos que no haya existido un error al insertar el producto.
	if err2 != nil {
		return 400, "Ocurrio Un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
	//return 200, "Ha llegado correctamente a Query Products"
	
	//return 200, "{ Se inserto correctamente CategID: " + strconv.Itoa(int(result)) + "}"
}