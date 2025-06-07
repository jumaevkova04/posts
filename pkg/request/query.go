package request

import (
	"github.com/go-playground/form"
	"github.com/shopspring/decimal"
	"net/http"
)

var formDecoder *form.Decoder

func DecodeFormQuery(r *http.Request, dst interface{}) error {
	formDecoder = form.NewDecoder()
	formDecoder.SetTagName("json")
	formDecoder.SetMode(form.ModeImplicit)
	formDecoder.RegisterCustomTypeFunc(func(values []string) (interface{}, error) {
		v := values[1]
		return decimal.NewFromString(v)
	}, decimal.Decimal{})

	err := formDecoder.Decode(dst, r.URL.Query())
	if err != nil {
		return err
	}

	return nil
}
