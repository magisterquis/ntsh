/* ntsh provides a small shell that gives the users a "Nice Try." */
package main

/*
 * ntsh.go
 * Next-to-useless shell, for a honeypot
 * by J. Stuart McMurray and Josiah Hamilton
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
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var prompt = flag.String(
	"p",
	"[root@localhost:~]# ",
	"Prompt, may be changed later",
)

func main() {
	var (
		logFile = flag.String(
			"l",
			"/tmp/ntsh/ntsh.log",
			"Logfile",
		)
		motdFile = flag.String(
			"motd",
			"/tmp/ntsh/motd",
			"MOTD to print on connect",
		)
		caddr = flag.String(
			"c",
			"",
			"Client's address",
		)
	)

	/* Figure out who's connected */
	if "" == *caddr {
		/* Try to get it from the environment */
		for _, v := range []string{
			"SOCAT_PEERADDR",
			"NCAT_REMOTE_ADDR",
			"SSH_CLIENT",
			"SSH_CONNECTION",
		} {
			*caddr = os.Getenv(v)
			if "" != *caddr {
				break
			}
		}
		/* Failing that, use the pid */
		if "" == *caddr {
			*caddr = fmt.Sprintf("%v", os.Getpid())
		}
	}

	/* Log file */
	lf, err := os.OpenFile(
		*logFile,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if nil != err {
		log.Printf("Sorry, the sysadmin borked the install: %v", err)
		return
	}
	log.SetOutput(lf)
	log.SetFlags(log.Lmicroseconds | log.LstdFlags)

	log.Printf("%v!: Start", *caddr)
	defer log.Printf("%v!: Stop", *caddr)

	/* Handle motd */
	if "" != *motdFile {
		motd, err := ioutil.ReadFile(*motdFile)
		if nil != err {
			log.Printf("%v!: Unale to get MOTD: %v", *caddr, err)
		}
		if nil != motd && 0 != len(motd) {
			fmt.Printf("%v", string(motd))
		}
	}

	/* Take lines of input, handle them */
	fmt.Printf("%v", *prompt)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := run(scanner.Text(), *caddr); nil != err {
			log.Printf("%v!: Run error: %v", *caddr, err)
			break
		}
		fmt.Printf("%v", *prompt)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%v!: Input error: %v", *caddr, err)
	}
}
