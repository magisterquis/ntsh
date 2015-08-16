package main

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

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

/* Auth does fake authentication, if user or pass are set, it'll only allow
that specific username or password.  Auth returns nil on successful auth. */
func Auth(in bufio.Scanner) (u, p string, err error) {
	u, err = NextString("Username", in)
	if err != nil {
		return
	}
	p, err = NextString("Password", in)
	return
}

/* Get the next string from a scanner after printing a prompt, or an error */
func NextString(p string, s bufio.Scanner) (string, error) {
	fmt.Printf("%v: ", p)
	/* See if there's anything to scan */
	if !s.Scan() {
		/* If not, return why */
		e := s.Err()
		if nil == e {
			e = io.EOF
		}
		return "", e
	}
	/* Return the last printable chunk */
	f := strings.FieldsFunc(
		s.Text(),
		func(r rune) bool { return !unicode.IsGraphic(r) },
	)
	if 0 == len(f) {
		return "", nil
	}
	return f[len(f)-1], nil
}
