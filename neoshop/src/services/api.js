// services/api.js
const BASE_URL = 'http://localhost:8080'

const mockPedidos = [
  { id: 'ORD-99231', cliente: 'Global Connect Ltda.', email: 'gc@global.com', valor_total: 12450.00, status: 'CONCLUIDO', data: '24 Mai, 2024', created_at: '2024-05-24T10:00:00Z', produtos: [] },
  { id: 'ORD-99230', cliente: 'Nexus Tech Soluções', email: 'nexus@tech.com', valor_total: 8120.50, status: 'PROCESSANDO_PAGAMENTO', data: '23 Mai, 2024', created_at: '2024-05-23T10:00:00Z', produtos: [] },
  { id: 'ORD-99229', cliente: 'Solaris Manufatura', email: 'sol@mfg.com', valor_total: 450.00, status: 'PENDENTE', data: '22 Mai, 2024', created_at: '2024-05-22T10:00:00Z', produtos: [] },
  { id: 'ORD-99228', cliente: 'Apex Analytics', email: 'apex@anal.com', valor_total: 32100.00, status: 'CONCLUIDO', data: '22 Mai, 2024', created_at: '2024-05-22T10:00:00Z', produtos: [] },
  { id: 'ORD-99227', cliente: 'V7 Media Group', email: 'v7@media.com', valor_total: 2145.20, status: 'CANCELADO', data: '21 Mai, 2024', created_at: '2024-05-21T10:00:00Z', produtos: [] },
]

export async function criarPedido(dados) {
  try {
    const response = await fetch(`${BASE_URL}/pedido`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dados),
    })
    
    if (!response.ok) {
      const err = await response.text()
      throw new Error(err || 'Erro ao criar pedido')
    }
    
    const pedido = await response.json()
    
    // Buscar o pedido completo com produtos
    const pedidoCompleto = await buscarPedido(pedido.id)
    return pedidoCompleto
  } catch (error) {
    console.error('Erro ao criar pedido:', error)
    throw error
  }
}

export async function listarPedidos() {
  try {
    const response = await fetch(`${BASE_URL}/pedidos`)
    
    if (!response.ok) {
      throw new Error('Erro ao listar pedidos')
    }
    
    const pedidos = await response.json()
    
    // Mapear os dados do backend para o formato esperado pelo frontend
    return pedidos.map(pedido => ({
      id: pedido.id,
      cliente: pedido.cliente,
      email: pedido.email,
      valor_total: pedido.valor_total,
      status: pedido.status,
      data: formatarData(pedido.created_at),
      created_at: pedido.created_at,
      produtos: pedido.produtos || []
    }))
  } catch (error) {
    console.error('Erro ao listar pedidos:', error)
    // Retornar array vazio em caso de erro
    return []
  }
}

export async function buscarPedido(id) {
  try {
    const response = await fetch(`${BASE_URL}/pedido/${id}`)
    
    if (!response.ok) {
      throw new Error('Pedido não encontrado')
    }
    
    const pedido = await response.json()
    
    return {
      id: pedido.id,
      cliente: pedido.cliente,
      documento: pedido.documento,
      email: pedido.email,
      valor_total: pedido.valor_total,
      status: pedido.status,
      produtos: pedido.produtos || [],
      created_at: pedido.created_at,
      nota_fiscal: pedido.nota_fiscal
    }
  } catch (error) {
    console.error('Erro ao buscar pedido:', error)
    throw error
  }
}

export async function buscarStatus(id) {
  try {
    const response = await fetch(`${BASE_URL}/pedido/${id}/status`)
    
    if (!response.ok) {
      throw new Error('Erro ao buscar status')
    }
    
    return await response.json()
  } catch (error) {
    console.error('Erro ao buscar status:', error)
    throw error
  }
}

export async function baixarNotaFiscal(id) {
  try {
    const response = await fetch(`${BASE_URL}/download/${id}`)
    
    if (!response.ok) {
      throw new Error('Nota fiscal não encontrada')
    }
    
    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `nota-fiscal-${id}.xml`
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    window.URL.revokeObjectURL(url)
    
    return true
  } catch (error) {
    console.error('Erro ao baixar nota fiscal:', error)
    throw error
  }
}

// Função auxiliar para formatar data
function formatarData(dataISO) {
  if (!dataISO) return '—'
  const data = new Date(dataISO)
  return data.toLocaleDateString('pt-BR', {
    day: '2-digit',
    month: 'short',
    year: 'numeric'
  })
}