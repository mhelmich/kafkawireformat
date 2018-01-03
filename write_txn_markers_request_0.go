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

type WriteTxnMarkersRequest0_TransactionMarkers struct {
    ProducerId int64
    ProducerEpoch int16
    TransactionResult bool
    Topics []*WriteTxnMarkersRequest0_Topics
    CoordinatorEpoch int32
}

func (that *WriteTxnMarkersRequest0_TransactionMarkers) Encode(enc *Encoder) error {
    enc.WriteInt64(that.ProducerId)
    enc.WriteInt16(that.ProducerEpoch)
    enc.WriteBool(that.TransactionResult)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    enc.WriteInt32(that.CoordinatorEpoch)
    return nil
}

func (that *WriteTxnMarkersRequest0_TransactionMarkers) Decode(dec *Decoder) error {
    that.ProducerId = dec.ReadInt64()
    that.ProducerEpoch = dec.ReadInt16()
    that.TransactionResult = dec.ReadBool()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*WriteTxnMarkersRequest0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(WriteTxnMarkersRequest0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    that.CoordinatorEpoch = dec.ReadInt32()
    return nil
}

type WriteTxnMarkersRequest0_Topics struct {
    Topic string
    Partitions []int32
}

func (that *WriteTxnMarkersRequest0_Topics) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32Array(that.Partitions)
    return nil
}

func (that *WriteTxnMarkersRequest0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partitions = dec.ReadInt32Array()
    return nil
}

type WriteTxnMarkersRequest0 struct {
    TransactionMarkers []*WriteTxnMarkersRequest0_TransactionMarkers
}

func (that *WriteTxnMarkersRequest0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.TransactionMarkers)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TransactionMarkers[i].Encode(enc)
        }
    }

    return nil
}

func (that *WriteTxnMarkersRequest0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TransactionMarkers = nil
        } else {
            buf := make([]*WriteTxnMarkersRequest0_TransactionMarkers, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(WriteTxnMarkersRequest0_TransactionMarkers)
                item.Decode(dec)
                buf[i] = item
            }
            that.TransactionMarkers = buf
        }
    }
    return nil
}

