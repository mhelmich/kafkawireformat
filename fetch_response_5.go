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

type FetchResponse5 struct {
    ThrottleTimeMs int32
    Responses []*FetchResponse5_Responses
}

func (that *FetchResponse5) Encode(enc *Encoder) error {
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

func (that *FetchResponse5) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*FetchResponse5_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse5_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    return nil
}

type FetchResponse5_PartitionHeader struct {
    Partition int32
    ErrorCode int16
    HighWatermark int64
    LastStableOffset int64
    LogStartOffset int64
    AbortedTransactions []*FetchResponse5_AbortedTransactions
}

func (that *FetchResponse5_PartitionHeader) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt64(that.HighWatermark)
    enc.WriteInt64(that.LastStableOffset)
    enc.WriteInt64(that.LogStartOffset)
    {
        arrayLength := len(that.AbortedTransactions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.AbortedTransactions[i].Encode(enc)
        }
    }

    return nil
}

func (that *FetchResponse5_PartitionHeader) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    that.HighWatermark = dec.ReadInt64()
    that.LastStableOffset = dec.ReadInt64()
    that.LogStartOffset = dec.ReadInt64()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.AbortedTransactions = nil
        } else {
            buf := make([]*FetchResponse5_AbortedTransactions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse5_AbortedTransactions)
                item.Decode(dec)
                buf[i] = item
            }
            that.AbortedTransactions = buf
        }
    }
    return nil
}

type FetchResponse5_AbortedTransactions struct {
    ProducerId int64
    FirstOffset int64
}

func (that *FetchResponse5_AbortedTransactions) Encode(enc *Encoder) error {
    enc.WriteInt64(that.ProducerId)
    enc.WriteInt64(that.FirstOffset)
    return nil
}

func (that *FetchResponse5_AbortedTransactions) Decode(dec *Decoder) error {
    that.ProducerId = dec.ReadInt64()
    that.FirstOffset = dec.ReadInt64()
    return nil
}

type FetchResponse5_Responses struct {
    Topic string
    PartitionResponses []*FetchResponse5_PartitionResponses
}

func (that *FetchResponse5_Responses) Encode(enc *Encoder) error {
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

func (that *FetchResponse5_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*FetchResponse5_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse5_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type FetchResponse5_PartitionResponses struct {
    PartitionHeader *FetchResponse5_PartitionHeader
    RecordSet *RecordBatch
}

func (that *FetchResponse5_PartitionResponses) Encode(enc *Encoder) error {
    that.PartitionHeader.Encode(enc)

    that.RecordSet.Encode(enc)

    return nil
}

func (that *FetchResponse5_PartitionResponses) Decode(dec *Decoder) error {
    that.PartitionHeader = new(FetchResponse5_PartitionHeader)
    that.PartitionHeader.Decode(dec)
    that.RecordSet = new(RecordBatch)
    that.RecordSet.Decode(dec)
    return nil
}

