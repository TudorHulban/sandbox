package resources

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

/*
There is no update activity, just use delete and add activity for now.
*/

// Activity Consolidates activities for resource.
type Activity struct {
	TimeFrame

	ActionID int64
}

// Resource Consolidates resource availability.
type Resource struct {
	Activities   []Activity
	BusyInterval []TimeFrame

	overlapIndex int
}

// NewResource Constructor for resources.
func NewResource() *Resource {
	return &Resource{
		Activities:   []Activity{},
		BusyInterval: []TimeFrame{},
	}
}

func (r *Resource) AddActivity(activities ...Activity) {
	r.Activities = append(r.Activities, activities...)
}

func (r *Resource) DeleteActivity(a Activity) error {
	return nil
}

func (r *Resource) sortActivities() {
	sort.Slice(r.Activities,
		func(i, j int) bool {
			return r.Activities[i].UnixStartTime < r.Activities[j].UnixStartTime
		},
	)
}

func (r *Resource) syncBusy() {
	r.BusyInterval = make([]TimeFrame, len(r.Activities)) // reset slice

	for ix, activity := range r.Activities {
		r.BusyInterval[ix] = activity.TimeFrame
	}

	log.Println("filled up busy intervals:", r.BusyInterval)
}

// hasOverlapping given ordered by start time slice gives indication if it has overlapping time frames.
func (r *Resource) hasOverlapping() bool {
	if len(r.Activities) < 2 {
		return true
	}

	for i := 0; i < len(r.Activities)-1; i++ {
		if r.Activities[i].TimeFrame.isOverLapping(r.Activities[i+1].TimeFrame) {
			r.overlapIndex = i

			return true
		}
	}

	return false
}

// updateBusyTime Method updates busy time once a change in activities occurs.
func (r *Resource) updateBusyTime() error {
	r.sortActivities()
	r.syncBusy()

	if len(r.BusyInterval) == 0 {
		return errors.New("no busy intervals")
	}

	if r.overlapIndex > len(r.BusyInterval) {
		return fmt.Errorf("overlapIndex is %d and number of busy intervals %d",
			r.overlapIndex, len(r.BusyInterval),
		)
	}

	for r.hasOverlapping() {
		log.Println("overlapping index:", r.overlapIndex)

		newTimeFrame := mergeOverlapping(
			r.BusyInterval[r.overlapIndex],
			r.BusyInterval[r.overlapIndex+1],
		)

		r.BusyInterval = r.removeTimeFrame(r.overlapIndex + 1)
		r.BusyInterval[r.overlapIndex] = newTimeFrame
	}

	log.Println("updated busy intervals:", r.BusyInterval)

	return nil
}

func (r *Resource) removeTimeFrame(index int) []TimeFrame {
	return append(r.BusyInterval[:index], r.BusyInterval[index+1:]...)
}
