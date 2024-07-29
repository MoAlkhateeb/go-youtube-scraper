# Go YouTube Search Scraper

This Go program scrapes YouTube search results and exports the data to a JSON file. It uses the Geziyor web crawling and scraping framework along with goquery for HTML parsing.

## Features

- Scrapes YouTube search results
- Extracts video title, URL, view count, upload date, channel name, and channel URL
- Exports data to a JSON file
- Configurable search query, output file, and concurrency

## Prerequisites

- Go 1.15 or higher
- GitHub.com account to download dependencies

## Installation

1. Clone the repository:
```bash
git clone https://github.com/MoAlkhateeb/go-youtube-scraper.git
cd go-youtube-scraper
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

Run the program with default settings:

```bash
go run main.go
```

Or use command-line flags to customize the behavior:

```bash
go run main.go -query "your search query" -output results.json -concurrency 2
```

Available flags:
- `-base-url`: Base URL for YouTube (default: "https://www.youtube.com")
- `-query`: Search query for YouTube (default: "geziyor golang")
- `-output`: Output JSON file name (default: "videos.json")
- `-concurrency`: Number of concurrent requests (default: 1)

## Output

The program generates a JSON file containing an array of video information objects. Each object includes:

- id: Index of the video in the search results
- title: Video title
- url: Full URL of the video
- views: Number of views
- relative_date: Relative upload date (e.g., "2 weeks ago")
- channel_name: Name of the channel
- channel_url: Full URL of the channel

## Limitations

- This scraper is for educational purposes only. Be sure to comply with YouTube's terms of service and robots.txt file.
- The program may break if YouTube changes its HTML structure.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).