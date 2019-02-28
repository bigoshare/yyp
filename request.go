package yyp

import (
	"encoding/binary"
)

type Header struct {
	Length  uint32
	Uri     uint32
	ResCode uint16
}

type Request struct {
	Header
	Data []byte
}

// 从buff中读取request
func LoadRequest(buff []byte) *Request {
	if len(buff) < 10 { // 包头为10个字节，则必须大于10个字节才可能是一个包
		return nil
	}

	packLen := binary.LittleEndian.Uint32(buff[0:4]) // 没有集齐包
	if uint32(len(buff)) < packLen {
		return nil
	}

	uri := binary.LittleEndian.Uint32(buff[4:8])
	resCode := binary.LittleEndian.Uint16(buff[8:10])

	return &Request{
		Header: Header{
			Length:  packLen,
			Uri:     uri,
			ResCode: resCode,
		},
		Data: buff[10:packLen],
	}
}

func (r *Request) SetLength(len uint32) {
	r.Length = len
}

func (r *Request) SetUri(uri uint32) {
	r.Uri = uri
}

func (r *Request) SetResCode(resCode uint16) {
	r.ResCode = resCode
}
