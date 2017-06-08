package less

import (
	duktape "gopkg.in/olebedev/go-duktape.v2"
	//"github.com/robertkrimen/otto"
	"net/http"
	"io"
	"bytes"
	"testing"
	"fmt"
)

func TestDownloadAndRun(t *testing.T){
	w := new(bytes.Buffer)
	err := downloadFile2(w ,"https://raw.githubusercontent.com/less/less.js/v2.7.2/dist/less.js")
	if err != nil {
		t.Error(err)
	}
	ctx := duktape.New()

	// Let's inject `setTimeout`, `setInterval`, `clearTimeout`,
	// `clearInterval` into global scope.
	ctx.PushTimers()

	ch := make(chan string)
	ctx.PushGlobalGoFunction("second", func(_ *duktape.Context) int {
		ch <- "second step"
		return 0
	})
	ctx.EvalStringNoresult(w.String())
	//ctx.EvalString(w.String())
	ctx.PevalString(`
    setTimeout(second, 0);
    print('first step');
    print(less);
  `)
	fmt.Println(<-ch)

}

func downloadFile2(w *bytes.Buffer, url string) (err error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Writer the body to file
	_, err = io.Copy(w, resp.Body)
	if err != nil  {
		return err
	}
	return nil
}