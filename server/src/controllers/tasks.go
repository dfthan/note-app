package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	type Task struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}
	tasks := []Task {
		{ID: 1, Description: "Task 1", Completed: false},
		{ID: 2, Description: "Task 2", Completed: false},
		{ID: 3, Description: "Task 3", Completed: false},
	}

	if r.URL.Path == "/api/tasks/" {	
		switch r.Method {
			case "GET":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(tasks)
			case "POST":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(tasks)
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"message": "Not found"})
		}
	} else {
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
    		http.Error(w, "enter /api/tasks/<id>", http.StatusBadRequest)
    		return
  		}
		id, err := strconv.Atoi(pathParts[2])
		if err != nil {
			http.Error(w, "enter /api/tasks/<id>", http.StatusBadRequest)
			return
		}

		// quick hack fix for the list position =)
		id -= 1
		if len(tasks) <= id {
			http.Error(w, "Task not found", http.StatusBadRequest)
			return
		}

		switch r.Method {
			case "GET":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(tasks[id])
			case "PUT":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(tasks[id])
			case "DELETE":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(tasks[id])
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"message": "Not found"})
		}
	

	}
}

