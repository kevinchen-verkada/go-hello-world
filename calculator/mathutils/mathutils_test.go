package mathutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang.kevinchen-verkada/calculator/mathutils"
)

func TestNewCalculator(t *testing.T) {
	calc := mathutils.NewCalculator()

	assert.NotEmpty(t, calc.ID)
	assert.Empty(t, calc.History)
	assert.False(t, calc.CreatedAt.IsZero())
}

func TestBasicOperations(t *testing.T) {
	calc := mathutils.NewCalculator()

	// Test addition
	result := calc.Add(5, 3)
	assert.Equal(t, 8.0, result)

	// Test subtraction
	result = calc.Subtract(10, 4)
	assert.Equal(t, 6.0, result)

	// Test multiplication
	result = calc.Multiply(6, 7)
	assert.Equal(t, 42.0, result)

	// Test division
	result, err := calc.Divide(15, 3)
	require.NoError(t, err)
	assert.Equal(t, 5.0, result)

	// Test power
	result = calc.Power(2, 3)
	assert.Equal(t, 8.0, result)
}

func TestDivisionByZero(t *testing.T) {
	calc := mathutils.NewCalculator()

	_, err := calc.Divide(10, 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "division by zero")
}

func TestHistory(t *testing.T) {
	calc := mathutils.NewCalculator()

	// Perform some operations
	calc.Add(1, 2)
	calc.Multiply(3, 4)

	history := calc.GetHistory()
	assert.Len(t, history, 2)

	// Check first operation
	assert.Equal(t, "add", history[0].Type)
	assert.Equal(t, []float64{1, 2}, history[0].Arguments)
	assert.Equal(t, 3.0, history[0].Result)

	// Check second operation
	assert.Equal(t, "multiply", history[1].Type)
	assert.Equal(t, []float64{3, 4}, history[1].Arguments)
	assert.Equal(t, 12.0, history[1].Result)
}

func TestGetLastResult(t *testing.T) {
	calc := mathutils.NewCalculator()

	// Test with no operations
	_, err := calc.GetLastResult()
	assert.Error(t, err)

	// Test with operations
	calc.Add(5, 5)
	result, err := calc.GetLastResult()
	require.NoError(t, err)
	assert.Equal(t, 10.0, result)
}

func TestClear(t *testing.T) {
	calc := mathutils.NewCalculator()

	// Add some operations
	calc.Add(1, 1)
	calc.Subtract(5, 2)

	assert.Len(t, calc.GetHistory(), 2)

	// Clear history
	calc.Clear()
	assert.Empty(t, calc.GetHistory())
}
