'use client'

import { useEffect, useState } from 'react'
import { TarotPageClient } from './TarotPageClient'

export function TarotClientIsland() {
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  if (!mounted) return null

  return <TarotPageClient />
}
