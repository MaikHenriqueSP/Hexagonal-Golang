package config

import "github.com/benbjohnson/clock"

func InitializeClock() (c clock.Clock) {
	return clock.New()
}