package ignite

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

type Runner struct {
	command string
	cmd     *exec.Cmd
	mu      sync.Mutex
}

func Runners(command string) *Runner {
	return &Runner{command: command}
}

func (r *Runner) Start() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cmd != nil && r.cmd.Process != nil {
		r.Stop()
	}

	log.Println("Starting:", r.command)
	cmd := exec.Command("sh", "-c", r.command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	r.cmd = cmd
	go cmd.Run()
}

func (r *Runner) Stop() {
	if r.cmd != nil && r.cmd.Process != nil {
		log.Println("Stopping process...")
		r.cmd.Process.Kill()
	}
}
