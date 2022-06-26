package models

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var BusinessErrors = []Error{
	{Code: "busErr1", Message: "Cannot create reservation in given interval."},
	{Code: "busErr2", Message: "The car's office is currently closed."},
	{Code: "busErr3", Message: "Desired location is not active."},
	{Code: "busErr4", Message: "There are no available cars for the given location and time."},
	{Code: "busErr5", Message: "There are no active locations."},
	{Code: "busErr6", Message: "There are no vacant cars in given time interval."},
	{Code: "busErr7", Message: "Car reservation list is empty, there are no reservations at this time."},
}

var TechnicalErrors = []Error{
	{Code: "tecErr1", Message: "Cannot parse request body."},
	{Code: "tecErr2", Message: "Database error."},
	{Code: "tecErr3", Message: "Cannot send JSON."},
}
