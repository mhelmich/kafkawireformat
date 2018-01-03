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

type ControlledShutdownResponse0 struct {
    ErrorCode int16
    PartitionsRemaining []*ControlledShutdownResponse0_PartitionsRemaining
}

func (that *ControlledShutdownResponse0) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    {
        arrayLength := len(that.PartitionsRemaining)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.PartitionsRemaining[i].Encode(enc)
        }
    }

    return nil
}

func (that *ControlledShutdownResponse0) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.PartitionsRemaining = nil
        } else {
            buf := make([]*ControlledShutdownResponse0_PartitionsRemaining, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ControlledShutdownResponse0_PartitionsRemaining)
                item.Decode(dec)
                buf[i] = item
            }
            that.PartitionsRemaining = buf
        }
    }
    return nil
}

type ControlledShutdownResponse0_PartitionsRemaining struct {
    Topic string
    Partition int32
}

func (that *ControlledShutdownResponse0_PartitionsRemaining) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32(that.Partition)
    return nil
}

func (that *ControlledShutdownResponse0_PartitionsRemaining) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partition = dec.ReadInt32()
    return nil
}

