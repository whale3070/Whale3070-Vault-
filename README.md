# Whale-Vault —— Web3驱动的数字出版与资产金库协议

### 1. 执行摘要 (Executive Summary)

Whale-Vault 是一个基于 Polkadot 和 Arweave 的去中心化数字出版与资产管理协议，专注解决传统出版业**防伪失效、结算黑箱、内容审查**三大痛点。

我们已完成**可运行的完整DApp原型**（前端 + Go Relay后端 + 合约交互），实现“一书一码、扫码Mint NFT、Gasless领取、即时分账、管理后台实时监控与提现”全流程。读者扫码即享正版数字权益，作者/出版社秒级到账并可一键提现，平台通过协议手续费与增值服务持续盈利。

**核心盈利逻辑**：  

- 每本正版书通过NFT铸造与动态内容门禁驱动80%以上读者转化  
- 即时分账 + 二手交易持续版税 → 作者/出版社高黏性  
- 协议抽成3-5% + 二手版税分成 → 平台高毛利、可规模化收入  

### 2. 市场背景与痛点分析 (Market & Pain Points)

#### 2.1 市场规模：万亿级的盗版黑洞

全球图书出版市场规模约为 **1270亿美元**，数字出版子市场增速更快，预计2025-2030年复合增长率超过8-13%。

盗版却造成巨大收益流失，现有防御体系已全面失效：

- **日本漫画与出版业**：2025年最新报告显示，海外+国内盗版造成约 **¥8.5万亿日元（约合510-550亿美元）** 的年度损失，其中权利人直接损失约 ¥7048亿。这已是日本出版业合法海外市场规模的数倍。

- **美国出版业**：电子书与数字内容盗版每年造成 **3亿美元以上** 直接损失，整体数字媒体盗版（含图书）对美国经济影响达 **292-710亿美元/年**。出版盗版网站访问量2024年达 **664亿次**，同比增长4.3%。

- **中国出版业**：在线文学与图书盗版历史损失巨大，近年虽有打击但亚洲仍是全球盗版重灾区。全球出版业整体盗版损失估计 **10-12亿美元/年**，中国贡献显著比例。

#### 2.2 行业五大刚需 (The 5 Rigid Demands)

1. **版权防伪**：传统激光标/防伪贴易复制，读者无法低成本验证真伪。数字化传播后，物理防伪彻底失效。

2. **经济结算**：传统版税结算周期长达6-18个月，数据不透明，作者常被拖欠或压低分成。

3. **动态内容**：纸质书内容静态，无法更新、迭代或与读者互动，附加价值低。

4. **创作者保护**：中心化平台存在审查、封号风险，创作者面临“数字死亡”威胁。

5. **分发效率**：缺乏裂变激励与合法二手流转机制，正版书“一次销售”后即失去持续收益。

### 3. 解决方案：Whale-Vault 协议 (The Solution)

Whale-Vault 采用 **“Web2 入口 + Web3 主权”** 的混合架构，重构出版全链路。

#### 3.1 核心产品逻辑

- **一书一码，链上确权**：每本实体书嵌入唯一哈希码。读者扫码 Mint NFT，合约自动验证真伪。盗版书无法通过验证。

- **智能分账收银台**：读者支付瞬间，资金通过 Ink! 合约原子化拆分（如：作者 80%, 出版社 20%），实现秒级到账，彻底打破年度结算黑箱。

- **逻辑掩体与永久存储**：
  - 内容上链至 Arweave，实现永久存储、抗篡改、抗删除。
  - 支持匿名钱包交互，创作者无需实名，规避审查风险。

#### 3.2 动态增值服务 (Value-Add)

- **Token Gating（NFT门禁）**：正版NFT持有者解锁隐藏章节、实时更新、作者私域社群、专属福利等。盗版书失去核心价值。

- **SBT 与可交易NFT**：
  - SBT（灵魂绑定）作为荣誉勋章与社区通行证。
  - 可交易NFT支持合法二手流转，每次转售自动触发版税分成。

### 4. 技术壁垒与架构 (Technical Architecture)

- **底层**：Polkadot + Scaffold DOT，提供高性能与跨链扩展（XCM）。
- **合约**：Solidity实现核心逻辑，兼容Solidity迁移。
- **零门槛体验**：Relayer 代付 Gas，普通读者扫码即用，无需钱包或加密货币。
- **隐私与安全**：内容非对称加密，仅NFT持有者可解密；Arweave 确保永久可验证。
- **竞争优势**：Polkadot生态中暂无直接对标实体书+NFT防伪项目，Whale-Vault 填补实体-数字桥接空白。

### 5. 团队介绍 (Team)

