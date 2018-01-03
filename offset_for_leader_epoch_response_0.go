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

type OffsetForLeaderEpochResponse0_Partitions struct {
    ErrorCode int16
    Partition int32
    EndOffset int64
}

func (that *OffsetForLeaderEpochResponse0_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt32(that.Partition)
    enc.WriteInt64(that.EndOffset)
    return nil
}

func (that *OffsetForLeaderEpochResponse0_Partitions) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.Partition = dec.ReadInt32()
    that.EndOffset = dec.ReadInt64()
    return nil
}

type OffsetForLeaderEpochResponse0 struct {
    Topics []*OffsetForLeaderEpochResponse0_Topics
}

func (that *OffsetForLeaderEpochResponse0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *OffsetForLeaderEpochResponse0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*OffsetForLeaderEpochResponse0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetForLeaderEpochResponse0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

type OffsetForLeaderEpochResponse0_Topics struct {
    Topic string
    Partitions []*OffsetForLeaderEpochResponse0_Partitions
}

func (that *OffsetForLeaderEpochResponse0_Topics) Encode(enc *Encoder) error {
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

func (that *OffsetForLeaderEpochResponse0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*OffsetForLeaderEpochResponse0_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(OffsetForLeaderEpochResponse0_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

