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

type WriteTxnMarkersResponse0_Partitions struct {
    Partition int32
    ErrorCode int16
}

func (that *WriteTxnMarkersResponse0_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    return nil
}

func (that *WriteTxnMarkersResponse0_Partitions) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    return nil
}

type WriteTxnMarkersResponse0_TransactionMarkers struct {
    ProducerId int64
    Topics []*WriteTxnMarkersResponse0_Topics
}

func (that *WriteTxnMarkersResponse0_TransactionMarkers) Encode(enc *Encoder) error {
    enc.WriteInt64(that.ProducerId)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *WriteTxnMarkersResponse0_TransactionMarkers) Decode(dec *Decoder) error {
    that.ProducerId = dec.ReadInt64()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*WriteTxnMarkersResponse0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(WriteTxnMarkersResponse0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

type WriteTxnMarkersResponse0_Topics struct {
    Topic string
    Partitions []*WriteTxnMarkersResponse0_Partitions
}

func (that *WriteTxnMarkersResponse0_Topics) Encode(enc *Encoder) error {
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

func (that *WriteTxnMarkersResponse0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*WriteTxnMarkersResponse0_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(WriteTxnMarkersResponse0_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

type WriteTxnMarkersResponse0 struct {
    TransactionMarkers []*WriteTxnMarkersResponse0_TransactionMarkers
}

func (that *WriteTxnMarkersResponse0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.TransactionMarkers)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TransactionMarkers[i].Encode(enc)
        }
    }

    return nil
}

func (that *WriteTxnMarkersResponse0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TransactionMarkers = nil
        } else {
            buf := make([]*WriteTxnMarkersResponse0_TransactionMarkers, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(WriteTxnMarkersResponse0_TransactionMarkers)
                item.Decode(dec)
                buf[i] = item
            }
            that.TransactionMarkers = buf
        }
    }
    return nil
}