- **Whale3070 (Team Lead)**：8年内容创作与社区运营经验，负责产品哲学与用户体验设计。
- **Hank (Backend & Relayer)**：分布式系统架构师，负责Gasless高并发与稳定性。
- **Evan (Asset & Metadata)**：Arweave生态专家，确保内容永久性与分布式存储。

团队兼具Web2产品直觉与Web3硬核技术，已完成核心合约开发与测试网部署。

### 6. 亮点

1. **完整闭环**：从防伪入口 → 即时分账 → 动态内容 → 二手持续版税，形成强商业正循环。

2. **刚需高频**：直击出版业每年数万亿日元规模的盗版损失，用区块链解决物理世界无法解决的问题。

3. **巨大社会价值**：在内容审查日益严格的时代，提供真正的“数字避风港”，类似ProtonMail在隐私领域的定位，具备强大社区号召力。

4. **技术可落地**：Polkadot + Arweave成熟组合，Gasless设计大幅降低门槛；2026年可快速试点实体书合作。

5. **市场时机**：数字出版高速增长，Web3+NFT图书仍处早期蓝海，Whale-Vault 是少数兼顾实体桥接的创新协议。

6. **已落地的可运行原型**：完整DApp已开发完毕，支持扫码领取、Gasless铸造、管理后台实时监控与一键提现，极大降低执行风险。

### 7.财务模型 (Financial Projections)

#### 7.1 收入来源（Revenue Streams）

Whale-Vault 作为协议层项目，采用轻资产、多流可持续模型（参考典型Web3协议：手续费+版税分成+潜在代币经济）：

1. **协议手续费**：每笔读者支付/分账抽成 **3-5%**（其中1-2%归协议宝库，用于生态激励）。
2. **NFT铸造微费**：Gasless下可选0.1-0.5 USD/次（覆盖Relayer成本）。
3. **二手交易持续版税**：每次NFT转售，协议+创作者分成 **10%**（协议占2-3%）。
4. **增值服务**：Token Gating高级内容订阅分成（未来扩展）。
5. **代币经济（可选Phase 4）**：发行治理/实用代币$VAULT，用于手续费折扣、staking奖励，参考Polkadot生态项目捕获价值。

### 8. 当前开发进度与原型演示（Proof of Execution）

Whale-Vault **已完成完整可运行DApp原型**，涵盖读者全流程、作者/出版社管理后台与Gasless中继后端，充分验证商业模式的可行性与盈利能力。

#### 8.1 功能总览

##### 8.1.1 收银台（读者侧）

完整用户流：**扫码打开领取链接 → 自动校验兑换码 → 填写波卡地址 → 确认领取 → 成功页 → 解锁内容**。

- 领取页 `/valut_mint_nft/:hashCode`
  - 实体书二维码内容为 URL：`http://Domain/valut_mint_nft/:hashCode`
  - 页面通过 Path Parameter 读取 `hashCode`，并在加载时自动调用后端校验接口：
    - `GET /secret/verify?codeHash={hashCode}`
  - 校验通过后显示：
    - 新手引导（钱包下载说明）
    - 波卡地址输入框：「请输入您的波卡钱包地址」
    - 按钮：「确认领取」
  - 校验失败时显示全屏错误提示，不展示输入框

- 成功展示页 `/success`
  - 展示一个 NFT 勋章占位图和简洁的祝贺文案
  - 展示当前钱包地址与 `book_id`
  - 点击「验证访问权限」调用合约只读方法 `has_access(address, book_id)`
  - 验证通过后展示：
    - Arweave 资源链接：`https://arweave.net/{TX_ID}`
    - Matrix 私域社群入口
  - 提供「返回首页」按钮

##### 8.1.2 管理后台（作者 / 出版社）

管理后台挂在 `/admin` 路由下，采用侧边栏布局，包含：

- 数据总览 `/admin/overview`
  - 卡片展示：
    - 累计销售额
    - 已 Mint 数量
    - 当前可提现余额
  - 使用 Mock 数据 + 可选链上统计进行组合展示

- 销量监控看板 `/admin/monitor`
  - 数据卡片：总销售额、已铸造 NFT 数量、待提现余额
  - 趋势图表：使用 Recharts 绘制近 30 天销售增长曲线
  - 明细表：展示最近 Mint 记录（时间、book_id、用户地址缩略、状态）
  - 可从后端 `/metrics/mint` 或合约只读方法拉取统计，并与 Mock 数据叠加

- 销量明细 `/admin/sales`
  - 表格列出每笔 Mint 记录
  - 包含时间戳、book_id、交易哈希等
  - 目前以 Mock 数据为主，可平滑切换到真实后端

- 财务提现 `/admin/withdraw`
  - 展示当前账户在合约中的「可提取余额」
  - 点击「立即提现」时：
    - 调用 Polkadot{.js} 钱包，签名执行合约 `pull_funds()`
    - 显示发送与区块确认状态
    - 提现成功后触发前端 Confetti（纸屑特效）
    - 自动刷新可提余额

