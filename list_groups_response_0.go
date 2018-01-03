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

type ListGroupsResponse0 struct {
    ErrorCode int16
    Groups []*ListGroupsResponse0_Groups
}

func (that *ListGroupsResponse0) Encode(enc *Encoder) error {
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

func (that *ListGroupsResponse0) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Groups = nil
        } else {
            buf := make([]*ListGroupsResponse0_Groups, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ListGroupsResponse0_Groups)
                item.Decode(dec)
                buf[i] = item
            }
            that.Groups = buf
        }
    }
    return nil
}

type ListGroupsResponse0_Groups struct {
    GroupId string
    ProtocolType string
}

func (that *ListGroupsResponse0_Groups) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    enc.WriteString(that.ProtocolType)
    return nil
}

func (that *ListGroupsResponse0_Groups) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    that.ProtocolType = dec.ReadString()
    return nil
}

