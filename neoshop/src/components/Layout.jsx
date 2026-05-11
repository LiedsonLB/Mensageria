import React from 'react'
import { useNavigate, useLocation } from 'react-router-dom'
import logo from '../assets/neoshop_icon.png'

const navItems = [
  { label: 'Novo Pedido', icon: 'add_circle', path: '/' },
  { label: 'Pedidos', icon: 'shopping_cart', path: '/pedidos' },
]

export default function Layout({ children }) {
  const navigate = useNavigate()
  const location = useLocation()

  return (
    <div className="app-layout cyber-bg">
      <aside className="sidebar">
        <div className="sidebar-logo">
          <img
            src={logo}
            height={40}
            width={40}
            alt="NeoShop Logo"
            style={{ 
              borderRadius: '10px',
              marginBottom: '8px'
            }}
          />
          <div className="sidebar-tagline">
            NeoShop
          </div>
        </div>

        <nav>
          {navItems.map(item => (
            <button
              key={item.path}
              className={`nav-item ${location.pathname === item.path ? 'active' : ''}`}
              onClick={() => navigate(item.path)}
            >
              <span className="material-symbols-outlined">{item.icon}</span>
              {item.label}
            </button>
          ))}
        </nav>

        <div className="sidebar-bottom">
          <button className="nav-item">
            <span className="material-symbols-outlined">help</span>
            Central de Ajuda
          </button>
          <button className="nav-item">
            <span className="material-symbols-outlined">logout</span>
            Sair
          </button>
        </div>
      </aside>

      <div className="main-content">
        <header className="topbar">
          <div className="topbar-search"></div>
          <div className="topbar-right">
            <button className="icon-btn">
              <span className="material-symbols-outlined">notifications</span>
              <span className="notif-dot" />
            </button>
            <div className="avatar-badge">
              <img
                src="https://api.dicebear.com/7.x/shapes/svg?seed=Alex"
                alt="Avatar"
              />
              <div>
                <div className="name">Francisco Liédson</div>
                <div className="role">Administrador</div>
              </div>
            </div>
          </div>
        </header>

        <main className="page fade-in">
          {children}
        </main>
      </div>
    </div>
  )
}