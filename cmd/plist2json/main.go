package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"howett.net/plist"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	decoder := plist.NewDecoder(bytes.NewReader(data))

	var result map[string]interface{}
	decoder.Decode(&result)
	data, err = json.Marshal(result)
	fmt.Println(string(data))
}
