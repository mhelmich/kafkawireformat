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

type Records interface {
}

type RecordBatch struct {
    FirstOffset          int64
    Length               int32
    PartitionLeaderEpoch int32
    Magic                int8
    CRC                  int32
    // lowest three bits contain the compression algorithm used for the message (0 -> none, 1 -> gzip, 2 -> snappy)
    // the fourth lowest bit represents the timestamp type (1 -> LogAppendType, producers should always set this to 0)
    // the fith lowest bit indicates whether this record batch is part of a transaction or not (0 -> not transactional, 1 -> transactional)
    // the sixth lowest bit indicates whether the record batch includes a control message (1 -> contains control message, 0 -> doesn't contain a control message)
    Attributes      int16
    LastOffsetDelta int32
    FirstTimestamp  int64
    ProducerId      int64
    ProducerEpoch   int16
    FirstSequence   int32
    Records         []*Record
}

func (rb *RecordBatch) Decode(dec *Decoder) error {
    rb.FirstOffset = dec.ReadInt64()
    rb.Length = dec.ReadInt32()
    rb.PartitionLeaderEpoch = dec.ReadInt32()
    rb.Magic = dec.ReadInt8()
    rb.CRC = dec.ReadInt32()
    rb.Attributes = dec.ReadInt16()
    rb.LastOffsetDelta = dec.ReadInt32()
    rb.FirstTimestamp = dec.ReadInt64()
    rb.ProducerId = dec.ReadInt64()
    rb.ProducerEpoch = dec.ReadInt16()
    rb.FirstSequence = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            rb.Records = nil
        } else {
            buf := make([]*Record, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(Record)
                item.Decode(dec)
                buf[i] = item
            }
            rb.Records = buf
        }
    }
    return nil
}

func (rb *RecordBatch) Encode(enc *Encoder) error {
    enc.WriteInt64(rb.FirstOffset)
    enc.WriteInt32(rb.Length)
    enc.WriteInt32(rb.PartitionLeaderEpoch)
    enc.WriteInt8(rb.Magic)
    enc.WriteInt32(rb.CRC)
    enc.WriteInt16(rb.Attributes)
    enc.WriteInt32(rb.LastOffsetDelta)
    enc.WriteInt64(rb.FirstTimestamp)
    enc.WriteInt64(rb.ProducerId)
    enc.WriteInt16(rb.ProducerEpoch)
    enc.WriteInt32(rb.FirstSequence)
    {
        l := len(rb.Records)
        enc.WriteInt32(int32(l))
        for i := 0; i < l; i++ {
            rb.Records[i].Encode(enc)
        }
    }
    return nil
}

type Record struct {
    Length int64 // varint
    // currently unused
    Attributes     int8
    TimestampDelta int64  // varint
    OffsetDelta    int64  // varint
    KeyLen         int64  // varint
    Key            []byte // data
    ValueLen       int64  // varint
    Value          []byte // data
    Headers        []*Header
}

func (r *Record) Decode(dec *Decoder) error {
    r.Length, _ = dec.ReadVarInt()
    r.Attributes = dec.ReadInt8()
    r.TimestampDelta, _ = dec.ReadVarInt()
    r.OffsetDelta, _ = dec.ReadVarInt()
    r.KeyLen, _ = dec.ReadVarInt()
    r.Key = dec.ReadByteArray()
    r.ValueLen, _ = dec.ReadVarInt()
    r.Value = dec.ReadByteArray()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            r.Headers = nil
        } else {
            buf := make([]*Header, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(Header)
                item.Decode(dec)
                buf[i] = item
            }
            r.Headers = buf
        }
    }
    return nil
}

func (r *Record) Encode(enc *Encoder) error {
    enc.WriteVarInt(r.Length)
    enc.WriteInt8(r.Attributes)
    enc.WriteVarInt(r.TimestampDelta)
    enc.WriteVarInt(r.OffsetDelta)
    enc.WriteVarInt(r.KeyLen)
    enc.WriteByteArray(r.Key)
    enc.WriteVarInt(r.ValueLen)
    enc.WriteByteArray(r.Value)
    {
        l := len(r.Headers)
        enc.WriteInt32(int32(l))
        for i := 0; i < l; i++ {
            r.Headers[i].Encode(enc)
        }
    }
    return nil
}

type Header struct {
    HeaderKeyLen   int64 // varint
    HeaderKey      string
    HeaderValueLen int64  // varint
    HeaderValue    []byte // data
}

func (h *Header) Decode(dec *Decoder) error {
    h.HeaderKeyLen, _ = dec.ReadVarInt()
    h.HeaderKey = dec.ReadString()
    h.HeaderValueLen, _ = dec.ReadVarInt()
    h.HeaderValue = dec.ReadByteArray()
    return nil
}

func (h *Header) Encode(enc *Encoder) error {
    enc.WriteVarInt(h.HeaderKeyLen)
    enc.WriteString(h.HeaderKey)
    enc.WriteVarInt(h.HeaderValueLen)
    enc.WriteByteArray(h.HeaderValue)
    return nil
}
