package main

//importar librerías
import (
	"context"
	"database/sql" //para importar una libreria de golang
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // para importar librerias que no proceden de golang pero que igualmente la vamos a utilizar
)

type Author struct {
	ID   int64
	Name string
}

type Book struct {
	ID          int64
	Title       string
	AuthorID    string
	PublishDate time.Time
	Author      Author
}

// función principal
func main() {

	//context para la llamada del query
	ctx := context.Background()

	db, err := createConnection()
	if err != nil {
		panic(err)
	}

	//nos pide un context... se lo creamos arriba y nos pide la base de datos, nos mostrara un error dependiendo y validamos

	//-----------------------------------------------------------------------------------------------------------------------------------------
	//Consultar todo el array
	err = queryBooks(ctx, db)
	if err != nil {
		panic(err)
	}

	//-----------------------------------------------------------------------------------------------------------------------------------------
	//Consultar un solo libro
	// err = queryBook(ctx, db)
	// if err != nil {
	// 	panic(err)
	// }

	//-----------------------------------------------------------------------------------------------------------------------------------------
	//Consultar por id de autor
	// err = queryBooksAuthor(ctx, db)
	// if err != nil {
	// 	panic(err)
	// }

	//-----------------------------------------------------------------------------------------------------------------------------------------
	//Agregar un libro
	// err = addBook(ctx, db, "The matrix comics", 1, time.Now())
	// if err != nil {
	// 	panic(err)
	// }

	db.Close() // cerrar la conexión
}

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Crear una conexion a la BD
func createConnection() (*sql.DB, error) {
	// conexión con MySQL y con la base de datos específica tenemos que agregar "?parseTime=True" para que no diera error de conversion []IUT8
	connectionString := "root:Shoudymella1986*@tcp(localhost:3306)/books?parseTime=True"

	db, err := sql.Open("mysql", connectionString) // puede conectar o dar un error
	if err != nil {                                //validar conexión
		return nil, err
	}

	//pool de conexiones
	db.SetMaxOpenConns(5) //la cantidad máxima de conexiones abiertas

	err = db.Ping() //validar ping
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Consultar un solo libro
func queryBook(ctx context.Context, db *sql.DB) error { //esta funcion hace una conexion a la base datos para realizar una consulta

	////query para consultar, damos un argumento con ? abajo damos numero 3 entonces este debería referenciar al id=3 que le pasamos, titulo "Dune" y el id del autor # "1"

	//////descomentar desde aquí
	qry := `select b.id, b.title, b.author_id, b.publish_date from books b where b.id = ?;` // para obtener uno en especial

	////hay diferentes formas de traer un query para la presentacion escogimos QueryRowContext, pero hay 4 formas de traerlo aquí le pusimos el argumento numero 3
	////pasamos el ctx = contexto, el qry = consulta, argumento = 3 la cual nos va a devolver una sola linea

	row := db.QueryRowContext(ctx, qry, 3) //para obtener uno específico

	////para mostrar los datos despues de scanear el qry
	var id int64 //como buena practica utilizar enteros64 para los ids de la base de datos
	var title string
	var authorID int64
	var publish_date time.Time

	err := row.Scan(&id, &title, &authorID, &publish_date) //metodo Scan nos es posible escanear la informacion que obtenemos del qry
	if err != nil {
		return err
	}

	fmt.Println("ROW: ", id, title, authorID, publish_date)
	//////hasta aquí
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Consultar todo el array de libros
func queryBooks(ctx context.Context, db *sql.DB) error {
	/////para consultarlos todos
	qry := `select b.id, b.title, b.author_id, b.publish_date from books b;` // para obtenerlos todos

	rows, err := db.QueryContext(ctx, qry)
	if err != nil {
		return err
	}

	books := []Book{}

	for rows.Next() {
		b := Book{}

		err = rows.Scan(&b.ID, &b.Title, &b.AuthorID, &b.PublishDate)
		if err != nil {
			return err
		}

		books = append(books, b)
	}

	fmt.Println(books)

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Consultar un autor
func queryBooksAuthor(ctx context.Context, db *sql.DB) error {
	/////para consultarlos todos
	qry := `SELECT b.id, b.title, b.author_id, a.name AS 'author', b.publish_date FROM books b INNER JOIN authors a ON a.id = b.author_id;` // para obtenerlos todos

	rows, err := db.QueryContext(ctx, qry)
	if err != nil {
		return err
	}

	books := []Book{}

	for rows.Next() {
		b := Book{}

		err = rows.Scan(&b.ID, &b.Title, &b.AuthorID, &b.Author.Name, &b.PublishDate)
		if err != nil {
			return err
		}

		books = append(books, b)
	}

	fmt.Println(books)

	return nil
}

//---------------------------------------------------------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------------------------------------------------------------------------------------
//Agregar un libro

func addBook(ctx context.Context, db *sql.DB, title string, authorID int, publishDate time.Time) error {
	qryadd := `insert into books(title, author_id, publish_date) values (?,?,?);`

	result, err := db.ExecContext(ctx, qryadd, title, authorID, publishDate)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("INSERT ID: ", id)

	return nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------------
