const BASE_URL = 'http://localhost:8080'

export async function criarPedido(dados) {
  const response = await fetch(`${BASE_URL}/pedido`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(dados),
  })
  if (!response.ok) {
    const err = await response.text()
    throw new Error(err || 'Erro ao criar pedido')
  }
  return response.json()
}

export async function listarPedidos() {
  // const response = await fetch(`${BASE_URL}/pedidos/`)
  // return response.json()

  // Dados simulados até a API de listagem estar disponível
  return [
    { id: 'ORD-99231', cliente: 'Global Connect Ltda.', email: 'gc@global.com', valor_total: 12450.00, status: 'CONCLUIDO', data: '24 Mai, 2024' },
    { id: 'ORD-99230', cliente: 'Nexus Tech Soluções', email: 'nexus@tech.com', valor_total: 8120.50, status: 'PROCESSANDO', data: '23 Mai, 2024' },
    { id: 'ORD-99229', cliente: 'Solaris Manufatura', email: 'sol@mfg.com', valor_total: 450.00, status: 'PENDENTE', data: '22 Mai, 2024' },
    { id: 'ORD-99228', cliente: 'Apex Analytics', email: 'apex@anal.com', valor_total: 32100.00, status: 'CONCLUIDO', data: '22 Mai, 2024' },
    { id: 'ORD-99227', cliente: 'V7 Media Group', email: 'v7@media.com', valor_total: 2145.20, status: 'CANCELADO', data: '21 Mai, 2024' },
    { id: 'ORD-99226', cliente: 'Francisco Bonfim', email: 'fb@email.com', valor_total: 4050.00, status: 'PENDENTE', data: '20 Mai, 2024' },
  ]
}
