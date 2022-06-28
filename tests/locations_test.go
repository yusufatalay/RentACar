package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yusufatalay/RentACar/models"
)

func TestThereIsNoActiveLocations(t *testing.T) {
	// clean out the database
	WipeDatabase()
	// at the start database is empty so there is no active locations
	result, err := models.GetActiveLocations()

	// err should be nil
	assert.Nil(t, err, "there should be no error but got %v", err)

	// length of the result should be 0
	assert.Equal(t, 0, len(result), "there should be no active locations")
}

func TestWhereThereAreActiveLocations(t *testing.T) {
	// clean out the database
	WipeDatabase()
	// seed the database with some active - non-active locations
	err := models.CreateLocation(&models.Location{Name: "Location 1", Active: "true"})
	assert.Nil(t, err, "there should be no error but got %v", err)

	err = models.CreateLocation(&models.Location{Name: "Location 2", Active: "false"})
	assert.Nil(t, err, "there should be no error but got %v", err)

	result, err := models.GetActiveLocations()
	assert.Nil(t, err, "there should be no error but got %v", err)

	// result should have one element
	assert.Equal(t, 1, len(result), "there should be one active location")

	// active locations should be Location 1
	assert.Equal(t, "Location 1", result[0].Name, "active location should be Location 1")

	// empty out the database
	WipeDatabase()
}
