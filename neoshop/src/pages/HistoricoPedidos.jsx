import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { listarPedidos } from '../services/api'

const statusMap = {
  PENDENTE: 'badge-amber',
  PROCESSANDO_PAGAMENTO: 'badge-cyan',
  PAGAMENTO_APROVADO: 'badge-green',
  GERANDO_NF: 'badge-cyan',
  NF_EMITIDA: 'badge-green',
  ENVIANDO_EMAIL: 'badge-cyan',
  CONCLUIDO: 'badge-green',
  PAGAMENTO_FALHOU: 'badge-red',
  ERRO_GERAR_NF: 'badge-red',
  CANCELADO: 'badge-red',
}

const statusLabel = {
  PENDENTE: 'Pendente',
  PROCESSANDO_PAGAMENTO: 'Processando Pagamento',
  PAGAMENTO_APROVADO: 'Pagamento Aprovado',
  GERANDO_NF: 'Gerando Nota Fiscal',
  NF_EMITIDA: 'NF Emitida',
  ENVIANDO_EMAIL: 'Enviando E-mail',
  CONCLUIDO: 'Concluído',
  PAGAMENTO_FALHOU: 'Pagamento Falhou',
  ERRO_GERAR_NF: 'Erro na Nota Fiscal',
  CANCELADO: 'Cancelado',
}

