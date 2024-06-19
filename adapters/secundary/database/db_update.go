package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
//	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/utils"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func UpdateCategoryQuery(c entities.Product) error {

	fmt.Println("Comienza Registro de UpdateCategory")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	
	fmt.Println("Update Category > Ejecución Exitosa !")

	return nil

}