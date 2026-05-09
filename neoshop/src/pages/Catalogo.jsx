import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { produtos, categorias, filtrarProdutos } from '../data/produtos'

const categoryIcons = {
  notebooks: '💻', perifericos: '🖱️', hardware: '🔧',
  monitores: '🖥️', audio: '🎧', armazenamento: '💾', redes: '🌐',
}

export default function Catalogo() {
  const navigate = useNavigate()
  const [busca, setBusca] = useState('')
  const [categoria, setCategoria] = useState('todos')

  const lista = filtrarProdutos(busca, categoria)

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 28 }}>
      {/* Header */}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end' }}>
        <div>
          <div style={{ fontFamily: 'var(--font-mono)', fontSize: 11, color: 'var(--text-faint)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: 6 }}>CATÁLOGO</div>
          <h1 style={{ fontFamily: 'var(--font-display)', fontSize: 28, fontWeight: 800, letterSpacing: '-0.02em' }}>Produtos Disponíveis</h1>
          <p style={{ color: 'var(--text-muted)', fontSize: 13, marginTop: 4 }}>{produtos.length} produtos em estoque · Pronto para integrar com API</p>
        </div>
        <button className="btn-primary" onClick={() => navigate('/novo-pedido')}>
          <span className="material-symbols-outlined" style={{ fontSize: 18 }}>add_shopping_cart</span>
          Criar Pedido
        </button>
      </div>

      {/* Filters */}
      <div style={{ display: 'flex', gap: 12, alignItems: 'center', flexWrap: 'wrap' }}>
        <div style={{ position: 'relative', flex: 1, minWidth: 220, maxWidth: 400 }}>
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
                padding: '7px 14px',
                borderRadius: 'var(--radius)',
                border: '1px solid',
                borderColor: categoria === c.id ? 'var(--cyan)' : 'var(--border)',
                background: categoria === c.id ? 'var(--cyan-dim)' : 'none',
                color: categoria === c.id ? 'var(--cyan)' : 'var(--text-muted)',
                fontSize: 12,
                fontFamily: 'var(--font-body)',
                cursor: 'pointer',
                transition: 'all 0.15s',
                display: 'flex',
                alignItems: 'center',
                gap: 6,
              }}
            >
              <span className="material-symbols-outlined" style={{ fontSize: 16 }}>{c.icon}</span>
              {c.label}
            </button>
          ))}
        </div>
      </div>

      {/* API notice */}
      <div style={{ background: 'rgba(0,210,255,0.05)', border: '1px solid rgba(0,210,255,0.15)', borderRadius: 'var(--radius)', padding: '10px 16px', display: 'flex', alignItems: 'center', gap: 10 }}>
        <span className="material-symbols-outlined text-cyan" style={{ fontSize: 18 }}>info</span>
        <span style={{ fontSize: 12, color: 'var(--text-muted)' }}>
          Dados mockados para desenvolvimento. Substitua a função <code style={{ fontFamily: 'var(--font-mono)', color: 'var(--cyan)', background: 'rgba(0,210,255,0.08)', padding: '1px 6px', borderRadius: 4 }}>filtrarProdutos()</code> em <code style={{ fontFamily: 'var(--font-mono)', color: 'var(--cyan)', background: 'rgba(0,210,255,0.08)', padding: '1px 6px', borderRadius: 4 }}>src/data/produtos.js</code> por uma chamada à sua API.
        </span>
      </div>

      {/* Product grid */}
      {lista.length === 0 ? (
        <div style={{ textAlign: 'center', padding: 60, color: 'var(--text-faint)' }}>
          <span className="material-symbols-outlined" style={{ fontSize: 48, display: 'block', marginBottom: 12 }}>search_off</span>
          <p style={{ fontFamily: 'var(--font-mono)', fontSize: 13 }}>Nenhum produto encontrado para "{busca}"</p>
        </div>
      ) : (
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(280px, 1fr))', gap: 16 }}>
          {lista.map(p => (
            <div key={p.id} className="card" style={{ padding: 20, display: 'flex', flexDirection: 'column', gap: 14, position: 'relative' }}>
              {p.destaque && (
                <span className="badge badge-cyan" style={{ position: 'absolute', top: 14, right: 14 }}>{p.destaque}</span>
              )}

              <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
                <div style={{ width: 48, height: 48, borderRadius: 'var(--radius)', background: 'var(--bg-elevated)', border: '1px solid var(--border)', display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: 24 }}>
                  {categoryIcons[p.categoria] || '📦'}
                </div>
                <div style={{ flex: 1, minWidth: 0 }}>
                  <div style={{ fontWeight: 700, fontSize: 14, lineHeight: 1.3, marginBottom: 2 }}>{p.nome}</div>
                  <div style={{ fontFamily: 'var(--font-mono)', fontSize: 10, color: 'var(--text-faint)', textTransform: 'uppercase', letterSpacing: '0.05em' }}>
                    {categorias.find(c => c.id === p.categoria)?.label}
                  </div>
                </div>
              </div>

              <p style={{ fontSize: 12, color: 'var(--text-muted)', lineHeight: 1.5 }}>{p.descricao}</p>

              <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
                <span className="material-symbols-outlined" style={{ fontSize: 14, color: 'var(--amber)', fontVariationSettings: "'FILL' 1" }}>star</span>
                <span style={{ fontSize: 12, fontWeight: 600 }}>{p.rating}</span>
                <span style={{ fontSize: 11, color: 'var(--text-faint)' }}>({p.reviews} avaliações)</span>
                <span style={{ marginLeft: 'auto', fontSize: 11, color: p.estoque > 10 ? 'var(--green)' : p.estoque > 0 ? 'var(--amber)' : 'var(--red)' }}>
                  {p.estoque > 0 ? `${p.estoque} em estoque` : 'Sem estoque'}
                </span>
              </div>

              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', borderTop: '1px solid var(--border)', paddingTop: 14 }}>
                <div>
                  <div style={{ fontFamily: 'var(--font-display)', fontSize: 20, fontWeight: 800, color: 'var(--cyan)', letterSpacing: '-0.02em' }}>
                    {p.preco.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                  </div>
                </div>
                <button
                  className="btn-primary"
                  style={{ padding: '8px 14px', fontSize: 12 }}
                  onClick={() => navigate('/novo-pedido')}
                >
                  <span className="material-symbols-outlined" style={{ fontSize: 16 }}>add_shopping_cart</span>
                  Pedir
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
