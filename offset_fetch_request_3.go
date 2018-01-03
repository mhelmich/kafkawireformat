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

type OffsetFetchRequest3_Partitions struct {
    Partition int32
}

func (that *OffsetFetchRequest3_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    return nil
}

func (that *OffsetFetchRequest3_Partitions) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    return nil
}

type OffsetFetchRequest3_Topics struct {
    Topic string
    Partitions []*OffsetFetchRequest3_Partitions
}

func (that *OffsetFetchRequest3_Topics) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    {
        arrayLength := len(that.Partitions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Partitions[i].Encode(enc)
        }
    }

    return nil
}

func (that *OffsetFetchRequest3_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*OffsetFetchRequest3_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetFetchRequest3_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

type OffsetFetchRequest3 struct {
    GroupId string
    Topics []*OffsetFetchRequest3_Topics
}

func (that *OffsetFetchRequest3) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *OffsetFetchRequest3) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*OffsetFetchRequest3_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetFetchRequest3_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

