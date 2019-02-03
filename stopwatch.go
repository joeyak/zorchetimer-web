package main

import "time"

type stopwatch struct {
	elapsed     time.Duration
	lastElapsed time.Duration
	start       time.Time
	lastStart   time.Time
	isRunning   bool
}

func (s *stopwatch) Start() {
	if !s.isRunning {
		s.start = time.Now()
		s.lastStart = s.start
		s.isRunning = true
	}
}

func (s *stopwatch) Stop() {
	if s.isRunning {
		s.lastElapsed = time.Since(s.lastStart)
		s.elapsed += time.Since(s.start)
		s.isRunning = false
	}
}

func (s *stopwatch) Flip() {
	if s.isRunning {
		s.Stop()
	} else {
		s.Start()
	}
}

func (s *stopwatch) Elapsed() time.Duration {
	if s.isRunning {
		s.elapsed += time.Since(s.start)
		s.start = time.Now()
	}
	return s.elapsed
}

func newStopwatch(start bool) stopwatch {
	var sw stopwatch
	if start {
		sw.Start()
	}
	return sw
}
