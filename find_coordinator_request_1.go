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

type FindCoordinatorRequest1 struct {
    CoordinatorKey string
    CoordinatorType int8
}

func (that *FindCoordinatorRequest1) Encode(enc *Encoder) error {
    enc.WriteString(that.CoordinatorKey)
    enc.WriteInt8(that.CoordinatorType)
    return nil
}

func (that *FindCoordinatorRequest1) Decode(dec *Decoder) error {
    that.CoordinatorKey = dec.ReadString()
    that.CoordinatorType = dec.ReadInt8()
    return nil
}

