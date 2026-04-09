package main

import (
	"embed"
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed web/*
var content embed.FS

type Company struct {
	LegalName            string `json:"legal_name"`
	Inn                  string `json:"inn"`
	Kpp                  string `json:"kpp"`
	LegalAddress         string `json:"legal_address"`
	PostalAddress        string `json:"postal_address"`
	SettlementAccount    string `json:"settlement_account"`
	BankName             string `json:"bank_name"`
	CorrespondentAccount string `json:"correspondent_account"`
	Bik                  string `json:"bik"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Director             string `json:"director"`
}

var companyData = Company{
	LegalName:            "ООО «СТРОИТЕЛЬНО-МОНТАЖНАЯ КОМПАНИЯ \"К2\"»",
	Inn:                  "7819046579",
	Kpp:                  "781901001",
	LegalAddress:         "198412, г. Санкт-Петербург, вн.тер.г. город Ломоносов, ул. Победы, д. 24, лит. А, помещ. 4-Н",
	PostalAddress:        "198412, г. Санкт-Петербург, вн.тер.г. город Ломоносов, ул. Победы, д. 24, лит. А, помещ. 4-Н",
	SettlementAccount:    "40702810955000093779",
	BankName:             "СЕВЕРО-ЗАПАДНЫЙ БАНК ПАО СБЕРБАНК",
	CorrespondentAccount: "30101810500000000653",
	Bik:                  "044030653",
	Phone:                "+7 (962) 684-28-69",
	Email:                "smkk2@bk.ru",
	Director:             "Потоцкий Артур Александрович",
}

func main() {
	addr := envOrDefault("ADDR", ":8080")

	staticFS, err := fs.Sub(content, "web/static")
	if err != nil {
		log.Fatalf("prepare static fs: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", cacheControl(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))))
	mux.HandleFunc("/api/company", handleCompany)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/", serveIndex)

	server := &http.Server{
		Addr:         addr,
		Handler:      withSecurityHeaders(logging(mux)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server started on %s", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed: %v", err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file, err := content.ReadFile("web/index.html")
	if err != nil {
		http.Error(w, "index not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(file)
}

func handleCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(companyData); err != nil {
		http.Error(w, "encode failed", http.StatusInternalServerError)
	}
}

func envOrDefault(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func withSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; img-src 'self' data:; script-src 'self'; style-src 'self' https://fonts.googleapis.com; font-src https://fonts.gstatic.com; connect-src 'self'; object-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=(), payment=(), usb=(), accelerometer=(), gyroscope=()")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		next.ServeHTTP(w, r)
	})
}

func cacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/img/") || strings.HasPrefix(r.URL.Path, "img/") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
		}
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
