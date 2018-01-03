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

type FetchResponse1_PartitionHeader struct {
    Partition int32
    ErrorCode int16
    HighWatermark int64
}

func (that *FetchResponse1_PartitionHeader) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt64(that.HighWatermark)
    return nil
}

func (that *FetchResponse1_PartitionHeader) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    that.HighWatermark = dec.ReadInt64()
    return nil
}

type FetchResponse1_Responses struct {
    Topic string
    PartitionResponses []*FetchResponse1_PartitionResponses
}

func (that *FetchResponse1_Responses) Encode(enc *Encoder) error {
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

func (that *FetchResponse1_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*FetchResponse1_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse1_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type FetchResponse1 struct {
    ThrottleTimeMs int32
    Responses []*FetchResponse1_Responses
}

func (that *FetchResponse1) Encode(enc *Encoder) error {
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

func (that *FetchResponse1) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*FetchResponse1_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchResponse1_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    return nil
}

type FetchResponse1_PartitionResponses struct {
    PartitionHeader *FetchResponse1_PartitionHeader
    RecordSet *RecordBatch
}

func (that *FetchResponse1_PartitionResponses) Encode(enc *Encoder) error {
    that.PartitionHeader.Encode(enc)

    that.RecordSet.Encode(enc)

    return nil
}

func (that *FetchResponse1_PartitionResponses) Decode(dec *Decoder) error {
    that.PartitionHeader = new(FetchResponse1_PartitionHeader)
    that.PartitionHeader.Decode(dec)
    that.RecordSet = new(RecordBatch)
    that.RecordSet.Decode(dec)
    return nil
}

