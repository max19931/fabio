package route

import (
    "log"
    "sync/atomic"
)

// HTML Wrapper struct so we can store the html string in an atomic.Value
type HTML struct {
    value string
}

// html stores the no route html string
var store atomic.Value

func init() {
    store.Store(HTML{""})
}

// GetHTML returns the HTML for not found routes. The function is safe to be
// called from multiple goroutines.
func GetHTML() string {
    return store.Load().(HTML).value
}

// SetHTML sets the current noroute html. The function is safe to be called from
// multiple goroutines.
func SetHTML(h string) {
    html := HTML{h}
    store.Store(html)

    if h == "" {
        log.Print("[INFO] Unset noroute HTML")
    } else {
        log.Printf("[INFO] Set noroute HTML (%d bytes)", len(h))
    }
}