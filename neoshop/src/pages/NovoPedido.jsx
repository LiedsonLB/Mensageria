import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { criarPedido } from '../services/api'
import { produtos as catalogo, categorias, filtrarProdutos } from '../data/produtos'

const categoryIcons = {
  notebooks: '💻', perifericos: '🖱️', hardware: '🔧',
  monitores: '🖥️', audio: '🎧', armazenamento: '💾', redes: '🌐',
}

// Chave para salvar no localStorage
const STORAGE_KEY = 'neoShop_dadosCliente'

export default function NovoPedido() {
  const navigate = useNavigate()

  // Carregar dados do cliente do localStorage ao iniciar
  const loadClienteData = () => {
    const savedData = localStorage.getItem(STORAGE_KEY)
    if (savedData) {
      try {
        return JSON.parse(savedData)
      } catch (e) {
        console.error('Erro ao carregar dados do cliente:', e)
        return { nome: '', documento: '', email: '' }
      }
    }
    return { nome: '', documento: '', email: '' }
  }

  const [cliente, setCliente] = useState(loadClienteData)
  const [itens, setItens] = useState([])
  const [busca, setBusca] = useState('')
  const [categoria, setCategoria] = useState('todos')
  const [loading, setLoading] = useState(false)
  const [erro, setErro] = useState('')
  const [mostrarNotificacao, setMostrarNotificacao] = useState(false)

  const produtosFiltrados = filtrarProdutos(busca, categoria)

  // Salvar dados do cliente no localStorage sempre que mudar
  const salvarDadosCliente = (novosDados) => {
    setCliente(novosDados)
    localStorage.setItem(STORAGE_KEY, JSON.stringify(novosDados))

    // Mostrar notificação de salvamento
    setMostrarNotificacao(true)
    setTimeout(() => setMostrarNotificacao(false), 2000)
  }

  const handleClienteChange = (campo, valor) => {
    const novosDados = { ...cliente, [campo]: valor }
    salvarDadosCliente(novosDados)
  }

  const limparDadosCliente = () => {
    const dadosVazios = { nome: '', documento: '', email: '' }
    salvarDadosCliente(dadosVazios)
  }

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

      console.log('ID completo do pedido:', result.id)

      // Opcional: Limpar o carrinho após enviar o pedido
      setItens([])

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

      {/* Notificação de salvamento */}
      {mostrarNotificacao && (
        <div style={{
          position: 'fixed',
          top: 20,
          right: 20,
          background: 'var(--green)',
          color: '#000',
          padding: '10px 16px',
          borderRadius: 'var(--radius)',
          display: 'flex',
          alignItems: 'center',
          gap: 8,
          zIndex: 1000,
          animation: 'slideIn 0.3s ease',
          fontSize: 13
        }}>
          <span className="material-symbols-outlined" style={{ fontSize: 18 }}>save</span>
          Dados do cliente salvos automaticamente
        </div>
      )}

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
            <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 20, paddingBottom: 16, borderBottom: '1px solid var(--border)' }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                <span className="material-symbols-outlined text-cyan">person</span>
                <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.05em', fontSize: 13 }}>Dados do Cliente</span>
              </div>
              <div style={{ display: 'flex', gap: 8 }}>
                <button
                  className="icon-btn"
                  onClick={limparDadosCliente}
                  style={{ fontSize: 12, padding: '4px 8px' }}
                  title="Limpar dados salvos"
                >
                  <span className="material-symbols-outlined" style={{ fontSize: 16 }}>delete_sweep</span>
                </button>
                {cliente.nome && (
                  <span style={{ fontSize: 10, color: 'var(--green)', display: 'flex', alignItems: 'center', gap: 4 }}>
                    <span className="material-symbols-outlined" style={{ fontSize: 12 }}>check_circle</span>
                    Salvo
                  </span>
                )}
              </div>
            </div>
            <div className="grid-2">
              <div>
                <label className="input-label">Nome / Razão Social</label>
                <input
                  className="input-field"
                  placeholder="ex: João da Silva Ltda."
                  value={cliente.nome}
                  onChange={e => handleClienteChange('nome', e.target.value)}
                />
              </div>
              <div>
                <label className="input-label">CPF / CNPJ</label>
                <input
                  className="input-field"
                  placeholder="000.000.000-00"
                  value={cliente.documento}
                  onChange={e => handleClienteChange('documento', e.target.value)}
                />
              </div>
              <div style={{ gridColumn: '1 / -1' }}>
                <label className="input-label">E-mail de Faturamento</label>
                <input
                  className="input-field"
                  type="email"
                  placeholder="cliente@empresa.com"
                  value={cliente.email}
                  onChange={e => handleClienteChange('email', e.target.value)}
                />
              </div>
            </div>
            <div style={{ marginTop: 12, fontSize: 10, color: 'var(--text-faint)', textAlign: 'center' }}>
              <span className="material-symbols-outlined" style={{ fontSize: 12, verticalAlign: 'middle' }}>info</span>
              Seus dados são salvos localmente para agilizar próximos pedidos
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
                <button
                  className="btn-ghost"
                  style={{ marginLeft: 'auto', padding: '4px 8px', fontSize: 11 }}
                  onClick={() => setItens([])}
                >
                  <span className="material-symbols-outlined" style={{ fontSize: 14 }}>delete_sweep</span>
                  Limpar tudo
                </button>
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

          {/* Product selection with catalog style */}
          <div className="card" style={{ padding: 24 }}>
            <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 20, paddingBottom: 16, borderBottom: '1px solid var(--border)' }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                <span className="material-symbols-outlined text-cyan">inventory_2</span>
                <span style={{ fontFamily: 'var(--font-display)', fontWeight: 700, textTransform: 'uppercase', letterSpacing: '0.05em', fontSize: 13 }}>Selecionar Produtos</span>
              </div>
              <span style={{ fontFamily: 'var(--font-mono)', fontSize: 10, color: 'var(--text-faint)' }}>{catalogo.length} produtos disponíveis</span>
            </div>

            {/* Search and Filters */}
            <div style={{ display: 'flex', gap: 12, alignItems: 'center', flexWrap: 'wrap', marginBottom: 20 }}>
              <div style={{ position: 'relative', flex: 1, minWidth: 200 }}>
                <span className="material-symbols-outlined" style={{ position: 'absolute', left: 10, top: '50%', transform: 'translateY(-50%)', fontSize: 18, color: 'var(--text-faint)' }}>search</span>
                <input
                  className="input-field"
                  style={{ paddingLeft: 36 }}
                  placeholder="Buscar produto..."
                  value={busca}
                  onChange={e => setBusca(e.target.value)}
                />
              </div>
              <div style={{ display: 'flex', gap: 6, flexWrap: 'wrap' }}>
                {categorias.map(c => (
                  <button
                    key={c.id}
                    onClick={() => setCategoria(c.id)}
                    style={{
                      padding: '6px 12px',
                      borderRadius: 'var(--radius)',
                      border: '1px solid',
                      borderColor: categoria === c.id ? 'var(--cyan)' : 'var(--border)',
                      background: categoria === c.id ? 'var(--cyan-dim)' : 'none',
                      color: categoria === c.id ? 'var(--cyan)' : 'var(--text-muted)',
                      fontSize: 11,
                      fontFamily: 'var(--font-body)',
                      cursor: 'pointer',
                      transition: 'all 0.15s',
                      display: 'flex',
                      alignItems: 'center',
                      gap: 4,
                    }}
                  >
                    <span className="material-symbols-outlined" style={{ fontSize: 14 }}>{c.icon}</span>
                    {c.label}
                  </button>
                ))}
              </div>
            </div>

            {/* Products Grid */}
            {produtosFiltrados.length === 0 ? (
              <div style={{ textAlign: 'center', padding: 40, color: 'var(--text-faint)' }}>
                <span className="material-symbols-outlined" style={{ fontSize: 36, display: 'block', marginBottom: 8 }}>search_off</span>
                <p style={{ fontSize: 12 }}>Nenhum produto encontrado</p>
              </div>
            ) : (
              <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(260px, 1fr))', gap: 12 }}>
                {produtosFiltrados.map(p => (
                  <div
                    key={p.id}
                    className="card"
                    style={{
                      padding: 14,
                      cursor: 'pointer',
                      transition: 'all 0.15s',
                      border: '1px solid var(--border)',
                    }}
                    onClick={() => adicionarItem(p)}
                    onMouseEnter={e => { e.currentTarget.style.borderColor = 'var(--cyan)'; e.currentTarget.style.transform = 'translateY(-2px)' }}
                    onMouseLeave={e => { e.currentTarget.style.borderColor = 'var(--border)'; e.currentTarget.style.transform = 'translateY(0)' }}
                  >
                    <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 10 }}>
                      <div style={{ width: 40, height: 40, borderRadius: 'var(--radius)', background: 'var(--bg-elevated)', border: '1px solid var(--border)', display: 'flex', alignItems: 'center', justifyContent: 'center', overflow: 'hidden' }}>
                        {p.image ? (
                          <img src={p.image} alt={p.nome} style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
                        ) : (
                          <span style={{ fontSize: 20 }}>{categoryIcons[p.categoria] || '📦'}</span>
                        )}
                      </div>
                      <div style={{ flex: 1 }}>
                        <div style={{ fontWeight: 700, fontSize: 13, marginBottom: 2 }}>{p.nome}</div>
                        <div style={{ fontSize: 10, color: 'var(--text-faint)' }}>
                          {categorias.find(c => c.id === p.categoria)?.label}
                        </div>
                      </div>
                    </div>

                    <p style={{ fontSize: 11, color: 'var(--text-muted)', marginBottom: 10, lineHeight: 1.4 }}>{p.descricao.substring(0, 60)}...</p>

                    <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 8 }}>
                      <span className="material-symbols-outlined" style={{ fontSize: 12, color: 'var(--amber)', fontVariationSettings: "'FILL' 1" }}>star</span>
                      <span style={{ fontSize: 11, fontWeight: 600 }}>{p.rating}</span>
                      <span style={{ fontSize: 10, color: 'var(--text-faint)' }}>({p.reviews})</span>
                      <span style={{ marginLeft: 'auto', fontSize: 10, color: p.estoque > 10 ? 'var(--green)' : p.estoque > 0 ? 'var(--amber)' : 'var(--red)' }}>
                        {p.estoque > 0 ? `${p.estoque} un` : 'Sem estoque'}
                      </span>
                    </div>

                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', borderTop: '1px solid var(--border)', paddingTop: 10 }}>
                      <div style={{ fontFamily: 'var(--font-display)', fontSize: 16, fontWeight: 800, color: 'var(--cyan)' }}>
                        {p.preco.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                      </div>
                      <button
                        className="btn-primary"
                        style={{ padding: '4px 10px', fontSize: 11 }}
                        onClick={(e) => {
                          e.stopPropagation()
                          adicionarItem(p)
                        }}
                      >
                        <span className="material-symbols-outlined" style={{ fontSize: 14 }}>add</span>
                        Adicionar
                      </button>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
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

      <style>{`
        @keyframes spin {
          from { transform: rotate(0deg); }
          to { transform: rotate(360deg); }
        }
        @keyframes slideIn {
          from {
            transform: translateX(100%);
            opacity: 0;
          }
          to {
            transform: translateX(0);
            opacity: 1;
          }
        }
        .spin {
          animation: spin 1s linear infinite;
        }
      `}</style>
    </div>
  )
}