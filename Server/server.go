package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

var students = make(map[string]map[string]float32)
var subjects = make(map[string]map[string]float32)

//Server is...
type Server struct{}

//AddNewScore is...
func (server *Server) AddNewScore(name string, subject string, score float32) error {
	if _, exists := subjects[subject][name]; !exists {
		students[name][subject] = score
		subjects[subject][name] = score
		return nil
	} else {
		return errors.New("Score already captured")
	}
}

func (server *Server) getStudentAverage(name string, reply *float32) error {
	if _, exists := students[name]; exists {
		var average float32 = 0

		for _, v := range students[name] {
			average += v
		}

		average /= (float32)(len(students[name]))
		*reply = average
	}
	return nil
}

func (server *Server) getAllStudentsAverage(reply *float32) error {
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

func (server *Server) getSubjectAverage(subject string, reply *float32) error {
	if _, exists := subjects[subject]; exists {
		var average float32 = 0

		for _, v := range subjects[subject] {
			average += v
		}

		average /= (float32)(len(subjects[subject]))
		*reply = average
	}
	return nil
}

func runServer() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go runServer()

	var input string
	fmt.Scanln(&input)
}
