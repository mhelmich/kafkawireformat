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

type AddOffsetsToTxnRequest0 struct {
    TransactionalId string
    ProducerId int64
    ProducerEpoch int16
    GroupId string
}

func (that *AddOffsetsToTxnRequest0) Encode(enc *Encoder) error {
    enc.WriteString(that.TransactionalId)
    enc.WriteInt64(that.ProducerId)
    enc.WriteInt16(that.ProducerEpoch)
    enc.WriteString(that.GroupId)
    return nil
}

func (that *AddOffsetsToTxnRequest0) Decode(dec *Decoder) error {
    that.TransactionalId = dec.ReadString()
    that.ProducerId = dec.ReadInt64()
    that.ProducerEpoch = dec.ReadInt16()
    that.GroupId = dec.ReadString()
    return nil
}

