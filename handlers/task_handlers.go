package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/matildez0/simple-go-mod/models"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

// GET /tasks
func (h *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := make([]models.Task, 0)
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Erro ao ler JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	err := h.DB.QueryRow(query, t.Title, t.Description, t.Status).Scan(&t.ID)

	if err != nil {
		http.Error(w, "Erro na BD: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Println("Erro ao codificar resposta:", err)
	}
}

// PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	query := "UPDATE tasks SET title=$1, description=$2, status=$3 WHERE id=$4"
	result, err := h.DB.Exec(query, t.Title, t.Description, t.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	result, err := h.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
