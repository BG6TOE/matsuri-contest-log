package adif

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type ADIFItem map[string]string
type ADIFFile struct {
	Header ADIFItem   `json:"header"`
	Record []ADIFItem `json:"records"`
}

const (
	nextItemState_Seeking = iota
	nextItemState_ReadingTag
	nextItemState_ReadingLength
	nextItemState_ReadingValue
)

func nextItem(r io.Reader) (string, uint, string, uint, error) {
	len := uint(0)
	key := strings.Builder{}
	chbuf := make([]byte, 1)
	byteread := uint(0)
	currentState := nextItemState_Seeking
	for {
		n, err := r.Read(chbuf)
		if n == 0 {
			return "", 0, "", byteread, errors.New("eof")
		}
		if err != nil {
			return "", 0, "", byteread, err
		}
		byteread += 1
		switch currentState {
		case nextItemState_Seeking:
			if chbuf[0] == '<' {
				currentState = nextItemState_ReadingTag
			}
		case nextItemState_ReadingTag:
			if chbuf[0] == ':' {
				currentState = nextItemState_ReadingLength
			} else if chbuf[0] == '>' {
				currentState = nextItemState_ReadingValue
			} else {
				if err := key.WriteByte(chbuf[0]); err != nil {
					return "", 0, "", byteread, err
				}
			}
		case nextItemState_ReadingLength:
			if chbuf[0] == '>' {
				currentState = nextItemState_ReadingValue
			} else if chbuf[0] >= '0' && chbuf[0] <= '9' {
				len = len*10 + uint(chbuf[0]-'0')
			} else {
				return "", 0, "", byteread, errors.New("illegal length value")
			}
		}
		if currentState == nextItemState_ReadingValue {
			break
		}
		chbuf = make([]byte, 1)
	}
	if len == 0 {
		return strings.ToLower(key.String()), 0, "", byteread, nil
	}
	value := make([]byte, len)
	{
		n, err := r.Read(value)
		byteread += uint(n)
		if uint(n) != len || err != nil {
			return key.String(), len, "", byteread, fmt.Errorf("failed to read value: %v", err)
		}
	}
	return strings.ToLower(key.String()), len, string(value), byteread, nil
}

func readRecord(r io.Reader) (map[string]string, string, uint, error) {
	ret := make(map[string]string)
	bytesread := uint(0)
	for {
		key, _, val, n, err := nextItem(r)
		bytesread += n
		if err != nil || n == 0 {
			return ret, "", bytesread, err
		}
		if key == "eor" || key == "eoh" {
			return ret, key, bytesread, nil
		}
		if _, ok := ret[key]; ok {
			return ret, "", bytesread, fmt.Errorf("duplicate key: %s", key)
		}
		ret[key] = val
	}
}

func ParseADIF(v []byte) (ADIFFile, error) {
	reader := bytes.NewReader(v)
	ret := ADIFFile{nil, make([]ADIFItem, 0)}
	for {
		record, endtag, _, err := readRecord(reader)
		if len(record) == 0 {
			if err != nil && err.Error() == "eof" {
				return ret, nil
			}
		}
		if err != nil {
			return ret, err
		}
		if endtag == "eoh" {
			if ret.Header != nil {
				return ret, errors.New("duplicate EOH tag")
			}
			ret.Header = record
		} else {
			ret.Record = append(ret.Record, record)
		}
	}
}

func (item *ADIFItem) write(w io.Writer) {
	for k, v := range *item {
		if len(v) != 0 {
			fmt.Fprintf(w, "<%s:%d>%s", k, len(v), v)
		} else {
			fmt.Fprintf(w, "<%s>", k)
		}
	}
}

func (f *ADIFFile) Write(w io.Writer) {
	fmt.Fprintln(w, "ADIF File Written by matsu-radio-log.adif")
	if len(f.Header) != 0 {
		f.Header.write(w)
		fmt.Fprintln(w, "<eoh>")
	}
	for _, vi := range f.Record {
		vi.write(w)
		fmt.Fprintln(w, "<eor>")
	}
}
