package p2p

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AuthMsgV4) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Signature":
			err = dc.ReadExactBytes((z.Signature)[:])
			if err != nil {
				return
			}
		case "InitiatorPubkey":
			err = dc.ReadExactBytes((z.InitiatorPubkey)[:])
			if err != nil {
				return
			}
		case "Nonce":
			err = dc.ReadExactBytes((z.Nonce)[:])
			if err != nil {
				return
			}
		case "Version":
			z.Version, err = dc.ReadUint()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AuthMsgV4) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Signature"
	err = en.Append(0x84, 0xa9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Signature)[:])
	if err != nil {
		return
	}
	// write "InitiatorPubkey"
	err = en.Append(0xaf, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.InitiatorPubkey)[:])
	if err != nil {
		return
	}
	// write "Nonce"
	err = en.Append(0xa5, 0x4e, 0x6f, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Nonce)[:])
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AuthMsgV4) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Signature"
	o = append(o, 0x84, 0xa9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o = msgp.AppendBytes(o, (z.Signature)[:])
	// string "InitiatorPubkey"
	o = append(o, 0xaf, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79)
	o = msgp.AppendBytes(o, (z.InitiatorPubkey)[:])
	// string "Nonce"
	o = append(o, 0xa5, 0x4e, 0x6f, 0x6e, 0x63, 0x65)
	o = msgp.AppendBytes(o, (z.Nonce)[:])
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AuthMsgV4) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Signature":
			bts, err = msgp.ReadExactBytes(bts, (z.Signature)[:])
			if err != nil {
				return
			}
		case "InitiatorPubkey":
			bts, err = msgp.ReadExactBytes(bts, (z.InitiatorPubkey)[:])
			if err != nil {
				return
			}
		case "Nonce":
			bts, err = msgp.ReadExactBytes(bts, (z.Nonce)[:])
			if err != nil {
				return
			}
		case "Version":
			z.Version, bts, err = msgp.ReadUintBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AuthMsgV4) Msgsize() (s int) {
	s = 1 + 10 + msgp.ArrayHeaderSize + (sigLen * (msgp.ByteSize)) + 16 + msgp.ArrayHeaderSize + (pubLen * (msgp.ByteSize)) + 6 + msgp.ArrayHeaderSize + (shaLen * (msgp.ByteSize)) + 8 + msgp.UintSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *AuthRespV4) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RandomPubkey":
			err = dc.ReadExactBytes((z.RandomPubkey)[:])
			if err != nil {
				return
			}
		case "Nonce":
			err = dc.ReadExactBytes((z.Nonce)[:])
			if err != nil {
				return
			}
		case "Version":
			z.Version, err = dc.ReadUint()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AuthRespV4) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "RandomPubkey"
	err = en.Append(0x83, 0xac, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.RandomPubkey)[:])
	if err != nil {
		return
	}
	// write "Nonce"
	err = en.Append(0xa5, 0x4e, 0x6f, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Nonce)[:])
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AuthRespV4) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "RandomPubkey"
	o = append(o, 0x83, 0xac, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79)
	o = msgp.AppendBytes(o, (z.RandomPubkey)[:])
	// string "Nonce"
	o = append(o, 0xa5, 0x4e, 0x6f, 0x6e, 0x63, 0x65)
	o = msgp.AppendBytes(o, (z.Nonce)[:])
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AuthRespV4) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "RandomPubkey":
			bts, err = msgp.ReadExactBytes(bts, (z.RandomPubkey)[:])
			if err != nil {
				return
			}
		case "Nonce":
			bts, err = msgp.ReadExactBytes(bts, (z.Nonce)[:])
			if err != nil {
				return
			}
		case "Version":
			z.Version, bts, err = msgp.ReadUintBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AuthRespV4) Msgsize() (s int) {
	s = 1 + 13 + msgp.ArrayHeaderSize + (pubLen * (msgp.ByteSize)) + 6 + msgp.ArrayHeaderSize + (shaLen * (msgp.ByteSize)) + 8 + msgp.UintSize
	return
}
