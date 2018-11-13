/*  This file is part of diff-matrix.
 *
 *  Copyright (C) 2018  Pablo M. Bermudo Garay
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var trees []string
	args := os.Args[1:]
	forest := make(map[string][]int)

	for _, dir := range args {
		file, err := os.Stat(dir)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else if !file.Mode().IsDir() {
			fmt.Printf("'%s' is not a directory\n", dir)
			os.Exit(1)
		}

		trees = append(trees, dir)
	}

	for i, tree := range trees {
		filepath.Walk(tree, func(path string, info os.FileInfo, err error) error {
			filename := info.Name()
			localpath := strings.Replace(path, tree, "", 1)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if info.IsDir() && string(filename[0]) == "." {
				return filepath.SkipDir
			}

			_, exists := forest[localpath]
			if !exists {
				forest[localpath] = make([]int, len(args))
			}
			forest[localpath][i] = 1

			return nil
		})
	}

	fmt.Println(trees)
	fmt.Println(forest)
}
