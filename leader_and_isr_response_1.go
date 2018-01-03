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

type LeaderAndIsrResponse1_Partitions struct {
    Topic string
    Partition int32
    ErrorCode int16
}

func (that *LeaderAndIsrResponse1_Partitions) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    return nil
}

func (that *LeaderAndIsrResponse1_Partitions) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    return nil
}

type LeaderAndIsrResponse1 struct {
    ErrorCode int16
    Partitions []*LeaderAndIsrResponse1_Partitions
}

func (that *LeaderAndIsrResponse1) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    {
        arrayLength := len(that.Partitions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Partitions[i].Encode(enc)
        }
    }

    return nil
}

func (that *LeaderAndIsrResponse1) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*LeaderAndIsrResponse1_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(LeaderAndIsrResponse1_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

