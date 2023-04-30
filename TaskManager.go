package main

import (
	"fmt"
	"sync"
	"time"
)

// Estructura para representar una tarea
type Task struct {
	ID       int
	Name     string
	DueDate  time.Time
	Complete bool
}

// Estructura para representar el administrador de tareas
type TaskManager struct {
	tasks   []Task
	mutex   sync.Mutex
	stop    chan struct{}
	eventCh chan string
}

// Función para crear una nueva instancia de TaskManager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:   make([]Task, 0),
		mutex:   sync.Mutex{},
		stop:    make(chan struct{}),
		eventCh: make(chan string),
	}
}

// Función para crear una nueva tarea
func (tm *TaskManager) CreateTask(name string, dueDate time.Time) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	taskID := len(tm.tasks) + 1
	task := Task{
		ID:       taskID,
		Name:     name,
		DueDate:  dueDate,
		Complete: false,
	}
	tm.tasks = append(tm.tasks, task)

	fmt.Println("Nueva tarea creada:", name)
}

// Función para obtener la lista de tareas
func (tm *TaskManager) GetTasks() []Task {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	return tm.tasks
}

// Función para marcar una tarea como completada
func (tm *TaskManager) CompleteTask(taskID int) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks[i].Complete = true
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("tarea no encontrada")
}

// Rutina de fondo para actualizar las tareas vencidas
func (tm *TaskManager) BackgroundTaskUpdater() {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			tm.mutex.Lock()

			now := time.Now()
			for i, task := range tm.tasks {
				if task.DueDate.Before(now) && !task.Complete {
					tm.tasks[i].Complete = true
					tm.eventCh <- fmt.Sprintf("La tarea '%s' ha vencido", task.Name)
				}
			}

			tm.mutex.Unlock()
		case <-tm.stop:
			return
		}
	}
}

// Función para detener el administrador de tareas
func (tm *TaskManager) Stop() {
	close(tm.stop)
}


func main() {
	// Crear una instancia del administrador de tareas
	taskManager := NewTaskManager()

	// Crear algunas tareas de ejemplo
	taskManager.CreateTask("Enviar correo de seguimiento", time.Now().Add(30*time.Second))
	taskManager.CreateTask("Completar informe", time.Now().Add(45*time.Second))
	taskManager.CreateTask("Realizar llamada de seguimiento", time.Now().Add(1*time.Minute))

	// Iniciar la rutina de fondo para actualizar tareas vencidas
	go taskManager.BackgroundTaskUpdater()

	// Crear un canal para recibir una señal de finalización
	finish := make(chan bool)

	// Goroutine para permitir la interacción del usuario
	go func() {
		for {
			// Mostrar opciones de acciones al usuario
			fmt.Println("Acciones disponibles:")
			fmt.Println("1. Listar tareas")
			fmt.Println("2. Marcar tarea como completada")
			fmt.Println("3. Salir")

			// Leer la opción del usuario
			var option int
			fmt.Print("Seleccione una opción: ")
			fmt.Scanln(&option)

			// Realizar la acción seleccionada
			switch option {
			case 1:
				fmt.Println("Tareas:")
				tasks := taskManager.GetTasks()
				for _, task := range tasks {
					fmt.Printf("- ID: %d, Nombre: %s (Vence en: %s)\n", task.ID, task.Name, time.Until(task.DueDate))
				}
				fmt.Println()
			
			case 2:
				fmt.Print("Ingrese el ID de la tarea a marcar como completada: ")
				var taskID int
				fmt.Scanln(&taskID)
				err := taskManager.CompleteTask(taskID)
				if err != nil {
					fmt.Printf("Error al completar la tarea: %v\n", err)
				} else {
					fmt.Println("Tarea completada correctamente.")
				}

			case 3:
				fmt.Println("Saliendo del programa...")
				close(finish)
				taskManager.Stop()
				return
			default:
				fmt.Println("Opción inválida.")
			}
		}
	}()

	// Esperar hasta que se reciba la señal de finalización o se cumpla el tiempo límite
	select {
	case <-finish:
		return
	case <-time.After(60 * time.Second):
		fmt.Println("Tiempo límite alcanzado. Saliendo del programa...")
		return
	}
}
