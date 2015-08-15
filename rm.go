package main

/*
 * rm.go
 * Handle rm commands
 * By J. Stuart McMurray
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
	"fmt"
	"io"
	"strings"
)

func init() {
	Register(
		"rm",
		func(c string, a []string, o io.Writer) {
			fmt.Fprintf(
				o,
				"Do you really think you can %v?\n"+
					"Nice try, though.\n",
				strings.Join(
					append([]string{c}, a...),
					" ",
				),
			)
		},
	)
}
