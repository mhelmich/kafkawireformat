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

type DeleteAclsRequest0 struct {
    Filters []*DeleteAclsRequest0_Filters
}

func (that *DeleteAclsRequest0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Filters)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Filters[i].Encode(enc)
        }
    }

    return nil
}

func (that *DeleteAclsRequest0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Filters = nil
        } else {
            buf := make([]*DeleteAclsRequest0_Filters, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DeleteAclsRequest0_Filters)
                item.Decode(dec)
                buf[i] = item
            }
            that.Filters = buf
        }
    }
    return nil
}

type DeleteAclsRequest0_Filters struct {
    ResourceType int8
    ResourceName string
    Principal string
    Host string
    Operation int8
    PermissionType int8
}

func (that *DeleteAclsRequest0_Filters) Encode(enc *Encoder) error {
    enc.WriteInt8(that.ResourceType)
    enc.WriteString(that.ResourceName)
    enc.WriteString(that.Principal)
    enc.WriteString(that.Host)
    enc.WriteInt8(that.Operation)
    enc.WriteInt8(that.PermissionType)
    return nil
}

func (that *DeleteAclsRequest0_Filters) Decode(dec *Decoder) error {
    that.ResourceType = dec.ReadInt8()
    that.ResourceName = dec.ReadString()
    that.Principal = dec.ReadString()
    that.Host = dec.ReadString()
    that.Operation = dec.ReadInt8()
    that.PermissionType = dec.ReadInt8()
    return nil
}

