// data/produtos.js
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

// Função para gerar URLs de imagens baseadas no nome do produto
const gerarImagem = (nome, categoria) => {
  // Usando LoremFlick (placeholder images) com categorias
  const imagens = {
    notebooks: {
      legion: 'https://images5.kabum.com.br/produtos/fotos/480655/notebook-gamer-lenovo-legion-slim-5i-intel-core-i5-13420h-16gb-ram-geforce-rtx3050-ssd-512gb-16-2k-qhd-win-11-cinza-83d60003br_1704722946_gg.jpg',
      creator: 'https://a-static.mlcdn.com.br/470x352/notebook-lenovo-x1-carbon-g13-intel-core-ultra-7-268v-vpro-32gb-1tb-ssd-windows-11-pro-14-21nt000hbr-preto/lenovooficial/21nt000hbr/2ba164ebcf537d1f6b8a417f51f3e850.jpeg',
      dell: 'https://i5.walmartimages.com/asr/4346747d-44a8-469a-97aa-e5c52ca143f7.2598ca152f660711428fd11eefc5bca0.jpeg',
      asus: 'https://images.kabum.com.br/produtos/fotos/480655/notebook-gamer-lenovo-legion-slim-5i-intel-core-i5-13420h-16gb-ram-geforce-rtx3050-ssd-512gb-16-2k-qhd-win-11-cinza-83d60003br_1704722946_gg.jpg',
      hp: 'https://images.unsplash.com/photo-1588872657578-7efd1f1555ed?w=150&h=150&fit=crop',
      lenovo: 'https://i.zst.com.br/thumbs/12/6/16/933325240.jpg',
    },
    perifericos: {
      mouse: 'https://images0.kabum.com.br/produtos/fotos/149990/mouse-sem-fio-gamer-logitech-g-pro-x-superlight-lightspeed-5-botoes-25000-dpi-branco-910-005941_1614261697_gg.jpg',
      teclado: 'https://images8.kabum.com.br/produtos/fotos/371598/teclado-hyperx-alloy-origins-65hkbo1t-rd-eua-n-4p5d6aa-aba_1659554473_gg.jpg',
      headset: 'https://images.kabum.com.br/produtos/fotos/128544/headset-gamer-razer-blackshark-v2-x-multi-platform-drivers-50mm-rz04-03240100-r3u1_1600956255_gg.jpg',
      mousepad: 'https://images.kabum.com.br/produtos/fotos/sync_mirakl/575535/Mousepad-Gamer-Steelseries-Qck-Heavy-Xxl-900x400mm_1713968149_gg.jpg',
      webcam: 'https://eletronicasantana.vteximg.com.br/arquivos/ids/112994-1000-1000/Webcam-Brio-4K-Pro-Ultra-HD-Logitech-0.jpg?v=638482016126270000',
      controle: 'https://m.media-amazon.com/images/I/717XTm0moDL.jpg',
    },
    hardware: {
      gpu: 'https://images0.kabum.com.br/produtos/fotos/537350/placa-de-video-rtx-4090-asus-rog-strix-o24g-btf-gaming-nvidia-geforce-24gb-gddr6x-rgb-dlss-ray-tracing-90yv0jt0-m0na00_1716573750_gg.jpg',
      cpu: 'https://images.kabum.com.br/produtos/fotos/sync_mirakl/390981/xlarge/Processador-AMD-Ryzen-Radeon-9-7950X-2200MHz-Cache-80MB-1Hexa-Core-AM5-V-deo-Integrado-100-100000514WOF_1777055974.jpg',
      ram: 'https://cdn.awsli.com.br/600x1000/2557/2557636/produto/289313589/mem-ddr5-32gb-6000mhz-corsair--dominator-platinum-rgb-preto--1-yro6wsa7xv.jpg',
      ssd: 'https://images.kabum.com.br/produtos/fotos/sync_mirakl/442823/xlarge/SSD-Samsung-990-Pro-2TB-Nvme-M-2-2280-Leitura-at-7450mb-s-e-Grava-o-at-6900mb-s-_1753470294.jpg',
      placa_mae: 'https://m.media-amazon.com/images/I/81CpgF-+P4L.jpg',
      fonte: 'https://m.media-amazon.com/images/I/81dwGXVwpgL.jpg',
    },
    monitores: {
      lg: 'https://images.kabum.com.br/produtos/fotos/sync_mirakl/649756/xlarge/Monitor-Gamer-Curvo-Ultrawide-LG-Ultragear-34-Wqhd-160Hz-1ms-AMD-Freesync-Premium-34gp63a-b_1772568517.jpg',
      asus: 'https://cdn.awsli.com.br/600x450/1271/1271561/produto/216493122/monitor-asus-rog-swift-pg27aqn-27-1440p-360hz-hdr-g-sync-quantum-dot--1--ahttjnd8gu.jpg',
      alienware: 'https://http2.mlstatic.com/D_NQ_NP_837865-MLA99868967327_112025-O.webp',
      samsung: 'https://images.kabum.com.br/produtos/fotos/129919/monitor-gamer-samsung-odyssey-led-49-curvo-dqhd-hdmi-displayport-usb-g-sync-freesync-240hz-1ms-altura-ajustavel-lc49g95tsslxzd_1603130064_original.jpg',
    },
    audio: {
      sony: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2MQo3bpCPTZlTDRieALPCMDo4NU-YwOANQw&s',
      fiio: 'https://m.media-amazon.com/images/I/51ctW8llHwL._AC_UF894,1000_QL80_.jpg',
      yeti: 'https://images.kabum.com.br/produtos/fotos/magalu/496394/Microfone-Condensador-Streaming-Blue-Yeti-USB_1695908359_gg.jpg',
      harman: 'https://http2.mlstatic.com/D_Q_NP_755931-CBT110020462608_042026-O.webp',
    },
    armazenamento: {
      hd: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcStb_jI_HkQmKc_28hpBYlgKUn8haUXLseOFw&s',
      ssd: 'https://media.pichau.com.br/media/catalog/product/cache/2f958555330323e505eba7ce930bdf27/m/z/mz-v9p1t0.jpg',
      nas: 'https://images.kabum.com.br/produtos/fotos/sync_mirakl/477627/xlarge/Synology-NAS-DS923-4-baias-Diskstation_1769536121.png',
    },
    redes: {
      roteador: 'https://m.media-amazon.com/images/I/71HzNl--8qL.jpg',
      switch: 'https://m.media-amazon.com/images/I/71VBiMA8cjL.jpg',
      ap: 'https://altaseg.cdn.magazord.com.br/img/2024/04/produto/1082/05-roteador-wi-fi-6-access-point-5-3-gbps-injetor-poe-u6-pro-ubiquiti-c-fonte-48v-30w.png?ims=600x600',
    }
  }

  // Retornar imagem baseada na categoria e nome
  const nomeLower = nome.toLowerCase()
  
  if (categoria === 'notebooks') {
    if (nomeLower.includes('legion')) return imagens.notebooks.legion
    if (nomeLower.includes('creator')) return imagens.notebooks.creator
    if (nomeLower.includes('dell')) return imagens.notebooks.dell
    if (nomeLower.includes('asus')) return imagens.notebooks.asus
    if (nomeLower.includes('hp')) return imagens.notebooks.hp
    return imagens.notebooks.lenovo
  }
  
  if (categoria === 'perifericos') {
    if (nomeLower.includes('mouse') && !nomeLower.includes('pad')) return imagens.perifericos.mouse
    if (nomeLower.includes('teclado')) return imagens.perifericos.teclado
    if (nomeLower.includes('headset')) return imagens.perifericos.headset
    if (nomeLower.includes('mouse pad')) return imagens.perifericos.mousepad
    if (nomeLower.includes('webcam')) return imagens.perifericos.webcam
    if (nomeLower.includes('controle')) return imagens.perifericos.controle
    return imagens.perifericos.mouse
  }
  
  if (categoria === 'hardware') {
    if (nomeLower.includes('gpu') || nomeLower.includes('rtx')) return imagens.hardware.gpu
    if (nomeLower.includes('cpu') || nomeLower.includes('ryzen')) return imagens.hardware.cpu
    if (nomeLower.includes('memória') || nomeLower.includes('ram')) return imagens.hardware.ram
    if (nomeLower.includes('ssd')) return imagens.hardware.ssd
    if (nomeLower.includes('placa-mãe')) return imagens.hardware.placa_mae
    if (nomeLower.includes('fonte')) return imagens.hardware.fonte
    return imagens.hardware.gpu
  }
  
  if (categoria === 'monitores') {
    if (nomeLower.includes('lg')) return imagens.monitores.lg
    if (nomeLower.includes('asus')) return imagens.monitores.asus
    if (nomeLower.includes('alienware')) return imagens.monitores.alienware
    if (nomeLower.includes('samsung')) return imagens.monitores.samsung
    return imagens.monitores.lg
  }
  
  if (categoria === 'audio') {
    if (nomeLower.includes('sony')) return imagens.audio.sony
    if (nomeLower.includes('fiio') || nomeLower.includes('dac')) return imagens.audio.fiio
    if (nomeLower.includes('blue') || nomeLower.includes('yeti')) return imagens.audio.yeti
    if (nomeLower.includes('harman')) return imagens.audio.harman
    return imagens.audio.sony
  }
  
  if (categoria === 'armazenamento') {
    if (nomeLower.includes('seagate') || nomeLower.includes('hd')) return imagens.armazenamento.hd
    if (nomeLower.includes('ssd')) return imagens.armazenamento.ssd
    if (nomeLower.includes('synology') || nomeLower.includes('nas')) return imagens.armazenamento.nas
    return imagens.armazenamento.ssd
  }
  
  if (categoria === 'redes') {
    if (nomeLower.includes('roteador') || nomeLower.includes('asus')) return imagens.redes.roteador
    if (nomeLower.includes('switch')) return imagens.redes.switch
    if (nomeLower.includes('access point') || nomeLower.includes('unifi')) return imagens.redes.ap
    return imagens.redes.roteador
  }
  
  return 'https://m.media-amazon.com/images/I/71HzNl--8qL.jpg'
}

