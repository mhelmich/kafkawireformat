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

type DescribeGroupsRequest1 struct {
    GroupIds []string
}

func (that *DescribeGroupsRequest1) Encode(enc *Encoder) error {
    enc.WriteStringArray(that.GroupIds)
    return nil
}

func (that *DescribeGroupsRequest1) Decode(dec *Decoder) error {
    that.GroupIds = dec.ReadStringArray()
    return nil
}

