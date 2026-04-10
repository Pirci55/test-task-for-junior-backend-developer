package transporthttp

import (
	"net/http"

	"github.com/gorilla/mux"

	swaggerdocs "example.com/taskservice/internal/transport/http/docs"
	httphandlers "example.com/taskservice/internal/transport/http/handlers"
)

func NewRouter(taskHandler *httphandlers.TaskHandler, scheduleHandler *httphandlers.ScheduleHandler, docsHandler *swaggerdocs.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/swagger/openapi.json", docsHandler.ServeSpec).Methods(http.MethodGet)
	router.HandleFunc("/swagger/", docsHandler.ServeUI).Methods(http.MethodGet)
	router.HandleFunc("/swagger", docsHandler.RedirectToUI).Methods(http.MethodGet)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/tasks", taskHandler.CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/tasks", taskHandler.TaskList).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{taskId:[0-9]+}", taskHandler.GetTaskByID).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{taskId:[0-9]+}", taskHandler.UpdateTask).Methods(http.MethodPut)
	api.HandleFunc("/tasks/{taskId:[0-9]+}", taskHandler.DeleteTask).Methods(http.MethodDelete)

	api.HandleFunc("/tasks/{taskId:[0-9]+}/schedule", scheduleHandler.CreateSchedule).Methods(http.MethodPost)
	api.HandleFunc("/tasks/{taskId:[0-9]+}/schedule", scheduleHandler.ScheduleList).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{taskId:[0-9]+}/schedule/{scheduleId:[0-9]+}", scheduleHandler.GetScheduleByID).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{taskId:[0-9]+}/schedule/{scheduleId:[0-9]+}", scheduleHandler.UpdateSchedule).Methods(http.MethodPut)
	api.HandleFunc("/tasks/{taskId:[0-9]+}/schedule/{scheduleId:[0-9]+}", scheduleHandler.DeleteSchedule).Methods(http.MethodDelete)

	return router
}
