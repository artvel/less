# less
golang less compiler

# Status: tried otto and duktape in vain

Current Version calls ruby gems lessc to compile less.
better version coming as soon as possible.

### install lessc
```
sudo apt-get install rubygems
sudo apt-get install ruby-dev
sudo gem install rubygems-update
sudo gem update rubygems
sudo gem install less
sudo ln -s /var/lib/gems/2.3.0/gems/less-2.6.0/bin/lessc /usr/bin/
```

### go code
```go
import (
"github.com/artvel/less"
"fmt"
"bytes"
)
func main(){
    w := new(bytes.Buffer)
    err := less.Compile("my.less", w)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println("my css")
    fmt.Println(w.String())
}
sudo apt-get install rubygems
sudo apt-get install ruby-dev
sudo gem install rubygems-update
sudo gem update rubygems
sudo gem install less
sudo ln -s /var/lib/gems/2.3.0/gems/less-2.6.0/bin/lessc /usr/bin/
```
