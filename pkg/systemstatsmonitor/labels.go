/*
Copyright 2020 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package systemstatsmonitor

// deviceNameLabel labels the monitored disk device, e.g.: "sda", "sda1".
const deviceNameLabel = "device_name"

// directionLabel labels the direction of the disk operations, e.g.: "read", "write".
const directionLabel = "direction"

// stateLabel labels the state of disk/memory/cpu usage, e.g.: "free", "used".
const stateLabel = "state"

// fstypeLabel labels the fst type of the disk, e.g.: "ext4", "ext2", "vfat"
const fstypeLabel = "fstype"

// mountPointLabel labels the mountpoint of the monitored disk device
const mountOptionLabel = "mountoption"