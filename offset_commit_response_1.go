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

type OffsetCommitResponse1 struct {
    Responses []*OffsetCommitResponse1_Responses
}

func (that *OffsetCommitResponse1) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Responses)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Responses[i].Encode(enc)
        }
    }

    return nil
}

func (that *OffsetCommitResponse1) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*OffsetCommitResponse1_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetCommitResponse1_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    return nil
}

type OffsetCommitResponse1_Responses struct {
    Topic string
    PartitionResponses []*OffsetCommitResponse1_PartitionResponses
}

func (that *OffsetCommitResponse1_Responses) Encode(enc *Encoder) error {
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

func (that *OffsetCommitResponse1_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*OffsetCommitResponse1_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetCommitResponse1_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type OffsetCommitResponse1_PartitionResponses struct {
    Partition int32
    ErrorCode int16
}

func (that *OffsetCommitResponse1_PartitionResponses) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    return nil
}

func (that *OffsetCommitResponse1_PartitionResponses) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    return nil
}

