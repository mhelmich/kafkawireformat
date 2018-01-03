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

type DeleteTopicsRequest0 struct {
    Topics []string
    Timeout int32
}

func (that *DeleteTopicsRequest0) Encode(enc *Encoder) error {
    enc.WriteStringArray(that.Topics)
    enc.WriteInt32(that.Timeout)
    return nil
}

func (that *DeleteTopicsRequest0) Decode(dec *Decoder) error {
    that.Topics = dec.ReadStringArray()
    that.Timeout = dec.ReadInt32()
    return nil
}

