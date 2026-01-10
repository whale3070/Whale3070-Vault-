/*
 * 版权所有 (C) [2026] [whale3070/ Whale-Valut-NFT团队]
 * 本项目基于 CC BY-NC 4.0 协议开源，禁止第三方商用（详见仓库 LICENSE 文件）。
 * 著作权人保留本项目的商业使用权。
 */
import React from 'react'
import SalesBoard from '../components/SalesBoard'

export default function Home() {
  return (
    <div className="mx-auto max-w-7xl px-4 py-10">
      <section className="grid gap-6 md:grid-cols-2">
        <div className="rounded-xl border border-white/10 bg-white/5 p-6">
          <h2 className="text-lg font-semibold mb-2">NFT 金库</h2>
          <p className="text-sm text-white/70 mb-4">
            使用微信或系统相机扫码实体书二维码，将自动打开领取页面。
          </p>
          <div className="text-sm text-white/60">
            领取入口示例：<span className="font-mono">/valut_mint_nft/&lt;hashCode&gt;</span>
          </div>
        </div>
        <div className="rounded-xl border border-white/10 bg-white/5 p-6">
          <SalesBoard />
        </div>
      </section>
    </div>
  )
}
