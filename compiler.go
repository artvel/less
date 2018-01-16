package less

import (
	"os/exec"
	"bytes"
	"fmt"
)

func Compile(path string, w *bytes.Buffer, paths...string)error {
	var cmd *exec.Cmd
	l := len(paths)
	if paths != nil && l > 0 {
		pathBuffer := new(bytes.Buffer)
		for i, p := range paths {
			pathBuffer.WriteString(p)
			if i >= 0 && l-1 != i {
				pathBuffer.WriteString(":")
			}
		}
		cmd = exec.Command("lessc", path, "--paths", pathBuffer.String())
	} else {
		cmd = exec.Command("lessc", path)
	}
	cmdOut, err := cmd.Output()
	if err != nil {
		return err
	}
	_, err = w.Write(cmdOut)
	return err
}