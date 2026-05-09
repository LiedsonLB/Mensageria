import React from 'react'
import { useNavigate } from 'react-router-dom'

const stats = [
  { label: 'Receita Mensal', value: 'R$ 428.590', delta: '+12,5%', positive: true, icon: 'trending_up', color: 'var(--cyan)' },
  { label: 'Pedidos Ativos', value: '1.842', delta: '+4,2%', positive: true, icon: 'shopping_bag', color: 'var(--green)' },
  { label: 'Ticket Médio', value: 'R$ 2.327', delta: '+8,1%', positive: true, icon: 'receipt_long', color: 'var(--purple)' },
  { label: 'Tempo Médio', value: '1,2 dias', delta: '-2,1%', positive: false, icon: 'timer', color: 'var(--amber)' },
]

const recentOrders = [
  { id: '#ORD-99231', cliente: 'Global Connect Ltda.', valor: 'R$ 12.450,00', status: 'CONCLUIDO' },
  { id: '#ORD-99230', cliente: 'Nexus Tech Soluções', valor: 'R$ 8.120,50', status: 'PROCESSANDO' },
  { id: '#ORD-99229', cliente: 'Solaris Manufatura', valor: 'R$ 450,00', status: 'PENDENTE' },
  { id: '#ORD-99228', cliente: 'Apex Analytics', valor: 'R$ 32.100,00', status: 'CONCLUIDO' },
]

const statusMap = {
  CONCLUIDO: 'badge-green',
  PROCESSANDO: 'badge-cyan',
  PENDENTE: 'badge-amber',
  CANCELADO: 'badge-red',
}

export default function Dashboard() {
  const navigate = useNavigate()

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 32 }}>
      {/* Header */}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end' }}>
        <div>
          <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-faint)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: 6 }}>
            PAINEL PRINCIPAL
          </div>
          <h1 style={{ fontFamily: 'var(--font-display)', fontSize: 28, fontWeight: 800, letterSpacing: '-0.02em' }}>
            Visão Geral
          </h1>
        </div>
        <button className="btn-primary" onClick={() => navigate('/novo-pedido')}>
          <span className="material-symbols-outlined" style={{ fontSize: 18 }}>add</span>
          Novo Pedido
        </button>
      </div>

      {/* Stats */}
      <div className="grid-4">
        {stats.map(s => (
          <div className="stat-card" key={s.label}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 14 }}>
              <div style={{ width: 38, height: 38, borderRadius: 'var(--radius)', background: `${s.color}18`, border: `1px solid ${s.color}30`, display: 'flex', alignItems: 'center', justifyContent: 'center', color: s.color }}>
                <span className="material-symbols-outlined" style={{ fontSize: 20 }}>{s.icon}</span>
              </div>
              <span className={`badge ${s.positive ? 'badge-green' : 'badge-red'}`}>{s.delta}</span>
            </div>
            <div className="stat-label">{s.label}</div>
            <div className="stat-value">{s.value}</div>
          </div>
        ))}
      </div>

      {/* Recent Orders */}
      <div className="card" style={{ padding: 0, overflow: 'hidden' }}>
        <div style={{ padding: '20px 24px', borderBottom: '1px solid var(--border)', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div className="section-title" style={{ fontSize: 16 }}>Pedidos Recentes</div>
          <button className="btn-ghost" style={{ padding: '6px 14px', fontSize: 12 }} onClick={() => navigate('/pedidos')}>
            Ver todos
            <span className="material-symbols-outlined" style={{ fontSize: 16 }}>arrow_forward</span>
          </button>
        </div>
        <table className="data-table">
          <thead>
            <tr>
              <th>ID do Pedido</th>
              <th>Cliente</th>
              <th>Valor</th>
              <th>Status</th>
              <th style={{ textAlign: 'right' }}>Ação</th>
            </tr>
          </thead>
          <tbody>
            {recentOrders.map(o => (
              <tr key={o.id}>
                <td><span className="mono text-cyan">{o.id}</span></td>
                <td style={{ fontWeight: 500 }}>{o.cliente}</td>
                <td style={{ fontFamily: 'var(--font-mono)', fontWeight: 600 }}>{o.valor}</td>
                <td><span className={`badge ${statusMap[o.status]}`}>{o.status}</span></td>
                <td style={{ textAlign: 'right' }}>
                  <button className="icon-btn" style={{ marginLeft: 'auto' }}>
                    <span className="material-symbols-outlined">visibility</span>
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Quick actions */}
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 20 }}>
        <div className="card" style={{ padding: 24 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 16 }}>
            <span className="material-symbols-outlined text-cyan">inventory_2</span>
            <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, fontSize: 15 }}>Catálogo de Produtos</span>
          </div>
          <p style={{ fontSize: 13, color: 'var(--text-muted)', marginBottom: 16, lineHeight: 1.6 }}>
            Explore 32+ produtos disponíveis para adicionar aos seus pedidos.
          </p>
          <button className="btn-ghost" style={{ fontSize: 12 }} onClick={() => navigate('/catalogo')}>
            Acessar Catálogo
            <span className="material-symbols-outlined" style={{ fontSize: 16 }}>arrow_forward</span>
          </button>
        </div>
        <div className="card" style={{ padding: 24, background: 'rgba(0,210,255,0.04)', borderColor: 'rgba(0,210,255,0.15)' }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 16 }}>
            <span className="material-symbols-outlined text-cyan">add_circle</span>
            <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, fontSize: 15 }}>Criar Novo Pedido</span>
          </div>
          <p style={{ fontSize: 13, color: 'var(--text-muted)', marginBottom: 16, lineHeight: 1.6 }}>
            Preencha os dados do cliente e selecione os produtos para gerar um pedido.
          </p>
          <button className="btn-primary" style={{ fontSize: 12, padding: '8px 16px' }} onClick={() => navigate('/novo-pedido')}>
            <span className="material-symbols-outlined" style={{ fontSize: 16 }}>add</span>
            Criar Pedido
          </button>
        </div>
      </div>
    </div>
  )
}
