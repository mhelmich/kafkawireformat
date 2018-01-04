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
    "bufio"
    "encoding/binary"
    "fmt"
)

type Decodable interface {
    Decode(dec *Decoder) error
}

func NewDecoder(r *bufio.Reader) *Decoder {
    return &Decoder{
        decoder: r,
    }
}

type Decoder struct {
    decoder *bufio.Reader
}

func (dec *Decoder) ReadByte() (byte, error) {
    return dec.decoder.ReadByte()
}

func (dec *Decoder) ReadVarInt() (int64, error) {
    return binary.ReadVarint(dec.decoder)
}

func (dec *Decoder) ReadBool() (ret bool) {
    var i int8
    binary.Read(dec.decoder, binary.BigEndian, &i)
    ret = (i != 0)
    return
}

func (dec *Decoder) ReadInt64() (ret int64) {
    binary.Read(dec.decoder, binary.BigEndian, &ret)
    return
}

func (dec *Decoder) ReadInt32() (ret int32) {
    binary.Read(dec.decoder, binary.BigEndian, &ret)
    return
}

func (dec *Decoder) ReadInt16() (ret int16) {
    binary.Read(dec.decoder, binary.BigEndian, &ret)
    return
}

func (dec *Decoder) ReadInt8() (ret int8) {
    binary.Read(dec.decoder, binary.BigEndian, &ret)
    return
}

func (dec *Decoder) ReadString() (ret string) {
    strLength := dec.ReadInt16()
    if int(strLength) == -1 {
        ret = ""
        return
    } else {
        buf := make([]byte, strLength)
        n, err := dec.decoder.Read(buf)
        if err != nil || n != int(strLength) {
            fmt.Printf("Couldn't read %d bytes but only %d\n", strLength, n)
        }
        ret = string(buf)
        return
    }
}

func (dec *Decoder) ReadVarIntByteArray() (ret []byte) {
    bitesLength, _ := dec.ReadVarInt()
    if int(bitesLength) == -1 {
        ret = make([]byte, 0)
        return
    } else {
        buf := make([]byte, bitesLength)
        n, err := dec.decoder.Read(buf)
        if err != nil || n != int(bitesLength) {
            fmt.Printf("Couldn't read %d bytes but only %d\n", bitesLength, n)
        }
        ret = buf
        return
    }
}

func (dec *Decoder) ReadByteArray() (ret []byte) {
    bitesLength := dec.ReadInt32()
    if int(bitesLength) == -1 {
        ret = make([]byte, 0)
        return
    } else {
        buf := make([]byte, bitesLength)
        n, err := dec.decoder.Read(buf)
        if err != nil || n != int(bitesLength) {
            fmt.Printf("Couldn't read %d bytes but only %d\n", bitesLength, n)
        }
        ret = buf
        return
    }
}

func (dec *Decoder) ReadInt32Array() (ret []int32) {
    arrayLength := dec.ReadInt32()
    if int(arrayLength) == -1 {
        ret = nil
        return
    } else {
        buf := make([]int32, arrayLength)
        var i int32
        for i = 0; i < arrayLength; i++ {
            buf[i] = dec.ReadInt32()
        }
        ret = buf
        return
    }
}

func (dec *Decoder) ReadInt64Array() (ret []int64) {
    arrayLength := dec.ReadInt32()
    if int(arrayLength) == -1 {
        ret = nil
        return
    } else {
        buf := make([]int64, arrayLength)
        var i int32
        for i = 0; i < arrayLength; i++ {
            buf[i] = dec.ReadInt64()
        }
        ret = buf
        return
    }
}

func (dec *Decoder) ReadStringArray() (ret []string) {
    arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
        ret = nil
        return
    } else {
        buf := make([]string, arrayLength)
        var i int32
        for i = 0; i < arrayLength; i++ {
            buf[i] = dec.ReadString()
        }
        ret = buf
        return
    }
}

func (dec *Decoder) Remaining() int {
    return dec.decoder.Buffered()
}
