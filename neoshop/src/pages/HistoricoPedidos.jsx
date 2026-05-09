import React, { useState, useEffect } from 'react'
import { listarPedidos } from '../services/api'

const statusMap = {
  CONCLUIDO: 'badge-green',
  PROCESSANDO: 'badge-cyan',
  PENDENTE: 'badge-amber',
  CANCELADO: 'badge-red',
}

const statusLabel = {
  CONCLUIDO: 'Concluído',
  PROCESSANDO: 'Processando',
  PENDENTE: 'Pendente',
  CANCELADO: 'Cancelado',
}

export default function HistoricoPedidos() {
  const [pedidos, setPedidos] = useState([])
  const [busca, setBusca] = useState('')
  const [filtroStatus, setFiltroStatus] = useState('TODOS')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    listarPedidos().then(data => {
      setPedidos(data)
      setLoading(false)
    })
  }, [])

  const filtrados = pedidos.filter(p => {
    const matchBusca = !busca || p.cliente.toLowerCase().includes(busca.toLowerCase()) || p.id.toLowerCase().includes(busca.toLowerCase())
    const matchStatus = filtroStatus === 'TODOS' || p.status === filtroStatus
    return matchBusca && matchStatus
  })

  const filtros = ['TODOS', 'PENDENTE', 'PROCESSANDO', 'CONCLUIDO', 'CANCELADO']
  const filtroLabel = { TODOS: 'Todos', PENDENTE: 'Pendentes', PROCESSANDO: 'Processando', CONCLUIDO: 'Concluídos', CANCELADO: 'Cancelados' }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 28 }}>
      {/* Header */}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end' }}>
        <div>
          <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-faint)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: 6 }}>GESTÃO DE PEDIDOS</div>
          <h1 style={{ fontFamily: 'var(--font-display)', fontSize: 28, fontWeight: 800, letterSpacing: '-0.02em' }}>Histórico de Pedidos</h1>
        </div>
        <button className="btn-ghost">
          <span className="material-symbols-outlined" style={{ fontSize: 18 }}>download</span>
          Exportar Dados
        </button>
      </div>

      {/* Stats strip */}
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(4,1fr)', gap: 16 }}>
        {[
          { label: 'Total de Pedidos', value: pedidos.length, color: 'var(--cyan)' },
          { label: 'Concluídos', value: pedidos.filter(p => p.status === 'CONCLUIDO').length, color: 'var(--green)' },
          { label: 'Pendentes', value: pedidos.filter(p => p.status === 'PENDENTE').length, color: 'var(--amber)' },
          { label: 'Cancelados', value: pedidos.filter(p => p.status === 'CANCELADO').length, color: 'var(--red)' },
        ].map(s => (
          <div className="stat-card" key={s.label} style={{ padding: 16 }}>
            <div className="stat-label">{s.label}</div>
            <div className="stat-value" style={{ fontSize: 28, color: s.color }}>{s.value}</div>
          </div>
        ))}
      </div>

      {/* Table */}
      <div className="card" style={{ padding: 0, overflow: 'hidden' }}>
        {/* Filters */}
        <div style={{ padding: '16px 20px', borderBottom: '1px solid var(--border)', display: 'flex', gap: 12, alignItems: 'center', flexWrap: 'wrap' }}>
          <div style={{ position: 'relative', flex: 1, minWidth: 200 }}>
            <span className="material-symbols-outlined" style={{ position: 'absolute', left: 10, top: '50%', transform: 'translateY(-50%)', fontSize: 18, color: 'var(--text-faint)' }}>search</span>
            <input
              className="input-field"
              style={{ paddingLeft: 36 }}
              placeholder="Buscar por cliente ou ID..."
              value={busca}
              onChange={e => setBusca(e.target.value)}
            />
          </div>
          <div style={{ display: 'flex', gap: 6 }}>
            {filtros.map(f => (
              <button
                key={f}
                onClick={() => setFiltroStatus(f)}
                style={{
                  padding: '6px 14px',
                  borderRadius: 'var(--radius)',
                  border: '1px solid',
                  borderColor: filtroStatus === f ? 'var(--cyan)' : 'var(--border)',
                  background: filtroStatus === f ? 'var(--cyan-dim)' : 'none',
                  color: filtroStatus === f ? 'var(--cyan)' : 'var(--text-muted)',
                  fontSize: 12,
                  fontFamily: 'var(--font-mono)',
                  cursor: 'pointer',
                  transition: 'all 0.15s',
                  textTransform: 'uppercase',
                  letterSpacing: '0.05em',
                }}
              >
                {filtroLabel[f]}
              </button>
            ))}
          </div>
        </div>

        {loading ? (
          <div style={{ padding: 60, textAlign: 'center', color: 'var(--text-faint)' }}>
            <span className="material-symbols-outlined spin" style={{ fontSize: 32, color: 'var(--cyan)' }}>autorenew</span>
            <p style={{ marginTop: 12, fontFamily: 'var(--font-mono)', fontSize: 12 }}>Carregando pedidos...</p>
          </div>
        ) : (
          <table className="data-table">
            <thead>
              <tr>
                <th>ID do Sistema</th>
                <th>Cliente</th>
                <th>E-mail</th>
                <th>Data</th>
                <th>Valor Total</th>
                <th>Status</th>
                <th style={{ textAlign: 'right' }}>Ações</th>
              </tr>
            </thead>
            <tbody>
              {filtrados.length === 0 ? (
                <tr>
                  <td colSpan={7} style={{ textAlign: 'center', color: 'var(--text-faint)', padding: 40, fontFamily: 'var(--font-mono)', fontSize: 12 }}>
                    Nenhum pedido encontrado
                  </td>
                </tr>
              ) : filtrados.map(p => (
                <tr key={p.id}>
                  <td><span className="mono text-cyan">#{p.id}</span></td>
                  <td style={{ fontWeight: 500 }}>{p.cliente}</td>
                  <td style={{ color: 'var(--text-muted)', fontSize: 12 }}>{p.email}</td>
                  <td style={{ color: 'var(--text-muted)', fontSize: 12 }}>{p.data || '—'}</td>
                  <td style={{ fontFamily: 'var(--font-mono)', fontWeight: 700, color: 'var(--cyan)' }}>
                    {p.valor_total.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                  </td>
                  <td><span className={`badge ${statusMap[p.status] || 'badge-amber'}`}>{statusLabel[p.status] || p.status}</span></td>
                  <td style={{ textAlign: 'right' }}>
                    <button className="icon-btn" style={{ display: 'inline-flex', marginLeft: 'auto' }}>
                      <span className="material-symbols-outlined">visibility</span>
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}

        <div style={{ padding: '14px 20px', borderTop: '1px solid var(--border)', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <span style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-faint)' }}>
            Exibindo <span style={{ color: 'var(--cyan)' }}>{filtrados.length}</span> de <span style={{ color: 'var(--text)' }}>{pedidos.length}</span> registros
          </span>
        </div>
      </div>
    </div>
  )
}
