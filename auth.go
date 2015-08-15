package main

import (
	"bufio"
	"fmt"
	"io"
)

/*
 * auth.go
 * Handle fake authentication
 * by J. Stuart McMurray
 * Created 20150815
 * Last Modified 20150815
 */

/*
 * ntsh -- The "Nice Try" shell
 * version 0.0.1, August 15, 2015
 *
 * Copyright (C) 2015 Stuart McMurray and Josiah Hamilton
 *
 * This software is provided 'as-is', without any express or implied
 * warranty.  In no event will the authors be held liable for any damages
 * arising from the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not
 *    claim that you wrote the original software. If you use this software
 *    in a product, an acknowledgment in the product documentation would be
 *    appreciated but is not required.
 * 2. Altered source versions must be plainly marked as such, and must not be
 *    misrepresented as being the original software.
 * 3. This notice may not be removed or altered from any source distribution.
 *
 * Stuart McMurray      Josiah Hamilton
 * kd5pbo@gmail.com     dev.x.josiah@mamber.net
 */

/* Auth does fake authentication, if user or pass are set, it'll only allow
that specific username or password.  Auth returns nil on successful auth. */
func Auth(uname, pass string, in bufio.Scanner) error {
	ue := authPrompt("Username", uname, in)
	pe := authPrompt("Password", pass, in)
	if ue != nil {
		return ue
	}
	return pe
}

/* Prompt prints p, gets input from in, and if it doesn't match w (if w isn't
empty), returns an error. */
func authPrompt(p, w string, in bufio.Scanner) error {
	fmt.Printf("%v: ", p)
	h, err := nextString(in)
	if nil != err {
		return err
	}
	if "" != w && h != w {
		return fmt.Errorf("Bad %v answer %v, expected %v", p, h, w)
	}
	return nil
}

/* Get the next string from a scanner, or an error */
func nextString(s bufio.Scanner) (string, error) {
	/* See if there's anything to scan */
	if !s.Scan() {
		/* If not, return why */
		e := s.Err()
		if nil == e {
			e = io.EOF
		}
		return "", e
	}
	/* Return the next string */
	return s.Text(), nil
}
