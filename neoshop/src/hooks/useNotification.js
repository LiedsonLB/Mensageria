// hooks/useNotification.js
import { useCallback, useRef } from 'react'

export const useNotification = () => {
  const audioRef = useRef(null)

  const playSound = useCallback((type = 'notification') => {
    const sounds = {
      notification: '/notification.mp3',
      success: '/notification.mp3',
      error: '/notification.mp3'
    }
    
    if (audioRef.current) {
      audioRef.current.pause()
      audioRef.current.currentTime = 0
    }
    
    audioRef.current = new Audio(sounds[type])
    audioRef.current.volume = 0.5
    audioRef.current.play().catch(e => console.log('Erro ao tocar som:', e))
  }, [])

  const showNotification = useCallback((title, body, type = 'info') => {
    // Play sound
    playSound(type === 'success' ? 'success' : type === 'error' ? 'error' : 'notification')
    
    // Show browser notification (se permitido)
    if (Notification.permission === 'granted') {
      new Notification(title, { body, icon: '/logo192.png' })
    } else if (Notification.permission !== 'denied') {
      Notification.requestPermission().then(permission => {
        if (permission === 'granted') {
          new Notification(title, { body, icon: '/logo192.png' })
        }
      })
    }
    
    // Também mostra toast na interface
    const toastEvent = new CustomEvent('show-toast', { detail: { title, body, type } })
    window.dispatchEvent(toastEvent)
  }, [playSound])

  return { showNotification, playSound }
}