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

type MetadataResponse0_Brokers struct {
    NodeId int32
    Host string
    Port int32
}

func (that *MetadataResponse0_Brokers) Encode(enc *Encoder) error {
    enc.WriteInt32(that.NodeId)
    enc.WriteString(that.Host)
    enc.WriteInt32(that.Port)
    return nil
}

func (that *MetadataResponse0_Brokers) Decode(dec *Decoder) error {
    that.NodeId = dec.ReadInt32()
    that.Host = dec.ReadString()
    that.Port = dec.ReadInt32()
    return nil
}

type MetadataResponse0 struct {
    Brokers []*MetadataResponse0_Brokers
    TopicMetadata []*MetadataResponse0_TopicMetadata
}

func (that *MetadataResponse0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Brokers)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Brokers[i].Encode(enc)
        }
    }

    {
        arrayLength := len(that.TopicMetadata)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TopicMetadata[i].Encode(enc)
        }
    }

    return nil
}

func (that *MetadataResponse0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Brokers = nil
        } else {
            buf := make([]*MetadataResponse0_Brokers, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(MetadataResponse0_Brokers)
                item.Decode(dec)
                buf[i] = item
            }
            that.Brokers = buf
        }
    }
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TopicMetadata = nil
        } else {
            buf := make([]*MetadataResponse0_TopicMetadata, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(MetadataResponse0_TopicMetadata)
                item.Decode(dec)
                buf[i] = item
            }
            that.TopicMetadata = buf
        }
    }
    return nil
}

type MetadataResponse0_PartitionMetadata struct {
    ErrorCode int16
    Partition int32
    Leader int32
    Replicas []int32
    Isr []int32
}

func (that *MetadataResponse0_PartitionMetadata) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt32(that.Partition)
    enc.WriteInt32(that.Leader)
    enc.WriteInt32Array(that.Replicas)
    enc.WriteInt32Array(that.Isr)
    return nil
}

func (that *MetadataResponse0_PartitionMetadata) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.Partition = dec.ReadInt32()
    that.Leader = dec.ReadInt32()
    that.Replicas = dec.ReadInt32Array()
    that.Isr = dec.ReadInt32Array()
    return nil
}

type MetadataResponse0_TopicMetadata struct {
    ErrorCode int16
    Topic string
    PartitionMetadata []*MetadataResponse0_PartitionMetadata
}

func (that *MetadataResponse0_TopicMetadata) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.Topic)
    {
        arrayLength := len(that.PartitionMetadata)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.PartitionMetadata[i].Encode(enc)
        }
    }

    return nil
}

func (that *MetadataResponse0_TopicMetadata) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionMetadata = nil
        } else {
            buf := make([]*MetadataResponse0_PartitionMetadata, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(MetadataResponse0_PartitionMetadata)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionMetadata = buf
        }
    }
    return nil
}

