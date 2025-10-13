package ignite

func Start(paths []string, command string) error {
	runner := Runners(command)
	runner.Start()

	watcher, err := Watcher(paths, func(_ string) {
		runner.Start()
	})
	if err != nil {
		return err
	}

	watcher.Start()
	return nil
}
