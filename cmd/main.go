package main

import (
	"fmt"
	"os"

	"github.com/idkmaybedeveloper/rustore-api-go/rustore"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "info":
		cmdInfo(args)
	case "search":
		cmdSearch(args)
	case "suggest":
		cmdSuggest(args)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("usage: rustore-api <command> [args]")
	fmt.Println("commands:")
	fmt.Println("  info <package_name>   fetch app info")
	fmt.Println("  search <query>        search apps")
	fmt.Println("  suggest <query>       get search suggestions")
}

func cmdInfo(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: rustore-api info <package_name>")
		fmt.Println("example: rustore-api info ru.ozon.app.android")
		os.Exit(1)
	}

	packageName := args[0]
	fmt.Printf("fetching info for: %s\n\n", packageName)

	summary, err := rustore.GetAppSummary(packageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("name: %s\n", summary.AppName)
	fmt.Printf("package: %s\n", summary.PackageName)
	fmt.Printf("appId: %d\n", summary.AppId)
	fmt.Printf("company: %s\n", summary.CompanyName)
	fmt.Printf("version: %s (%d)\n", summary.Version, summary.VersionCode)
	fmt.Printf("size: %s\n", summary.FileSize)
	fmt.Printf("downloads: %s\n", summary.Downloads)
	fmt.Printf("rating: %.1f (%d votes)\n", summary.Rating, summary.RatingVotes)
	fmt.Printf("icon: %s\n", summary.IconUrl)

	if summary.DownloadUrl != nil {
		fmt.Printf("\ndownload url: %s\n", *summary.DownloadUrl)
		if summary.DownloadSize != nil {
			fmt.Printf("download size: %s\n", rustore.FormatFileSize(*summary.DownloadSize))
		}
	} else {
		fmt.Println("\ndownload url: not available")
	}
}

func cmdSearch(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: rustore-api search <query>")
		fmt.Println("example: rustore-api search telegram")
		os.Exit(1)
	}

	query := args[0]
	fmt.Printf("searching for: %s\n\n", query)

	results, err := rustore.SearchApps(query, 0, 20)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("found %d apps (page %d/%d)\n\n", results.TotalElements, results.PageNumber+1, results.TotalPages)

	for _, app := range results.Content {
		fmt.Printf("[%d] %s\n", app.AppId, app.AppName)
		fmt.Printf("package: %s\n", app.PackageName)
		fmt.Printf("rating: %.1f (%d votes)\n", app.AverageUserRating, app.TotalRatings)
		fmt.Printf("downloads: %s\n", app.RoundedDownloadsText)
		fmt.Printf("type: %s\n", app.AppType)
		fmt.Println()
	}
}

func cmdSuggest(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: rustore-api suggest <query>")
		fmt.Println("example: rustore-api suggest yande")
		os.Exit(1)
	}

	query := args[0]
	results, err := rustore.GetSearchSuggestions(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if len(results.Suggests) == 0 {
		fmt.Println("no suggestions found")
		return
	}

	for _, suggest := range results.Suggests {
		fmt.Println(suggest.Text)
	}
}
