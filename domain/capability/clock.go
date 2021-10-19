package capability

import "time"

type Clock interface {
	Now() time.Time
}
