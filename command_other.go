//go:build !linux

/*
   Copyright The containerd Authors.

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

package runc

import (
	"context"
	"os"
	"os/exec"
)

func (r *Runc) command(context context.Context, args ...string) *exec.Cmd {
	var cmd *exec.Cmd
	command := r.Command
	if r.Debug {
		command = "strace"
		args = append([]string{"-f", "-o", "/tmp/runc.strace", "runc"}, append(r.args(), args...)...)
		cmd = exec.CommandContext(context, command, args...)
	} else if command == "" {
		command = DefaultCommand
		cmd = exec.CommandContext(context, command, append(r.args(), args...)...)
	}
	cmd.Env = os.Environ()
	return cmd
}
