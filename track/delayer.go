package track

import (
	"time"
)

type Delayer interface {
	Delay() time.Duration
}