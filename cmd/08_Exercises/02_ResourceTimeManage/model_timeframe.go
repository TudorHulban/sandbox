package resources

// TimeFrame Defines a timeframe using Unix Epoch time start and end seconds.
type TimeFrame struct {
	UnixStartTime int64
	UnixEndTime   int64
}

func (t TimeFrame) isOverLapping(frameTime TimeFrame) bool {
	return t.UnixEndTime >= frameTime.UnixStartTime
}

func (t TimeFrame) mergeOverLapping(frameTime TimeFrame) *TimeFrame {
	if !t.isOverLapping(frameTime) {
		return nil
	}

	max := func(x, y int64) int64 {
		if x > y {
			return x
		}

		return y
	}

	return &TimeFrame{
		UnixStartTime: t.UnixStartTime,
		UnixEndTime: max(
			t.UnixEndTime,
			frameTime.UnixEndTime,
		),
	}
}
