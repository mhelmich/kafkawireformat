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

type SyncGroupResponse0 struct {
    ErrorCode int16
    MemberAssignment []byte
}

func (that *SyncGroupResponse0) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteByteArray(that.MemberAssignment)
    return nil
}

func (that *SyncGroupResponse0) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.MemberAssignment = dec.ReadByteArray()
    return nil
}

