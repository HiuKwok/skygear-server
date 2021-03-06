// Copyright 2015-present Oursky Ltd.
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

package event

import (
	"github.com/skygeario/skygear-server/pkg/server/logging"
	"github.com/skygeario/skygear-server/pkg/server/plugin"
)

type defaultEventSender struct {
	pluginContext *plugin.Context
}

// NewSender creates an event sender
func NewSender(ctx *plugin.Context) Sender {
	return &defaultEventSender{
		pluginContext: ctx,
	}
}

func (s defaultEventSender) Send(name string, data []byte, async bool) {
	logger := logging.LoggerEntry("plugin")
	logger.
		WithField("name", name).
		WithField("date", string(data)).
		Info("Sending plugin event")

	s.pluginContext.SendEvent(name, data, async)
}
