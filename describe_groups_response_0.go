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

type DescribeGroupsResponse0_Groups struct {
    ErrorCode int16
    GroupId string
    State string
    ProtocolType string
    Protocol string
    Members []*DescribeGroupsResponse0_Members
}

func (that *DescribeGroupsResponse0_Groups) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.GroupId)
    enc.WriteString(that.State)
    enc.WriteString(that.ProtocolType)
    enc.WriteString(that.Protocol)
    {
        arrayLength := len(that.Members)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Members[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeGroupsResponse0_Groups) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.GroupId = dec.ReadString()
    that.State = dec.ReadString()
    that.ProtocolType = dec.ReadString()
    that.Protocol = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Members = nil
        } else {
            buf := make([]*DescribeGroupsResponse0_Members, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeGroupsResponse0_Members)
                item.Decode(dec)
                buf[i] = item
            }
            that.Members = buf
        }
    }
    return nil
}

type DescribeGroupsResponse0_Members struct {
    MemberId string
    ClientId string
    ClientHost string
    MemberMetadata []byte
    MemberAssignment []byte
}

func (that *DescribeGroupsResponse0_Members) Encode(enc *Encoder) error {
    enc.WriteString(that.MemberId)
    enc.WriteString(that.ClientId)
    enc.WriteString(that.ClientHost)
    enc.WriteByteArray(that.MemberMetadata)
    enc.WriteByteArray(that.MemberAssignment)
    return nil
}

func (that *DescribeGroupsResponse0_Members) Decode(dec *Decoder) error {
    that.MemberId = dec.ReadString()
    that.ClientId = dec.ReadString()
    that.ClientHost = dec.ReadString()
    that.MemberMetadata = dec.ReadByteArray()
    that.MemberAssignment = dec.ReadByteArray()
    return nil
}

type DescribeGroupsResponse0 struct {
    Groups []*DescribeGroupsResponse0_Groups
}

func (that *DescribeGroupsResponse0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.Groups)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Groups[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeGroupsResponse0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Groups = nil
        } else {
            buf := make([]*DescribeGroupsResponse0_Groups, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeGroupsResponse0_Groups)
                item.Decode(dec)
                buf[i] = item
            }
            that.Groups = buf
        }
    }
    return nil
}

