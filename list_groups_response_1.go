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

type ListGroupsResponse1 struct {
    ThrottleTimeMs int32
    ErrorCode int16
    Groups []*ListGroupsResponse1_Groups
}

func (that *ListGroupsResponse1) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    enc.WriteInt16(that.ErrorCode)
    {
        arrayLength := len(that.Groups)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Groups[i].Encode(enc)
        }
    }

    return nil
}

func (that *ListGroupsResponse1) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    that.ErrorCode = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Groups = nil
        } else {
            buf := make([]*ListGroupsResponse1_Groups, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ListGroupsResponse1_Groups)
                item.Decode(dec)
                buf[i] = item
            }
            that.Groups = buf
        }
    }
    return nil
}

type ListGroupsResponse1_Groups struct {
    GroupId string
    ProtocolType string
}

func (that *ListGroupsResponse1_Groups) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    enc.WriteString(that.ProtocolType)
    return nil
}

func (that *ListGroupsResponse1_Groups) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    that.ProtocolType = dec.ReadString()
    return nil
}

