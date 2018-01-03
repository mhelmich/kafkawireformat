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

type RequestHeader struct {
    ApiKey int16
    ApiVersion int16
    CorrelationId int32
    ClientId string
}

func (that *RequestHeader) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ApiKey)
    enc.WriteInt16(that.ApiVersion)
    enc.WriteInt32(that.CorrelationId)
    enc.WriteString(that.ClientId)
    return nil
}

func (that *RequestHeader) Decode(dec *Decoder) error {
    that.ApiKey = dec.ReadInt16()
    that.ApiVersion = dec.ReadInt16()
    that.CorrelationId = dec.ReadInt32()
    that.ClientId = dec.ReadString()
    return nil
}

