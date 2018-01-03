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

type ListOffsetsRequest0_Partitions struct {
    Partition int32
    Timestamp int64
    MaxNumOffsets int32
}

func (that *ListOffsetsRequest0_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt64(that.Timestamp)
    enc.WriteInt32(that.MaxNumOffsets)
    return nil
}

func (that *ListOffsetsRequest0_Partitions) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.Timestamp = dec.ReadInt64()
    that.MaxNumOffsets = dec.ReadInt32()
    return nil
}

type ListOffsetsRequest0 struct {
    ReplicaId int32
    Topics []*ListOffsetsRequest0_Topics
}

func (that *ListOffsetsRequest0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ReplicaId)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *ListOffsetsRequest0) Decode(dec *Decoder) error {
    that.ReplicaId = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*ListOffsetsRequest0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ListOffsetsRequest0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

type ListOffsetsRequest0_Topics struct {
    Topic string
    Partitions []*ListOffsetsRequest0_Partitions
}

func (that *ListOffsetsRequest0_Topics) Encode(enc *Encoder) error {
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

func (that *ListOffsetsRequest0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*ListOffsetsRequest0_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ListOffsetsRequest0_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

