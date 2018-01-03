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

type SaslAuthenticateResponse0 struct {
    ErrorCode int16
    ErrorMessage string
    SaslAuthBytes []byte
}

func (that *SaslAuthenticateResponse0) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.ErrorMessage)
    enc.WriteByteArray(that.SaslAuthBytes)
    return nil
}

func (that *SaslAuthenticateResponse0) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.ErrorMessage = dec.ReadString()
    that.SaslAuthBytes = dec.ReadByteArray()
    return nil
}

