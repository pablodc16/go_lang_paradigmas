package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Conexion Server ---------------------------------------------------------------------------------------------------------------------------------------------------------------
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API, thank's...")

}

// Function main -----------------------------------------------------------------------------------------------------------------------------------------------------------------------
func main() {

	// Server ------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)

	// CRUD DB ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	// Listening port Server ----------------------------------------------------------------------------------------------------------------------------------------------------------
	log.Fatal(http.ListenAndServe(":3000", router))
}

// Struc DB ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------
type task struct {
	ID      int    `json:ID` //alt 96 comillas invertidas
	Name    string `json:Name`
	Content string `jason:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task class Paradigmas",
		Content: "Some Content in class",
	},
}

// Create -----------------------------------------------------------------------------------------------------------------------------------------------------------------------
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(newTask)
}

// Read -----------------------------------------------------------------------------------------------------------------------------------------------------------------------
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range tasks {

		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
		}

	}
}

// delete -----------------------------------------------------------------------------------------------------------------------------------------------------------------------
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...) // para eliminar buscamos el indice de la tarea y declaramos la lista sin el indice encontrado
			fmt.Fprintf(w, "The task with ID %v has been remove succesfully", taskID)
		}
	}
}
