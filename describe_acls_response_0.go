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

type DescribeAclsResponse0_Acls struct {
    Principal string
    Host string
    Operation int8
    PermissionType int8
}

func (that *DescribeAclsResponse0_Acls) Encode(enc *Encoder) error {
    enc.WriteString(that.Principal)
    enc.WriteString(that.Host)
    enc.WriteInt8(that.Operation)
    enc.WriteInt8(that.PermissionType)
    return nil
}

func (that *DescribeAclsResponse0_Acls) Decode(dec *Decoder) error {
    that.Principal = dec.ReadString()
    that.Host = dec.ReadString()
    that.Operation = dec.ReadInt8()
    that.PermissionType = dec.ReadInt8()
    return nil
}

type DescribeAclsResponse0_Resources struct {
    ResourceType int8
    ResourceName string
    Acls []*DescribeAclsResponse0_Acls
}

func (that *DescribeAclsResponse0_Resources) Encode(enc *Encoder) error {
    enc.WriteInt8(that.ResourceType)
    enc.WriteString(that.ResourceName)
    {
        arrayLength := len(that.Acls)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Acls[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeAclsResponse0_Resources) Decode(dec *Decoder) error {
    that.ResourceType = dec.ReadInt8()
    that.ResourceName = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Acls = nil
        } else {
            buf := make([]*DescribeAclsResponse0_Acls, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeAclsResponse0_Acls)
                item.Decode(dec)
                buf[i] = item
            }
            that.Acls = buf
        }
    }
    return nil
}

type DescribeAclsResponse0 struct {
    ThrottleTimeMs int32
    ErrorCode int16
    ErrorMessage string
    Resources []*DescribeAclsResponse0_Resources
}

func (that *DescribeAclsResponse0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.ErrorMessage)
    {
        arrayLength := len(that.Resources)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Resources[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeAclsResponse0) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    that.ErrorMessage = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Resources = nil
        } else {
            buf := make([]*DescribeAclsResponse0_Resources, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeAclsResponse0_Resources)
                item.Decode(dec)
                buf[i] = item
            }
            that.Resources = buf
        }
    }
    return nil
}

