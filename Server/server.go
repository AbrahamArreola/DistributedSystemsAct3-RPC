package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

//Data is...
type Data struct {
	Student string
	Subject string
	Score   float32
}

//Server is...
type Server struct{}

var students = make(map[string]map[string]float32)
var subjects = make(map[string]map[string]float32)
var scholarData Data

//AddNewScore is...
func (this *Server) AddNewScore(data Data, reply *string) error {
	if _, exists := subjects[data.Subject]; !exists {
		subjects[data.Subject] = make(map[string]float32)
	}

	if _, exists := students[data.Student]; !exists {
		students[data.Student] = make(map[string]float32)
	}

	if _, exists := subjects[data.Subject][data.Student]; !exists {
		students[data.Student][data.Subject] = data.Score
		subjects[data.Subject][data.Student] = data.Score
		*reply = "Calificación agregada!"
		return nil
	} else {
		return errors.New("Error: calificación existente")
	}
}

//GetStudentAverage is...
func (this *Server) GetStudentAverage(name string, reply *float32) error {
	if _, exists := students[name]; exists {
		var average float32 = 0

		for _, v := range students[name] {
			average += v
		}

		average /= float32(len(students[name]))
		*reply = average
	}
	return nil
}

//GetAllStudentsAverage is...
func (this *Server) GetAllStudentsAverage(data float32, reply *float32) error {
	if len(students) > 0 {
		var studentAverage float32 = 0
		var totalAverage float32 = 0

		for _, value := range students {
			for _, v := range value {
				studentAverage += v
			}

			studentAverage /= (float32)(len(value))
			totalAverage += studentAverage
			studentAverage = 0
		}

		totalAverage /= (float32)(len(students))
		*reply = totalAverage
	}
	return nil
}

//GetSubjectAverage is...
func (this *Server) GetSubjectAverage(subject string, reply *float32) error {
	if _, exists := subjects[subject]; exists {
		var average float32 = 0

		for _, v := range subjects[subject] {
			average += v
		}

		average /= float32(len(subjects[subject]))
		*reply = average
	}
	return nil
}

func runServer() {
	fmt.Println("Servidor corriendo...")
	rpc.Register(new(Server))
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(connection)
	}
}

func main() {
	go runServer()

	var input string
	fmt.Scanln(&input)
}
