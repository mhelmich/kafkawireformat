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

type ProduceResponse1_Responses struct {
    Topic string
    PartitionResponses []*ProduceResponse1_PartitionResponses
}

func (that *ProduceResponse1_Responses) Encode(enc *Encoder) error {
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

func (that *ProduceResponse1_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*ProduceResponse1_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ProduceResponse1_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type ProduceResponse1 struct {
    Responses []*ProduceResponse1_Responses
    ThrottleTimeMs int32
}

func (that *ProduceResponse1) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Responses)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Responses[i].Encode(enc)
        }
    }

    enc.WriteInt32(that.ThrottleTimeMs)
    return nil
}

func (that *ProduceResponse1) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*ProduceResponse1_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ProduceResponse1_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    that.ThrottleTimeMs = dec.ReadInt32()
    return nil
}

type ProduceResponse1_PartitionResponses struct {
    Partition int32
    ErrorCode int16
    BaseOffset int64
}

func (that *ProduceResponse1_PartitionResponses) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt64(that.BaseOffset)
    return nil
}

func (that *ProduceResponse1_PartitionResponses) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    that.BaseOffset = dec.ReadInt64()
    return nil
}

