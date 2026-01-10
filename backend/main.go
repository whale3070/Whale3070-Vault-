/*
 * 版权所有 (C) [2026] [whale3070/ Whale-Valut-NFT团队]
 * 本项目基于 CC BY-NC 4.0 协议开源，禁止第三方商用（详见仓库 LICENSE 文件）。
 * 著作权人保留本项目的商业使用权。
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

const (
	MatrixHomeserver = "https://matrix.org"
	MatrixRoomID     = "!jOcJpAxdUNYvaMZuqJ:matrix.org"
	hashCodeFilePath = "hash-code.txt"
)

type RelayRequest struct {
	Dest                string  `json:"dest"`
	Value               string  `json:"value"`
	GasLimit            string  `json:"gasLimit"`
	StorageDepositLimit *string `json:"storageDepositLimit"`
	DataHex             string  `json:"dataHex"`
	Signer              string  `json:"signer"`
	CodeHash            string  `json:"codeHash"`
}

type RelayResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type VerifyResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type Limiter struct {
	mu       sync.Mutex
	visitors map[string]*rate.Limiter
}

var (
	codeStatusMu sync.Mutex
	codeStatus   = map[string]string{}

	validCodesMu sync.Mutex
	validCodes   = map[string]struct{}{}

	usedCodesMu sync.Mutex
	usedCodes   = map[string]struct{}{}

	mintLogsMu sync.Mutex
	mintLogs   []map[string]any

	fileMu sync.Mutex
)

func NewLimiter() *Limiter {
	return &Limiter{visitors: make(map[string]*rate.Limiter)}
}

func (l *Limiter) get(ip string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()
	if lim, ok := l.visitors[ip]; ok {
		return lim
	}
	lim := rate.NewLimiter(rate.Every(2*time.Second), 3)
	l.visitors[ip] = lim
	return lim
}

func lockCode(hash string) (string, error) {
	if hash == "" {
		return "", fmt.Errorf("empty code hash")
	}
	codeStatusMu.Lock()
	defer codeStatusMu.Unlock()
	if v, ok := codeStatus[hash]; ok {
		return v, nil
	}
	codeStatus[hash] = "PENDING"
	return "OK", nil
}

func setCodeSuccess(hash string) {
	if hash == "" {
		return
	}
	codeStatusMu.Lock()
	defer codeStatusMu.Unlock()
	codeStatus[hash] = "SUCCESS"
	validCodesMu.Lock()
	delete(validCodes, hash)
	validCodesMu.Unlock()
	usedCodesMu.Lock()
	usedCodes[hash] = struct{}{}
	usedCodesMu.Unlock()
	if err := markCodeUsed(hash); err != nil {
		log.Printf("mark code used failed: %v", err)
	}
}

func loadValidCodes(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("cannot read hash code file: %v", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	valid := map[string]struct{}{}
	used := map[string]struct{}{}
	for _, line := range lines {
		s := strings.TrimSpace(line)
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "#") {
			continue
		}
		if strings.HasPrefix(s, "USED:") {
			code := strings.TrimSpace(strings.TrimPrefix(s, "USED:"))
			if code != "" {
				used[code] = struct{}{}
			}
			continue
		}
		valid[s] = struct{}{}
	}
	validCodesMu.Lock()
	validCodes = valid
	validCodesMu.Unlock()
	usedCodesMu.Lock()
	usedCodes = used
	usedCodesMu.Unlock()
	log.Printf("loaded %d valid codes, %d used codes from %s", len(valid), len(used), path)
}

func markCodeUsed(hash string) error {
	fileMu.Lock()
	defer fileMu.Unlock()
	data, err := os.ReadFile(hashCodeFilePath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	found := false
	for i, line := range lines {
		s := strings.TrimSpace(line)
		if s == hash {
			lines[i] = "USED:" + hash
			found = true
			break
		}
	}
	if !found {
		lines = append(lines, "USED:"+hash)
	}
	out := strings.Join(lines, "\n")
	return os.WriteFile(hashCodeFilePath, []byte(out), 0644)
}

func isCodeValid(hash string) bool {
	if hash == "" {
		return false
	}
	validCodesMu.Lock()
	_, ok := validCodes[hash]
	validCodesMu.Unlock()
	return ok
}

func isCodeUsed(hash string) bool {
	if hash == "" {
		return false
	}
	usedCodesMu.Lock()
	_, ok := usedCodes[hash]
	usedCodesMu.Unlock()
	return ok
}

func appendMintLog(entry map[string]any) {
	mintLogsMu.Lock()
	defer mintLogsMu.Unlock()
	mintLogs = append(mintLogs, entry)
	if len(mintLogs) > 1000 {
		mintLogs = mintLogs[len(mintLogs)-1000:]
	}
}

func getMintLogs(limit int) []map[string]any {
	mintLogsMu.Lock()
	defer mintLogsMu.Unlock()
	if limit <= 0 || limit > len(mintLogs) {
		limit = len(mintLogs)
	}
	out := make([]map[string]any, 0, limit)
	for i := 0; i < limit; i++ {
		out = append(out, mintLogs[i])
	}
	return out
}

func main() {
	limiter := NewLimiter()
	loadValidCodes(hashCodeFilePath)

	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/secret/verify", func(w http.ResponseWriter, r *http.Request) {
		codeHash := r.URL.Query().Get("codeHash")
		if codeHash == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(VerifyResponse{Ok: false, Error: "missing codeHash"})
			return
		}
		if isCodeUsed(codeHash) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(VerifyResponse{Ok: false, Error: "code used"})
			return
		}
		if !isCodeValid(codeHash) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(VerifyResponse{Ok: false, Error: "invalid code"})
			return
		}
		json.NewEncoder(w).Encode(VerifyResponse{Ok: true})
	}).Methods("GET")

	router.HandleFunc("/relay/mint", func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
		if !limiter.get(ip).Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "rate limited"})
			return
		}

		var req RelayRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "invalid json"})
			return
		}

		if req.CodeHash == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "missing code hash"})
			return
		}

		if isCodeUsed(req.CodeHash) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "此书已经生成过 NFT 了"})
			return
		}
		if !isCodeValid(req.CodeHash) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "无效的兑换码"})
			return
		}
		st, err := lockCode(req.CodeHash)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "lock error"})
			return
		}
		if st == "PENDING" {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "正在铸造中"})
			return
		}
		if st == "SUCCESS" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RelayResponse{Status: "error", Error: "此书已经生成过 NFT 了"})
			return
		}

		txHash := fmt.Sprintf("0x%x", time.Now().UnixNano())
		logEntry := map[string]any{
			"timestamp": time.Now().Unix(),
			"tx_hash":   txHash,
			"book_id":   r.URL.Query().Get("book_id"),
		}
		appendMintLog(logEntry)
		setCodeSuccess(req.CodeHash)
		json.NewEncoder(w).Encode(RelayResponse{Status: "submitted", TxHash: txHash})
	}).Methods("POST")

	// Matrix invitation endpoint
	router.HandleFunc("/api/matrix/test-invite", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		var req struct {
			MatrixID string `json:"matrixId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request format"})
			return
		}

		token := "mat_ZVZBVzMyxjn1IMKXMCSIpKyhPuz0qS_86XDZ3"
		if token == "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Backend not configured with MATRIX_ACCESS_TOKEN"})
			return
		}

		// Build Matrix API request
		url := fmt.Sprintf("%s/_matrix/client/v3/rooms/%s/invite", MatrixHomeserver, MatrixRoomID)
		payload, _ := json.Marshal(map[string]string{"user_id": req.MatrixID})

		log.Printf("Inviting user %s to Matrix room...\n", req.MatrixID)

		matrixReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create request"})
			return
		}
		matrixReq.Header.Set("Authorization", "Bearer "+token)
		matrixReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(matrixReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Connection to Matrix server timed out"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			json.NewEncoder(w).Encode(map[string]string{"error": "Matrix API returned error, please check token permissions or room ID"})
			return
		}

		log.Printf("Successfully invited user: %s\n", req.MatrixID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "success"})
	}).Methods("POST")

	// Root status endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		status := map[string]interface{}{
			"status": "Whale Vault Backend is Running",
			"services": map[string]string{
				"relay":  "active",
				"matrix": "active",
			},
			"endpoints": map[string]string{
				"relay":         "/relay/mint",
				"verify":        "/secret/verify",
				"metrics":       "/metrics/mint",
				"matrix_invite": "/api/matrix/test-invite",
			},
		}
		json.NewEncoder(w).Encode(status)
	}).Methods("GET")

	// Metrics endpoint for frontend
	router.HandleFunc("/metrics/mint", func(w http.ResponseWriter, r *http.Request) {
		out := getMintLogs(50)
		json.NewEncoder(w).Encode(out)
	}).Methods("GET")

	addr := ":8080"
	log.Printf("🚀 Whale Vault Backend is starting...")
	log.Printf("📍 Listening on %s", addr)
	log.Printf("🔧 Services: Relay API, Matrix Integration")
	log.Printf("🔗 Endpoints: /relay/mint, /api/matrix/test-invite, /metrics/mint")

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
