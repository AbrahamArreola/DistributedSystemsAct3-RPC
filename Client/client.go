package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

//Data is...
type Data struct {
	Student string
	Subject string
	Score   float32
}

var scanner = bufio.NewReader(os.Stdin)
var dataByte = make([]byte, 100)

func runClient() {
	connection, err := rpc.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	var menuOpc int
	var data Data
	for menuOpc != 5 {
		fmt.Println("1. Agregar calificación\n2. Obtener promedio de alumno\n3. Obtener promedio general" +
			"\n4. Obtener promedio de materia\n5. Salir")

		fmt.Print("Seleccione un opción: ")
		fmt.Scanln(&menuOpc)

		switch menuOpc {
		case 1:
			var message string

			fmt.Println("\n********** Agregar calificación **********")
			fmt.Print("Ingrese nombre del alumno: ")
			dataByte, _, _ = scanner.ReadLine()
			data.Student = string(dataByte)
			fmt.Print("Ingrese nombre de la materia: ")
			dataByte, _, _ = scanner.ReadLine()
			data.Subject = string(dataByte)
			fmt.Print("Ingrese calificación: ")
			fmt.Scanln(&data.Score)

			err = connection.Call("Server.AddNewScore", data, &message)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.AddNewsCore:", message)
			}
			fmt.Println("******************************************")

		case 2:
			fmt.Println("\n********** Obtener promedio de alumno **********")
			fmt.Print("Ingrese nombre del alumno: ")
			dataByte, _, _ = scanner.ReadLine()
			data.Student = string(dataByte)

			err = connection.Call("Server.GetStudentAverage", data.Student, &data.Score)
			if err != nil {
				fmt.Println(err)
			} else {
				if data.Score != 0 {
					fmt.Println("Promedio: ", data.Score)
				} else {
					fmt.Println("Alumno inexistente")
				}
			}
			fmt.Println("************************************************")

		case 3:
			fmt.Println("\n********** Obtener promedio general **********")
			err = connection.Call("Server.GetAllStudentsAverage", data.Score, &data.Score)
			if err != nil {
				fmt.Println(err)
			} else {
				if data.Score != 0 {
					fmt.Println("Promedio general: ", data.Score)
				} else {
					fmt.Println("No hay alumnos capturados")
				}
			}
			fmt.Println("************************************************")

		case 4:
			fmt.Println("\n********** Obtener promedio de materia **********")
			fmt.Print("Ingrese nombre de la materia: ")
			dataByte, _, _ = scanner.ReadLine()
			data.Subject = string(dataByte)

			err = connection.Call("Server.GetSubjectAverage", data.Subject, &data.Score)
			if err != nil {
				fmt.Println(err)
			} else {
				if data.Score != 0 {
					fmt.Println("Promedio: ", data.Score)
				} else {
					fmt.Println("Materia inexistente")
				}
			}
			fmt.Println("*************************************************")
		}
	}
}

func main() {
	runClient()
}
