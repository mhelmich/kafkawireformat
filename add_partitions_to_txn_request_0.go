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

type AddPartitionsToTxnRequest0 struct {
    TransactionalId string
    ProducerId int64
    ProducerEpoch int16
    Topics []*AddPartitionsToTxnRequest0_Topics
}

func (that *AddPartitionsToTxnRequest0) Encode(enc *Encoder) error {
    enc.WriteString(that.TransactionalId)
    enc.WriteInt64(that.ProducerId)
    enc.WriteInt16(that.ProducerEpoch)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *AddPartitionsToTxnRequest0) Decode(dec *Decoder) error {
    that.TransactionalId = dec.ReadString()
    that.ProducerId = dec.ReadInt64()
    that.ProducerEpoch = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*AddPartitionsToTxnRequest0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AddPartitionsToTxnRequest0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

type AddPartitionsToTxnRequest0_Topics struct {
    Topic string
    Partitions []int32
}

func (that *AddPartitionsToTxnRequest0_Topics) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32Array(that.Partitions)
    return nil
}

func (that *AddPartitionsToTxnRequest0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partitions = dec.ReadInt32Array()
    return nil
}

