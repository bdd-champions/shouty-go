package main

import (
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

type TestContext struct {
	T *testing.T
}

var shouty *Shouty
var godogTags string

const ArbitraryMessage = "Hello, world"

func init() {
	flag.StringVar(&godogTags, "godog.tags", "", "tag filter for godog features")
}

func lucyIsAt(x, y int) {
	shouty.SetLocation("Lucy", NewCoordinate(float64(x), float64(y)))
}

func seanIsAt(x, y int) {
	shouty.SetLocation("Sean", NewCoordinate(float64(x), float64(y)))
}

func seanShouts() {
	shouty.Shout("Sean", ArbitraryMessage)
}

func (tc *TestContext) lucyShouldHearNothing() error {
	if !assert.Empty(tc.T, shouty.GetShoutsHeardBy("Lucy")) {
		return fmt.Errorf("Lucy should hear nothing")
	}
	return nil
}

func (tc *TestContext) lucyShouldHearSean() error {
	if !assert.Equal(tc.T, 1, len(shouty.GetShoutsHeardBy("Lucy"))) {
		return fmt.Errorf("Lucy should hear Sean")
	}
	return nil
}

func TestFeatures(t *testing.T) {

	var opts = godog.Options{
		Format:   "pretty", // your existing options
		Paths:    []string{"features"},
		TestingT: t,
	}

	if godogTags != "" {
		opts.Tags = godogTags
	}

	suite := godog.TestSuite{
		Name:                 "shouty",
		TestSuiteInitializer: InitializeSuite,
		ScenarioInitializer:  InitializeScenario(t),
		Options:              &opts,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(t *testing.T) func(*godog.ScenarioContext) {
	return func(ctx *godog.ScenarioContext) {

		tc := &TestContext{
			T: t,
		}

		ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
			shouty = NewShouty()
			return ctx, nil
		})

		ctx.Step(`^Lucy is at (\d+), (\d+)$`, lucyIsAt)
		ctx.Step(`^Sean is at (\d+), (\d+)$`, seanIsAt)
		ctx.Step(`^Sean shouts$`, seanShouts)
		ctx.Step(`^Lucy should hear nothing$`, tc.lucyShouldHearNothing)
		ctx.Step(`^Lucy should hear Sean$`, tc.lucyShouldHearSean)
	}
}

func InitializeSuite(ctx *godog.TestSuiteContext) {

}
