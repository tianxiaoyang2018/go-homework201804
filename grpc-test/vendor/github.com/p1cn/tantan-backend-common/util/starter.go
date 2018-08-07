package util

type Starter interface {
	Start() error
}

func RunMultiStarter(starters ...Starter) error {
	errChan := make(chan error, len(starters))
	for _, s := range starters {
		go func(s Starter) {
			errChan <- s.Start()
		}(s)
	}

	return <-errChan
}
