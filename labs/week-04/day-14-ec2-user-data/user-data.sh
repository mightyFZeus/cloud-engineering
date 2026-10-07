#!/bin/bash
set -euo pipefail

dnf install -y golang

if ! id cloudeng >/dev/null 2>&1; then
    useradd \
        --system \
        --home-dir /nonexistent \
        --shell /sbin/nologin \
        cloudeng
fi

cat > /tmp/cloud-eng-health.go <<'GO_SOURCE'
package main

import (
	"log"
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("ok\n")); err != nil {
		log.Printf("write health response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("health service listening on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("health service failed: %v", err)
	}
}
GO_SOURCE

gofmt -w /tmp/cloud-eng-health.go
go build -trimpath -o /tmp/cloud-eng-health /tmp/cloud-eng-health.go

install -o root -g root -m 0755 \
    /tmp/cloud-eng-health \
    /usr/local/bin/cloud-eng-health

cat > /etc/systemd/system/cloud-eng-health.service <<'SYSTEMD_UNIT'
[Unit]
Description=Cloud Engineering Health Service
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
User=cloudeng
Group=cloudeng
ExecStart=/usr/local/bin/cloud-eng-health
Restart=on-failure
RestartSec=5s
NoNewPrivileges=true
PrivateTmp=true
ProtectHome=true
ProtectSystem=strict

[Install]
WantedBy=multi-user.target
SYSTEMD_UNIT

systemctl daemon-reload
systemctl enable --now cloud-eng-health.service
systemctl is-active --quiet cloud-eng-health.service
