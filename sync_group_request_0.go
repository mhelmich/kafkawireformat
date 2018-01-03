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

type SyncGroupRequest0_GroupAssignment struct {
    MemberId string
    MemberAssignment []byte
}

func (that *SyncGroupRequest0_GroupAssignment) Encode(enc *Encoder) error {
    enc.WriteString(that.MemberId)
    enc.WriteByteArray(that.MemberAssignment)
    return nil
}

func (that *SyncGroupRequest0_GroupAssignment) Decode(dec *Decoder) error {
    that.MemberId = dec.ReadString()
    that.MemberAssignment = dec.ReadByteArray()
    return nil
}

type SyncGroupRequest0 struct {
    GroupId string
    GenerationId int32
    MemberId string
    GroupAssignment []*SyncGroupRequest0_GroupAssignment
}

func (that *SyncGroupRequest0) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    enc.WriteInt32(that.GenerationId)
    enc.WriteString(that.MemberId)
    {
        arrayLength := len(that.GroupAssignment)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.GroupAssignment[i].Encode(enc)
        }
    }

    return nil
}

func (that *SyncGroupRequest0) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    that.GenerationId = dec.ReadInt32()
    that.MemberId = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.GroupAssignment = nil
        } else {
            buf := make([]*SyncGroupRequest0_GroupAssignment, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(SyncGroupRequest0_GroupAssignment)
                item.Decode(dec)
                buf[i] = item
            }
            that.GroupAssignment = buf
        }
    }
    return nil
}

