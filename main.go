package main

import (
	"time"
)

type bucket struct {
	points []int
	count  int
	next   *bucket
}

func (b *bucket) reset() {
	b.points = []int{}
	b.count = 0
}

type window struct {
	size  int
	slots []*bucket
}

func (w *window) resetBucket(index int) {
	w.slots[index]
}

type rollingCounter struct {
	bktDuration    time.Duration
	size           int
	offset         int
	lastAppendTime time.Time
	window         *window
}

func (rc *rollingCounter) timeSpan() int {
	return int(time.Since(rc.lastAppendTime) / rc.bktDuration)
}

func (rc *rollingCounter) add() {
	span := rc.timeSpan()

	offset := rc.offset
	s := rc.offset + 1
	end := s + span

	for i := s; i < end; i++ {

	}
}
