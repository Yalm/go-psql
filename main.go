package main

import (
	"log"
	"net/http"

	"github.com/Yalm/go-psql/config"
	"github.com/Yalm/go-psql/controllers"
	"github.com/Yalm/go-psql/databases"
	"github.com/Yalm/go-psql/repositories"
	"github.com/Yalm/go-psql/routers"
	"github.com/Yalm/go-psql/services"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

func main() {
	router := mux.NewRouter()

	container := dig.New()
	container.Provide(config.New)
	container.Provide(databases.CreateConnectionToPostgreSql)
	container.Provide(repositories.NewUserRepository)
	container.Provide(services.NewUserService)
	container.Provide(controllers.NewUserController)

	routers.New(router, container)

	log.Printf("server listening at: %s", "8080")
	log.Fatalln(http.ListenAndServe(":8080", router))
	// xlsx := excelize.NewFile()
	// streamWriter, err := xlsx.NewStreamWriter("Sheet1")

	// products := []product{
	// 	{"Pizza", 10.6},
	// 	{"Ham", 5.5},
	// }

	// salt := 2
	// streamWriter.SetRow("A1", []interface{}{
	// 	excelize.Cell{Value: "Nombre"},
	// 	excelize.Cell{Value: "Precio"},
	// })

	// file, err := os.Create("Book1.xlsx")
	// check(err)
	// defer file.Close()

	// for index, element := range products {
	// 	index += salt
	// 	if err := streamWriter.SetRow(fmt.Sprintf("A%d", index), []interface{}{
	// 		excelize.Cell{Value: element.name},
	// 		excelize.Cell{Value: element.price},
	// 	}); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	// if err := streamWriter.Flush(); err != nil {
	// 	fmt.Println(err)
	// }

	// if err := xlsx.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
}
