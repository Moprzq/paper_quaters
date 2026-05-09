package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "HTTP server address")
	dir := flag.String("dir", "web", "directory to serve")
	build := flag.Bool("build", false, "rebuild browser wasm before serving")
	open := flag.Bool("open", true, "open the served URL in the default browser")
	language := flag.String("lang", "ru", "browser UI language: ru or eng")
	flag.Parse()

	normalizedLanguage, err := normalizeLanguage(*language)
	if err != nil {
		log.Fatal(err)
	}
	if err := ensureWASM(*dir, *build); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", noCache(http.FileServer(http.Dir(*dir))))

	serveURL := browserURL(*addr, normalizedLanguage)
	fmt.Printf("Serving %s at %s\n", *dir, serveURL)
	if *open {
		go func() {
			if err := openBrowser(serveURL); err != nil {
				log.Printf("open browser: %v", err)
			}
		}()
	}
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func ensureWASM(dir string, forceBuild bool) error {
	wasmPath := filepath.Join(dir, "paper_quarters.wasm")
	wasmExecPath := filepath.Join(dir, "wasm_exec.js")
	if !forceBuild && fileExists(wasmPath) && fileExists(wasmExecPath) {
		return nil
	}

	return buildWASM(dir)
}

func buildWASM(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create web dir: %w", err)
	}

	wasmPath := filepath.Join(dir, "paper_quarters.wasm")
	cmd := exec.Command("go", "build", "-o", wasmPath, "./cmd/paper-quarters")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build wasm: %w", err)
	}

	source, err := os.Open(filepath.Join(runtime.GOROOT(), "lib", "wasm", "wasm_exec.js"))
	if err != nil {
		return fmt.Errorf("open wasm_exec.js: %w", err)
	}
	defer source.Close()

	targetPath := filepath.Join(dir, "wasm_exec.js")
	if err := os.Chmod(targetPath, 0o666); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("make wasm_exec.js writable: %w", err)
	}

	target, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("create wasm_exec.js: %w", err)
	}
	defer target.Close()

	if _, err := io.Copy(target, source); err != nil {
		return fmt.Errorf("copy wasm_exec.js: %w", err)
	}

	return nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}

func normalizeLanguage(value string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "ru":
		return "ru", nil
	case "eng", "en":
		return "eng", nil
	default:
		return "", fmt.Errorf("unsupported language %q: use ru or eng", value)
	}
}

func browserURL(addr, language string) string {
	u := url.URL{
		Scheme: "http",
		Host:   addr,
		Path:   "/",
	}
	if language == "eng" {
		q := u.Query()
		q.Set("lang", language)
		u.RawQuery = q.Encode()
	}
	return u.String()
}
