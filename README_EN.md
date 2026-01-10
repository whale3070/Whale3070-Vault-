# Whale-Vault — Web3-Powered Digital Publishing and Asset Vault Protocol

> Copyright Notice: This project's source code is open-source and developed for the Polkadot Hackathon. Commercial use of any code, documents, or derivatives from this project is prohibited without written permission from the project team, and violators will be held legally liable.

### 1. Executive Summary

Whale-Vault is a decentralized digital publishing and asset management protocol built on Polkadot and Arweave, focused on solving the three core pain points of the traditional publishing industry: **counterfeiting ineffectiveness, opaque settlement, and content censorship**.

We have completed a **fully functional DApp prototype** (frontend + Go Relay backend + contract interaction), implementing the complete flow of “one code per book, scan-to-mint NFT, gasless claiming, instant settlement, real-time monitoring in the admin dashboard, and one-click withdrawal.” Readers scan to instantly access exclusive digital rights, authors/publishers receive payments in seconds and can withdraw with one click, and the platform generates sustainable revenue through protocol fees and value-added services.

**Core Profit Logic**:

- Each legitimate physical book drives over 80% reader conversion through NFT minting and dynamic content gating.
- Instant settlement + ongoing royalties from secondary sales → high retention for authors and publishers.
- Protocol takes a 3-5% fee + share of secondary royalties → high margins and scalable revenue for the platform.

### 2. Market Background and Pain Points

#### 2.1 Market Size: A Trillion-Yen Piracy Black Hole

The global book publishing market is valued at approximately **$127 billion**, with the digital publishing sub-sector growing rapidly and projected to achieve a CAGR exceeding 8-13% from 2025 to 2030.

Piracy, however, causes massive revenue loss, and existing anti-counterfeiting systems have completely failed:

- **Japanese manga and publishing industry**: 2025 reports show that domestic and overseas piracy causes approximately **¥8.5 trillion JPY (about $51-55 billion USD)** in annual losses, with direct losses to rights holders around ¥704.8 billion. This is several times the size of Japan’s legitimate overseas publishing market.

- **U.S. publishing industry**: Digital book and content piracy causes more than **$300 million** in direct annual losses, with the broader impact of digital media piracy (including books) on the U.S. economy reaching **$29.2-71 billion/year**. Publishing piracy websites recorded **66.4 billion visits** in 2024, up 4.3% year-over-year.

- **Chinese publishing industry**: Historical losses from online literature and book piracy have been enormous. Despite crackdowns in recent years, Asia remains the global epicenter of piracy. Global publishing piracy losses are estimated at **$1-1.2 billion/year**, with China contributing a significant share.

#### 2.2 The Five Core Industry Needs

1. **Copyright Anti-Counterfeiting**: Traditional laser labels and security stickers are easily replicated, and readers have no low-cost way to verify authenticity. Once content is digitized, physical anti-counterfeiting becomes completely ineffective.

2. **Economic Settlement**: Traditional royalty settlements take 6-18 months, with opaque data; authors are often delayed or shortchanged.

3. **Dynamic Content**: Printed books are static, cannot be updated or iterated, and lack interactive channels with readers, resulting in low added value.

4. **Creator Protection**: Centralized platforms carry risks of account bans and content deletion, leaving creators vulnerable to “digital death.”

5. **Distribution Efficiency**: Lack of viral incentives and legitimate secondary resale mechanisms means legitimate books lose ongoing revenue after the initial sale.

### 3. Solution: Whale-Vault Protocol

Whale-Vault adopts a **“Web2 entry + Web3 sovereignty”** hybrid architecture to restructure the entire publishing value chain.

#### 3.1 Core Product Logic

- **One Code Per Book, On-Chain Provenance**: Each physical book embeds a unique hash code. Readers scan to mint an NFT, and the contract automatically verifies authenticity. Pirated copies fail verification.

- **Intelligent Settlement Cashier**: Payments are instantly split atomically via Ink! contracts (e.g., 80% to author, 20% to publisher), achieving second-level settlement and completely eliminating annual settlement black boxes.

- **Logical Sanctuary and Permanent Storage**:
  - Content is stored on Arweave for permanent, tamper-proof, undeletable preservation.
  - Supports anonymous wallet interaction; creators need no real-name registration, avoiding censorship risks.

#### 3.2 Value-Added Services

- **Token Gating (NFT Access Control)**: Legitimate NFT holders unlock hidden chapters, real-time updates, private author communities, and exclusive perks—rendering pirated copies worthless.

- **SBT and Tradable NFTs**:
  - Soulbound Tokens (SBTs) serve as badges of honor and community passes.
  - Tradable NFTs enable legitimate secondary sales, with automatic royalty distribution on each resale.

### 4. Technical Barriers and Architecture

