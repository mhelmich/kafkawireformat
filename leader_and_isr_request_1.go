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

type LeaderAndIsrRequest1_PartitionStates struct {
    Topic string
    Partition int32
    ControllerEpoch int32
    Leader int32
    LeaderEpoch int32
    Isr []int32
    ZkVersion int32
    Replicas []int32
    IsNew bool
}

func (that *LeaderAndIsrRequest1_PartitionStates) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32(that.Partition)
    enc.WriteInt32(that.ControllerEpoch)
    enc.WriteInt32(that.Leader)
    enc.WriteInt32(that.LeaderEpoch)
    enc.WriteInt32Array(that.Isr)
    enc.WriteInt32(that.ZkVersion)
    enc.WriteInt32Array(that.Replicas)
    enc.WriteBool(that.IsNew)
    return nil
}

func (that *LeaderAndIsrRequest1_PartitionStates) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partition = dec.ReadInt32()
    that.ControllerEpoch = dec.ReadInt32()
    that.Leader = dec.ReadInt32()
    that.LeaderEpoch = dec.ReadInt32()
    that.Isr = dec.ReadInt32Array()
    that.ZkVersion = dec.ReadInt32()
    that.Replicas = dec.ReadInt32Array()
    that.IsNew = dec.ReadBool()
    return nil
}

type LeaderAndIsrRequest1 struct {
    ControllerId int32
    ControllerEpoch int32
    PartitionStates []*LeaderAndIsrRequest1_PartitionStates
    LiveLeaders []*LeaderAndIsrRequest1_LiveLeaders
}

func (that *LeaderAndIsrRequest1) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ControllerId)
    enc.WriteInt32(that.ControllerEpoch)
    {
        arrayLength := len(that.PartitionStates)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.PartitionStates[i].Encode(enc)
        }
    }

    {
        arrayLength := len(that.LiveLeaders)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.LiveLeaders[i].Encode(enc)
        }
    }

    return nil
}

func (that *LeaderAndIsrRequest1) Decode(dec *Decoder) error {
    that.ControllerId = dec.ReadInt32()
    that.ControllerEpoch = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionStates = nil
        } else {
            buf := make([]*LeaderAndIsrRequest1_PartitionStates, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(LeaderAndIsrRequest1_PartitionStates)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionStates = buf
        }
    }
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.LiveLeaders = nil
        } else {
            buf := make([]*LeaderAndIsrRequest1_LiveLeaders, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(LeaderAndIsrRequest1_LiveLeaders)
                item.Decode(dec)
                buf[i] = item
            }
            that.LiveLeaders = buf
        }
    }
    return nil
}

type LeaderAndIsrRequest1_LiveLeaders struct {
    Id int32
    Host string
    Port int32
}

func (that *LeaderAndIsrRequest1_LiveLeaders) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Id)
    enc.WriteString(that.Host)
    enc.WriteInt32(that.Port)
    return nil
}

func (that *LeaderAndIsrRequest1_LiveLeaders) Decode(dec *Decoder) error {
    that.Id = dec.ReadInt32()
    that.Host = dec.ReadString()
    that.Port = dec.ReadInt32()
    return nil
}

