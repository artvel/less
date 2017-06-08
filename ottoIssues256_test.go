package less

import (
	v8 "github.com/lazytiger/go-v8"
	"net/http"
	"io"
	"bytes"
	"testing"
)

func TestDownload(t *testing.T){
	w := new(bytes.Buffer)
	err := downloadFile(w ,"https://raw.githubusercontent.com/less/less.js/3.x/dist/less.min.js")
	if err != nil {
		t.Error(err)
	}
	//vm := otto.New()
	//v, err := vm.Run(w.String())
	//fmt.Println(v)
	//fmt.Println(err)
	//vm.Run(w.String())
	engine := v8.NewEngine()
	script := engine.Compile([]byte("'Hello ' + 'World!'"), nil, nil)
	context := engine.NewContext(nil)

	context.Scope(func(cs v8.ContextScope) {
		result := script.Run()
		println(result.ToString())
	})
//	v, err = vm.Run(`
//	    abc = 2 + 2;
//	    console.log("The value of abc is " + abc); // 4
//
//	    //var less = require('less');
//console.log(less);
//		less.render('.class { width: (1 + 1) }',
//		    {
//		      paths: ['.', './lib'],  // Specify search paths for @import directives
//		      filename: 'signature.less', // Specify a filename, for better error messages
//		    },
//		    function (e, output) {
//		       console.log(output.css);
//		    });
//
//	`)
	//fmt.Println(v)
	//fmt.Println(err)

}

func downloadFile(w *bytes.Buffer, url string) (err error) {
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