- **Underlying Layer**: Polkadot + Scaffold DOT for high performance and cross-chain expansion (XCM).
- **Contracts**: Solidity implements core logic.
- **Zero-Barrier Experience**: Relayer covers gas fees; ordinary readers can scan and use without a wallet or cryptocurrency.
- **Privacy and Security**: Content is asymmetrically encrypted and decryptable only by NFT holders; Arweave ensures permanent verifiability.
- **Competitive Advantage**: No direct competitor in the Polkadot ecosystem for physical book + NFT anti-counterfeiting; Whale-Vault uniquely bridges physical and digital worlds.

### 5. Team

- **Whale3070 (Team Lead)**: 8 years of content creation and community operations experience, responsible for product philosophy and UX design.
- **Hank (Backend & Relayer)**: Distributed systems architect, responsible for gasless high-concurrency and stability.
- **Evan (Asset & Metadata)**: Arweave ecosystem expert, ensuring true permanence and decentralized sovereignty of vault assets.

The team combines Web2 product intuition with hardcore Web3 technical expertise and has completed core contract development and testnet deployment.

### 6. Highlights

1. **Complete Closed Loop**: From anti-counterfeiting entry → instant settlement → dynamic content → ongoing secondary royalties, forming a powerful positive commercial cycle.

2. **High-Frequency Rigid Demand**: Directly addresses trillions of yen in annual publishing piracy losses, solving problems the physical world cannot with blockchain.

3. **Significant Social Value**: In an era of increasingly strict content censorship, provides a true “digital sanctuary,” similar to ProtonMail’s positioning in privacy, with strong community appeal.

4. **Proven Technical Feasibility**: Mature Polkadot + Arweave stack, gasless design dramatically lowers barriers; rapid physical book pilots possible in 2026.

5. **Perfect Market Timing**: Digital publishing is growing rapidly, Web3 + NFT books remain an early blue ocean; Whale-Vault is one of the few innovations bridging physical and digital.

6. **Fully Functional Deployed Prototype**: Complete DApp is ready, supporting scan-to-claim, gasless minting, real-time admin monitoring, and one-click withdrawal—greatly reducing execution risk.

### 7. Financial Model (Financial Projections)

#### 7.1 Revenue Streams

Whale-Vault operates as a lightweight protocol-layer project with a multi-stream sustainable model (referencing typical Web3 protocols: fees + royalty shares + potential token economics):

1. **Protocol Fees**: 3-5% on each reader payment/settlement (1-2% directed to protocol treasury for ecosystem incentives).
2. **NFT Minting Micro-Fees**: Optional 0.1-0.5 USD per mint under gasless model (covers Relayer costs).
3. **Secondary Sale Royalties**: 10% total on each NFT resale (protocol takes 2-3%).
4. **Value-Added Services**: Revenue share from token-gated premium content subscriptions (future expansion).
5. **Token Economics (Optional in Phase 4)**: Issue governance/utility token $VAULT for fee discounts and staking rewards, capturing value like other Polkadot ecosystem projects.

### 8. Current Development Progress and Prototype Demo (Proof of Execution)

Whale-Vault has completed a **fully functional DApp prototype** covering the entire reader flow, author/publisher admin dashboard, and gasless relay backend—fully validating the feasibility and profitability of the business model.

#### 8.1 Feature Overview

##### 8.1.1 Reader-Side Cashier

Complete user journey: **Scan code → open claim link → automatic code verification → enter Polkadot address → confirm claim → success page → unlock content**.

- Claim Page `/valut_mint_nft/:hashCode`
  - QR code in physical book contains URL: `http://Domain/valut_mint_nft/:hashCode`
  - Page reads `hashCode` from path parameter and automatically calls backend verification endpoint on load:
    - `GET /secret/verify?codeHash={hashCode}`
  - On success: displays beginner guide (wallet download instructions), Polkadot address input field, and “Confirm Claim” button.
  - On failure: full-screen error, no input field shown.

- Success Page `/success`
  - Shows NFT badge placeholder and congratulatory copy.
  - Displays current wallet address and `book_id`.
  - “Verify Access” button calls contract readonly method `has_access(address, book_id)`.
  - On success: displays Arweave content link (`https://arweave.net/{TX_ID}`) and Matrix private community entry.
  - Includes “Return Home” button.

##### 8.1.2 Admin Dashboard (Author / Publisher)

Admin dashboard under `/admin` route with sidebar layout, including:

- Overview `/admin/overview`
  - Cards showing cumulative sales, minted NFTs, and withdrawable balance (mock + optional on-chain data).

- Sales Monitoring Dashboard `/admin/monitor`
  - Cards: total sales, minted NFTs, pending withdrawal balance.
  - Trend chart: 30-day sales growth curve using Recharts.
  - Recent mint table: timestamp, book_id, truncated user address, status.
  - Data sourced from backend `/metrics/mint` or contract queries, blended with mock data.

