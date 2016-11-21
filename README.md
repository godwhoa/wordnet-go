WordNet library for Go
======================
It uses a sqlite database.


## Install
`go get -v github.com/godwhoa/wordnet-go`<br>
Requires sqlite pre-installed

## Usage
```go
package main
import(
	"fmt"
	"github.com/godwhoa/wordnet-go"
)
func main() {
	wn := &WordNet{}
	wn.Init("") // defaults to $GOPATH/src/github.com/godwhoa/wordnet-go/wordnet.db if left empty
	result, err := wn.ByType("eat", "noun", 1)
	fmt.Println(result, err)
}
```
## Credits
[TMiguelT](https://github.com/TMiguelT/) for [wordnet-sqlite](https://github.com/TMiguelT/wordnet-sqlite)<br>
Princeton University for [WordNet](https://wordnet.princeton.edu/wordnet/)