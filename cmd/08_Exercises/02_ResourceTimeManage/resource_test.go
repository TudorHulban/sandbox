package resources

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBasics(t *testing.T) {
	activ1 := Activity{
		ActionID: 1,
		TimeFrame: TimeFrame{
			UnixStartTime: 1000,
			UnixEndTime:   2000,
		},
	}

	activ2 := Activity{
		ActionID: 2,
		TimeFrame: TimeFrame{
			UnixStartTime: 1500,
			UnixEndTime:   2500,
		},
	}

	activ3 := Activity{
		ActionID: 3,
		TimeFrame: TimeFrame{
			UnixStartTime: 2400,
			UnixEndTime:   5500,
		},
	}

	r := NewResource()
	r.AddActivity(
		activ1,
		activ2,
		activ3,
	)
	assert.True(t, r.hasOverlapping())

	// maybe r.updateBusyTime() ?

	log.Println(r.overlapIndex)
	log.Println("activities:", r.Activities)
	log.Println("busy:", r.BusyInterval)
}

func TestRemoveTimeFrame(t *testing.T) {
	activ1 := Activity{
		ActionID: 1,
		TimeFrame: TimeFrame{
			UnixStartTime: 1000,
			UnixEndTime:   2000,
		},
	}

	activ2 := Activity{
		ActionID: 2,
		TimeFrame: TimeFrame{
			UnixStartTime: 1500,
			UnixEndTime:   2500,
		},
	}

	r := NewResource()
	r.AddActivity(activ1, activ2)

	r.updateBusyTime()

	assert.Equal(t,
		1,
		len(r.BusyInterval),
	)
}
