/*
 * Copyright 2018 Marco Helmich
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kafkawireformat

import (
    "bytes"
    "encoding/binary"
)

type Encodable interface {
    Encode(enc *Encoder) error
}

type Encoder struct {
    buffer *bytes.Buffer
}

func NewEncoder() *Encoder {
    buf := new(bytes.Buffer)
    return &Encoder{
        buffer: buf,
    }
}

func (enc *Encoder) WriteVarInt(i int64) {
    buf := make([]byte, binary.MaxVarintLen64)
    n := binary.PutVarint(buf, i)
    binary.Write(enc.buffer, binary.BigEndian, buf[:n])
}

func (enc *Encoder) WriteInt8(i int8) {
    binary.Write(enc.buffer, binary.BigEndian, i)
}

func (enc *Encoder) WriteInt16(i int16) {
    binary.Write(enc.buffer, binary.BigEndian, i)
}

func (enc *Encoder) WriteInt32(i int32) {
    binary.Write(enc.buffer, binary.BigEndian, i)
}

func (enc *Encoder) WriteInt64(i int64) {
    binary.Write(enc.buffer, binary.BigEndian, i)
}

func (enc *Encoder) WriteString(s string) {
    if s == "" {
        enc.WriteInt16(-1)
    } else {
        bites := []byte(s)
        enc.WriteInt16(int16(len(bites)))
        binary.Write(enc.buffer, binary.BigEndian, bites)
    }
}

// this was incredible difficult to figure out
// in kafka land a boolean is a one byte integer !!!
// a value of 0 means false
// everything else is true
func (enc *Encoder) WriteBool(b bool) {
    if b {
        enc.WriteInt8(1)
    } else {
        enc.WriteInt8(0)
    }
}

func (enc *Encoder) WriteInt32Array(a []int32) {
    arrayLength := len(a)
    enc.WriteInt32(int32(arrayLength))
    for i := 0; i < arrayLength; i++ {
        enc.WriteInt32(a[i])
    }
}

func (enc *Encoder) WriteInt64Array(a []int64) {
    arrayLength := len(a)
    enc.WriteInt32(int32(arrayLength))
    for i := 0; i < arrayLength; i++ {
        enc.WriteInt64(a[i])
    }
}

func (enc *Encoder) WriteStringArray(strings []string) {
    arrayLength := len(strings)
    enc.WriteInt32(int32(arrayLength))
    for i := 0; i < arrayLength; i++ {
        enc.WriteString(strings[i])
    }
}

func (enc *Encoder) WriteByteArray(bites []byte) {
    arrayLength := len(bites)
    enc.WriteInt32(int32(arrayLength))
    binary.Write(enc.buffer, binary.BigEndian, bites)
}

func (enc *Encoder) WriteBoolArray(bs []bool) {
    arrayLength := len(bs)
    enc.WriteInt32(int32(arrayLength))
    for i := 0; i < arrayLength; i++ {
        enc.WriteBool(bs[i])
    }
}

func (enc *Encoder) Len() int {
    return enc.buffer.Len()
}

func (enc *Encoder) Bytes() []byte {
    return enc.buffer.Bytes()
}