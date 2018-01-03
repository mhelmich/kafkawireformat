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

type AlterReplicaLogDirsRequest0_Topics struct {
    Topic string
    Partitions []int32
}

func (that *AlterReplicaLogDirsRequest0_Topics) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt32Array(that.Partitions)
    return nil
}

func (that *AlterReplicaLogDirsRequest0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.Partitions = dec.ReadInt32Array()
    return nil
}

type AlterReplicaLogDirsRequest0 struct {
    LogDirs []*AlterReplicaLogDirsRequest0_LogDirs
}

func (that *AlterReplicaLogDirsRequest0) Encode(enc *Encoder) error {
    {
        arrayLength := len(that.LogDirs)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.LogDirs[i].Encode(enc)
        }
    }

    return nil
}

func (that *AlterReplicaLogDirsRequest0) Decode(dec *Decoder) error {
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.LogDirs = nil
        } else {
            buf := make([]*AlterReplicaLogDirsRequest0_LogDirs, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AlterReplicaLogDirsRequest0_LogDirs)
                item.Decode(dec)
                buf[i] = item
            }
            that.LogDirs = buf
        }
    }
    return nil
}

type AlterReplicaLogDirsRequest0_LogDirs struct {
    LogDir string
    Topics []*AlterReplicaLogDirsRequest0_Topics
}

func (that *AlterReplicaLogDirsRequest0_LogDirs) Encode(enc *Encoder) error {
    enc.WriteString(that.LogDir)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *AlterReplicaLogDirsRequest0_LogDirs) Decode(dec *Decoder) error {
    that.LogDir = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*AlterReplicaLogDirsRequest0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(AlterReplicaLogDirsRequest0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

