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

type DescribeConfigsRequest0 struct {
    Resources []*DescribeConfigsRequest0_Resources
}

func (that *DescribeConfigsRequest0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Resources)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Resources[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeConfigsRequest0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Resources = nil
        } else {
            buf := make([]*DescribeConfigsRequest0_Resources, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeConfigsRequest0_Resources)
                item.Decode(dec)
                buf[i] = item
            }
            that.Resources = buf
        }
    }
    return nil
}

type DescribeConfigsRequest0_Resources struct {
    ResourceType int8
    ResourceName string
    ConfigNames []string
}

func (that *DescribeConfigsRequest0_Resources) Encode(enc *Encoder) error {
    enc.WriteInt8(that.ResourceType)
    enc.WriteString(that.ResourceName)
    enc.WriteStringArray(that.ConfigNames)
    return nil
}

func (that *DescribeConfigsRequest0_Resources) Decode(dec *Decoder) error {
    that.ResourceType = dec.ReadInt8()
    that.ResourceName = dec.ReadString()
    that.ConfigNames = dec.ReadStringArray()
    return nil
}

