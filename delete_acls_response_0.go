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

type DeleteAclsResponse0_MatchingAcls struct {
    ErrorCode int16
    ErrorMessage string
    ResourceType int8
    ResourceName string
    Principal string
    Host string
    Operation int8
    PermissionType int8
}

func (that *DeleteAclsResponse0_MatchingAcls) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.ErrorMessage)
    enc.WriteInt8(that.ResourceType)
    enc.WriteString(that.ResourceName)
    enc.WriteString(that.Principal)
    enc.WriteString(that.Host)
    enc.WriteInt8(that.Operation)
    enc.WriteInt8(that.PermissionType)
    return nil
}

func (that *DeleteAclsResponse0_MatchingAcls) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.ErrorMessage = dec.ReadString()
    that.ResourceType = dec.ReadInt8()
    that.ResourceName = dec.ReadString()
    that.Principal = dec.ReadString()
    that.Host = dec.ReadString()
    that.Operation = dec.ReadInt8()
    that.PermissionType = dec.ReadInt8()
    return nil
}

type DeleteAclsResponse0_FilterResponses struct {
    ErrorCode int16
    ErrorMessage string
    MatchingAcls []*DeleteAclsResponse0_MatchingAcls
}

func (that *DeleteAclsResponse0_FilterResponses) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.ErrorMessage)
    {
        arrayLength := len(that.MatchingAcls)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.MatchingAcls[i].Encode(enc)
        }
    }

    return nil
}

func (that *DeleteAclsResponse0_FilterResponses) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.ErrorMessage = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.MatchingAcls = nil
        } else {
            buf := make([]*DeleteAclsResponse0_MatchingAcls, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DeleteAclsResponse0_MatchingAcls)
                item.Decode(dec)
                buf[i] = item
            }
            that.MatchingAcls = buf
        }
    }
    return nil
}

type DeleteAclsResponse0 struct {
    ThrottleTimeMs int32
    FilterResponses []*DeleteAclsResponse0_FilterResponses
}

func (that *DeleteAclsResponse0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    {
        arrayLength := len(that.FilterResponses)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.FilterResponses[i].Encode(enc)
        }
    }

    return nil
}

func (that *DeleteAclsResponse0) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.FilterResponses = nil
        } else {
            buf := make([]*DeleteAclsResponse0_FilterResponses, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DeleteAclsResponse0_FilterResponses)
                item.Decode(dec)
                buf[i] = item
            }
            that.FilterResponses = buf
        }
    }
    return nil
}

