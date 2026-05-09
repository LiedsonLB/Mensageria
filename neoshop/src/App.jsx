import React from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Layout from './components/Layout'
import Dashboard from './pages/Dashboard'
import HistoricoPedidos from './pages/HistoricoPedidos'
import NovoPedido from './pages/NovoPedido'
import Confirmacao from './pages/Confirmacao'
import Catalogo from './pages/Catalogo'

export default function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/pedidos" element={<HistoricoPedidos />} />
          <Route path="/novo-pedido" element={<NovoPedido />} />
          <Route path="/confirmacao" element={<Confirmacao />} />
          <Route path="/catalogo" element={<Catalogo />} />
          <Route path="/clientes" element={<PlaceholderPage titulo="Clientes" icone="group" />} />
          <Route path="/configuracoes" element={<PlaceholderPage titulo="Configurações" icone="settings" />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  )
}

function PlaceholderPage({ titulo, icone }) {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: 400, gap: 16, color: 'var(--text-faint)' }}>
      <span className="material-symbols-outlined" style={{ fontSize: 56, color: 'var(--border-hover)' }}>{icone}</span>
      <div style={{ fontFamily: 'var(--font-display)', fontSize: 22, fontWeight: 700, color: 'var(--text-muted)' }}>{titulo}</div>
      <div style={{ fontFamily: 'var(--font-mono)', fontSize: 12 }}>Em desenvolvimento</div>
    </div>
  )
}
