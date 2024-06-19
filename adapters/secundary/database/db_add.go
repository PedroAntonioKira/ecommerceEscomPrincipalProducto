package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	//"strconv"
	//"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
)

func AddCategoryDatabase(c entities.Category) (int64, error) {
	fmt.Println("Comienza Registro de InsertCategory")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return 0, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	//Nos ayudara a guardar el resultado cuando ejecutemos la sentencia SQL (trae filas afectadas y ultima inserción)
	var result sql.Result

	//Ejecutamos la sentencia SQL
	result, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	//nos regresa el ultimo ID insertado en la base
	LastInsertId, err2 := result.LastInsertId()

	//Verificamos no exista un error al haber preguntado cual era el ultimo ID insertado
	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err
	}

	fmt.Println("Insert Category > Ejecución Exitosa")

	return LastInsertId, nil

}