export default function HistoricoPedidos() {
  const navigate = useNavigate()
  const [pedidos, setPedidos] = useState([])
  const [busca, setBusca] = useState('')
  const [filtroStatus, setFiltroStatus] = useState('TODOS')
  const [loading, setLoading] = useState(true)
  const [erro, setErro] = useState('')

  useEffect(() => {
    carregarPedidos()
    
    const interval = setInterval(() => {
      if (!loading) {
        carregarPedidos()
      }
    }, 10000)
    
    return () => clearInterval(interval)
  }, [])

  const carregarPedidos = async () => {
    try {
      setLoading(true)
      const data = await listarPedidos()
      setPedidos(data)
      setErro('')
    } catch (error) {
      console.error('Erro ao carregar pedidos:', error)
      setErro('Erro ao carregar pedidos. Verifique se o servidor está rodando.')
    } finally {
      setLoading(false)
    }
  }

  const filtrados = pedidos.filter(p => {
    const matchBusca = !busca || 
      p.cliente?.toLowerCase().includes(busca.toLowerCase()) || 
      p.id?.toLowerCase().includes(busca.toLowerCase())
    const matchStatus = filtroStatus === 'TODOS' || p.status === filtroStatus
    return matchBusca && matchStatus
  })

  const filtros = [
    'TODOS', 
    'PENDENTE', 
    'PROCESSANDO_PAGAMENTO', 
    'PAGAMENTO_APROVADO',
    'GERANDO_NF',
    'NF_EMITIDA',
    'ENVIANDO_EMAIL',
    'CONCLUIDO', 
    'PAGAMENTO_FALHOU',
    'CANCELADO'
  ]
  
  const filtroLabel = { 
    TODOS: 'Todos', 
    PENDENTE: 'Pendentes', 
    PROCESSANDO_PAGAMENTO: 'Processando Pagamento',
    PAGAMENTO_APROVADO: 'Pagamento Aprovado',
    GERANDO_NF: 'Gerando NF',
    NF_EMITIDA: 'NF Emitida',
    ENVIANDO_EMAIL: 'Enviando Email',
    CONCLUIDO: 'Concluídos', 
    PAGAMENTO_FALHOU: 'Pagamento Falhou',
    CANCELADO: 'Cancelados'
  }

  const formatarData = (dataISO) => {
    if (!dataISO) return '—'
    const data = new Date(dataISO)
    return data.toLocaleDateString('pt-BR', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  const verDetalhes = (pedido) => {
    navigate('/confirmacao', { state: { pedido } })
  }

  const recarregar = async () => {
    await carregarPedidos()
  }

  const totalPedidos = pedidos.length
  const concluidos = pedidos.filter(p => p.status === 'CONCLUIDO').length
  const pendentes = pedidos.filter(p => p.status === 'PENDENTE' || p.status === 'PROCESSANDO_PAGAMENTO').length
  const falhas = pedidos.filter(p => p.status === 'PAGAMENTO_FALHOU' || p.status === 'ERRO_GERAR_NF' || p.status === 'CANCELADO').length
  const processando = pedidos.filter(p => p.status === 'PROCESSANDO_PAGAMENTO' || p.status === 'GERANDO_NF' || p.status === 'ENVIANDO_EMAIL').length

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 28 }}>
      {/* Header */}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end' }}>
        <div>
          <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-faint)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: 6 }}>
            GESTÃO DE PEDIDOS
          </div>
          <h1 style={{ fontFamily: 'var(--font-display)', fontSize: 28, fontWeight: 800, letterSpacing: '-0.02em' }}>
            Histórico de Pedidos
          </h1>
          <p style={{ color: 'var(--text-muted)', fontSize: 13, marginTop: 4 }}>
            {pedidos.length} pedidos no total · Atualização automática a cada 10s
          </p>
        </div>
        <div style={{ display: 'flex', gap: 8 }}>
          <button className="btn-ghost" onClick={recarregar}>
            <span className="material-symbols-outlined" style={{ fontSize: 18 }}>refresh</span>
            Atualizar
          </button>
          <button className="btn-primary" onClick={() => navigate('/novo-pedido')}>
            <span className="material-symbols-outlined" style={{ fontSize: 18 }}>add</span>
            Novo Pedido
          </button>
        </div>
      </div>

      {/* Stats strip - 5 cards para mostrar mais informações */}
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(5,1fr)', gap: 16 }}>
        {[
          { label: 'Total de Pedidos', value: totalPedidos, color: 'var(--cyan)' },
          { label: 'Concluídos', value: concluidos, color: 'var(--green)' },
          { label: 'Em Processamento', value: processando, color: 'var(--cyan)' },
          { label: 'Pendentes', value: pendentes, color: 'var(--amber)' },
          { label: 'Falhas/Cancelados', value: falhas, color: 'var(--red)' },
        ].map(s => (
          <div className="stat-card" key={s.label} style={{ padding: 16, background: 'var(--bg-elevated)', borderRadius: 'var(--radius)', border: '1px solid var(--border)' }}>
            <div className="stat-label" style={{ fontSize: 12, color: 'var(--text-muted)', textTransform: 'uppercase', fontFamily: 'var(--font-mono)', marginBottom: 8 }}>{s.label}</div>
            <div className="stat-value" style={{ fontSize: 32, fontWeight: 700, color: s.color }}>{s.value}</div>
          </div>
        ))}
      </div>

      {/* Error message */}
      {erro && (
        <div style={{ 
          background: 'rgba(255,82,82,0.08)', 
          border: '1px solid rgba(255,82,82,0.25)', 
          borderRadius: 'var(--radius)', 
          padding: '12px 16px', 
          color: 'var(--red)', 
          fontSize: 13, 
          display: 'flex', 
          gap: 10, 
          alignItems: 'center' 
        }}>
          <span className="material-symbols-outlined" style={{ fontSize: 20 }}>error</span>
          {erro}
        </div>
      )}

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
          <div style={{ display: 'flex', gap: 6, flexWrap: 'wrap', maxHeight: 100, overflowY: 'auto' }}>
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
                  fontSize: 11,
                  fontFamily: 'var(--font-mono)',
                  cursor: 'pointer',
                  transition: 'all 0.15s',
                  textTransform: 'uppercase',
                  letterSpacing: '0.05em',
                  whiteSpace: 'nowrap',
                }}
              >
                {filtroLabel[f]}
              </button>
            ))}
          </div>
        </div>

        {loading ? (
          <div style={{ padding: 60, textAlign: 'center', color: 'var(--text-faint)' }}>
            <span className="material-symbols-outlined spin" style={{ fontSize: 32, color: 'var(--cyan)', animation: 'spin 1s linear infinite' }}>autorenew</span>
            <p style={{ marginTop: 12, fontFamily: 'var(--font-mono)', fontSize: 12 }}>Carregando pedidos...</p>
          </div>
        ) : (
          <div style={{ overflowX: 'auto' }}>
            <table style={{ width: '100%', borderCollapse: 'collapse', minWidth: 800 }}>
              <thead>
                <tr style={{ borderBottom: '1px solid var(--border)', background: 'var(--bg-elevated)' }}>
                  <th style={{ padding: 16, textAlign: 'left', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>ID do Pedido</th>
                  <th style={{ padding: 16, textAlign: 'left', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>Cliente</th>
                  <th style={{ padding: 16, textAlign: 'left', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>E-mail</th>
                  <th style={{ padding: 16, textAlign: 'left', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>Data</th>
                  <th style={{ padding: 16, textAlign: 'right', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>Valor Total</th>
                  <th style={{ padding: 16, textAlign: 'center', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>Status</th>
                  <th style={{ padding: 16, textAlign: 'center', fontSize: 12, fontWeight: 600, color: 'var(--text-muted)' }}>Ações</th>
                </tr>
              </thead>
              <tbody>
                {filtrados.length === 0 ? (
                  <tr>
                    <td colSpan={7} style={{ textAlign: 'center', color: 'var(--text-faint)', padding: 40, fontFamily: 'var(--font-mono)', fontSize: 12 }}>
                      <span className="material-symbols-outlined" style={{ fontSize: 48, marginBottom: 12 }}>inbox</span>
                      <div>Nenhum pedido encontrado</div>
                    </td>
                  </tr>
                ) : (
                  filtrados.map(p => (
                    <tr key={p.id} style={{ borderBottom: '1px solid var(--border)' }}>
                      <td style={{ padding: 16 }}>
                        <span className="mono text-cyan" style={{ fontFamily: 'monospace', fontSize: 12, color: 'var(--cyan)' }}>
                          #{p.id?.slice(0, 8).toUpperCase()}
                        </span>
                      </td>
                      <td style={{ padding: 16, fontSize: 13, fontWeight: 500 }}>{p.cliente}</td>
                      <td style={{ padding: 16, color: 'var(--text-muted)', fontSize: 12 }}>{p.email}</td>
                      <td style={{ padding: 16, color: 'var(--text-muted)', fontSize: 12 }}>{formatarData(p.created_at)}</td>
                      <td style={{ padding: 16, textAlign: 'right', fontFamily: 'monospace', fontWeight: 700, color: 'var(--cyan)', fontSize: 13 }}>
                        {p.valor_total?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                      </td>
                      <td style={{ padding: 16, textAlign: 'center' }}>
                        <span className={`badge ${statusMap[p.status] || 'badge-amber'}`} style={{ fontSize: 11 }}>
                          {statusLabel[p.status] || p.status}
                        </span>
                      </td>
                      <td style={{ padding: 16, textAlign: 'center' }}>
                        <button 
                          className="icon-btn" 
                          onClick={() => verDetalhes(p)}
                          style={{ background: 'none', border: 'none', cursor: 'pointer', color: 'var(--cyan)' }}
                          title="Ver detalhes"
                        >
                          <span className="material-symbols-outlined" style={{ fontSize: 20 }}>visibility</span>
                        </button>
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        )}

        <div style={{ padding: '14px 20px', borderTop: '1px solid var(--border)', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <span style={{ fontFamily: 'monospace', fontSize: 11, color: 'var(--text-faint)' }}>
            Exibindo <span style={{ color: 'var(--cyan)' }}>{filtrados.length}</span> de <span style={{ color: 'var(--text)' }}>{pedidos.length}</span> registros
          </span>
          {filtrados.length === pedidos.length && pedidos.length > 0 && (
            <span style={{ fontFamily: 'monospace', fontSize: 10, color: 'var(--green)' }}>
              ✅ Dados atualizados em tempo real
            </span>
          )}
        </div>
      </div>

      <style>{`
        @keyframes spin {
          from { transform: rotate(0deg); }
          to { transform: rotate(360deg); }
        }
      `}</style>
    </div>
  )
}