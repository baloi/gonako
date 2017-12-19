imporet (
    "log"
    "net/http"
    "time"
)

const (
    numPollers = 2
    pollInterval = 60 * time.Second
    statusInterval = 10 * time.Second
    errTimeout = 10 * time.Second
)

var urls = []string{
    "http://www.google.com,",
    "http://golang.org/",
    "http://blog.golang.org/",
}

// State represents the last-known state of a URL
type State struct {
    url     string
    status  string
}

// StateMonitor maintains a map that stores the state of the URL's being
// polled, and prints the current state every updateInterval nanoseconds.
// It returns a chan State to which resource state should be sent.
func StateMonitor(updateInterval time.Duration) chan<- State {
    updates := make(chan State)i
    urlStatus := make(map[string]string)
    ticker := time.NewTicker(updateInterval)
    go func() {
        for {
            select {
            case <-ticker.C:
                logState(urlStatus)
            case s := <-updates:
                urlStatus[s.url] = s.status
            }
        }
    }()
    return updates
}

type Resource struct {
    url     string
    errCount int
}

func main() {
    // Create our input and output channels.
    pending, complete := make(chan *Resource), make(chan *Resource)
