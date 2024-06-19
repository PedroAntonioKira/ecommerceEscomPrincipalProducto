package queries_category

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

func AddCategoryQuery(body string, User string) (int, string) {
	//declaramos una variable con la estructura que tiene una categoria de acuerdo a como esta en nuestra base de datos.
	var t entities.Category

	//descontraturamos lo que viene en el body en la estructura
	err := json.Unmarshal([]byte(body), &t)

	//verificamos que no haya un error (que venga en un formato tipo json)
	if err != nil {
		return 400, "Error en los datos recibidos"
	}

	//verificamos que en el json recibido tegamos el campo categName (nombre de la categoria)
	if len(t.CategName) == 0 {
		return 400, "debe especificar el Nombre (Title) de la Categoría"
	}

	//verificamos que en el json recibido tegamos el campo pathName (ruta de la categoria)
	if len(t.CategPath) == 0 {
		return 400, "debe especificar el Path (Ruta) de la Categoría"
	}

	//Verificamos si User Is Admin
	isAdmin, msg := secundary.UserIsAdmin(User)

	//Verificamos si efectivamente no es admin
	if !isAdmin {
		return 400, msg
	}

	//Se realiza el registro de la categoria
	result, err2 := database.AddCategoryDatabase(t)
	//result, err2 := bd.InsertCategory(t)

	fmt.Println(result)
	fmt.Println(err2)

	//Verificamos que no exista un error al intentar realizar el registro de la categoria
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoria " + t.CategName + " > " + err2.Error()
	}
	
	return 200, "{ Se inserto correctamente CategID: " + strconv.Itoa(int(result)) + "}"
}