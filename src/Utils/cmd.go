package Utils

import (
	"os"
	"os/exec"
)

func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	PanicIfNotNil(err)

	_ = cmd.Wait()
}
