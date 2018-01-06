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
    "hash/crc32"
)

type Records interface {
}

// the length field of a RecordBatch denotes the length of the entire message
// this is the byte size of the static "overhead" (the size of the metadata fields)
// the number of actual bytes that contain Records (or the payload) of this message
// can be calculated by (Length - constantBatchMessageHeaderSize)
const constantBatchMessageHeaderSize = 57

type RecordBatch struct {
    FirstOffset int64
    // this is the byte size of this entire RecordBatch
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
    MaxTimestamp    int64
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
    rb.MaxTimestamp = dec.ReadInt64()
    rb.ProducerId = dec.ReadInt64()
    rb.ProducerEpoch = dec.ReadInt16()
    rb.FirstSequence = dec.ReadInt32()
    numRecords := dec.ReadInt32()
    // read all records as raw bytes first
    recordsAsBytes := dec.ReadRawBytes(int(rb.Length) - constantBatchMessageHeaderSize)
    innerDec := NewDecoderFromBytes(recordsAsBytes)
    {
        if int(numRecords) == -1 {
            rb.Records = nil
        } else {
            buf := make([]*Record, numRecords)
            var i int32
            for i = 0; i < numRecords; i++ {
                item := new(Record)
                item.Decode(innerDec)
                buf[i] = item
            }
            rb.Records = buf
        }
    }

    return nil
}

func (rb *RecordBatch) Encode(enc *Encoder) error {
    numRecords := int32(len(rb.Records))
    recordsEnc := NewEncoder()
    {
        for i := int32(0); i < numRecords; i++ {
            rb.Records[i].Encode(recordsEnc)
        }
    }

    crcCheckSumEnc := NewEncoder()
    crcCheckSumEnc.WriteInt16(rb.Attributes)
    crcCheckSumEnc.WriteInt32(rb.LastOffsetDelta)
    crcCheckSumEnc.WriteInt64(rb.FirstTimestamp)
    crcCheckSumEnc.WriteInt64(rb.MaxTimestamp)
    crcCheckSumEnc.WriteInt64(rb.ProducerId)
    crcCheckSumEnc.WriteInt16(rb.ProducerEpoch)
    crcCheckSumEnc.WriteInt32(rb.FirstSequence)
    // yes, this is the number of records (not the byte size!!!)
    crcCheckSumEnc.WriteInt32(numRecords)
    crcCheckSumEnc.WriteRawBytes(recordsEnc.Bytes())

    // write the byte size of the RecordBatch structure first
    // that'll be my calculation plus these magical 4 bytes
    enc.WriteInt32(int32(constantBatchMessageHeaderSize + recordsEnc.Len() + 4))
    enc.WriteInt64(rb.FirstOffset)
    // this is the byte size of the remainder of the RecordBatch (basically the size of the message from here on including me [the size int32])
    enc.WriteInt32(int32(constantBatchMessageHeaderSize + recordsEnc.Len() - 8)) // this might be -8 (the first offset (8 for the int64) that we skipped)
    enc.WriteInt32(rb.PartitionLeaderEpoch)
    enc.WriteInt8(int8(2))

    crcCheckSum := crc32.Checksum(crcCheckSumEnc.Bytes(), crc32.MakeTable(crc32.Castagnoli))
    enc.WriteInt32(int32(crcCheckSum))
    enc.WriteRawBytes(crcCheckSumEnc.Bytes())
    return nil
}

type Record struct {
    // Length int64 // varint byte size of the Record
    // currently unused
    Attributes     int8
    TimestampDelta int64  // varint
    OffsetDelta    int64  // varint
    Key            []byte // data
    Value          []byte // data
    Headers        []*Header
}

func (r *Record) Decode(dec *Decoder) error {
    // This looks a little different than any other
    // message because a Record is lead with its byte size
    // as varint. That means I need to serialize the message
    // first in order to calculate its byte size.
    innerDec := NewDecoderFromBytes(dec.ReadVarIntByteArray())
    r.Attributes = innerDec.ReadInt8()
    r.TimestampDelta, _ = innerDec.ReadVarInt()
    r.OffsetDelta, _ = innerDec.ReadVarInt()
    r.Key = innerDec.ReadVarIntByteArray()
    r.Value = innerDec.ReadVarIntByteArray()
    {
        arrayLength := innerDec.ReadInt32()
        if int(arrayLength) == -1 {
            r.Headers = nil
        } else {
            buf := make([]*Header, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(Header)
                item.Decode(innerDec)
                buf[i] = item
            }
            r.Headers = buf
        }
    }
    return nil
}

func (r *Record) Encode(enc *Encoder) error {
    // This looks a little different than any other
    // message because a Record is lead with its byte size
    // as varint. That means I need to serialize the message
    // first in order to calculate its byte size.
    innerEnc := NewEncoder()
    innerEnc.WriteInt8(r.Attributes)
    innerEnc.WriteVarInt(r.TimestampDelta)
    innerEnc.WriteVarInt(r.OffsetDelta)
    innerEnc.WriteVarIntByteArray(r.Key)
    innerEnc.WriteVarIntByteArray(r.Value)
    {
        arrayLength := len(r.Headers)
        innerEnc.WriteVarInt(int64(arrayLength))
        for i := 0; i < arrayLength; i++ {
            r.Headers[i].Encode(innerEnc)
        }
    }
    enc.WriteVarIntByteArray(innerEnc.Bytes())
    return nil
}

type Header struct {
    HeaderKey   string
    HeaderValue []byte // this is just bytes -- the consumer needs to figure out what this is
}

func (h *Header) Decode(dec *Decoder) error {
    h.HeaderKey = string(dec.ReadVarIntByteArray())
    h.HeaderValue = dec.ReadVarIntByteArray()
    return nil
}

func (h *Header) Encode(enc *Encoder) error {
    enc.WriteVarIntByteArray([]byte(h.HeaderKey))
    enc.WriteVarIntByteArray(h.HeaderValue)
    return nil
}
