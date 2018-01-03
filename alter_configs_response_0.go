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

type AlterConfigsResponse0 struct {
    ThrottleTimeMs int32
    Resources []*AlterConfigsResponse0_Resources
}

func (that *AlterConfigsResponse0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    {
        arrayLength := len(that.Resources)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Resources[i].Encode(enc)
        }
    }

    return nil
}

func (that *AlterConfigsResponse0) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Resources = nil
        } else {
            buf := make([]*AlterConfigsResponse0_Resources, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AlterConfigsResponse0_Resources)
                item.Decode(dec)
                buf[i] = item
            }
            that.Resources = buf
        }
    }
    return nil
}

type AlterConfigsResponse0_Resources struct {
    ErrorCode int16
    ErrorMessage string
    ResourceType int8
    ResourceName string
}

func (that *AlterConfigsResponse0_Resources) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.ErrorMessage)
    enc.WriteInt8(that.ResourceType)
    enc.WriteString(that.ResourceName)
    return nil
}

func (that *AlterConfigsResponse0_Resources) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.ErrorMessage = dec.ReadString()
    that.ResourceType = dec.ReadInt8()
    that.ResourceName = dec.ReadString()
    return nil
}

