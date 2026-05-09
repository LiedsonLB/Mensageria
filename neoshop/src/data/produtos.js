// TODO: substituir por chamada real à API de produtos

export const categorias = [
  { id: 'todos', label: 'Todos os Produtos', icon: 'grid_view' },
  { id: 'notebooks', label: 'Notebooks', icon: 'laptop' },
  { id: 'perifericos', label: 'Periféricos', icon: 'mouse' },
  { id: 'hardware', label: 'Hardware', icon: 'memory' },
  { id: 'monitores', label: 'Monitores', icon: 'desktop_windows' },
  { id: 'audio', label: 'Áudio', icon: 'headphones' },
  { id: 'armazenamento', label: 'Armazenamento', icon: 'storage' },
  { id: 'redes', label: 'Redes', icon: 'router' },
]

export const produtos = [
  // Notebooks
  { id: 1, nome: 'Notebook Gamer Legion 5i', preco: 7499.99, categoria: 'notebooks', rating: 4.9, reviews: 312, destaque: 'Mais Vendido', estoque: 12, descricao: 'Intel Core i7-13700H, RTX 4060, 16GB DDR5, 512GB NVMe' },
  { id: 2, nome: 'Notebook Pro Creator X1', preco: 9999.99, categoria: 'notebooks', rating: 4.8, reviews: 198, destaque: 'Premium', estoque: 5, descricao: 'Apple M3 Pro, 18GB RAM, 512GB SSD, Tela Liquid Retina XDR' },
  { id: 3, nome: 'Ultrabook DellXPS 15', preco: 8249.00, categoria: 'notebooks', rating: 4.7, reviews: 256, estoque: 8, descricao: 'Intel i9-13900H, 32GB DDR5, 1TB NVMe, OLED Touch 3.5K' },
  { id: 4, nome: 'Notebook Asus ROG Zephyrus', preco: 11999.00, categoria: 'notebooks', rating: 4.8, reviews: 144, destaque: 'Top Gamer', estoque: 3, descricao: 'AMD Ryzen 9 7940HS, RTX 4090, 32GB, 2TB SSD, 240Hz' },
  { id: 5, nome: 'Notebook HP Spectre x360', preco: 6899.00, categoria: 'notebooks', rating: 4.6, reviews: 189, estoque: 7, descricao: 'Intel i7-1255U, 16GB, 512GB, 13.5" OLED Touch, conversível' },
  { id: 6, nome: 'Notebook Lenovo ThinkPad X1 Carbon', preco: 8799.00, categoria: 'notebooks', rating: 4.9, reviews: 421, estoque: 10, descricao: 'Intel i7-1365U, 16GB LPDDR5, 512GB, 14" IPS, ultraleve 1.12kg' },

  // Periféricos
  { id: 7, nome: 'Mouse Gamer Logitech G Pro X', preco: 549.90, categoria: 'perifericos', rating: 5.0, reviews: 2100, destaque: 'Best Seller', estoque: 50, descricao: 'Sensor HERO 25K, 63g, sem fio Lightspeed, 70h bateria' },
  { id: 8, nome: 'Teclado Mecânico HyperX Alloy Origins', preco: 699.00, categoria: 'perifericos', rating: 4.8, reviews: 876, estoque: 30, descricao: 'Switches Red, RGB per-key, full-size, alumínio escovado' },
  { id: 9, nome: 'Headset Gamer Razer BlackShark V2 Pro', preco: 899.00, categoria: 'perifericos', rating: 4.9, reviews: 1430, destaque: 'Escolha Pro', estoque: 22, descricao: 'Drivers TriForce 50mm, THX Spatial Audio, 70h, sem fio' },
  { id: 10, nome: 'Mouse Pad XXL SteelSeries QcK', preco: 189.90, categoria: 'perifericos', rating: 4.7, reviews: 3200, estoque: 80, descricao: '90x40cm, superfície micro-texturada, base antiderrapante' },
  { id: 11, nome: 'Webcam Logitech Brio 4K', preco: 1299.00, categoria: 'perifericos', rating: 4.8, reviews: 567, estoque: 15, descricao: '4K Ultra HD, HDR, autofoco, 90° FOV, USB-C' },
  { id: 12, nome: 'Controle Xbox Elite Series 2', preco: 1099.00, categoria: 'perifericos', rating: 4.9, reviews: 892, destaque: 'Pro Controller', estoque: 18, descricao: 'Botões configuráveis, vibração háptica, 40h, sem fio' },

  // Hardware
  { id: 13, nome: 'GPU RTX 4090 Asus ROG Strix', preco: 12499.00, categoria: 'hardware', rating: 5.0, reviews: 234, destaque: 'Top de Linha', estoque: 2, descricao: '24GB GDDR6X, 16384 Cuda cores, DLSS 3, Ray Tracing' },
  { id: 14, nome: 'CPU AMD Ryzen 9 7950X', preco: 4299.00, categoria: 'hardware', rating: 4.9, reviews: 318, estoque: 9, descricao: '16 núcleos / 32 threads, até 5.7GHz, TDP 170W, AM5' },
  { id: 15, nome: 'Memória RAM DDR5 32GB Corsair Dominator', preco: 999.00, categoria: 'hardware', rating: 4.8, reviews: 445, estoque: 25, descricao: '2x16GB, 6200MHz CL36, RGB, dissipador alumínio' },
  { id: 16, nome: 'SSD NVMe Samsung 990 Pro 2TB', preco: 1199.00, categoria: 'hardware', rating: 4.9, reviews: 782, destaque: 'Mais Rápido', estoque: 35, descricao: 'PCIe 4.0, Leitura 7450MB/s, Escrita 6900MB/s, M.2' },
  { id: 17, nome: 'Placa-Mãe ASUS ROG Maximus Z790', preco: 3899.00, categoria: 'hardware', rating: 4.8, reviews: 167, estoque: 6, descricao: 'LGA1700, DDR5, PCIe 5.0, Wi-Fi 6E, Thunderbolt 4' },
  { id: 18, nome: 'Fonte Corsair RM1000x 1000W 80+ Gold', preco: 1299.00, categoria: 'hardware', rating: 4.9, reviews: 534, estoque: 14, descricao: '80+ Gold, modular, 135mm FDB fan, 10 anos garantia' },

  // Monitores
  { id: 19, nome: 'Monitor LG UltraWide 34" QHD Curved', preco: 3499.00, categoria: 'monitores', rating: 4.8, reviews: 672, estoque: 11, descricao: '3440x1440, IPS, 160Hz, FreeSync Premium, HDR10, 1ms GtG' },
  { id: 20, nome: 'Monitor Asus ROG Swift 27" 360Hz', preco: 5299.00, categoria: 'monitores', rating: 4.9, reviews: 289, destaque: 'E-Sports', estoque: 7, descricao: '1080p, IPS, 360Hz, 1ms, G-Sync, HDR400' },
  { id: 21, nome: 'Monitor Dell Alienware 32" QD-OLED', preco: 7999.00, categoria: 'monitores', rating: 5.0, reviews: 148, destaque: 'OLED', estoque: 4, descricao: '4K, QD-OLED, 240Hz, 0.1ms, HDR True Black 400' },
  { id: 22, nome: 'Monitor Samsung Odyssey G9 49"', preco: 9999.00, categoria: 'monitores', rating: 4.7, reviews: 213, estoque: 3, descricao: 'Ultrawide Dualscreen, QLED, 240Hz, 1000R, HDR2000' },

  // Áudio
  { id: 23, nome: 'Fone Sony WH-1000XM5', preco: 1899.00, categoria: 'audio', rating: 4.9, reviews: 4200, destaque: 'ANC Premium', estoque: 28, descricao: 'ANC líder de mercado, 30h, LDAC, Hi-Res Audio, USB-C' },
  { id: 24, nome: 'DAC/AMP FiiO K7', preco: 1299.00, categoria: 'audio', rating: 4.8, reviews: 345, estoque: 12, descricao: 'ESS9038Q2M, saída balanceada 4.4mm, USB/RCA/Optical' },
  { id: 25, nome: 'Microfone Blue Yeti X', preco: 899.00, categoria: 'audio', rating: 4.7, reviews: 1876, estoque: 20, descricao: 'USB, 4 padrões polares, monitor mix, Blue Voice FX' },
  { id: 26, nome: 'Caixas Harman Kardon Aura Studio 3', preco: 2499.00, categoria: 'audio', rating: 4.6, reviews: 567, estoque: 9, descricao: '360° de som, Bluetooth 5.0, design dome transparente' },

  // Armazenamento
  { id: 27, nome: 'HD Externo Seagate IronWolf 8TB', preco: 1299.00, categoria: 'armazenamento', rating: 4.7, reviews: 893, estoque: 16, descricao: 'NAS, CMR, 7200RPM, SATA, AgileArray, IHD Guard' },
  { id: 28, nome: 'SSD Externo Samsung T7 Shield 2TB', preco: 699.00, categoria: 'armazenamento', rating: 4.9, reviews: 2134, destaque: 'IP65', estoque: 40, descricao: '1050MB/s, USB 3.2 Gen 2, resistente a quedas e poeira' },
  { id: 29, nome: 'NAS Synology DS923+ 4 Baias', preco: 3499.00, categoria: 'armazenamento', rating: 4.8, reviews: 234, estoque: 5, descricao: 'AMD Ryzen R1600, 4GB ECC, 2x 1GbE, PCIe 3.0, DSM 7' },

  // Redes
  { id: 30, nome: 'Roteador ASUS ROG Rapture GT-AXE16000', preco: 5299.00, categoria: 'redes', rating: 4.8, reviews: 312, destaque: 'Wi-Fi 6E', estoque: 8, descricao: 'Quad-band, 16Gbps, 2.5G WAN, 8 portas LAN, VPN Gaming' },
  { id: 31, nome: 'Switch TP-Link TL-SG108E 8 Portas', preco: 399.00, categoria: 'redes', rating: 4.8, reviews: 1234, estoque: 30, descricao: 'Gigabit gerenciável, VLAN 802.1Q, QoS, plug-and-play' },
  { id: 32, nome: 'Access Point Ubiquiti UniFi U6 Pro', preco: 1599.00, categoria: 'redes', rating: 4.9, reviews: 567, estoque: 15, descricao: 'Wi-Fi 6, até 5700 Mbps, 300+ clientes, PoE, 6 antenas' },
]

export function filtrarProdutos(busca = '', categoria = 'todos') {
  return produtos.filter(p => {
    const matchCategoria = categoria === 'todos' || p.categoria === categoria
    const matchBusca = !busca || p.nome.toLowerCase().includes(busca.toLowerCase())
    return matchCategoria && matchBusca
  })
}
