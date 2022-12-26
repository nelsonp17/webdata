package app

import (
	//"fmt"
	"github.com/gorilla/mux"
	//"github.com/Nelson2017-8/webdata/app/controllers"
	"log"
	"net/http"
	"time"
)

func Init() {
	//GetArray(env)
	// Crear el enrutador y definir las rutas en la función definirRutas
	enrutador := mux.NewRouter()
	Routes(enrutador)

	// archivos estatichos (js, css, img, ...)
	//enrutador.PathPrefix("/public/assets/").Handler( http.StripPrefix("/public/assets/", http.FileServer(http.Dir("./public/assets")) ) )
	//enrutador.PathPrefix("/public/assets/").Handler(http.StripPrefix("/public/assets/", http.FileServer(http.Dir(env.Assets))))
	//enrutador.PathPrefix("/public/node_modules/").Handler(http.StripPrefix("/public/node_modules/", http.FileServer(http.Dir("./public/node_modules"))))

	// Dirección del servidor. En este caso solo indicamos el puerto
	// pero podría ser algo como "127.0.0.1:8000" "https://nelson2017-8.github.io"
	direccion := "https://nelson2017-8.github.io"
	//http.ListenAndServe(direccion, enrutador)

	servidor := &http.Server{
		Handler: enrutador,
		Addr:    direccion,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir \n", direccion)
	log.Fatal(servidor.ListenAndServe())
}
