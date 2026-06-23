package main

import (
	
	"monitoring-edc/internal/database"
	"monitoring-edc/internal/monitoring"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	repo := monitoring.NewRepository(database.DB)

	service := monitoring.NewService(repo)
	//handler := monitoring.NewHandler(service)
	importService := importer.NewService(service)
	handler := monitoring.NewHandler(
		service,
		importService
	)

	r := gin.Default()

	r.GET("/monitoring", handler.GetAll)

	if err := r.Run(":8989"); err != nil {
		panic(err)
	}
	// importService := importer.NewService(
	// 	service,
	// )

	// ctx := context.Background()
	// err := importService.ImportFile(
	// 	ctx,
	// 	`d:\edc0206.xlsx`,
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Import Berhasil")

	// total, err := service.Count()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Total Data :", total)
}
