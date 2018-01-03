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

type MetadataResponse1_Brokers struct {
    NodeId int32
    Host string
    Port int32
    Rack string
}

func (that *MetadataResponse1_Brokers) Encode(enc *Encoder) error {
    enc.WriteInt32(that.NodeId)
    enc.WriteString(that.Host)
    enc.WriteInt32(that.Port)
    enc.WriteString(that.Rack)
    return nil
}

func (that *MetadataResponse1_Brokers) Decode(dec *Decoder) error {
    that.NodeId = dec.ReadInt32()
    that.Host = dec.ReadString()
    that.Port = dec.ReadInt32()
    that.Rack = dec.ReadString()
    return nil
}

type MetadataResponse1 struct {
    Brokers []*MetadataResponse1_Brokers
    ControllerId int32
    TopicMetadata []*MetadataResponse1_TopicMetadata
}

func (that *MetadataResponse1) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Brokers)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Brokers[i].Encode(enc)
        }
    }

    enc.WriteInt32(that.ControllerId)
    {
        arrayLength := len(that.TopicMetadata)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TopicMetadata[i].Encode(enc)
        }
    }

    return nil
}

func (that *MetadataResponse1) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Brokers = nil
        } else {
            buf := make([]*MetadataResponse1_Brokers, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(MetadataResponse1_Brokers)
                item.Decode(dec)
                buf[i] = item
            }
            that.Brokers = buf
        }
    }
    that.ControllerId = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TopicMetadata = nil
        } else {
            buf := make([]*MetadataResponse1_TopicMetadata, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(MetadataResponse1_TopicMetadata)
                item.Decode(dec)
                buf[i] = item
            }
            that.TopicMetadata = buf
        }
    }
    return nil
}

type MetadataResponse1_PartitionMetadata struct {
    ErrorCode int16
    Partition int32
    Leader int32
    Replicas []int32
    Isr []int32
}

func (that *MetadataResponse1_PartitionMetadata) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt32(that.Partition)
    enc.WriteInt32(that.Leader)
    enc.WriteInt32Array(that.Replicas)
    enc.WriteInt32Array(that.Isr)
    return nil
}

func (that *MetadataResponse1_PartitionMetadata) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.Partition = dec.ReadInt32()
    that.Leader = dec.ReadInt32()
    that.Replicas = dec.ReadInt32Array()
    that.Isr = dec.ReadInt32Array()
    return nil
}

type MetadataResponse1_TopicMetadata struct {
    ErrorCode int16
    Topic string
    IsInternal bool
    PartitionMetadata []*MetadataResponse1_PartitionMetadata
}

func (that *MetadataResponse1_TopicMetadata) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.Topic)
    enc.WriteBool(that.IsInternal)
    {
        arrayLength := len(that.PartitionMetadata)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.PartitionMetadata[i].Encode(enc)
        }
    }

    return nil
}

func (that *MetadataResponse1_TopicMetadata) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.Topic = dec.ReadString()
    that.IsInternal = dec.ReadBool()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionMetadata = nil
        } else {
            buf := make([]*MetadataResponse1_PartitionMetadata, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(MetadataResponse1_PartitionMetadata)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionMetadata = buf
        }
    }
    return nil
}

