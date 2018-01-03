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

type UpdateMetadataRequest4 struct {
    ControllerId int32
    ControllerEpoch int32
    PartitionStates []*UpdateMetadataRequest4_PartitionStates
    LiveBrokers []*UpdateMetadataRequest4_LiveBrokers
}

func (that *UpdateMetadataRequest4) Encode(enc *Encoder) error {
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
        arrayLength := len(that.LiveBrokers)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.LiveBrokers[i].Encode(enc)
        }
    }

    return nil
}

func (that *UpdateMetadataRequest4) Decode(dec *Decoder) error {
    that.ControllerId = dec.ReadInt32()
    that.ControllerEpoch = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionStates = nil
        } else {
            buf := make([]*UpdateMetadataRequest4_PartitionStates, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(UpdateMetadataRequest4_PartitionStates)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionStates = buf
        }
    }
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.LiveBrokers = nil
        } else {
            buf := make([]*UpdateMetadataRequest4_LiveBrokers, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(UpdateMetadataRequest4_LiveBrokers)
                item.Decode(dec)
                buf[i] = item
            }
            that.LiveBrokers = buf
        }
    }
    return nil
}

type UpdateMetadataRequest4_LiveBrokers struct {
    Id int32
    EndPoints []*UpdateMetadataRequest4_EndPoints
    Rack string
}

func (that *UpdateMetadataRequest4_LiveBrokers) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Id)
    {
        arrayLength := len(that.EndPoints)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.EndPoints[i].Encode(enc)
        }
    }

    enc.WriteString(that.Rack)
    return nil
}

func (that *UpdateMetadataRequest4_LiveBrokers) Decode(dec *Decoder) error {
    that.Id = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.EndPoints = nil
        } else {
            buf := make([]*UpdateMetadataRequest4_EndPoints, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(UpdateMetadataRequest4_EndPoints)
                item.Decode(dec)
                buf[i] = item
            }
            that.EndPoints = buf
        }
    }
    that.Rack = dec.ReadString()
    return nil
}

type UpdateMetadataRequest4_PartitionStates struct {
    Topic string
    Partition int32
    ControllerEpoch int32
    Leader int32
    LeaderEpoch int32
    Isr []int32
    ZkVersion int32
    Replicas []int32
    OfflineReplicas []int32
}

func (that *UpdateMetadataRequest4_PartitionStates) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32(that.Partition)
    enc.WriteInt32(that.ControllerEpoch)
    enc.WriteInt32(that.Leader)
    enc.WriteInt32(that.LeaderEpoch)
    enc.WriteInt32Array(that.Isr)
    enc.WriteInt32(that.ZkVersion)
    enc.WriteInt32Array(that.Replicas)
    enc.WriteInt32Array(that.OfflineReplicas)
    return nil
}

func (that *UpdateMetadataRequest4_PartitionStates) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partition = dec.ReadInt32()
    that.ControllerEpoch = dec.ReadInt32()
    that.Leader = dec.ReadInt32()
    that.LeaderEpoch = dec.ReadInt32()
    that.Isr = dec.ReadInt32Array()
    that.ZkVersion = dec.ReadInt32()
    that.Replicas = dec.ReadInt32Array()
    that.OfflineReplicas = dec.ReadInt32Array()
    return nil
}

type UpdateMetadataRequest4_EndPoints struct {
    Port int32
    Host string
    ListenerName string
    SecurityProtocolType int16
}

func (that *UpdateMetadataRequest4_EndPoints) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Port)
    enc.WriteString(that.Host)
    enc.WriteString(that.ListenerName)
    enc.WriteInt16(that.SecurityProtocolType)
    return nil
}

func (that *UpdateMetadataRequest4_EndPoints) Decode(dec *Decoder) error {
    that.Port = dec.ReadInt32()
    that.Host = dec.ReadString()
    that.ListenerName = dec.ReadString()
    that.SecurityProtocolType = dec.ReadInt16()
    return nil
}

