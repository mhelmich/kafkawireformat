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

type FetchRequest4_Partitions struct {
    Partition int32
    FetchOffset int64
    MaxBytes int32
}

func (that *FetchRequest4_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt64(that.FetchOffset)
    enc.WriteInt32(that.MaxBytes)
    return nil
}

func (that *FetchRequest4_Partitions) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.FetchOffset = dec.ReadInt64()
    that.MaxBytes = dec.ReadInt32()
    return nil
}

type FetchRequest4_Topics struct {
    Topic string
    Partitions []*FetchRequest4_Partitions
}

func (that *FetchRequest4_Topics) Encode(enc *Encoder) error {
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

func (that *FetchRequest4_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*FetchRequest4_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchRequest4_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

type FetchRequest4 struct {
    ReplicaId int32
    MaxWaitTime int32
    MinBytes int32
    MaxBytes int32
    IsolationLevel int8
    Topics []*FetchRequest4_Topics
}

func (that *FetchRequest4) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ReplicaId)
    enc.WriteInt32(that.MaxWaitTime)
    enc.WriteInt32(that.MinBytes)
    enc.WriteInt32(that.MaxBytes)
    enc.WriteInt8(that.IsolationLevel)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *FetchRequest4) Decode(dec *Decoder) error {
    that.ReplicaId = dec.ReadInt32()
    that.MaxWaitTime = dec.ReadInt32()
    that.MinBytes = dec.ReadInt32()
    that.MaxBytes = dec.ReadInt32()
    that.IsolationLevel = dec.ReadInt8()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*FetchRequest4_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(FetchRequest4_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

