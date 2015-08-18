package echo

/*
 * echo.go
 * Command to echo things, kinda
 * By J. Stuart McMurray
 * Created 20150817
 * Last Modified 20150817
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
	"log"
	"strings"
)

/* Default response, when nothing in ans fits */
var def = "Did you hear something?  No?  Well, nice try anyways.\n"

/* Arguments to echo and answers */
var ans [][2]string = [][2]string{
	{"-e '\\x67\\x61\\x79\\x66\\x67\\x74'", "gayfgt\n"},
	{"-e '\\\\x67\\\\x61\\\\x79\\\\x66\\\\x67\\\\x74'", "gayfgt\n"},
}

/* Things to echo and their responses, populated by initAr on the first call
to Echo. */
var ar [][2]string

func initAr() {
	/* Make enough space */
	ar = make([][2]string, len(ans))
	/* Populate ar */
	for i, v := range [][2]string{} {
		/* Normalize string we're looking for */
		ar[i] = [2]string{
			strings.Join(strings.Fields(v[0]), " "),
			v[1],
		}
	}
}

/* Echo a response, based on a */
func Echo(c string, a []string, out io.Writer) {
	res := def
	/* Make sure we have an array of matches */
	if nil == ar {
		initAr()
	}
	/* Command, proccessed as in ar */
	cmd := strings.Join(append([]string{c}, a...), " ")
	/* If we have a match, send that */
	for i, v := range ar {
		if cmd == v[0] {
			log.Printf("Echoing answer %v", i)
			res = v[1]
			break
		}
	}
	/* If no match, a sensible default */
	fmt.Fprintf(out, "%v", res)
}
