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

type AddPartitionsToTxnResponse0 struct {
    ThrottleTimeMs int32
    Errors []*AddPartitionsToTxnResponse0_Errors
}

func (that *AddPartitionsToTxnResponse0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    {
        arrayLength := len(that.Errors)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Errors[i].Encode(enc)
        }
    }

    return nil
}

func (that *AddPartitionsToTxnResponse0) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Errors = nil
        } else {
            buf := make([]*AddPartitionsToTxnResponse0_Errors, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AddPartitionsToTxnResponse0_Errors)
                item.Decode(dec)
                buf[i] = item
            }
            that.Errors = buf
        }
    }
    return nil
}

type AddPartitionsToTxnResponse0_PartitionErrors struct {
    Partition int32
    ErrorCode int16
}

func (that *AddPartitionsToTxnResponse0_PartitionErrors) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt16(that.ErrorCode)
    return nil
}

func (that *AddPartitionsToTxnResponse0_PartitionErrors) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    return nil
}

type AddPartitionsToTxnResponse0_Errors struct {
    Topic string
    PartitionErrors []*AddPartitionsToTxnResponse0_PartitionErrors
}

func (that *AddPartitionsToTxnResponse0_Errors) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    {
        arrayLength := len(that.PartitionErrors)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.PartitionErrors[i].Encode(enc)
        }
    }

    return nil
}

func (that *AddPartitionsToTxnResponse0_Errors) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionErrors = nil
        } else {
            buf := make([]*AddPartitionsToTxnResponse0_PartitionErrors, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AddPartitionsToTxnResponse0_PartitionErrors)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionErrors = buf
        }
    }
    return nil
}

