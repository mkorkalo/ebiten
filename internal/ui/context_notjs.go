// Copyright 2026 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !js

package ui

// maxSkipCount is the number of successive frames that don't modify the offscreen, after which presenting the
// offscreen is skipped.
//
// A swap chain can have several framebuffers and the screen can show any of them, so the unchanged offscreen is
// presented for some more frames until every framebuffer has it (#3055).
const maxSkipCount = 4
