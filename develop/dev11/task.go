package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var lastID = 0
var lastIDMutex = sync.Mutex{}

type Calendar struct {
	Events map[int]Event
	mu     sync.RWMutex
}

type Event struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

type Result struct {
	Result []Event `json:"result"`
}

// SerializeEventSlice преобразует срез событий в JSON.
func SerializeEventSlice(events []Event) ([]byte, error) {
	data := Result{events}
	result, err := json.Marshal(data)
	return result, err
}

// NewCalendar создает новый календарь.
func NewCalendar() *Calendar {
	return &Calendar{Events: make(map[int]Event), mu: sync.RWMutex{}}
}

// NewEvent создает новое событие с уникальным идентификатором.
func NewEvent(time time.Time, name string) *Event {
	lastIDMutex.Lock()
	lastID++
	lastIDMutex.Unlock()
	return &Event{ID: lastID, Name: name, Time: time}
}

// CreateEvent добавляет событие в календарь.
func (c *Calendar) CreateEvent(event *Event) {
	c.mu.Lock()
	c.Events[event.ID] = *event
	c.mu.Unlock()
}

// UpdateEvent обновляет событие по ID.
func (c *Calendar) UpdateEvent(id int, time time.Time, name string) error {
	c.mu.RLock()
	event, ok := c.Events[id]
	if !ok {
		c.mu.RUnlock()
		return errors.New("no such event")
	}
	c.mu.RUnlock()

	if !time.IsZero() {
		event.Time = time
	}
	if name != "" {
		event.Name = name
	}
	c.mu.Lock()
	c.Events[id] = event
	c.mu.Unlock()
	return nil
}

// DeleteEvent удаляет событие по ID.
func (c *Calendar) DeleteEvent(id int) (*Event, error) {
	c.mu.RLock()
	if _, ok := c.Events[id]; !ok {
		c.mu.RUnlock()
		return nil, errors.New("no such event")
	}
	c.mu.RUnlock()
	c.mu.Lock()
	deleted := c.Events[id]
	delete(c.Events, id)
	c.mu.Unlock()
	return &deleted, nil
}

