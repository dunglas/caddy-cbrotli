// Package caddycbrotli provides provides support for the Brotli compression format
// using a fast and efficient implementation written in C.
package caddycbrotli

import (
	"errors"
	"strconv"

	"github.com/google/brotli/go/cbrotli"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/encode"
)

func init() {
	caddy.RegisterModule(Br{})
}

// Br can create Brotli encoders.
type Br struct {
	// Quality controls the compression-speed vs compression-density trade-offs.
	// The higher the quality, the slower the compression. Range is 0 to 11. Defaults to 6.
	Quality *int `json:"quality,omitempty"`
	// LGWin is the base 2 logarithm of the sliding window size.
	// Range is 10 to 24. 0 indicates automatic configuration based on Quality.
	LGWin int `json:"lgwin,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Br) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.encoders.br",
		New: func() caddy.Module { return new(Br) },
	}
}

// UnmarshalCaddyfile sets up the handler from Caddyfile tokens.
func (b *Br) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.NextArg() {
			continue
		}
		qualityStr := d.Val()
		quality, err := strconv.Atoi(qualityStr)
		if err != nil {
			return err
		}
		b.Quality = &quality

		if !d.NextArg() {
			continue
		}
		lgwinStr := d.Val()
		lgwin, err := strconv.Atoi(lgwinStr)
		if err != nil {
			return err
		}
		b.LGWin = lgwin
	}

	return nil
}

// Validate validates b's configuration.
func (b Br) Validate() error {
	if b.LGWin != 0 {
		if b.LGWin < 10 {
			return errors.New("logarithm of the sliding window size too low; must be >= 10")
		}
		if b.LGWin > 24 {
			return errors.New("logarithm of the sliding window size too high; must be <= 24")
		}
	}

	if b.Quality != nil {
		if *b.Quality < 0 {
			return errors.New("quality too low; must be >= 0")
		}
		if *b.Quality > 11 {
			return errors.New("quality too high; must be <= 11")
		}
	}

	return nil
}

// AcceptEncoding returns the name of the encoding as
// used in the Accept-Encoding request headers.
func (Br) AcceptEncoding() string { return "br" }

// NewEncoder returns a new Brotli writer.
func (b Br) NewEncoder() encode.Encoder {
	q := 6
	if b.Quality != nil {
		q = *b.Quality
	}

	return newEncoder(cbrotli.WriterOptions{Quality: q, LGWin: b.LGWin})
}

// Interface guards
var (
	_ encode.Encoding       = (*Br)(nil)
	_ caddy.Validator       = (*Br)(nil)
	_ caddyfile.Unmarshaler = (*Br)(nil)
)
