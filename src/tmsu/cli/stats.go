// Copyright 2011-2015 Paul Ruane.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cli

import (
	"fmt"
	"math"
	"tmsu/storage"
)

var StatsCommand = Command{
	Name:        "stats",
	Synopsis:    "Show database statistics",
	Usages:      []string{"tmsu stats"},
	Description: "Shows the database statistics.",
	Options:     Options{Option{"--usage", "-u", "show tag usage breakdown", false, ""}},
	Exec:        statsExec,
}

func statsExec(store *storage.Storage, options Options, args []string) error {
	usage := options.HasOption("--usage")

	tx, err := store.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	tagCount, err := store.TagCount(tx)
	if err != nil {
		return fmt.Errorf("could not retrieve tag count: %v", err)
	}

	valueCount, err := store.ValueCount(tx)
	if err != nil {
		return fmt.Errorf("could not retrieve value count: %v", err)
	}

	fileCount, err := store.FileCount(tx)
	if err != nil {
		return fmt.Errorf("could not retrieve file count: %v", err)
	}

	fileTagCount, err := store.FileTagCount(tx)
	if err != nil {
		return fmt.Errorf("could not retrieve taggings count: %v", err)
	}

	var averageTagsPerFile float32
	if fileCount > 0 {
		averageTagsPerFile = float32(fileTagCount) / float32(fileCount)
	}

	var averageFilesPerTag float32
	if tagCount > 0 {
		averageFilesPerTag = float32(fileTagCount) / float32(tagCount)
	}

	fmt.Println("DATABASE")
	fmt.Println()
	fmt.Printf("  Path: %v\n", store.DbPath)
	fmt.Printf("  Root: %v\n", store.RootPath)
	fmt.Println()
	fmt.Println("COUNTS")
	fmt.Println()
	fmt.Printf("  Tags:     %v\n", tagCount)
	fmt.Printf("  Values:   %v\n", valueCount)
	fmt.Printf("  Files:    %v\n", fileCount)
	fmt.Printf("  Taggings: %v\n", fileTagCount)
	fmt.Println()

	fmt.Println("AVERAGES")
	fmt.Println()
	fmt.Printf("  Tags per file: %1.2f\n", averageTagsPerFile)
	fmt.Printf("  Files per tag: %1.2f\n", averageFilesPerTag)
	fmt.Println()

	if usage {
		tagUsages, err := store.TagUsage(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve tag usage: %v", err)
		}

		fmt.Println("TAG USAGE")
		fmt.Println()
		maxLength := 0
		maxCountWidth := 0
		for _, tagUsage := range tagUsages {
			countWidth := int(math.Log(float64(tagUsage.FileCount)))
			if countWidth > maxCountWidth {
				maxCountWidth = countWidth
			}
			if len(tagUsage.Name) > maxLength {
				maxLength = len(tagUsage.Name)
			}
		}
		for _, tagUsage := range tagUsages {
			fmt.Printf("  %*s %*v\n", -maxLength, tagUsage.Name, maxCountWidth, tagUsage.FileCount)
		}
		fmt.Println()
	}

	return nil
}
