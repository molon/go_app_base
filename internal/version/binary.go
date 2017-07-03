package version

import (
	"fmt"
	"runtime"
)

const VERSION = "0.0.1"

func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, VERSION, runtime.Version())
}
