package shorty

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
    "strings"
)

// Config the plugin configuration.docker-compose up

type Config struct {
    Links map[string]string `json:"links,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
    return &Config{
        Links: make(map[string]string),
    }
}

// Shorty a Shorty plugin.
type Shorty struct {
    next   http.Handler
    links  map[string]string
    name   string
}

// New creates a new Shorty plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    if len(config.Links) == 0 {
        return nil, fmt.Errorf("links cannot be empty")
    }

    // Check for duplicate links
    seen := make(map[string]bool)
    for key := range config.Links {
        if _, found := seen[key]; found {
            // Duplicate key found, return an error
            return nil, fmt.Errorf("duplicate link key found: %s", key)
        }
        seen[key] = true
    }

    return &Shorty{
        links: config.Links,
        next:  next,
        name:  name,
    }, nil
}

func (s *Shorty) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    // Remove the leading '/' from req.URL.Path
    trimmedPath := strings.TrimPrefix(req.URL.Path, "/")
    
    // Check if the path is '/all'
    if trimmedPath == "all" {
        s.handleAllLinks(rw)
        return
    }
    // Check if the path is in the links map
    if target, ok := s.links[trimmedPath]; ok {
        http.Redirect(rw, req, target, http.StatusFound)
        return
    }

    // If the path is not in the map, pass the request to the next handler
    s.next.ServeHTTP(rw, req)
}

// handleAllLinks sends a JSON response with all the links
func (s *Shorty) handleAllLinks(rw http.ResponseWriter) {
    rw.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(rw)
    if err := encoder.Encode(s.links); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }
}
