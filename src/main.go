package main // Folder name

import ( // Libraries
	"fmt"
	"github.com/fatih/color"
	"net/http"
    "log"
    "github.com/gorilla/mux"
)
func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Listoooo!!\n"))
}

type Persona struct {
	Nombre     string
	Apellido   string
	DNI        int
	Telefono   string
	Ejercicios bool
}

var subdito = new(Persona)

// Colors
var (
	greenFmt    = color.New(color.FgGreen)
	cyanFmt     = color.New(color.FgCyan)
	boldFmt     = color.New(color.Bold)
	cyanBoldFmt = color.New(color.FgCyan, color.Bold)
)

func main() {

	r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/health", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8080", r))
}

/* func FuncName ( Parameters ) Returns */
func (p Persona) despedirse() {
	fmt.Println("Ahora... Manos a la obra", p.Nombre)
	fmt.Println("")

}

func continuar() {
	cyanFmt.Println("\nPresione una tecla para continuar...")
	fmt.Scanf("%s", &subdito.Ejercicios)
}

func genericExample(example string, function func()) {
	greenFmt.Println("#################################################")
	greenFmt.Println("######### Ejecutando " + example + " Example #########")
	greenFmt.Println("#################################################")
	function()
	continuar()
}
