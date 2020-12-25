package main

import "fmt"

func main(){
	var num1 float64
	var num2 float64
	var opcion float64
	
	fmt.Println("Calculadora\n", "Opciones:")
	fmt.Println(`
		1) Suma
		2) Resta
		3) Multiplicación
		4) División
		5) Salir
		`)
	fmt.Println("Ingrese la opción que desea")	
	fmt.Scanln(&opcion)
	fmt.Println("\n--------------------------------------")
	

	for opcion >0 && opcion <5 {
	
		if opcion == 1 {	
		fmt.Println("Ingrese un número: ")
		fmt.Scanln(&num1)
		fmt.Println("\n--------------------------------------")
		fmt.Println("Ingrese otro número: \n")
		fmt.Scanln(&num2)
		fmt.Println("\n--------------------------------------")
		res := num1 + num2
		fmt.Println("El resultado es: ", res)

		fmt.Println(`
		1) Suma
		2) Resta
		3) Multiplicación
		4) División
		5) Salir
		`)
		fmt.Println("Ingrese la opción que desea")	
		fmt.Scanln(&opcion)
		fmt.Println("\n--------------------------------------")

	} else if opcion == 2 {
		fmt.Println("Ingrese un número: ")
		fmt.Scanln(&num1)
		fmt.Println("\n--------------------------------------")
		fmt.Println("Ingrese otro número: \n")
		fmt.Scanln(&num2)
		fmt.Println("\n--------------------------------------")
		res := num1 - num2
		fmt.Println("El resultado es: ", res)

		fmt.Println(`
		1) Suma
		2) Resta
		3) Multiplicación
		4) División
		5) Salir
		`)
		fmt.Println("Ingrese la opción que desea")	
		fmt.Scanln(&opcion)
		fmt.Println("\n--------------------------------------")
	} else if opcion == 3 {
		fmt.Println("Ingrese un número: ")
		fmt.Scanln(&num1)
		fmt.Println("\n--------------------------------------")
		fmt.Println("Ingrese otro número: \n")
		fmt.Scanln(&num2)
		fmt.Println("\n--------------------------------------")
		res := num1 * num2
		fmt.Println("El resultado es: ", res)

		fmt.Println(`
		1) Suma
		2) Resta
		3) Multiplicación
		4) División
		5) Salir
		`)
		fmt.Println("Ingrese la opción que desea")	
		fmt.Scanln(&opcion)
		fmt.Println("\n--------------------------------------")

	}else if opcion == 4 {
		fmt.Println("Ingrese un número: ")
		fmt.Scanln(&num1)
		fmt.Println("\n--------------------------------------")
		fmt.Println("Ingrese otro número: ")
		fmt.Scanln(&num2)
			if num2 == 0{
			fmt.Println("No se puede dividir por cero")
				}else{
				res := num1/num2
				fmt.Println("El resultado es: ", res)

				fmt.Println(`
		1) Suma
		2) Resta
		3) Multiplicación
		4) División
		5) Salir
		`)
		fmt.Println("Ingrese la opción que desea")	
		fmt.Scanln(&opcion)
		fmt.Println("\n--------------------------------------")

			}
		
	}else if opcion == 5 {
		fmt.Println("Bye World")
	}
}
	


}