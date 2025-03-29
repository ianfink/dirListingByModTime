/*
 * Copyright (C) 2025 Ian M. Fink.  All rights reserved.
 *
 * This program is free software:  you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
 * more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program.  If not, please see: https://www.gnu.org/licenses.
 */

package main

/*
 * Imports
 */

import (
	"fmt"
	"os"
	"path"
	"log"
	"slices"
)

/**********************************************************************/

func compareModTimesASC(a, b os.FileInfo) int {
	// For ascending modification time order
	return int(a.ModTime().Unix() - b.ModTime().Unix())
} /* compareModTimes */

/**********************************************************************/

func compareModTimesDSC(a, b os.FileInfo) int {
	// For descending modification time order
	return int(b.ModTime().Unix() - a.ModTime().Unix())
} /* compareModTimes */

/**********************************************************************/

func main() {
	var (
		theDirEntries		[]os.DirEntry
		theDirEntry			os.DirEntry
		theOsFileInfos		[]os.FileInfo
		theOsFileInfo		os.FileInfo
		targetDirectory		string
		err					error
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage:  %s <target directory>\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	targetDirectory = os.Args[1]
	theDirEntries, err = os.ReadDir(targetDirectory)

	if err != nil {
		log.Fatal(err)
	}

	theOsFileInfos = make([]os.FileInfo, 0, len(theDirEntries))

	for _, theDirEntry = range theDirEntries {
		theOsFileInfo, err = theDirEntry.Info()
		if err != nil {
			continue
		}

		theOsFileInfos = append(theOsFileInfos, theOsFileInfo)
	}

	// Sort the theFileInfos slice by file modification times
	slices.SortFunc(theOsFileInfos, compareModTimesASC)

	for _, theOsFileInfo = range theOsFileInfos {
		fmt.Println(theOsFileInfo.Name())
	}

} /* main */

/**********************************************************************/

/*
 * End of file:	main.go
 */

