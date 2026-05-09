import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { criarPedido } from '../services/api'
import { produtos as catalogo } from '../data/produtos'

export default function NovoPedido() {
  const navigate = useNavigate()
  const [cliente, setCliente] = useState({ nome: '', documento: '', email: '' })
  const [itens, setItens] = useState([])
  const [busca, setBusca] = useState('')
  const [loading, setLoading] = useState(false)
  const [erro, setErro] = useState('')

  const produtosFiltrados = catalogo.filter(p =>
    !busca || p.nome.toLowerCase().includes(busca.toLowerCase())
  ).slice(0, 8)

  const adicionarItem = (produto) => {
    setItens(prev => {
      const existe = prev.find(i => i.id === produto.id)
      if (existe) return prev.map(i => i.id === produto.id ? { ...i, quantidade: i.quantidade + 1 } : i)
      return [...prev, { ...produto, quantidade: 1 }]
    })
  }

  const removerItem = (id) => setItens(prev => prev.filter(i => i.id !== id))

  const alterarQtd = (id, delta) => {
    setItens(prev => prev
      .map(i => i.id === id ? { ...i, quantidade: i.quantidade + delta } : i)
      .filter(i => i.quantidade > 0)
    )
  }

  const total = itens.reduce((acc, i) => acc + i.preco * i.quantidade, 0)

  const handleSubmit = async () => {
    setErro('')
    if (!cliente.nome || !cliente.documento || !cliente.email) {
      setErro('Preencha todos os dados do cliente.')
      return
    }
    if (itens.length === 0) {
      setErro('Adicione pelo menos um produto ao pedido.')
      return
    }
    setLoading(true)
    try {
      const payload = {
        cliente: cliente.nome,
        documento: cliente.documento,
        email: cliente.email,
        produtos: itens.map(i => ({ nome: i.nome, preco: i.preco, quantidade: i.quantidade })),
      }
      const result = await criarPedido(payload)
      navigate('/confirmacao', { state: { pedido: result } })
    } catch (e) {
      setErro(e.message || 'Erro ao conectar com o servidor.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 28 }}>
      {/* Header */}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end' }}>
        <div>
          <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-faint)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: 6 }}>
            CRIAÇÃO DE PEDIDO
          </div>
          <h1 style={{ fontFamily: 'var(--font-display)', fontSize: 28, fontWeight: 800, letterSpacing: '-0.02em' }}>
            Nova Ordem de Serviço
          </h1>
        </div>
        <div style={{ display: 'flex', gap: 10 }}>
          <button className="btn-ghost" onClick={() => navigate('/pedidos')}>Cancelar</button>
          <button className="btn-primary" onClick={handleSubmit} disabled={loading}>
            {loading
              ? <span className="material-symbols-outlined spin" style={{ fontSize: 18 }}>autorenew</span>
              : <span className="material-symbols-outlined" style={{ fontSize: 18 }}>send</span>
            }
            {loading ? 'Enviando...' : 'Transmitir Pedido'}
          </button>
        </div>
      </div>

      {erro && (
        <div style={{ background: 'rgba(255,82,82,0.08)', border: '1px solid rgba(255,82,82,0.25)', borderRadius: 'var(--radius)', padding: '12px 16px', color: 'var(--red)', fontSize: 13, display: 'flex', gap: 10, alignItems: 'center' }}>
          <span className="material-symbols-outlined" style={{ fontSize: 20 }}>error</span>
          {erro}
        </div>
      )}

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 380px', gap: 24 }}>
        {/* Left */}
        <div style={{ display: 'flex', flexDirection: 'column', gap: 20 }}>
          {/* Client data */}
          <div className="card" style={{ padding: 24 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 20, paddingBottom: 16, borderBottom: '1px solid var(--border)' }}>
              <span className="material-symbols-outlined text-cyan">person</span>
              <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.05em', fontSize: 13 }}>Dados do Cliente</span>
            </div>
            <div className="grid-2">
              <div>
                <label className="input-label">Nome / Razão Social</label>
                <input className="input-field" placeholder="ex: João da Silva Ltda." value={cliente.nome} onChange={e => setCliente(c => ({ ...c, nome: e.target.value }))} />
              </div>
              <div>
                <label className="input-label">CPF / CNPJ</label>
                <input className="input-field" placeholder="000.000.000-00" value={cliente.documento} onChange={e => setCliente(c => ({ ...c, documento: e.target.value }))} />
              </div>
              <div style={{ gridColumn: '1 / -1' }}>
                <label className="input-label">E-mail de Faturamento</label>
                <input className="input-field" type="email" placeholder="cliente@empresa.com" value={cliente.email} onChange={e => setCliente(c => ({ ...c, email: e.target.value }))} />
              </div>
            </div>
          </div>

          {/* Product search */}
          <div className="card" style={{ padding: 24 }}>
            <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 16, paddingBottom: 16, borderBottom: '1px solid var(--border)' }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                <span className="material-symbols-outlined text-cyan">inventory_2</span>
                <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.05em', fontSize: 13 }}>Adicionar Produtos</span>
              </div>
              <span style={{ fontFamily: 'var(--font-mono)', fontSize: 10, color: 'var(--text-faint)' }}>{catalogo.length} produtos disponíveis</span>
            </div>
            <div style={{ position: 'relative', marginBottom: 16 }}>
              <span className="material-symbols-outlined" style={{ position: 'absolute', left: 10, top: '50%', transform: 'translateY(-50%)', fontSize: 18, color: 'var(--text-faint)' }}>search</span>
              <input className="input-field" style={{ paddingLeft: 36 }} placeholder="Buscar produto pelo nome..." value={busca} onChange={e => setBusca(e.target.value)} />
            </div>
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 10 }}>
              {produtosFiltrados.map(p => (
                <button
                  key={p.id}
                  onClick={() => adicionarItem(p)}
                  style={{
                    background: 'var(--bg-elevated)',
                    border: '1px solid var(--border)',
                    borderRadius: 'var(--radius)',
                    padding: '12px 14px',
                    cursor: 'pointer',
                    textAlign: 'left',
                    transition: 'all 0.15s',
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    gap: 8,
                  }}
                  onMouseEnter={e => { e.currentTarget.style.borderColor = 'var(--border-hover)'; e.currentTarget.style.background = 'var(--bg-high)' }}
                  onMouseLeave={e => { e.currentTarget.style.borderColor = 'var(--border)'; e.currentTarget.style.background = 'var(--bg-elevated)' }}
                >
                  <div>
                    <div style={{ fontSize: 13, fontWeight: 600, color: 'var(--text)', marginBottom: 2 }}>{p.nome}</div>
                    <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--cyan)' }}>
                      {p.preco.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                    </div>
                  </div>
                  <span className="material-symbols-outlined" style={{ fontSize: 20, color: 'var(--text-faint)' }}>add_circle</span>
                </button>
              ))}
            </div>
          </div>

          {/* Items in order */}
          {itens.length > 0 && (
            <div className="card" style={{ padding: 24 }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 16, paddingBottom: 16, borderBottom: '1px solid var(--border)' }}>
                <span className="material-symbols-outlined text-cyan">shopping_cart</span>
                <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.05em', fontSize: 13 }}>
                  Itens do Pedido ({itens.length})
                </span>
              </div>
              <div style={{ display: 'flex', flexDirection: 'column', gap: 10 }}>
                {itens.map(item => (
                  <div key={item.id} style={{ display: 'flex', alignItems: 'center', gap: 14, padding: '12px 14px', background: 'var(--bg-elevated)', borderRadius: 'var(--radius)', border: '1px solid var(--border)' }}>
                    <div style={{ flex: 1 }}>
                      <div style={{ fontSize: 13, fontWeight: 600 }}>{item.nome}</div>
                      <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-muted)' }}>
                        {item.preco.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })} / un
                      </div>
                    </div>
                    <div style={{ display: 'flex', alignItems: 'center', gap: 2, border: '1px solid var(--border)', borderRadius: 'var(--radius)', overflow: 'hidden' }}>
                      <button onClick={() => alterarQtd(item.id, -1)} style={{ background: 'none', border: 'none', padding: '6px 10px', cursor: 'pointer', color: 'var(--text-muted)', fontSize: 18, lineHeight: 1 }}>−</button>
                      <span style={{ minWidth: 28, textAlign: 'center', fontFamily: 'var(--font-mono)', fontSize: 13, color: 'var(--cyan)', fontWeight: 700 }}>{item.quantidade}</span>
                      <button onClick={() => alterarQtd(item.id, 1)} style={{ background: 'none', border: 'none', padding: '6px 10px', cursor: 'pointer', color: 'var(--text-muted)', fontSize: 18, lineHeight: 1 }}>+</button>
                    </div>
                    <div style={{ fontFamily: 'var(--font-mono)', fontWeight: 700, color: 'var(--cyan)', minWidth: 90, textAlign: 'right' }}>
                      {(item.preco * item.quantidade).toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                    </div>
                    <button className="icon-btn" onClick={() => removerItem(item.id)} style={{ color: 'var(--red)' }}>
                      <span className="material-symbols-outlined" style={{ fontSize: 18 }}>delete</span>
                    </button>
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>

        {/* Right — Summary */}
        <div style={{ position: 'sticky', top: 80, display: 'flex', flexDirection: 'column', gap: 16 }}>
          <div className="card" style={{ overflow: 'hidden' }}>
            <div style={{ padding: '16px 20px', background: 'linear-gradient(135deg, var(--cyan) 0%, #0099cc 100%)', color: 'var(--bg)' }}>
              <div style={{ fontFamily: 'var(--font-display)', fontWeight: 800, fontSize: 15, textTransform: 'uppercase', letterSpacing: '0.05em' }}>Resumo do Pedido</div>
              <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, opacity: 0.7, marginTop: 2 }}>TX: #ORD-NOVO</div>
            </div>
            <div style={{ padding: '20px', display: 'flex', flexDirection: 'column', gap: 14 }}>
              <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13, color: 'var(--text-muted)' }}>
                <span>Subtotal</span>
                <span style={{ fontFamily: 'var(--font-mono)' }}>{total.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</span>
              </div>
              <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13, color: 'var(--text-muted)' }}>
                <span>Frete</span>
                <span style={{ fontFamily: 'var(--font-mono)', color: 'var(--cyan)', fontStyle: 'italic' }}>A calcular</span>
              </div>
              <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13, color: 'var(--text-muted)' }}>
                <span>Desconto</span>
                <span style={{ fontFamily: 'var(--font-mono)', color: 'var(--red)' }}>-R$ 0,00</span>
              </div>
              <div style={{ borderTop: '1px solid var(--border)', paddingTop: 14 }}>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: 10, color: 'var(--cyan)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: 4 }}>Valor Total</div>
                <div style={{ fontFamily: 'var(--font-display)', fontSize: 32, fontWeight: 800, color: 'var(--cyan)', letterSpacing: '-0.02em', textShadow: '0 0 20px rgba(0,210,255,0.4)' }}>
                  {total.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                </div>
              </div>
            </div>

            <div style={{ padding: '12px 20px', borderTop: '1px solid var(--border)', background: 'var(--bg-elevated)', display: 'flex', gap: 8, alignItems: 'center' }}>
              <span className="material-symbols-outlined text-cyan spin" style={{ fontSize: 18 }}>sync</span>
              <span style={{ fontFamily: 'var(--font-mono)', fontSize: 10, color: 'var(--text-muted)', letterSpacing: '0.05em' }}>
                Conectado ao servidor <span style={{ color: 'var(--cyan)', fontWeight: 700 }}>ATIVO</span>
              </span>
            </div>
          </div>

          {/* Payment */}
          <div className="card" style={{ padding: 20 }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: 10, textTransform: 'uppercase', letterSpacing: '0.1em', color: 'var(--text-muted)', marginBottom: 14 }}>Forma de Pagamento</div>
            <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
              {['Transferência Bancária', 'Boleto 30 dias', 'Cartão de Crédito'].map((m, i) => (
                <label key={m} style={{ display: 'flex', alignItems: 'center', gap: 10, padding: '10px 12px', border: '1px solid var(--border)', borderRadius: 'var(--radius)', cursor: 'pointer', fontSize: 13 }}>
                  <input type="radio" name="pagamento" defaultChecked={i === 0} style={{ accentColor: 'var(--cyan)' }} />
                  {m}
                </label>
              ))}
            </div>
          </div>

          <button
            className="btn-primary"
            style={{ width: '100%', justifyContent: 'center', padding: '14px' }}
            onClick={handleSubmit}
            disabled={loading}
          >
            {loading
              ? <><span className="material-symbols-outlined spin" style={{ fontSize: 18 }}>autorenew</span> Enviando...</>
              : <><span className="material-symbols-outlined" style={{ fontSize: 18 }}>send</span> Confirmar Pedido</>
            }
          </button>
        </div>
      </div>
    </div>
  )
}
