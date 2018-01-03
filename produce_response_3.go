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

type ProduceResponse3_Responses struct {
    Topic string
    PartitionResponses []*ProduceResponse3_PartitionResponses
}

func (that *ProduceResponse3_Responses) Encode(enc *Encoder) error {
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

func (that *ProduceResponse3_Responses) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionResponses = nil
        } else {
            buf := make([]*ProduceResponse3_PartitionResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ProduceResponse3_PartitionResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionResponses = buf
        }
    }
    return nil
}

type ProduceResponse3 struct {
    Responses []*ProduceResponse3_Responses
    ThrottleTimeMs int32
}

func (that *ProduceResponse3) Encode(enc *Encoder) error {
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

func (that *ProduceResponse3) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Responses = nil
        } else {
            buf := make([]*ProduceResponse3_Responses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ProduceResponse3_Responses)
                item.Decode(dec)
                buf[i] = item
            }
            that.Responses = buf
        }
    }
    that.ThrottleTimeMs = dec.ReadInt32()
    return nil
}

type ProduceResponse3_PartitionResponses struct {
    Partition int32
    ErrorCode int16
    BaseOffset int64
    LogAppendTime int64
}

func (that *ProduceResponse3_PartitionResponses) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt64(that.BaseOffset)
    enc.WriteInt64(that.LogAppendTime)
    return nil
}

func (that *ProduceResponse3_PartitionResponses) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    that.BaseOffset = dec.ReadInt64()
    that.LogAppendTime = dec.ReadInt64()
    return nil
}

