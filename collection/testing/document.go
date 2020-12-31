package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TestDocument struct {
	ID     string  `json:"_id" bson:"_id"`
	REV    string  `json:"_rev,omitempty" bson:"_rev,omitempty"`
	Title  string  `json:"title"`
	Score  float32 `json:"score"`
	Year   int     `json:"year"`
	Oscars bool    `json:"oscars"`
}

func GetSingleRecord() (TestDocument, []TestDocument) {
	b, err := ioutil.ReadFile("../../data/testdata.json")
	if err != nil {
		fmt.Println(err)
	}

	tfile := struct {
		Single     TestDocument   `json:"single"`
		Collection []TestDocument `json:"collection"`
	}{}

	err = json.Unmarshal(b, &tfile)
	if err != nil {
		fmt.Println(err)
	}

	return tfile.Single, tfile.Collection
}
