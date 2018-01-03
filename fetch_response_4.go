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
import ()

type FetchResponse4 struct {
    ThrottleTimeMs int32
    Responses []*FetchResponse4_Responses
}

func (that *FetchResponse4) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    {
        arrayLength := len(that.Responses)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Responses[i].Encode(enc)
        }
    }

    return nil
}

func (that *FetchResponse4) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*FetchResponse4_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse4_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    return nil
}

type FetchResponse4_PartitionHeader struct {
    Partition int32
    ErrorCode int16
    HighWatermark int64
    LastStableOffset int64
    AbortedTransactions []*FetchResponse4_AbortedTransactions
}

func (that *FetchResponse4_PartitionHeader) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt64(that.HighWatermark)
    enc.WriteInt64(that.LastStableOffset)
    {
        arrayLength := len(that.AbortedTransactions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.AbortedTransactions[i].Encode(enc)
        }
    }

    return nil
}

func (that *FetchResponse4_PartitionHeader) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    that.HighWatermark = dec.ReadInt64()
    that.LastStableOffset = dec.ReadInt64()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.AbortedTransactions = nil
        } else {
            buf := make([]*FetchResponse4_AbortedTransactions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse4_AbortedTransactions)
                item.Decode(dec)
                buf[i] = item
            }
            that.AbortedTransactions = buf
        }
    }
    return nil
}

type FetchResponse4_AbortedTransactions struct {
    ProducerId int64
    FirstOffset int64
}

func (that *FetchResponse4_AbortedTransactions) Encode(enc *Encoder) error {
    enc.WriteInt64(that.ProducerId)
    enc.WriteInt64(that.FirstOffset)
    return nil
}

func (that *FetchResponse4_AbortedTransactions) Decode(dec *Decoder) error {
    that.ProducerId = dec.ReadInt64()
    that.FirstOffset = dec.ReadInt64()
    return nil
}

type FetchResponse4_Responses struct {
    Topic string
    PartitionResponses []*FetchResponse4_PartitionResponses
}

func (that *FetchResponse4_Responses) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    {
        arrayLength := len(that.PartitionResponses)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.PartitionResponses[i].Encode(enc)
        }
    }

    return nil
}

func (that *FetchResponse4_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*FetchResponse4_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse4_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type FetchResponse4_PartitionResponses struct {
    PartitionHeader *FetchResponse4_PartitionHeader
    RecordSet *RecordBatch
}

func (that *FetchResponse4_PartitionResponses) Encode(enc *Encoder) error {
    that.PartitionHeader.Encode(enc)

    that.RecordSet.Encode(enc)

    return nil
}

func (that *FetchResponse4_PartitionResponses) Decode(dec *Decoder) error {
    that.PartitionHeader = new(FetchResponse4_PartitionHeader)
    that.PartitionHeader.Decode(dec)
    that.RecordSet = new(RecordBatch)
    that.RecordSet.Decode(dec)
    return nil
}