- 批次创建 `/admin/batch`
  - 支持导入 CSV 文件批量配置授权：
    - 每行结构：`book_id, secret_code`
  - 前端自动对 Secret Code 计算 SHA-256 哈希（使用浏览器 `crypto.subtle.digest`）
  - 组装参数并调用合约 `add_book_batch(ids, hashes)`，实现一次性写入链上

#### 8.1.3 后端（Middle-layer / Relay Server）

位置：`backend/main.go`

当前 Demo 职责：

- 接收前端的领取请求（包含 `dest` / `dataHex` / `signer` / `codeHash`）
- 对每个 IP 做简单限流与封禁检查
- 基于 `codeHash`（兑换码 Hash Code）做**唯一性锁**，防止同一兑换码被并发或重复刷接口
- 写入成功的 Mint 日志到 Redis（`mint:logs`），供前端销量看板查询

依赖：

- `github.com/gorilla/mux`
- `golang.org/x/time/rate`
- 标准库 `net/http` / `encoding/json` 等

---

#### 8.2安装与运行

##### 8.2.1 前端 DApp

##### 安装依赖

```bash
npm install
```

##### 开发模式

```bash
npm run dev
```

默认通过 Vite 启动开发服务器（通常是 `http://localhost:5173`）。

##### 打包构建

```bash
npm run build
```

产物输出到 `dist/` 目录，可使用任意静态服务器托管。

##### 8.2.2 后端 Relay Server

##### 前置依赖

- Go 1.20+（或与你 go.mod 对应的版本）
- 可读写的 `backend/hash-code.txt` 文件（用于维护兑换码状态）

##### 启动后端

```bash
cd backend
go run main.go
```

默认监听：`http://localhost:8080`

#### 8.3 前端主要页面说明

##### 8.3.1首页 `/`

- 文件：`src/pages/Home.tsx`
- 功能：
  - 左侧卡片：项目简介 + 领取入口说明（扫码后自动进入领取页）
  - 右侧卡片：可扩展的销量展示组件（`SalesBoard`）

##### 8.3.2领取页 `/valut_mint_nft/:hashCode`

- 文件：`src/pages/MintConfirm.tsx`
- 功能：
  - 通过 Path Parameter 读取 `hashCode`
  - 页面加载时自动调用 `GET /secret/verify?codeHash={hashCode}` 校验兑换码
  - 校验通过后展示波卡地址输入框与「确认领取」按钮
  - 点击「确认领取」后调用 `POST /relay/mint`，成功后跳转 `/success`

##### 8.3.3成功展示页 `/success`

- 文件：`src/pages/Success.tsx`
- 功能：
  - 展示 NFT 勋章占位图与成功文案
  - 从 URL 中读取：
    - `book_id`：书籍编号
    - `ar`：Arweave 交易 ID（如存在则优先使用）
  - 读取当前钱包地址（优先使用 `localStorage.selectedAddress`）
  - Arweave 映射与网关：
    - 常量 `ARWEAVE_GATEWAY = "https://arweave.net/"`
    - 内置 `BOOKS` 映射：`book_id → txId`（例如 Book 1 映射为 `uxtt46m7gTAAcS9pnyh8LkPErCr4PFJiqYjQnWcbzBI`）
    - 计算规则：`arTxId ? ARWEAVE_GATEWAY + arTxId : ARWEAVE_GATEWAY + BOOKS[book_id].txId`
  - 按钮：
    - 「验证访问权限」：调用合约 `has_access(address, book_id)`
    - 验证通过后：
      - 显示「打开 Arweave 内容」按钮（新窗口）
      - 显示「进入 Matrix 私域社群」入口
    - 验证失败时给出错误提示
  - 「返回首页」：返回用户端首页

##### 8.3.4 管理后台

- 布局：
  - `src/admin/AdminLayout.tsx`：侧边栏 + 嵌套路由结构

- 子页面：
  - `src/admin/OverviewPage.tsx`：数据总览
  - `src/admin/MonitorPage.tsx`：销量监控看板（图表 + 明细）
  - `src/admin/SalesPage.tsx`：销量明细列表
  - `src/admin/WithdrawPage.tsx`：财务提现
  - `src/admin/BatchPage.tsx`：批次创建（CSV 导入 + SHA-256 + add_book_batch）

---

#### 8.3 钱包与链配置

##### 8.3.1 链参数配置

- 页面：`/settings`
- Hook：`useChainConfig`（`src/state/useChainConfig.ts`）
- 字段：
  - `endpoint`：节点 WebSocket 地址
  - `contractAddress`：Ink! 合约地址
  - `abiUrl`：ABI JSON 文件地址
