package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Config defines scan parameters
const (
	UserAgent = "Mozilla/5.0 (Compatible; GoSecurityScanner/1.0)"
	Timeout   = 10 * time.Second
)

type ScanResult struct {
	URL           string
	WPVersion     string
	WooVersion    string
	ExposedUsers  []string
	OpenEndpoints []string
	Plugins       []string
}

func main() {
	// Example Targets
	targets := []string{
		"https://yourtargetsite.com", // Replace with actual targets you own/have permission to test
	}

	var wg sync.WaitGroup

	fmt.Println("Starting Security Assessment...")
	fmt.Println("-------------------------------")

	for _, target := range targets {
		wg.Add(1)
		go func(site string) {
			defer wg.Done()
			assessSite(site)
		}(target)
	}

	wg.Wait()
	fmt.Println("Scan Complete.")
}

func assessSite(targetURL string) {
	client := &http.Client{
		Timeout: Timeout,
		// Ignore SSL errors for broader scanning capability
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	result := ScanResult{URL: targetURL}

	// 1. Fetch Homepage for Passive Analysis
	req, _ := http.NewRequest("GET", targetURL, nil)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[!] Error scanning %s: %v\n", targetURL, err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// 2. Passive Version Detection (Meta Generator & Scripts)
	result.WPVersion = detectWPVersion(doc)
	result.WooVersion = detectWooVersion(doc)
	result.Plugins = detectPlugins(doc)

	// 3. Active Endpoint Checks
	result.OpenEndpoints = checkEndpoints(client, targetURL)

	// 4. User Enumeration (REST API)
	result.ExposedUsers = enumerateUsers(client, targetURL)

	printReport(result)
}

// detectWPVersion looks for generator tags or common link versioning
func detectWPVersion(doc *goquery.Document) string {
	version := "Unknown"

	// Check meta tag
	doc.Find("meta[name='generator']").Each(func(i int, s *goquery.Selection) {
		content, _ := s.Attr("content")
		if strings.Contains(strings.ToLower(content), "wordpress") {
			version = content
		}
	})

	// Fallback: Check common CSS/JS query params (e.g., ?ver=6.1.1)
	if version == "Unknown" {
		doc.Find("link[href*='wp-includes']").First().Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			if strings.Contains(href, "?ver=") {
				parts := strings.Split(href, "?ver=")
				if len(parts) > 1 {
					version = "Inferred: " + parts[1]
				}
			}
		})
	}
	return version
}

// detectWooVersion looks for WooCommerce specific version indicators
func detectWooVersion(doc *goquery.Document) string {
	version := "Not Detected"

	// WooCommerce often puts version in a specific meta tag
	doc.Find("meta[name='generator']").Each(func(i int, s *goquery.Selection) {
		content, _ := s.Attr("content")
		if strings.Contains(strings.ToLower(content), "woocommerce") {
			version = content
		}
	})

	// Fallback: Scripts
	if version == "Not Detected" {
		doc.Find("script[src*='woocommerce']").First().Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr("src")
			if strings.Contains(src, "?ver=") {
				parts := strings.Split(src, "?ver=")
				if len(parts) > 1 {
					version = "Inferred: " + parts[1]
				}
			}
		})
	}
	return version
}

// detectPlugins scrapes asset URLs to find installed plugins
func detectPlugins(doc *goquery.Document) []string {
	pluginSet := make(map[string]bool)
	pluginRegex := regexp.MustCompile(`/wp-content/plugins/([^/]+)/`)

	doc.Find("link, script, img").Each(func(i int, s *goquery.Selection) {
		var urlStr string
		if href, exists := s.Attr("href"); exists {
			urlStr = href
		} else if src, exists := s.Attr("src"); exists {
			urlStr = src
		}

		matches := pluginRegex.FindStringSubmatch(urlStr)
		if len(matches) > 1 {
			pluginSet[matches[1]] = true
		}
	})

	var plugins []string
	for p := range pluginSet {
		plugins = append(plugins, p)
	}
	return plugins
}

// checkEndpoints probes for sensitive files
func checkEndpoints(client *http.Client, baseURL string) []string {
	endpoints := []string{"xmlrpc.php", "wp-json/", "readme.html", "license.txt"}
	var found []string

	for _, ep := range endpoints {
		fullURL := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), ep)
		req, _ := http.NewRequest("HEAD", fullURL, nil)
		req.Header.Set("User-Agent", UserAgent)

		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == 200 {
			found = append(found, ep)
		}
	}
	return found
}

// enumerateUsers attempts to list users via the WP REST API
func enumerateUsers(client *http.Client, baseURL string) []string {
	// Endpoint: /wp-json/wp/v2/users
	apiURL := fmt.Sprintf("%s/wp-json/wp/v2/users", strings.TrimRight(baseURL, "/"))
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil
	}
	defer resp.Body.Close()

	// Simple check: does the response contain "slug":"admin" or similar?
	// In a real tool, you would decode the JSON struct.
	// This is a quick heuristic scan.
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	text := doc.Text()

	var users []string
	if strings.Contains(text, "\"slug\":\"") {
		// Very basic regex to grab slugs from JSON (Use encoding/json for production)
		re := regexp.MustCompile(`"slug":"([^"]+)"`)
		matches := re.FindAllStringSubmatch(text, 5) // limit to 5
		for _, m := range matches {
			if len(m) > 1 {
				users = append(users, m[1])
			}
		}
	}
	return users
}

func printReport(r ScanResult) {
	fmt.Printf("\n[+] Report for: %s\n", r.URL)
	fmt.Printf("    WP Version:  %s\n", r.WPVersion)
	fmt.Printf("    Woo Version: %s\n", r.WooVersion)

	if len(r.ExposedUsers) > 0 {
		fmt.Printf("    [!] Exposed Users: %v (REST API Open)\n", r.ExposedUsers)
	} else {
		fmt.Println("    [+] User Enumeration: Blocked/Secure")
	}

	if len(r.OpenEndpoints) > 0 {
		fmt.Printf("    [!] Sensitive Files: %v\n", r.OpenEndpoints)
	}

	if len(r.Plugins) > 0 {
		fmt.Printf("    [*] Detected Plugins (%d): %v\n", len(r.Plugins), r.Plugins[:min(5, len(r.Plugins))]) // Show top 5
	}
	fmt.Println("------------------------------------------------")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
