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

type CreatePartitionsRequest0_NewPartitions struct {
    Count int32
    Assignment []int32
}

func (that *CreatePartitionsRequest0_NewPartitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Count)
    enc.WriteInt32Array(that.Assignment)
    return nil
}

func (that *CreatePartitionsRequest0_NewPartitions) Decode(dec *Decoder) error {
    that.Count = dec.ReadInt32()
    that.Assignment = dec.ReadInt32Array()
    return nil
}

type CreatePartitionsRequest0_TopicPartitions struct {
    Topic string
    NewPartitions *CreatePartitionsRequest0_NewPartitions
}

func (that *CreatePartitionsRequest0_TopicPartitions) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    that.NewPartitions.Encode(enc)

    return nil
}

func (that *CreatePartitionsRequest0_TopicPartitions) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.NewPartitions = new(CreatePartitionsRequest0_NewPartitions)
    that.NewPartitions.Decode(dec)
    return nil
}

type CreatePartitionsRequest0 struct {
    TopicPartitions []*CreatePartitionsRequest0_TopicPartitions
    Timeout int32
    ValidateOnly bool
}

func (that *CreatePartitionsRequest0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.TopicPartitions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TopicPartitions[i].Encode(enc)
        }
    }

    enc.WriteInt32(that.Timeout)
    enc.WriteBool(that.ValidateOnly)
    return nil
}

func (that *CreatePartitionsRequest0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TopicPartitions = nil
        } else {
            buf := make([]*CreatePartitionsRequest0_TopicPartitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(CreatePartitionsRequest0_TopicPartitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.TopicPartitions = buf
        }
    }
    that.Timeout = dec.ReadInt32()
    that.ValidateOnly = dec.ReadBool()
    return nil
}

