'use client'
import { cn } from '@/lib/utils'
import { useRef, useState } from 'react'

interface GlowCardProps {
  children: React.ReactNode
  className?: string
  variant?: 'gold' | 'purple' | 'default'
  ornament?: boolean
}

export function GlowCard({ children, className, variant = 'default', ornament = false }: GlowCardProps) {
  const ref = useRef<HTMLDivElement>(null)
  const [tilt, setTilt] = useState({ x: 0, y: 0 })

  const handleMouseMove = (e: React.MouseEvent<HTMLDivElement>) => {
    if (!ref.current) return
    const rect = ref.current.getBoundingClientRect()
    const x = ((e.clientX - rect.left) / rect.width - 0.5) * 8
    const y = ((e.clientY - rect.top) / rect.height - 0.5) * -8
    setTilt({ x, y })
  }

  const variantStyles = {
    gold:    'border-[rgba(201,162,39,0.25)] shadow-[0_0_40px_rgba(201,162,39,0.08)]',
    purple:  'border-[rgba(91,63,160,0.3)]  shadow-[0_0_40px_rgba(91,63,160,0.08)]',
    default: 'border-[rgba(201,162,39,0.1)] shadow-[0_4px_30px_rgba(0,0,0,0.4)]',
  }

  return (
    <div
      ref={ref}
      onMouseMove={handleMouseMove}
      onMouseLeave={() => setTilt({ x: 0, y: 0 })}
      style={{
        transform: `perspective(600px) rotateX(${tilt.y}deg) rotateY(${tilt.x}deg)`,
        transition: 'transform 0.15s ease',
      }}
      className={cn(
        'relative glass rounded-sm',
        variantStyles[variant],
        className
      )}
    >
      {ornament && (
        <>
          {/* Corner ornaments */}
          {['top-0 left-0', 'top-0 right-0 rotate-90', 'bottom-0 right-0 rotate-180', 'bottom-0 left-0 -rotate-90'].map((pos, i) => (
            <span key={i} className={`absolute ${pos} w-4 h-4 pointer-events-none`}>
              <svg viewBox="0 0 16 16" fill="none" className="w-full h-full">
                <path d="M1 8 L1 1 L8 1" stroke="rgba(201,162,39,0.4)" strokeWidth="0.75" />
              </svg>
            </span>
          ))}
        </>
      )}
      {children}
    </div>
  )
}
