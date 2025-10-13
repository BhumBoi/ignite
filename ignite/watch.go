package ignite

import (
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	watcher  *fsnotify.Watcher
	paths    []string
	onChange func(string)
}

func watch(paths []string, onChange func(string)) (*Watcher, error) {

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for _, p := range paths {
		matches, _ := filepath.Glob(p)
		for _, m := range matches {
			if err := w.Add(m); err != nil {
				log.Println("watch error", err)
			}
		}
	}

	return &Watcher{watcher: w, paths: paths, onChange: onChange}, nil
}

func (w *Watcher) Start() {
	log.Println("Ignite watching for file....")

	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
				log.Println("Change detected = ", event.Name)
				if w.onChange != nil {
					go func(name string) {
						time.Sleep(300 * time.Millisecond)
						w.onChange(name)
					}(event.Name)
				}
			}
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			log.Println("watch error:", err)
		}
	}
}
