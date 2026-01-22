package main

import "testing"
import "github.com/stretchr/testify/assert"

// Simplest
func TestSimplest(t *testing.T) {
	result, err := sort(1, 1, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, Standard, result)
}

// Special Cases
func TestBulkyButNotHeavy(t *testing.T) {
	result, err := sort(151, 1, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, Special, result)
}

func TestBulkyBecauseOfVolume(t *testing.T) {
	result, err := sort(1000, 1000, 2, 0)
	assert.NoError(t, err)
	assert.Equal(t, Special, result)
}

func TestHeavyButNotBulky(t *testing.T) {
	result, err := sort(1, 1, 1, 21)
	assert.NoError(t, err)
	assert.Equal(t, Special, result)
}

func TestBothBulkyAndHeavy(t *testing.T) {
	result, err := sort(151, 1, 1, 21)
	assert.NoError(t, err)
	assert.Equal(t, Rejected, result)
}

// Argument Boundaries
func TestSpecialTooLargeWidthDimension(t *testing.T) {
	result, err := sort(151, 1, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, Special, result)
}

func TestSpecialTooLargeHeightDimension(t *testing.T) {
	result, err := sort(1, 151, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, Special, result)
}

func TestSpecialTooLargeLengthDimension(t *testing.T) {
	result, err := sort(1, 1, 151, 0)
	assert.NoError(t, err)
	assert.Equal(t, Special, result)
}

func TestInvalidZeroWidthDimension(t *testing.T) {
	_, err := sort(0, 1, 1, 0)
	assert.Error(t, err)
}
func TestInvalidNegativeWidthDimension(t *testing.T) {
	_, err := sort(-1, 1, 1, 0)
	assert.Error(t, err)
}

func TestInvalidZeroHeightDimension(t *testing.T) {
	_, err := sort(1, 0, 1, 0)
	assert.Error(t, err)
}

func TestInvalidNegativeHeightDimension(t *testing.T) {
	_, err := sort(1, -1, 1, 0)
	assert.Error(t, err)
}

func TestInvalidZeroLengthDimension(t *testing.T) {
	_, err := sort(1, 1, 0, 0)
	assert.Error(t, err)
}

func TestInvalidNegativeLengthDimension(t *testing.T) {
	_, err := sort(1, 1, -1, 0)
	assert.Error(t, err)
}

func TestInvalidNegativeMass(t *testing.T) {
	_, err := sort(1, 1, 1, -1)
	assert.Error(t, err)
}

// Some edge cases
func TestEdgeCaseMaxDimensionsNotBulky(t *testing.T) {
	result, err := sort(150, 150, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, Standard, result)
}

func TestEdgeCaseMaxMassNotHeavy(t *testing.T) {
	result, err := sort(1, 1, 1, 20)
	assert.NoError(t, err)
	assert.Equal(t, Standard, result)
}
