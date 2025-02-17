//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

//go:build linux && !android
// +build linux,!android

package serial

import (
	"golang.org/x/sys/unix"
)

func (p *Port) retrieveTermSettings() (s *settings, err error) {
	s = &settings{
		termios: new(unix.Termios),
	}

	if s.termios, err = unix.IoctlGetTermios(p.internal.handle, unix.TCGETS); err != nil {
		return nil, newPortOSError(err)
	}

	if s.termios.Cflag&unix.BOTHER == unix.BOTHER {
		if s.termios, err = unix.IoctlGetTermios(p.internal.handle, unix.TCGETS2); err != nil {
			return nil, newPortOSError(err)
		}
	}

	return s, nil
}

func (p *Port) applyTermSettings(s *settings) error {
	req := uint(unix.TCSETS)
	if s.termios.Cflag&unix.BOTHER == unix.BOTHER {
		req = unix.TCSETS2
	}

	if err := unix.IoctlSetTermios(p.internal.handle, req, s.termios); err != nil {
		return newPortOSError(err)
	}

	return nil
}
