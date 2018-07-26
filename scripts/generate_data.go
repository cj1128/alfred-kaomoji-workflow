// generate emoji data for golang using `emojiimages` library
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Kaomoji struct {
	Kaomoji  string   `json:"entry"`
	Keywords []string `json:"keywords"`
}

const inputPath = "lib.json"
const outputPath = "kaomojis.go"

func main() {
	kaomojis, err := parseKaomojis()

	if err != nil {
		log.Fatal(err)
	}

	if err := generate(kaomojis); err != nil {
		log.Fatal(err)
	}

	fmt.Println(`All done. \\\(۶•̀ᴗ•́)۶//／／`)
}

func parseKaomojis() (map[string]*Kaomoji, error) {
	var result map[string]*Kaomoji
	buf, err := ioutil.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal json")
	}

	return result, nil
}

func generate(kaomojis map[string]*Kaomoji) error {
	out := bytes.Buffer{}

	out.WriteString(`package main

type Kaomoji struct {
  name string
  kaomoji string
  keywords string
}

var kaomojis = []*Kaomoji{
`)

	for name, kaomoji := range kaomojis {
		keywords := append([]string{name}, kaomoji.Keywords...)

		line := fmt.Sprintf(
			"  {%s, %s, %s},\n",
			strconv.Quote(name),
			strconv.Quote(kaomoji.Kaomoji),
			strconv.Quote(strings.Join(keywords, ",")),
		)
		out.WriteString(line)
	}
	out.WriteString("}")
	if err := writeData(out.Bytes()); err != nil {
		return err
	}
	return nil
}

func writeData(buf []byte) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "could not create file")
	}
	defer f.Close()
	_, err = f.Write(buf)
	if err != nil {
		return errors.Wrap(err, "could not write to file")
	}
	return nil
}
