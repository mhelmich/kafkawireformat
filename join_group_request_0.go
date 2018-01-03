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

type JoinGroupRequest0_GroupProtocols struct {
    ProtocolName string
    ProtocolMetadata []byte
}

func (that *JoinGroupRequest0_GroupProtocols) Encode(enc *Encoder) error {
    enc.WriteString(that.ProtocolName)
    enc.WriteByteArray(that.ProtocolMetadata)
    return nil
}

func (that *JoinGroupRequest0_GroupProtocols) Decode(dec *Decoder) error {
    that.ProtocolName = dec.ReadString()
    that.ProtocolMetadata = dec.ReadByteArray()
    return nil
}

type JoinGroupRequest0 struct {
    GroupId string
    SessionTimeout int32
    MemberId string
    ProtocolType string
    GroupProtocols []*JoinGroupRequest0_GroupProtocols
}

func (that *JoinGroupRequest0) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    enc.WriteInt32(that.SessionTimeout)
    enc.WriteString(that.MemberId)
    enc.WriteString(that.ProtocolType)
    {
        arrayLength := len(that.GroupProtocols)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.GroupProtocols[i].Encode(enc)
        }
    }

    return nil
}

func (that *JoinGroupRequest0) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    that.SessionTimeout = dec.ReadInt32()
    that.MemberId = dec.ReadString()
    that.ProtocolType = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.GroupProtocols = nil
        } else {
            buf := make([]*JoinGroupRequest0_GroupProtocols, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(JoinGroupRequest0_GroupProtocols)
                item.Decode(dec)
                buf[i] = item
            }
            that.GroupProtocols = buf
        }
    }
    return nil
}

