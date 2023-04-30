# Administrador de Tareas

El Administrador de Tareas es una aplicación que te permite crear y gestionar tareas con fechas de vencimiento. Puedes agregar nuevas tareas, listarlas, marcarlas como completadas y recibir notificaciones cuando una tarea ha vencido.

## Requisitos

Go (versión 1.16 o superior)

## Instalación

1. Clona este repositorio en tu máquina local:

```bash
git clone <URL del repositorio>
```

2. Navega hasta el directorio del proyecto:

```bash
cd TaskManager
```

3. Compila y ejecuta el programa:

```bash
go run TaskManager.go
```

1. Sigue las instrucciones en la terminal para interactuar con el Administrador de Tareas.

## Uso

Al ejecutar el programa, se mostrará un menú con las siguientes opciones:

1. Listar tareas: Muestra la lista de tareas existentes con sus respectivas fechas de vencimiento.
2. Marcar tarea como completada: Permite marcar una tarea como completada ingresando su ID.
3. Salir: Cierra el programa y detiene el Administrador de Tareas.

Puedes seleccionar una opción ingresando el número correspondiente y presionando Enter.

## Funcionalidades

- Crear tareas: Puedes crear nuevas tareas proporcionando un nombre y una fecha de vencimiento.
- Listar tareas: Obtén una lista de todas las tareas existentes y visualiza su ID, nombre y tiempo restante hasta la fecha de vencimiento.
- Marcar tarea como completada: Marca una tarea específica como completada mediante su ID.
- Actualización automática: El sistema actualiza automáticamente las tareas vencidas y envía notificaciones cuando una tarea ha vencido.

## Contribución

Si deseas contribuir a este proyecto, puedes realizar los siguientes pasos:

1. Haz un fork del repositorio.
2. Crea una rama con la nueva funcionalidad o solución de problemas: git checkout -b nueva-funcionalidad.
3. Realiza los cambios necesarios y realiza los commits: git commit -m "Agregar nueva funcionalidad".
4. Haz push a la rama creada en tu repositorio: git push origin nueva-funcionalidad.
5. Crea una solicitud de extracción en el repositorio original.

Asegúrate de seguir las buenas prácticas de desarrollo y agregar pruebas adecuadas si es necesario.

## Autor

Jenifer Mendoza
