package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/utils"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func UpdateCategoryQuery(c entities.Category) error {

	fmt.Println("Comienza Registro de UpdateCategory")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "UPDATE category SET "

	//Verificamos si estamos recibiendo "nombre de la cateria" para actualizar
	if len(c.CategName) > 0 {
		sentencia += " Categ_Name = '" + utils.EscapeString(c.CategName) + "'"
	}

	//Verificamos si estamos recibiendo "ruta de la cateria" para actualizar
	if len(c.CategPath) > 0 {
		//Verificamos si previamente ya le habiamos establecido un "nombre de la cateria"
		if !strings.HasSuffix(sentencia, "SET ") {
			//En caso de no termine con "SET", entonces almacenamos una coma para separar las sentencias.
			sentencia += ", "
		}
		sentencia += "Categ_Path = '" + utils.EscapeString(c.CategPath) + "'"
	}

	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	//Ejecutamos la sentencia SQL
	_, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución Exitosa !")

	return nil

}