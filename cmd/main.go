/*
Copyright 2025 Alexander <alexander.antoschuk.dev@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <www.gnu.org>.
*/
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aantoschuk/feed/internal/domain"
	"github.com/aantoschuk/feed/internal/engine"
	"github.com/aantoschuk/feed/internal/extractors"
)

func main() {

	ign := &extractors.IGNExtractor{
		URL:      "https://www.ign.com/news/",
		WaitTime: 1 * time.Second,
	}

	gamespot := &extractors.GamespotExtractor{
		URL:      "https://www.gamespot.com/news",
		WaitTime: 1 * time.Second,
	}

	params := engine.CreateEngineParams{
		Extractors: []domain.Extractor{ign, gamespot},
	}

	en := engine.CreateEngine(params)

	articles, err := en.Extract()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, a := range articles {
		if a.Title == "" {
			continue
		}
		fmt.Println(a)
		fmt.Println()
	}
}

// TODO: debug more not working properly.
// Creates process but not opening the actually window
