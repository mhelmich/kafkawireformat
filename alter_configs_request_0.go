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

type AlterConfigsRequest0_ConfigEntries struct {
    ConfigName string
    ConfigValue string
}

func (that *AlterConfigsRequest0_ConfigEntries) Encode(enc *Encoder) error {
    enc.WriteString(that.ConfigName)
    enc.WriteString(that.ConfigValue)
    return nil
}

func (that *AlterConfigsRequest0_ConfigEntries) Decode(dec *Decoder) error {
    that.ConfigName = dec.ReadString()
    that.ConfigValue = dec.ReadString()
    return nil
}

type AlterConfigsRequest0 struct {
    Resources []*AlterConfigsRequest0_Resources
    ValidateOnly bool
}

func (that *AlterConfigsRequest0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Resources)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Resources[i].Encode(enc)
        }
    }

    enc.WriteBool(that.ValidateOnly)
    return nil
}

func (that *AlterConfigsRequest0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Resources = nil
        } else {
            buf := make([]*AlterConfigsRequest0_Resources, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AlterConfigsRequest0_Resources)
                item.Decode(dec)
                buf[i] = item
            }
            that.Resources = buf
        }
    }
    that.ValidateOnly = dec.ReadBool()
    return nil
}

type AlterConfigsRequest0_Resources struct {
    ResourceType int8
    ResourceName string
    ConfigEntries []*AlterConfigsRequest0_ConfigEntries
}

func (that *AlterConfigsRequest0_Resources) Encode(enc *Encoder) error {
    enc.WriteInt8(that.ResourceType)
    enc.WriteString(that.ResourceName)
    {
        arrayLength := len(that.ConfigEntries)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.ConfigEntries[i].Encode(enc)
        }
    }

    return nil
}

func (that *AlterConfigsRequest0_Resources) Decode(dec *Decoder) error {
    that.ResourceType = dec.ReadInt8()
    that.ResourceName = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.ConfigEntries = nil
        } else {
            buf := make([]*AlterConfigsRequest0_ConfigEntries, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AlterConfigsRequest0_ConfigEntries)
                item.Decode(dec)
                buf[i] = item
            }
            that.ConfigEntries = buf
        }
    }
    return nil
}

