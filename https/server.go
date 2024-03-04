package https

import (
	"fmt"
	"log"
	"net/http"

	"binod210/task_remainder/helper"
	"binod210/task_remainder/routes"
)

func StartServer(port string, isDebug bool) {
	router := routes.InitRoutes()
	log.Println("server running in :", port)
	helper.StartCron()
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
