package handlers

import (
	"net/http"

	scheduleusecase "example.com/taskservice/internal/usecase/schedule"
)

type ScheduleHandler struct {
	usecase scheduleusecase.Usecase
}

func NewScheduleHandler(usecase scheduleusecase.Usecase) *ScheduleHandler {
	return &ScheduleHandler{usecase: usecase}
}

func (h *ScheduleHandler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	taskId, err := getTaskIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var req scheduleMutationDTO
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	created, err := h.usecase.CreateSchedule(r.Context(), taskId, scheduleusecase.CreateScheduleInput{
		RuleType:      req.RuleType,
		StartFrom:     req.StartFrom,
		Interval:      req.Interval,
		MonthDay:      req.MonthDay,
		IsEven:        req.IsEven,
		SpecificDates: req.SpecificDates,
	})
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, newScheduleDTO(created))
}

func (h *ScheduleHandler) GetScheduleByID(w http.ResponseWriter, r *http.Request) {
	scheduleId, err := getScheduleIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	schedule, err := h.usecase.GetScheduleByID(r.Context(), scheduleId)
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, newScheduleDTO(schedule))
}

func (h *ScheduleHandler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleId, err := getScheduleIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var req scheduleMutationDTO
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	updated, err := h.usecase.UpdateSchedule(r.Context(), scheduleId, scheduleusecase.UpdateScheduleInput{
		RuleType:      req.RuleType,
		StartFrom:     req.StartFrom,
		Interval:      req.Interval,
		MonthDay:      req.MonthDay,
		IsEven:        req.IsEven,
		SpecificDates: req.SpecificDates,
	})
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, newScheduleDTO(updated))
}

func (h *ScheduleHandler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleId, err := getScheduleIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.usecase.DeleteSchedule(r.Context(), scheduleId); err != nil {
		writeUsecaseError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ScheduleHandler) ScheduleList(w http.ResponseWriter, r *http.Request) {
	taskId, err := getTaskIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	schedules, err := h.usecase.ScheduleList(r.Context(), taskId)
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	response := make([]scheduleDTO, 0, len(schedules))
	for i := range schedules {
		response = append(response, newScheduleDTO(&schedules[i]))
	}

	writeJSON(w, http.StatusOK, response)
}
