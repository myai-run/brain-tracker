import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import AppResolver from './components/app-resolver.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AppResolver />
  </StrictMode>,
)
