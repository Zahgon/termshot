// Copyright © 2020 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ptexec

import (
	"io"
	"os"
	"os/exec"
)

// PseudoTerminal defines the setup for a command to be run in a pseudo
// terminal, e.g. terminal size, or output settings
type PseudoTerminal struct {
	name string
	args []string

	shell string

	cols   uint16
	rows   uint16
	resize bool

	stdout io.Writer
}

// New creates a new pseudo terminal builder
func New() *PseudoTerminal { _ = "STUB: not implemented"; return nil }

// Cols sets the width/columns for the pseudo terminal
func (c *PseudoTerminal) Cols(cols uint16) *PseudoTerminal { _ = "STUB: not implemented"; return nil }

// Rows sets the lines/rows for the pseudo terminal
func (c *PseudoTerminal) Rows(rows uint16) *PseudoTerminal { _ = "STUB: not implemented"; return nil }

// Stdout sets the writer to be used for the standard output
func (c *PseudoTerminal) Stdout(stdout io.Writer) *PseudoTerminal {
	_ = "STUB: not implemented"
	return nil
}

// Command sets the command and arguments to be used
func (c *PseudoTerminal) Command(name string, args ...string) *PseudoTerminal {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the provided command/script with the given arguments in a pseudo
// terminal (PTY) so that the behavior is the same if it would be executed
// in a terminal
func (c *PseudoTerminal) Run() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Convenience hack in case command contains a space, for example in case
// typical construct like "foo | grep" are used.

// Set RAW mode for Stdin

// And make sure to restore the original mode eventually

// collect all errors along the way

// #nosec G204 -- since this is exactly what we want, arbitrary commands

// Support terminal resizing

func (c *PseudoTerminal) pseudoTerminal(cmd *exec.Cmd) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Obtaining terminal size is prone to error in CI systems, e.g. in
// GitHub Action setup or similar, so only fail if CI is not set

// For CI systems, assume a reasonable default even if the terminal
// size cannot be obtained through ioctl

// Overwrite rows if fixed value is configured

// Overwrite columns if fixed value is configured

// With fixed rows/cols, terminal resizing support is not useful

func copy(dst io.Writer, src io.Reader) error { _ = "STUB: not implemented"; return nil }

//nolint:gocritic

// Workaround for issue https://github.com/creack/pty/issues/100
// where on Linux systems it can happen that the pseudo terminal
// process finishes while termshot is trying to read. Assuming
// that the content is already read, this error is treated the
// same as if it would be an EOF.

func isTerminal(f *os.File) bool { _ = "STUB: not implemented"; return false }

func isCI() bool { _ = "STUB: not implemented"; return false }