- 配置保存到 `localStorage.chainConfig`，前端所有合约调用均依赖此配置。

##### 8.3.2  钱包连接与账户选择

- Hook：`usePolkadotWallet`（`src/hooks/usePolkadotWallet.ts`）
- 功能：
  - 调用 `web3Enable('Whale Vault DApp')` 与 `web3Accounts()` 拉取扩展账户
  - 管理 `accounts` 列表与当前选中账户 `selected`
  - 支持切换账户与断开连接
  - 将选中的地址持久化到 `localStorage.selectedAddress`
  - 提供 `isConnected` 等布尔状态

导航栏组件 `NavBar` 通过该 Hook 实现：

- 「连接钱包」按钮：触发扩展授权
- 「断开」按钮：清空选择
- `AccountSelector`：展示并切换当前账户

#### 8.4 与合约的主要交互点

> 以下为前端使用到的合约接口名称，具体参数类型与链上实现需与实际 Ink! 合约保持一致。

- `has_access(address, book_id) -> bool`
  - 成功页中用于验证用户是否拥有访问该书籍内容的权限

- `pull_funds()`
  - 管理后台财务页调用
  - 将作者 / 出版社地址在合约中的可提现余额转出到当前账户

- `add_book_batch(ids: Vec<BookId>, hashes: Vec<Hash>)`
  - 批次创建页调用
  - 一次性写入多本书籍的授权哈希

- `get_withdrawable(address) -> Balance`（假定名称）
  - 用于查询某地址当前可提取余额
  - 前端在提现页只读调用，展示「当前地址可提取余额」

#### 8.5 典型使用流程（读者视角）

1. 通过实体书上的二维码或分享链接进入 DApp：
   - 二维码内容为 URL：`http://Domain/valut_mint_nft/{hashCode}`
2. 在领取页 `/valut_mint_nft/:hashCode` 中：
   - 页面自动调用 `GET /secret/verify?codeHash=...` 校验兑换码
   - 校验通过后输入波卡地址，点击「确认领取」提交领取请求到 `/relay/mint`
3. 成功后自动跳转 `/success`，在成功页点击「验证访问权限」，通过后：
   - 打开 Arweave 正文内容（优先使用 URL 参数 `ar`，否则使用 BOOKS 映射）。
   - 加入 Matrix 私域社群。

#### 8.6 部署与上线

##### 8.6.1 构建前端静态资源

```bash
npm install
npm run build
```

构建完成后，前端静态文件位于 `dist/` 目录。

##### 8.6.2 启动 Go Relay Server

在服务器上准备好 Go 环境，然后：

```bash
# 1. 进入后端目录
cd backend

# 2. 确认 `hash-code.txt` 中已经写入待发放的兑换码（每行一个，已使用的行以 `USED:` 前缀标记）

# 3. 启动后端（开发/测试时可以直接用 go run）
go run main.go
```

默认监听 `:8080`，即本机 `http://127.0.0.1:8080`。

##### 8.6.3 使用 Nginx 托管前端并反向代理后端

将 `dist/` 上传到服务器（例如 `/var/www/whale-vault`），站点配置示例（以 `198.55.109.102` 为例）：

```nginx
server {
    listen 80;
    server_name 198.55.109.102;

    root /var/www/whale-vault;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    # 转发中继与统计接口到本机 8080 端口
    location /relay/ {
        proxy_pass http://127.0.0.1:8080/relay/;
    }

    location /metrics/ {
        proxy_pass http://127.0.0.1:8080/metrics/;
    }

    # 转发兑换码校验接口
    location /secret/ {
        proxy_pass http://127.0.0.1:8080/secret/;
    }
}
```

这样，前端配置 `BACKEND_URL = 'http://198.55.109.102'` 后，会访问：

- `http://198.55.109.102/secret/verify`
- `http://198.55.109.102/relay/mint`
- `http://198.55.109.102/metrics/mint`

Nginx 会将这些请求转发到同一台机器上的 Go 后端（端口 8080），整个部署结构清晰且避免跨域问题。

#### 8.7 典型使用流程（作者 / 出版社视角）

1. 打开 DApp，连接钱包，使用作者 / 出版社账户登录
2. 进入管理后台 `/admin/overview` 查看数据总览
3. 在 `/admin/monitor` 看板中观察近 30 天销量趋势与最新 Mint 明细
4. 在 `/admin/batch` 导入 CSV 文件，批量配置书籍授权：
   - `book_id,secret_code` → 前端计算 SHA-256 哈希 → 调用 `add_book_batch`
5. 在 `/admin/withdraw` 查看可提取余额，并点击「立即提现」调用 `pull_funds()` 将收益转入当前地址

  

请马上扫码，来亲身来感受“扫码即盈利”。