export const produtos = [
  // Notebooks
  { id: 1, nome: 'Notebook Gamer Legion 5i', preco: 7499.99, categoria: 'notebooks', rating: 4.9, reviews: 312, destaque: 'Mais Vendido', estoque: 12, descricao: 'Intel Core i7-13700H, RTX 4060, 16GB DDR5, 512GB NVMe', image: gerarImagem('Legion 5i', 'notebooks') },
  { id: 2, nome: 'Notebook Pro Creator X1', preco: 9999.99, categoria: 'notebooks', rating: 4.8, reviews: 198, destaque: 'Premium', estoque: 5, descricao: 'Apple M3 Pro, 18GB RAM, 512GB SSD, Tela Liquid Retina XDR', image: gerarImagem('Creator X1', 'notebooks') },
  { id: 3, nome: 'Ultrabook DellXPS 15', preco: 8249.00, categoria: 'notebooks', rating: 4.7, reviews: 256, estoque: 8, descricao: 'Intel i9-13900H, 32GB DDR5, 1TB NVMe, OLED Touch 3.5K', image: gerarImagem('DellXPS 15', 'notebooks') },
  { id: 4, nome: 'Notebook Asus ROG Zephyrus', preco: 11999.00, categoria: 'notebooks', rating: 4.8, reviews: 144, destaque: 'Top Gamer', estoque: 3, descricao: 'AMD Ryzen 9 7940HS, RTX 4090, 32GB, 2TB SSD, 240Hz', image: gerarImagem('ROG Zephyrus', 'notebooks') },
  { id: 5, nome: 'Notebook HP Spectre x360', preco: 6899.00, categoria: 'notebooks', rating: 4.6, reviews: 189, estoque: 7, descricao: 'Intel i7-1255U, 16GB, 512GB, 13.5" OLED Touch, conversível', image: gerarImagem('Spectre x360', 'notebooks') },
  { id: 6, nome: 'Notebook Lenovo ThinkPad X1 Carbon', preco: 8799.00, categoria: 'notebooks', rating: 4.9, reviews: 421, estoque: 10, descricao: 'Intel i7-1365U, 16GB LPDDR5, 512GB, 14" IPS, ultraleve 1.12kg', image: gerarImagem('ThinkPad X1 Carbon', 'notebooks') },

  // Periféricos
  { id: 7, nome: 'Mouse Gamer Logitech G Pro X', preco: 549.90, categoria: 'perifericos', rating: 5.0, reviews: 2100, destaque: 'Best Seller', estoque: 50, descricao: 'Sensor HERO 25K, 63g, sem fio Lightspeed, 70h bateria', image: gerarImagem('Mouse Gamer', 'perifericos') },
  { id: 8, nome: 'Teclado Mecânico HyperX Alloy Origins', preco: 699.00, categoria: 'perifericos', rating: 4.8, reviews: 876, estoque: 30, descricao: 'Switches Red, RGB per-key, full-size, alumínio escovado', image: gerarImagem('Teclado Mecânico', 'perifericos') },
  { id: 9, nome: 'Headset Gamer Razer BlackShark V2 Pro', preco: 899.00, categoria: 'perifericos', rating: 4.9, reviews: 1430, destaque: 'Escolha Pro', estoque: 22, descricao: 'Drivers TriForce 50mm, THX Spatial Audio, 70h, sem fio', image: gerarImagem('Headset Gamer', 'perifericos') },
  { id: 10, nome: 'Mouse Pad XXL SteelSeries QcK', preco: 189.90, categoria: 'perifericos', rating: 4.7, reviews: 3200, estoque: 80, descricao: '90x40cm, superfície micro-texturada, base antiderrapante', image: gerarImagem('Mouse Pad', 'perifericos') },
  { id: 11, nome: 'Webcam Logitech Brio 4K', preco: 1299.00, categoria: 'perifericos', rating: 4.8, reviews: 567, estoque: 15, descricao: '4K Ultra HD, HDR, autofoco, 90° FOV, USB-C', image: gerarImagem('Webcam', 'perifericos') },
  { id: 12, nome: 'Controle Xbox Elite Series 2', preco: 1099.00, categoria: 'perifericos', rating: 4.9, reviews: 892, destaque: 'Pro Controller', estoque: 18, descricao: 'Botões configuráveis, vibração háptica, 40h, sem fio', image: gerarImagem('Controle Xbox', 'perifericos') },

  // Hardware
  { id: 13, nome: 'GPU RTX 4090 Asus ROG Strix', preco: 12499.00, categoria: 'hardware', rating: 5.0, reviews: 234, destaque: 'Top de Linha', estoque: 2, descricao: '24GB GDDR6X, 16384 Cuda cores, DLSS 3, Ray Tracing', image: gerarImagem('RTX 4090', 'hardware') },
  { id: 14, nome: 'CPU AMD Ryzen 9 7950X', preco: 4299.00, categoria: 'hardware', rating: 4.9, reviews: 318, estoque: 9, descricao: '16 núcleos / 32 threads, até 5.7GHz, TDP 170W, AM5', image: gerarImagem('Ryzen 9', 'hardware') },
  { id: 15, nome: 'Memória RAM DDR5 32GB Corsair Dominator', preco: 999.00, categoria: 'hardware', rating: 4.8, reviews: 445, estoque: 25, descricao: '2x16GB, 6200MHz CL36, RGB, dissipador alumínio', image: gerarImagem('Memória RAM', 'hardware') },
  { id: 16, nome: 'SSD NVMe Samsung 990 Pro 2TB', preco: 1199.00, categoria: 'hardware', rating: 4.9, reviews: 782, destaque: 'Mais Rápido', estoque: 35, descricao: 'PCIe 4.0, Leitura 7450MB/s, Escrita 6900MB/s, M.2', image: gerarImagem('SSD', 'hardware') },
  { id: 17, nome: 'Placa-Mãe ASUS ROG Maximus Z790', preco: 3899.00, categoria: 'hardware', rating: 4.8, reviews: 167, estoque: 6, descricao: 'LGA1700, DDR5, PCIe 5.0, Wi-Fi 6E, Thunderbolt 4', image: gerarImagem('Placa-Mãe', 'hardware') },
  { id: 18, nome: 'Fonte Corsair RM1000x 1000W 80+ Gold', preco: 1299.00, categoria: 'hardware', rating: 4.9, reviews: 534, estoque: 14, descricao: '80+ Gold, modular, 135mm FDB fan, 10 anos garantia', image: gerarImagem('Fonte', 'hardware') },

  // Monitores
  { id: 19, nome: 'Monitor LG UltraWide 34" QHD Curved', preco: 3499.00, categoria: 'monitores', rating: 4.8, reviews: 672, estoque: 11, descricao: '3440x1440, IPS, 160Hz, FreeSync Premium, HDR10, 1ms GtG', image: gerarImagem('LG UltraWide', 'monitores') },
  { id: 20, nome: 'Monitor Asus ROG Swift 27" 360Hz', preco: 5299.00, categoria: 'monitores', rating: 4.9, reviews: 289, destaque: 'E-Sports', estoque: 7, descricao: '1080p, IPS, 360Hz, 1ms, G-Sync, HDR400', image: gerarImagem('Asus ROG', 'monitores') },
  { id: 21, nome: 'Monitor Dell Alienware 32" QD-OLED', preco: 7999.00, categoria: 'monitores', rating: 5.0, reviews: 148, destaque: 'OLED', estoque: 4, descricao: '4K, QD-OLED, 240Hz, 0.1ms, HDR True Black 400', image: gerarImagem('Alienware', 'monitores') },
  { id: 22, nome: 'Monitor Samsung Odyssey G9 49"', preco: 9999.00, categoria: 'monitores', rating: 4.7, reviews: 213, estoque: 3, descricao: 'Ultrawide Dualscreen, QLED, 240Hz, 1000R, HDR2000', image: gerarImagem('Samsung Odyssey', 'monitores') },

  // Áudio
  { id: 23, nome: 'Fone Sony WH-1000XM5', preco: 1899.00, categoria: 'audio', rating: 4.9, reviews: 4200, destaque: 'ANC Premium', estoque: 28, descricao: 'ANC líder de mercado, 30h, LDAC, Hi-Res Audio, USB-C', image: gerarImagem('Sony', 'audio') },
  { id: 24, nome: 'DAC/AMP FiiO K7', preco: 1299.00, categoria: 'audio', rating: 4.8, reviews: 345, estoque: 12, descricao: 'ESS9038Q2M, saída balanceada 4.4mm, USB/RCA/Optical', image: gerarImagem('FiiO', 'audio') },
  { id: 25, nome: 'Microfone Blue Yeti X', preco: 899.00, categoria: 'audio', rating: 4.7, reviews: 1876, estoque: 20, descricao: 'USB, 4 padrões polares, monitor mix, Blue Voice FX', image: gerarImagem('Blue Yeti', 'audio') },
  { id: 26, nome: 'Caixas Harman Kardon Aura Studio 3', preco: 2499.00, categoria: 'audio', rating: 4.6, reviews: 567, estoque: 9, descricao: '360° de som, Bluetooth 5.0, design dome transparente', image: gerarImagem('Harman Kardon', 'audio') },

  // Armazenamento
  { id: 27, nome: 'HD Externo Seagate IronWolf 8TB', preco: 1299.00, categoria: 'armazenamento', rating: 4.7, reviews: 893, estoque: 16, descricao: 'NAS, CMR, 7200RPM, SATA, AgileArray, IHD Guard', image: gerarImagem('Seagate', 'armazenamento') },
  { id: 28, nome: 'SSD Externo Samsung T7 Shield 2TB', preco: 699.00, categoria: 'armazenamento', rating: 4.9, reviews: 2134, destaque: 'IP65', estoque: 40, descricao: '1050MB/s, USB 3.2 Gen 2, resistente a quedas e poeira', image: gerarImagem('SSD Samsung', 'armazenamento') },
  { id: 29, nome: 'NAS Synology DS923+ 4 Baias', preco: 3499.00, categoria: 'armazenamento', rating: 4.8, reviews: 234, estoque: 5, descricao: 'AMD Ryzen R1600, 4GB ECC, 2x 1GbE, PCIe 3.0, DSM 7', image: gerarImagem('Synology', 'armazenamento') },

  // Redes
  { id: 30, nome: 'Roteador ASUS ROG Rapture GT-AXE16000', preco: 5299.00, categoria: 'redes', rating: 4.8, reviews: 312, destaque: 'Wi-Fi 6E', estoque: 8, descricao: 'Quad-band, 16Gbps, 2.5G WAN, 8 portas LAN, VPN Gaming', image: gerarImagem('Roteador ASUS', 'redes') },
  { id: 31, nome: 'Switch TP-Link TL-SG108E 8 Portas', preco: 399.00, categoria: 'redes', rating: 4.8, reviews: 1234, estoque: 30, descricao: 'Gigabit gerenciável, VLAN 802.1Q, QoS, plug-and-play', image: gerarImagem('Switch', 'redes') },
  { id: 32, nome: 'Access Point Ubiquiti UniFi U6 Pro', preco: 1599.00, categoria: 'redes', rating: 4.9, reviews: 567, estoque: 15, descricao: 'Wi-Fi 6, até 5700 Mbps, 300+ clientes, PoE, 6 antenas', image: gerarImagem('Access Point', 'redes') },
]

export function filtrarProdutos(busca = '', categoria = 'todos') {
  return produtos.filter(p => {
    const matchCategoria = categoria === 'todos' || p.categoria === categoria
    const matchBusca = !busca || p.nome.toLowerCase().includes(busca.toLowerCase())
    return matchCategoria && matchBusca
  })
}