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

type CreateTopicsRequest2_ReplicaAssignment struct {
    Partition int32
    Replicas []int32
}

func (that *CreateTopicsRequest2_ReplicaAssignment) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt32Array(that.Replicas)
    return nil
}

func (that *CreateTopicsRequest2_ReplicaAssignment) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.Replicas = dec.ReadInt32Array()
    return nil
}

type CreateTopicsRequest2_CreateTopicRequests struct {
    Topic string
    NumPartitions int32
    ReplicationFactor int16
    ReplicaAssignment []*CreateTopicsRequest2_ReplicaAssignment
    ConfigEntries []*CreateTopicsRequest2_ConfigEntries
}

func (that *CreateTopicsRequest2_CreateTopicRequests) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32(that.NumPartitions)
    enc.WriteInt16(that.ReplicationFactor)
    {
        arrayLength := len(that.ReplicaAssignment)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.ReplicaAssignment[i].Encode(enc)
        }
    }

    {
        arrayLength := len(that.ConfigEntries)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.ConfigEntries[i].Encode(enc)
        }
    }

    return nil
}

func (that *CreateTopicsRequest2_CreateTopicRequests) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.NumPartitions = dec.ReadInt32()
    that.ReplicationFactor = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.ReplicaAssignment = nil
        } else {
            buf := make([]*CreateTopicsRequest2_ReplicaAssignment, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(CreateTopicsRequest2_ReplicaAssignment)
                item.Decode(dec)
                buf[i] = item
            }
            that.ReplicaAssignment = buf
        }
    }
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.ConfigEntries = nil
        } else {
            buf := make([]*CreateTopicsRequest2_ConfigEntries, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(CreateTopicsRequest2_ConfigEntries)
                item.Decode(dec)
                buf[i] = item
            }
            that.ConfigEntries = buf
        }
    }
    return nil
}

type CreateTopicsRequest2_ConfigEntries struct {
    ConfigName string
    ConfigValue string
}

func (that *CreateTopicsRequest2_ConfigEntries) Encode(enc *Encoder) error {
    enc.WriteString(that.ConfigName)
    enc.WriteString(that.ConfigValue)
    return nil
}

func (that *CreateTopicsRequest2_ConfigEntries) Decode(dec *Decoder) error {
    that.ConfigName = dec.ReadString()
    that.ConfigValue = dec.ReadString()
    return nil
}

type CreateTopicsRequest2 struct {
    CreateTopicRequests []*CreateTopicsRequest2_CreateTopicRequests
    Timeout int32
    ValidateOnly bool
}

func (that *CreateTopicsRequest2) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.CreateTopicRequests)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.CreateTopicRequests[i].Encode(enc)
        }
    }

    enc.WriteInt32(that.Timeout)
    enc.WriteBool(that.ValidateOnly)
    return nil
}

func (that *CreateTopicsRequest2) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.CreateTopicRequests = nil
        } else {
            buf := make([]*CreateTopicsRequest2_CreateTopicRequests, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(CreateTopicsRequest2_CreateTopicRequests)
                item.Decode(dec)
                buf[i] = item
            }
            that.CreateTopicRequests = buf
        }
    }
    that.Timeout = dec.ReadInt32()
    that.ValidateOnly = dec.ReadBool()
    return nil
}

