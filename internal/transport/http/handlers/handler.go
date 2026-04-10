package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	taskdomain "example.com/taskservice/internal/domain/task"
	taskusecase "example.com/taskservice/internal/usecase/task"
	scheduledomain "example.com/taskservice/internal/domain/schedule"
	scheduleusecase "example.com/taskservice/internal/usecase/schedule"
)

func getScheduleIDFromRequest(r *http.Request) (int64, error) {
	rawID := mux.Vars(r)["scheduleId"]

	return validateIDFromRequest(rawID)
}

func getTaskIDFromRequest(r *http.Request) (int64, error) {
	rawID := mux.Vars(r)["taskId"]

	return validateIDFromRequest(rawID)
}

func validateIDFromRequest(rawID string) (int64, error) {
	if rawID == "" {
		return 0, errors.New("missing id")
	}

	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	if id <= 0 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}

func writeUsecaseError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, taskdomain.ErrNotFound), errors.Is(err, scheduledomain.ErrNotFound):
		writeError(w, http.StatusNotFound, err)
	case errors.Is(err, taskusecase.ErrInvalidInput), errors.Is(err, scheduleusecase.ErrInvalidInput):
		writeError(w, http.StatusBadRequest, err)
	default:
		writeError(w, http.StatusInternalServerError, err)
	}
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(payload)
}
