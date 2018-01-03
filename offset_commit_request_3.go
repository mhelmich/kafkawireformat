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

type OffsetCommitRequest3_Partitions struct {
    Partition int32
    Offset int64
    Metadata string
}

func (that *OffsetCommitRequest3_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt64(that.Offset)
    enc.WriteString(that.Metadata)
    return nil
}

func (that *OffsetCommitRequest3_Partitions) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.Offset = dec.ReadInt64()
    that.Metadata = dec.ReadString()
    return nil
}

type OffsetCommitRequest3_Topics struct {
    Topic string
    Partitions []*OffsetCommitRequest3_Partitions
}

func (that *OffsetCommitRequest3_Topics) Encode(enc *Encoder) error {
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

func (that *OffsetCommitRequest3_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*OffsetCommitRequest3_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetCommitRequest3_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

type OffsetCommitRequest3 struct {
    GroupId string
    GenerationId int32
    MemberId string
    RetentionTime int64
    Topics []*OffsetCommitRequest3_Topics
}

func (that *OffsetCommitRequest3) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    enc.WriteInt32(that.GenerationId)
    enc.WriteString(that.MemberId)
    enc.WriteInt64(that.RetentionTime)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *OffsetCommitRequest3) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    that.GenerationId = dec.ReadInt32()
    that.MemberId = dec.ReadString()
    that.RetentionTime = dec.ReadInt64()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*OffsetCommitRequest3_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetCommitRequest3_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

