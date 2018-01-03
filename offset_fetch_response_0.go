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

type OffsetFetchResponse0_Responses struct {
    Topic string
    PartitionResponses []*OffsetFetchResponse0_PartitionResponses
}

func (that *OffsetFetchResponse0_Responses) Encode(enc *Encoder) error {
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

func (that *OffsetFetchResponse0_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*OffsetFetchResponse0_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetFetchResponse0_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type OffsetFetchResponse0 struct {
    Responses []*OffsetFetchResponse0_Responses
}

func (that *OffsetFetchResponse0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Responses)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Responses[i].Encode(enc)
        }
    }

    return nil
}

func (that *OffsetFetchResponse0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*OffsetFetchResponse0_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetFetchResponse0_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    return nil
}

type OffsetFetchResponse0_PartitionResponses struct {
    Partition int32
    Offset int64
    Metadata string
    ErrorCode int16
}

func (that *OffsetFetchResponse0_PartitionResponses) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt64(that.Offset)
    enc.WriteString(that.Metadata)
    enc.WriteInt16(that.ErrorCode)
    return nil
}

func (that *OffsetFetchResponse0_PartitionResponses) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.Offset = dec.ReadInt64()
    that.Metadata = dec.ReadString()
    that.ErrorCode = dec.ReadInt16()
    return nil
}

