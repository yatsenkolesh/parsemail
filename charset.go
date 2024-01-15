package parsemail

import (
	"errors"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"io"
	"strings"
)

type (
	Charset string
)

var (
	ErrUnsupported = errors.New("unsupported charset provided")
)

type (
	CharsetDecoder interface {
		Bytes(b []byte) ([]byte, error)
		Reader(r io.Reader) io.Reader
	}
	DefaultDecoder struct{}
)

const (
	Iso88591    Charset = "iso-8859-1"
	Iso88592    Charset = "iso-8859-2"
	Iso88593    Charset = "iso-8859-3"
	Iso88594    Charset = "iso-8859-4"
	Iso88595    Charset = "iso-8859-5"
	Iso88596    Charset = "iso-8859-6"
	Iso88597    Charset = "iso-8859-7"
	Iso88598    Charset = "iso-8859-8"
	Iso88599    Charset = "iso-8859-9"
	Iso885910   Charset = "iso-8859-10"
	Iso885913   Charset = "iso-8859-13"
	Iso885914   Charset = "iso-8859-14"
	Iso885915   Charset = "iso-8859-15"
	Iso885916   Charset = "iso-8859-16"
	Utf8        Charset = "utf-8"
	UsAscii     Charset = "us-ascii"
	Windows1252 Charset = "windows-1252"
)

var (
	decoders = map[Charset]func() *encoding.Decoder{
		Iso88591:  charmap.ISO8859_1.NewDecoder,
		Iso88592:  charmap.ISO8859_2.NewDecoder,
		Iso88593:  charmap.ISO8859_3.NewDecoder,
		Iso88594:  charmap.ISO8859_4.NewDecoder,
		Iso88595:  charmap.ISO8859_5.NewDecoder,
		Iso88596:  charmap.ISO8859_6.NewDecoder,
		Iso88597:  charmap.ISO8859_7.NewDecoder,
		Iso88598:  charmap.ISO8859_8.NewDecoder,
		Iso88599:  charmap.ISO8859_9.NewDecoder,
		Iso885910: charmap.ISO8859_10.NewDecoder,
		Iso885913: charmap.ISO8859_13.NewDecoder,
		Iso885914: charmap.ISO8859_14.NewDecoder,
		Iso885915: charmap.ISO8859_15.NewDecoder,
		Iso885916: charmap.ISO8859_16.NewDecoder,
	}
)

func (c Charset) String() string {
	return string(c)
}

func (d DefaultDecoder) Bytes(b []byte) ([]byte, error) {
	return b, nil
}
func (d DefaultDecoder) Reader(r io.Reader) io.Reader {
	return r
}

func charsetFromParams(params map[string]string) Charset {
	var (
		charset string
		ok      bool
	)
	if charset, ok = params["charset"]; !ok {
		return Utf8
	}
	return Charset(strings.ToLower(charset))
}

func charsetDecoder(c Charset) (CharsetDecoder, error) {
	if dec, ok := decoders[c]; ok {
		return dec(), nil
	}
	return DefaultDecoder{}, nil
}

func decodeCharsetFromParams(params map[string]string) (CharsetDecoder, error) {
	ch := charsetFromParams(params)
	return charsetDecoder(ch)
}
