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

type StopReplicaRequest0_Partitions struct {
    Topic string
    Partition int32
}

func (that *StopReplicaRequest0_Partitions) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32(that.Partition)
    return nil
}

func (that *StopReplicaRequest0_Partitions) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partition = dec.ReadInt32()
    return nil
}

type StopReplicaRequest0 struct {
    ControllerId int32
    ControllerEpoch int32
    DeletePartitions bool
    Partitions []*StopReplicaRequest0_Partitions
}

func (that *StopReplicaRequest0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ControllerId)
    enc.WriteInt32(that.ControllerEpoch)
    enc.WriteBool(that.DeletePartitions)
    {
        arrayLength := len(that.Partitions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Partitions[i].Encode(enc)
        }
    }

    return nil
}

func (that *StopReplicaRequest0) Decode(dec *Decoder) error {
    that.ControllerId = dec.ReadInt32()
    that.ControllerEpoch = dec.ReadInt32()
    that.DeletePartitions = dec.ReadBool()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*StopReplicaRequest0_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(StopReplicaRequest0_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

