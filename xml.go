package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type XmlEncoder struct {
	encoder *xml.Encoder
}

func NewXmlEncoder() *XmlEncoder {
	e := xml.NewEncoder(os.Stdout)
	e.Indent("  ", "    ")
	return &XmlEncoder{e}
}

func (e *XmlEncoder) Encode(values map[string]interface{}) error {
	row := xml.StartElement{Name: xml.Name{"", "row"}}
	tokens := []xml.Token{row}
	for key, value := range values {
		var charData xml.CharData

		t := xml.StartElement{Name: xml.Name{"", key}}

		switch value := (value).(type) {
		case []byte:
			charData = xml.CharData(string(value))
		case int64:
			charData = xml.CharData(fmt.Sprintf("%d", value))
		}
		tokens = append(tokens, t, charData, t.End())
	}
	tokens = append(tokens, row.End())

	for _, t := range tokens {
		err := e.encoder.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := e.encoder.Flush()
	if err != nil {
		return err
	}

	return nil
}
