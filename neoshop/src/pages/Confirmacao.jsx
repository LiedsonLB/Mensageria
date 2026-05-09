import React from 'react'
import { useLocation, useNavigate } from 'react-router-dom'

const statusMap = {
  PENDENTE: { badge: 'badge-amber', label: 'Pendente' },
  PROCESSANDO: { badge: 'badge-cyan', label: 'Processando' },
  CONCLUIDO: { badge: 'badge-green', label: 'Concluído' },
  CANCELADO: { badge: 'badge-red', label: 'Cancelado' },
}

const etapas = [
  { id: 'received', label: 'Pedido Recebido', icon: 'inventory', desc: 'O sistema NeoShop confirmou o recebimento da solicitação.', status: 'done' },
  { id: 'payment', label: 'Pagamento em Análise', icon: 'payments', desc: 'Validação de crédito e conciliação bancária em andamento.', status: 'done' },
  { id: 'invoice', label: 'Nota Fiscal sendo emitida', icon: 'description', desc: 'A NF-e está sendo assinada digitalmente pelos nossos servidores.', status: 'active' },
  { id: 'email', label: 'Envio por E-mail', icon: 'mail', desc: 'O pacote digital completo será enviado ao e-mail cadastrado.', status: 'pending' },
]

export default function Confirmacao() {
  const location = useLocation()
  const navigate = useNavigate()
  const pedido = location.state?.pedido

  if (!pedido) {
    return (
      <div style={{ textAlign: 'center', padding: 60 }}>
        <p style={{ color: 'var(--text-muted)', marginBottom: 20 }}>Nenhum pedido encontrado.</p>
        <button className="btn-primary" onClick={() => navigate('/novo-pedido')}>Criar Novo Pedido</button>
      </div>
    )
  }

  const status = statusMap[pedido.status] || statusMap['PENDENTE']

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 32 }}>
      {/* Success Hero */}
      <div style={{ textAlign: 'center', padding: '40px 0 24px' }}>
        <div className="success-icon-wrap" style={{ marginBottom: 20 }}>
          <span className="material-symbols-outlined" style={{ fontSize: 44, color: 'var(--green)', fontVariationSettings: "'FILL' 1" }}>check_circle</span>
        </div>
        <h1 style={{ fontFamily: 'var(--font-display)', fontSize: 36, fontWeight: 800, letterSpacing: '-0.03em', marginBottom: 10 }}>
          Pedido criado com sucesso!
        </h1>
        <p style={{ color: 'var(--text-muted)', maxWidth: 500, margin: '0 auto', fontSize: 14, lineHeight: 1.7 }}>
          Sua ordem de serviço foi processada pelo sistema NeoShop Enterprise. Os detalhes e o rastreamento estão disponíveis abaixo.
        </p>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '340px 1fr', gap: 24 }}>
        {/* Invoice card */}
        <div className="card" style={{ padding: 24, display: 'flex', flexDirection: 'column', gap: 20 }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
            <div>
              <div style={{ fontFamily: 'var(--font-mono)', fontSize: 10, textTransform: 'uppercase', letterSpacing: '0.1em', color: 'var(--text-faint)' }}>ID do Pedido</div>
              <div style={{ fontFamily: 'var(--font-display)', fontSize: 18, fontWeight: 800, color: 'var(--cyan)', letterSpacing: '-0.02em', marginTop: 4 }}>
                #{pedido.id?.slice(0, 8).toUpperCase()}
              </div>
            </div>
            <span className={`badge ${status.badge}`}>{status.label}</span>
          </div>

          <div style={{ background: 'var(--bg-elevated)', borderRadius: 'var(--radius)', padding: 16, display: 'flex', flexDirection: 'column', gap: 12, border: '1px solid var(--border)' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13 }}>
              <span style={{ color: 'var(--text-muted)' }}>Cliente</span>
              <span style={{ fontWeight: 600, maxWidth: 160, textAlign: 'right', fontSize: 12 }}>{pedido.cliente}</span>
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13 }}>
              <span style={{ color: 'var(--text-muted)' }}>Documento</span>
              <span style={{ fontFamily: 'var(--font-mono)', fontSize: 12 }}>{pedido.documento}</span>
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13 }}>
              <span style={{ color: 'var(--text-muted)' }}>E-mail</span>
              <span style={{ fontSize: 12, color: 'var(--cyan)' }}>{pedido.email}</span>
            </div>
            <div style={{ borderTop: '1px solid var(--border)', paddingTop: 12, display: 'flex', justifyContent: 'space-between' }}>
              <span style={{ fontWeight: 700 }}>Total do Pedido</span>
              <span style={{ fontFamily: 'var(--font-mono)', fontWeight: 800, color: 'var(--cyan)', fontSize: 16 }}>
                {pedido.valor_total.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
              </span>
            </div>
          </div>

          {/* Products */}
          <div>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: 10, textTransform: 'uppercase', letterSpacing: '0.1em', color: 'var(--text-faint)', marginBottom: 10 }}>Produtos</div>
            <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
              {pedido.produtos?.map((p, i) => (
                <div key={i} style={{ display: 'flex', justifyContent: 'space-between', fontSize: 12, padding: '8px 12px', background: 'var(--bg-elevated)', borderRadius: 'var(--radius)', border: '1px solid var(--border)' }}>
                  <span style={{ color: 'var(--text-muted)' }}>{p.nome} <span style={{ color: 'var(--text-faint)' }}>×{p.quantidade}</span></span>
                  <span style={{ fontFamily: 'var(--font-mono)', color: 'var(--cyan)' }}>
                    {(p.preco * p.quantidade).toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                  </span>
                </div>
              ))}
            </div>
          </div>

          <div style={{ display: 'flex', flexDirection: 'column', gap: 8, marginTop: 'auto' }}>
            <button className="btn-primary" style={{ justifyContent: 'center' }}>
              <span className="material-symbols-outlined" style={{ fontSize: 18 }}>download</span>
              Baixar XML
            </button>
            <button className="btn-ghost" style={{ justifyContent: 'center' }}>
              <span className="material-symbols-outlined" style={{ fontSize: 18 }}>print</span>
              Imprimir Recibo
            </button>
          </div>
        </div>

        {/* Status tracker */}
        <div className="card" style={{ padding: 28 }}>
          <div style={{ fontFamily: 'var(--font-display)', fontWeight: 700, fontSize: 18, marginBottom: 28 }}>Status do Processamento</div>
          <div style={{ position: 'relative', paddingLeft: 56 }}>
            {/* Line */}
            <div style={{ position: 'absolute', left: 23, top: 0, bottom: 0, width: 2, background: 'var(--border)' }} />
            <div style={{ position: 'absolute', left: 23, top: 0, height: '50%', width: 2, background: 'var(--green)', boxShadow: '0 0 8px var(--green)' }} />

            <div style={{ display: 'flex', flexDirection: 'column', gap: 32 }}>
              {etapas.map(etapa => (
                <div key={etapa.id} style={{ opacity: etapa.status === 'pending' ? 0.35 : 1 }}>
                  <div style={{ position: 'absolute', left: 0, width: 46, height: 46, borderRadius: '50%', display: 'flex', alignItems: 'center', justifyContent: 'center', zIndex: 1,
                    background: etapa.status === 'done' ? 'var(--green)' : etapa.status === 'active' ? 'var(--cyan-dim)' : 'var(--bg-elevated)',
                    border: `2px solid ${etapa.status === 'done' ? 'var(--green)' : etapa.status === 'active' ? 'var(--cyan)' : 'var(--border)'}`,
                    boxShadow: etapa.status === 'done' ? '0 0 12px rgba(0,229,160,0.4)' : etapa.status === 'active' ? '0 0 15px rgba(0,210,255,0.4)' : 'none',
                    color: etapa.status === 'done' ? 'var(--bg)' : etapa.status === 'active' ? 'var(--cyan)' : 'var(--text-faint)',
                  }}>
                    <span className={`material-symbols-outlined ${etapa.status === 'active' ? 'pulse-glow' : ''}`} style={{ fontSize: 22, fontVariationSettings: etapa.status === 'done' ? "'FILL' 1" : "'FILL' 0" }}>
                      {etapa.icon}
                    </span>
                  </div>
                  <div style={{ paddingBottom: 4 }}>
                    <div style={{ fontWeight: 700, fontSize: 14, color: etapa.status === 'active' ? 'var(--cyan)' : 'var(--text)' }}>{etapa.label}</div>
                    <div style={{ fontSize: 12, color: 'var(--text-muted)', marginTop: 4, lineHeight: 1.5 }}>{etapa.desc}</div>
                    <div style={{ fontSize: 11, fontFamily: 'var(--font-mono)', marginTop: 6, color: etapa.status === 'active' ? 'var(--cyan)' : 'var(--text-faint)' }}>
                      {etapa.status === 'done' ? `Hoje, ${new Date().getHours()}:${String(new Date().getMinutes()).padStart(2,'0')}` : etapa.status === 'active' ? 'Em andamento...' : 'Aguardando etapa anterior'}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* Actions */}
      <div style={{ display: 'flex', gap: 12, justifyContent: 'center' }}>
        <button className="btn-primary" onClick={() => navigate('/novo-pedido')}>
          <span className="material-symbols-outlined" style={{ fontSize: 18 }}>add</span>
          Novo Pedido
        </button>
        <button className="btn-ghost" onClick={() => navigate('/pedidos')}>
          <span className="material-symbols-outlined" style={{ fontSize: 18 }}>list</span>
          Ver Todos os Pedidos
        </button>
      </div>
    </div>
  )
}
