# Gaming News Aggregator

A Go-based CLI application that scrapes and aggregates gaming news articles from multiple sources including IGN and GameSpot.
![Go Version](https://img.shields.io/badge/Go-1.25.1+-blue.svg)
![License](https://img.shields.io/badge/License-GPL%20v3-green.svg)

## About 

THIS IS NOT COMMERCIAL PROJECT. Everything for the learning and practicing as well, as self use.

Currently supported sources:

- IGN
- GameSpot

## Features

- Web scraping from multiple gaming news websites;
- Concurrent processing with configurable worker pools;
- Flexible logging;
- Extensible architecture for adding new extractors;
- Clean CLI output format;
- Robust error handling with custom error types;

## Usage

### Basic Usage

Just execute binary:

```bash
./game-news-aggregator
```
### Verbose mode 
Verbose mode is prints messages when while executing:

```bash
./game-news-aggregator -v
```
### Debug Mode
Debug mode is primarly for the troubleshooting, as it prints debug messages, 
and runs browser with headless mode off:

```bash
./game-news-aggregator -d
```

### Combined Mode

It is also possible to combine flags together, although, 
there is no need to do it with `-d` and `-v` flags, as `-d` cover everything.


## Sample Output:

```
title: New-Gen console planned release date leaked by AMD
url: https://www.ign.com/articles/new-gen-console-leaked

title: New-Gen console planned release date leaked by AMD
url: https://www.gamespot.com/articles/new-gen-console-leaked
```

## Development 
1. Create new extractors, which will capture data form the html document
2. Implement `domain.Extractor` interface
3. Add new extractor to the slice in `main.go` file

## How it works

- Engine creates a new browser process
- then orchestrates multiple extractors concurrently
- each extractor implements the `Extract(page *rod.Page) ([]Article, error)` method.
- Engine runs extractors and collects articles from every extractor
- Rod browser automation handles JavaScript-heavy content

## Project Structure

- `cmd/main,go` - Entry point.
- `internal/engine/` - Runner and orchestrator of different extractors with worker pool logic.
- `internal/extractors/` - Site-specific extractors.
- `internal/domain/` - Models and interfaces.
- `internal/apperr/` - Structured error handling
- `internal/app_logger/` - Configurable logging
- `internal/browser/`  - Row browser creation and other functions which help to work with pages.

## License
- [GPL](./LICENSE)