// EventsForDay возвращает события на текущий день.
func (c *Calendar) EventsForDay() []Event {
	var result []Event
	tYear, tMonth, tDay := time.Now().Date() // today
	c.mu.RLock()
	for _, v := range c.Events {
		year, month, day := v.Time.Date()
		if tYear == year && tMonth == month && tDay == day {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

// EventsForWeek возвращает события на текущую неделю.
func (c *Calendar) EventsForWeek() []Event {
	var result []Event
	tYear, tWeek := time.Now().ISOWeek()
	c.mu.RLock()
	for _, v := range c.Events {
		year, week := v.Time.ISOWeek()
		if tYear == year && tWeek == week {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

// EventsForMonth возвращает события на текущий месяц.
func (c *Calendar) EventsForMonth() []Event {
	var result []Event
	tYear, tMonth, _ := time.Now().Date() // today
	c.mu.RLock()
	for _, v := range c.Events {
		year, month, _ := v.Time.Date()
		if tYear == year && tMonth == month {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

// CalendarHandler структура для обработчика HTTP запросов, работающего с календарем.
type CalendarHandler struct {
	calendar *Calendar
}

// NewCalendarHandler конструктор для обработчика HTTP запросов, работающего с календарем.
func NewCalendarHandler() *CalendarHandler {
	return &CalendarHandler{calendar: NewCalendar()}
}

// CreateEventRequest обработчик запроса на создание события.
func (c *CalendarHandler) CreateEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SendResult(w, []byte("method not allowed"))
		return
	}

	newEvent, err := ParseCreateRequest(r)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	c.calendar.CreateEvent(newEvent)
	SendResult(w, []byte("new event created"))
}

// UpdateEventRequest обработчик запроса на обновление события.
func (c *CalendarHandler) UpdateEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SendResult(w, []byte("method not allowed"))
		return
	}

	id, time, name, err := ParseUpdateRequest(r)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	err = c.calendar.UpdateEvent(id, time, name)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	SendResult(w, []byte(fmt.Sprintf("event #%d updated", id)))
}

// DeleteEventRequest обработчик запроса на удаление события.
func (c *CalendarHandler) DeleteEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SendResult(w, []byte("method not allowed"))
		return
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		SendError(w, errors.New("invalid data"), http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	deleted, err := c.calendar.DeleteEvent(id)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	SendResult(w, []byte(fmt.Sprintf("event #%d (%s, %v) removed", deleted.ID, deleted.Name, deleted.Time)))
}

// EventsForDayRequest обработчик запроса на получение событий на день.
func (c *CalendarHandler) EventsForDayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		SendResult(w, []byte("method not allowed"))
		return
	}

	data, err := SerializeEventSlice(c.calendar.EventsForDay())
	if err != nil {
		SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	SendResult(w, data)
}

// EventsForWeekRequest обработчик запроса на получение событий на неделю.
func (c *CalendarHandler) EventsForWeekRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		SendResult(w, []byte("method not allowed"))
		return
	}

	data, err := SerializeEventSlice(c.calendar.EventsForWeek())
	if err != nil {
		SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	SendResult(w, data)
}

// EventsForMonthRequest обработчик запроса на получение событий на месяц.
func (c *CalendarHandler) EventsForMonthRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		SendResult(w, []byte("method not allowed"))
		return
	}

	data, err := SerializeEventSlice(c.calendar.EventsForMonth())
	if err != nil {
		SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	SendResult(w, data)
}

// Logger структура для логгирования.
type Logger struct {
	handler http.Handler
}

// ServeHTTP обработчик для логгирования HTTP запросов.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger конструктор для логгирования HTTP запросов.
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

// ErrorResponse структура для ответа с ошибкой.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ResultResponse структура для ответа с результатом.
type ResultResponse struct {
	Result []byte `json:"result"`
}

// ParseCreateRequest парсит запрос на создание события.
func ParseCreateRequest(r *http.Request) (*Event, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		return nil, errors.New("invalid data")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02 15:04", r.FormValue("time"))
	if err != nil {
		return nil, err
	}

	name := r.FormValue("name")
	if name == "" {
		return nil, errors.New("name can't be blank")
	}

	newEvent := NewEvent(t, name)
	return newEvent, nil
}

// ParseUpdateRequest парсит запрос на обновление события.
func ParseUpdateRequest(r *http.Request) (int, time.Time, string, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		return -1, time.Time{}, "", errors.New("invalid data")
	}
	err := r.ParseForm()
	if err != nil {
		return -1, time.Time{}, "", err
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return -1, time.Time{}, "", err
	}

	timeStr := r.FormValue("time")
	parsedTime, err := time.Parse("2006-01-02 15:04", timeStr)
	if !(timeStr == "" || err == nil) {
		return -1, time.Time{}, "", err
	}

	name := r.FormValue("name")

	return id, parsedTime, name, nil
}

// SendError отправляет ответ с ошибкой.
func SendError(w http.ResponseWriter, err error, statusCode int) {
	data := ErrorResponse{Error: err.Error()}
	result, _ := json.Marshal(data)
	w.WriteHeader(statusCode)
	w.Write(result)
}

// SendResult отправляет ответ с результатом.
func SendResult(w http.ResponseWriter, response []byte) {
	data := ResultResponse{Result: response}
	result, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func main() {
	ch := NewCalendarHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", ch.CreateEventRequest)
	mux.HandleFunc("/update_event", ch.UpdateEventRequest)
	mux.HandleFunc("/delete_event", ch.DeleteEventRequest)
	mux.HandleFunc("/events_for_day", ch.EventsForDayRequest)
	mux.HandleFunc("/events_for_week", ch.EventsForWeekRequest)
	mux.HandleFunc("/events_for_month", ch.EventsForMonthRequest)

	wrappedMux := NewLogger(mux)

	http.ListenAndServe("localhost:8080", wrappedMux)
}
