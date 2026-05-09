// components/Toast.jsx
import React, { useState, useEffect } from 'react'

export default function Toast() {
  const [toasts, setToasts] = useState([])

  useEffect(() => {
    const handleShowToast = (event) => {
      const id = Date.now()
      setToasts(prev => [...prev, { id, ...event.detail }])
      
      // Auto-remove após 5 segundos
      setTimeout(() => {
        setToasts(prev => prev.filter(t => t.id !== id))
      }, 5000)
    }
    
    window.addEventListener('show-toast', handleShowToast)
    return () => window.removeEventListener('show-toast', handleShowToast)
  }, [])

  if (toasts.length === 0) return null

  return (
    <div style={{
      position: 'fixed',
      bottom: 20,
      right: 20,
      zIndex: 9999,
      display: 'flex',
      flexDirection: 'column',
      gap: 10
    }}>
      {toasts.map(toast => (
        <div
          key={toast.id}
          style={{
            background: toast.type === 'success' ? 'var(--green)' : 
                        toast.type === 'error' ? 'var(--red)' : 
                        'var(--cyan)',
            color: '#000',
            padding: '12px 20px',
            borderRadius: 'var(--radius)',
            minWidth: 280,
            boxShadow: '0 4px 12px rgba(0,0,0,0.3)',
            animation: 'slideIn 0.3s ease'
          }}
        >
          <div style={{ fontWeight: 700, marginBottom: 4 }}>{toast.title}</div>
          <div style={{ fontSize: 13 }}>{toast.body}</div>
        </div>
      ))}
    </div>
  )
}