package less

import (
	"os/exec"
	"bytes"
)

func Compile(path string, w *bytes.Buffer)error {
	dateCmd := exec.Command("lessc", path)
	dateOut, err := dateCmd.Output()
	if err != nil {
		return err
	}
	_, err = w.Write(dateOut)
	return err
}