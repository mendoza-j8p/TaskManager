package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)
func TestTaskManager(t *testing.T) {
	// Crear una nueva instancia del administrador de tareas
	taskManager := NewTaskManager()

	// Prueba para crear una tarea y verificar que se agregue correctamente
	taskManager.CreateTask("Enviar correo de seguimiento", time.Now().Add(30*time.Second))
	tasks := taskManager.GetTasks()
	if len(tasks) != 1 {
		t.Errorf("Se esperaba una tarea, pero se encontraron %d", len(tasks))
	}

	// Prueba para marcar una tarea como completada y verificar que se elimine de la lista
	err := taskManager.CompleteTask(1)
	if err != nil {
		t.Errorf("Error al marcar la tarea como completada: %v", err)
	}
	tasks = taskManager.GetTasks()
	if len(tasks) != 0 {
		t.Errorf("Se esperaba que la lista de tareas estuviera vacía, pero se encontraron %d tareas", len(tasks))
	}

	// Prueba para listar las tareas
	taskManager.CreateTask("Completar informe", time.Now().Add(45*time.Second))
	taskManager.CreateTask("Realizar llamada de seguimiento", time.Now().Add(1*time.Minute))
	tasks = taskManager.GetTasks()
	if len(tasks) != 2 {
		t.Errorf("Se esperaban dos tareas, pero se encontraron %d", len(tasks))
	}
	err = taskManager.CompleteTask(1)
	if err == nil {
    	t.Error("Se esperaba un error al marcar una tarea inexistente como completada, pero no se produjo ningún error")
	} else {
	expectedError := "tarea no encontrada"
	if err.Error() != expectedError {
    	t.Errorf("Se esperaba el mensaje de error '%s', pero se obtuvo '%s'", expectedError, err.Error())
	}
}
}
func TestExpiredTasks(t *testing.T) {
	// Crear una nueva instancia del administrador de tareas
	taskManager := NewTaskManager()

	// Crear una tarea vencida
	taskManager.CreateTask("Tarea vencida", time.Now().Add(-1*time.Hour))

	// Verificar que la tarea se haya marcado como completada automáticamente
	time.Sleep(2 * time.Second) // Esperar un momento para que se procese la tarea vencida
	tasks := taskManager.GetTasks()
	if len(tasks) != 0 {
		t.Errorf("Se esperaba que la tarea vencida se marcara como completada, pero aún se encontró en la lista")
	}
}

func TestConcurrentTasks(t *testing.T) {
	// Crear una nueva instancia del administrador de tareas
	taskManager := NewTaskManager()

	// Crear tareas concurrentemente
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			taskManager.CreateTask(fmt.Sprintf("Tarea concurrente %d", id), time.Now().Add(30*time.Second))
		}(i)
	}

	wg.Wait()

	// Obtener las tareas y verificar su cantidad
	tasks := taskManager.GetTasks()
	if len(tasks) != 10 {
		t.Errorf("Se esperaban 10 tareas, pero se encontraron %d", len(tasks))
	}

	// Marcar una tarea inexistente como completada y verificar el error
	err := taskManager.CompleteTask(10)
	if err == nil {
		t.Errorf("Se esperaba un error al marcar la tarea 10 como completada, pero no se produjo ningún error")
	} else {
		expectedError := "tarea no encontrada"
		if err.Error() != expectedError {
			t.Errorf("Se esperaba el mensaje de error '%s', pero se obtuvo '%s'", expectedError, err.Error())
		}
	}
}

