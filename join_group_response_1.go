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

type JoinGroupResponse1_Members struct {
    MemberId string
    MemberMetadata []byte
}

func (that *JoinGroupResponse1_Members) Encode(enc *Encoder) error {
    enc.WriteString(that.MemberId)
    enc.WriteByteArray(that.MemberMetadata)
    return nil
}

func (that *JoinGroupResponse1_Members) Decode(dec *Decoder) error {
    that.MemberId = dec.ReadString()
    that.MemberMetadata = dec.ReadByteArray()
    return nil
}

type JoinGroupResponse1 struct {
    ErrorCode int16
    GenerationId int32
    GroupProtocol string
    LeaderId string
    MemberId string
    Members []*JoinGroupResponse1_Members
}

func (that *JoinGroupResponse1) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteInt32(that.GenerationId)
    enc.WriteString(that.GroupProtocol)
    enc.WriteString(that.LeaderId)
    enc.WriteString(that.MemberId)
    {
        arrayLength := len(that.Members)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Members[i].Encode(enc)
        }
    }

    return nil
}

func (that *JoinGroupResponse1) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.GenerationId = dec.ReadInt32()
    that.GroupProtocol = dec.ReadString()
    that.LeaderId = dec.ReadString()
    that.MemberId = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Members = nil
        } else {
            buf := make([]*JoinGroupResponse1_Members, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(JoinGroupResponse1_Members)
                item.Decode(dec)
                buf[i] = item
            }
            that.Members = buf
        }
    }
    return nil
}

