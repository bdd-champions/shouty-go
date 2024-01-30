package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

var shouty *Shouty

const ArbitraryMessage = "Hello, world"

func lucyIsAt(x, y int) {
	shouty.SetLocation("Lucy", NewCoordinate(float64(x), float64(y)))
}

func seanIsAt(x, y int) {
	shouty.SetLocation("Sean", NewCoordinate(float64(x), float64(y)))
}

func seanShouts() {
	shouty.Shout("Sean", ArbitraryMessage)
}

func lucyShouldHearSean() error {
	return assertExpectedAndActual(assert.Equal, 1, len(shouty.GetShoutsHeardBy("Lucy")))
}

func lucyShouldHearNothing() error {
	return assertActual(
		assert.Empty,
		shouty.GetShoutsHeardBy("Lucy"),
	)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		shouty = NewShouty()
		return ctx, nil
	})

	sc.Step(`^Lucy is at (\d+), (\d+)$`, lucyIsAt)
	sc.Step(`^Sean is at (\d+), (\d+)$`, seanIsAt)
	sc.Step(`^Sean shouts$`, seanShouts)
	sc.Step(`^Lucy should hear Sean$`, lucyShouldHearSean)
	sc.Step(`^Lucy should hear nothing$`, lucyShouldHearNothing)
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predined state like nil, empty or true/false.
func assertActual(a actualAssertion, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, actual, msgAndArgs...)
	return t.err
}

type actualAssertion func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
