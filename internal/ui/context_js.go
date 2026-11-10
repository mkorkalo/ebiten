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

package ui

// maxSkipCount is the number of successive frames that don't modify the offscreen, after which presenting the
// offscreen is skipped.
//
// On browsers, a canvas keeps showing what was presented last until something is drawn onto it again, and a browser
// composites a canvas only when it is drawn. Presenting the unchanged offscreen again only redraws the same pixels
// and makes the browser composite them again, so presenting is skipped from the first frame that doesn't modify the
// offscreen. This saves GPU power when a game draws less often than the display refreshes, e.g. a game that draws
// at 30 FPS on a 120 Hz display.
//
// A resized canvas loses its content, but resizing forces a draw, which is never skipped.
const maxSkipCount = 1
