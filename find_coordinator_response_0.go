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

type FindCoordinatorResponse0_Coordinator struct {
    NodeId int32
    Host string
    Port int32
}

func (that *FindCoordinatorResponse0_Coordinator) Encode(enc *Encoder) error {
    enc.WriteInt32(that.NodeId)
    enc.WriteString(that.Host)
    enc.WriteInt32(that.Port)
    return nil
}

func (that *FindCoordinatorResponse0_Coordinator) Decode(dec *Decoder) error {
    that.NodeId = dec.ReadInt32()
    that.Host = dec.ReadString()
    that.Port = dec.ReadInt32()
    return nil
}

type FindCoordinatorResponse0 struct {
    ErrorCode int16
    Coordinator *FindCoordinatorResponse0_Coordinator
}

func (that *FindCoordinatorResponse0) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    that.Coordinator.Encode(enc)

    return nil
}

func (that *FindCoordinatorResponse0) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.Coordinator = new(FindCoordinatorResponse0_Coordinator)
    that.Coordinator.Decode(dec)
    return nil
}

