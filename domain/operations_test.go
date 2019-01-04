package domain

import "testing"
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {

	service := OperationsService{}

	n := service.Add(1, 2, 3, 4)

	assert.Equal(t, 10, n)

}

func TestSubtract(t *testing.T) {

	service := OperationsService{}

	n := service.Subtract(1, 2, 3, 4)

	assert.Equal(t, -10, n)

}
