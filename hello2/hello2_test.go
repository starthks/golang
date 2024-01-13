package hello2

import "testing"

type Race struct {
	Value string
}

func TestRaceCondition(t *testing.T) {
	c := make(chan bool)
	race := &Race{
		Value: "init",
	}
	go func() {
		race.Value = "first" // First conflicting access.
		c <- true
	}()
	race.Value = "second" // Second conflicting access.
	<-c

	t.Log("race value: ", race.Value)
}
