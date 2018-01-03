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

type ApiVersionsResponse1_ApiVersions struct {
    ApiKey int16
    MinVersion int16
    MaxVersion int16
}

func (that *ApiVersionsResponse1_ApiVersions) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ApiKey)
    enc.WriteInt16(that.MinVersion)
    enc.WriteInt16(that.MaxVersion)
    return nil
}

func (that *ApiVersionsResponse1_ApiVersions) Decode(dec *Decoder) error {
    that.ApiKey = dec.ReadInt16()
    that.MinVersion = dec.ReadInt16()
    that.MaxVersion = dec.ReadInt16()
    return nil
}

type ApiVersionsResponse1 struct {
    ErrorCode int16
    ApiVersions []*ApiVersionsResponse1_ApiVersions
    ThrottleTimeMs int32
}

func (that *ApiVersionsResponse1) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    {
        arrayLength := len(that.ApiVersions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.ApiVersions[i].Encode(enc)
        }
    }

    enc.WriteInt32(that.ThrottleTimeMs)
    return nil
}

func (that *ApiVersionsResponse1) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.ApiVersions = nil
        } else {
            buf := make([]*ApiVersionsResponse1_ApiVersions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ApiVersionsResponse1_ApiVersions)
                item.Decode(dec)
                buf[i] = item
            }
            that.ApiVersions = buf
        }
    }
    that.ThrottleTimeMs = dec.ReadInt32()
    return nil
}

