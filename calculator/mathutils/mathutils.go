package mathutils

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

// Calculator provides basic mathematical operations with tracking
type Calculator struct {
	ID        string
	History   []Operation
	CreatedAt time.Time
}

// Operation represents a mathematical operation
type Operation struct {
	Type      string
	Arguments []float64
	Result    float64
	Timestamp time.Time
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{
		ID:        uuid.New().String(),
		History:   make([]Operation, 0),
		CreatedAt: time.Now(),
	}
}

// Add performs addition and records the operation
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.recordOperation("add", []float64{a, b}, result)
	return result
}

// Subtract performs subtraction and records the operation
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.recordOperation("subtract", []float64{a, b}, result)
	return result
}

// Multiply performs multiplication and records the operation
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.recordOperation("multiply", []float64{a, b}, result)
	return result
}

// Divide performs division and records the operation
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.recordOperation("divide", []float64{a, b}, result)
	return result, nil
}

// Power calculates a^b and records the operation
func (c *Calculator) Power(base, exponent float64) float64 {
	result := math.Pow(base, exponent)
	c.recordOperation("power", []float64{base, exponent}, result)
	return result
}

// GetHistory returns the operation history
func (c *Calculator) GetHistory() []Operation {
	return c.History
}

// GetLastResult returns the result of the last operation
func (c *Calculator) GetLastResult() (float64, error) {
	if len(c.History) == 0 {
		return 0, fmt.Errorf("no operations performed")
	}
	return c.History[len(c.History)-1].Result, nil
}

// Clear clears the operation history
func (c *Calculator) Clear() {
	c.History = make([]Operation, 0)
}

// recordOperation adds an operation to the history
func (c *Calculator) recordOperation(opType string, args []float64, result float64) {
	op := Operation{
		Type:      opType,
		Arguments: args,
		Result:    result,
		Timestamp: time.Now(),
	}
	c.History = append(c.History, op)
}

// String returns a string representation of the calculator
func (c *Calculator) String() string {
	return fmt.Sprintf("Calculator{ID: %s, Operations: %d, Created: %s}",
		c.ID, len(c.History), c.CreatedAt.Format(time.RFC3339))
}
