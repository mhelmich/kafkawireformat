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

type HeartbeatRequest0 struct {
    GroupId string
    GenerationId int32
    MemberId string
}

func (that *HeartbeatRequest0) Encode(enc *Encoder) error {
    enc.WriteString(that.GroupId)
    enc.WriteInt32(that.GenerationId)
    enc.WriteString(that.MemberId)
    return nil
}

func (that *HeartbeatRequest0) Decode(dec *Decoder) error {
    that.GroupId = dec.ReadString()
    that.GenerationId = dec.ReadInt32()
    that.MemberId = dec.ReadString()
    return nil
}