- Sales Details `/admin/sales`
  - Table listing all mint records (timestamp, book_id, tx hash, etc.), currently mock-based with seamless switch to real backend.

- Withdrawal `/admin/withdraw`
  - Shows withdrawable balance in contract.
  - “Withdraw Now” calls Polkadot.js wallet to sign and execute `pull_funds()`.
  - Displays transaction status; success triggers confetti animation and auto-refresh.

- Batch Creation `/admin/batch`
  - Supports CSV import for bulk authorization:
    - Row format: `book_id, secret_code`
  - Frontend computes SHA-256 hash of secret code (`crypto.subtle.digest`).
  - Calls contract `add_book_batch(ids, hashes)` for bulk on-chain write.

##### 8.1.3 Backend (Middle-layer / Relay Server)

Location: `backend/main.go`

Current demo responsibilities:

- Handles frontend claim requests (`dest`, `dataHex`, `signer`, `codeHash`).
- Rate limiting and IP blocking.
- Unique lock on `codeHash` to prevent concurrent/repeated claims.
- Logs successful mints to Redis (`mint:logs`) for dashboard queries.

Dependencies:

- `github.com/gorilla/mux`
- `golang.org/x/time/rate`
- Standard library `net/http`, `encoding/json`, etc.

---

#### 8.2 Installation and Running

##### 8.2.1 Frontend DApp

**Install dependencies**

```bash
npm install
```

**Development mode**

```bash
npm run dev
```

Starts Vite dev server (typically `http://localhost:5173`).

**Production build**

```bash
npm run build
```

Outputs to `dist/` directory for static hosting.

##### 8.2.2 Backend Relay Server

**Prerequisites**

- Go 1.20+ (or version matching go.mod)
- Writable `backend/hash-code.txt` file (for code state tracking)

**Start backend**

```bash
cd backend
go run main.go
```

Listens on `http://localhost:8080` by default.

#### 8.3 Key Frontend Pages

##### 8.3.1 Home `/`

- File: `src/pages/Home.tsx`
- Features: Left card with project intro and claim instructions; right card with extensible sales board.

##### 8.3.2 Claim Page `/valut_mint_nft/:hashCode`

- File: `src/pages/MintConfirm.tsx`
- Reads `hashCode` from path, auto-verifies via backend, shows address input and claim button, redirects to `/success` on success.

##### 8.3.3 Success Page `/success`

- File: `src/pages/Success.tsx`
- Shows NFT badge and success message.
- Reads `book_id` and optional `ar` (Arweave TX ID) from URL.
- Uses built-in `BOOKS` mapping for Arweave content.
- “Verify Access” calls contract; success reveals Arweave link and Matrix community entry.

##### 8.3.4 Admin Dashboard

- Layout: `src/admin/AdminLayout.tsx` (sidebar + nested routes)
- Pages: Overview, Monitor (charts + details), Sales list, Withdrawal, Batch creation.

---

#### 8.4 Wallet and Chain Configuration

##### 8.4.1 Chain Settings

- Page: `/settings`
- Hook: `useChainConfig` (`src/state/useChainConfig.ts`)
- Persisted to `localStorage.chainConfig`.

##### 8.4.2 Wallet Connection

- Hook: `usePolkadotWallet` (`src/hooks/usePolkadotWallet.ts`)
- Handles extension authorization, account selection, persistence to `localStorage.selectedAddress`.

#### 8.5 Key Contract Interactions

- `has_access(address, book_id) -> bool`: Verify content access.
- `pull_funds()`: Withdraw earnings.
- `add_book_batch(ids: Vec<BookId>, hashes: Vec<Hash>)`: Bulk authorization.
- `get_withdrawable(address) -> Balance`: Query withdrawable balance.

#### 8.6 Reader Journey

1. Scan QR code or link from physical book → enters DApp.
2. Claim page verifies code, user enters address, submits to `/relay/mint`.
3. Success page → verify access → unlock Arweave content and Matrix community.

#### 8.7 Deployment and Launch

##### 8.7.1 Build Frontend

```bash
npm install
npm run build
```

##### 8.7.2 Start Go Relay Server

Prepare Go environment, ensure `hash-code.txt` is populated, then:

```bash
cd backend
go run main.go
```

##### 8.7.3 Nginx Hosting + Proxy Example

(Provided configuration for static hosting + backend proxy on port 8080.)

#### 8.8 Author / Publisher Journey

1. Connect wallet with author/publisher account.
2. View overview, monitor trends, batch upload authorizations, and withdraw earnings instantly.

Please scan the code now to experience the “scan-to-profit” flow firsthand.
