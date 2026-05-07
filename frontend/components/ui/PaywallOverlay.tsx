'use client'
import { useState } from 'react'
import { MysticButton } from './MysticButton'

interface PaywallOverlayProps {
  onUnlock: () => void
  loading?: boolean
  price?: string
}

export function PaywallOverlay({ onUnlock, loading, price = '49.000đ' }: PaywallOverlayProps) {
  const [hover, setHover] = useState(false)

  return (
    <div className="absolute inset-0 z-10 flex items-center justify-center rounded-sm overflow-hidden">
      {/* Gradient fog from bottom */}
      <div
        className="absolute inset-0"
        style={{
          background: 'linear-gradient(to bottom, transparent 0%, rgba(3,3,8,0.6) 25%, rgba(3,3,8,0.97) 55%)',
        }}
      />

      {/* Seal */}
      <div className="relative flex flex-col items-center gap-5 pb-2">
        {/* Rotating seal SVG */}
        <div
          className="relative"
          onMouseEnter={() => setHover(true)}
          onMouseLeave={() => setHover(false)}
        >
          <svg
            width="100" height="100"
            viewBox="0 0 100 100"
            className="transition-all duration-700"
            style={{ transform: hover ? 'scale(1.08)' : 'scale(1)' }}
          >
            {/* Outer ring — spins slowly */}
            <circle cx="50" cy="50" r="46" fill="none" stroke="rgba(201,162,39,0.15)" strokeWidth="0.5" />
            <g style={{ animation: 'rotateSlowCW 20s linear infinite', transformOrigin: '50px 50px' }}>
              <circle cx="50" cy="50" r="46" fill="none"
                stroke="rgba(201,162,39,0.35)" strokeWidth="0.75"
                strokeDasharray="4 8" />
            </g>

            {/* Inner ring — spins CCW */}
            <g style={{ animation: 'rotateSlowCCW 12s linear infinite', transformOrigin: '50px 50px' }}>
              <circle cx="50" cy="50" r="35" fill="none"
                stroke="rgba(201,162,39,0.2)" strokeWidth="0.5"
                strokeDasharray="2 6" />
            </g>

            {/* Seal center */}
            <circle cx="50" cy="50" r="22" fill="rgba(3,3,8,0.9)"
              stroke="rgba(201,162,39,0.5)" strokeWidth="1" />

            {/* Lock symbol */}
            <rect x="43" y="50" width="14" height="11" rx="1.5"
              fill="none" stroke="rgba(201,162,39,0.8)" strokeWidth="1.2" />
            <path d="M44.5 50 C44.5 43 55.5 43 55.5 50"
              fill="none" stroke="rgba(201,162,39,0.8)" strokeWidth="1.2" />
            <circle cx="50" cy="55.5" r="1.5" fill="rgba(201,162,39,0.8)" />

            {/* 4 corner marks */}
            {[0, 90, 180, 270].map((deg) => (
              <line key={deg}
                x1="50" y1="5" x2="50" y2="10"
                stroke="rgba(201,162,39,0.5)" strokeWidth="1"
                style={{ transformOrigin: '50px 50px', transform: `rotate(${deg}deg)` }}
              />
            ))}
          </svg>
        </div>

        {/* Text */}
        <div className="text-center space-y-1.5">
          <p
            className="text-xs tracking-[0.25em] uppercase"
            style={{ color: 'var(--gold-dim)', fontFamily: 'var(--font-mono)' }}
          >
            Ấn Phong Tỏa
          </p>
          <p
            className="text-base"
            style={{ color: 'var(--text-bright)', fontFamily: 'var(--font-display)', fontStyle: 'italic' }}
          >
            Mở ấn để khám phá toàn bộ vận mệnh
          </p>
          <p style={{ color: 'var(--text-dim)', fontSize: '0.72rem', fontFamily: 'var(--font-body)' }}>
            Luận giải đầy đủ · Vận trình từng tháng · AI cá nhân hóa
          </p>
        </div>

        <MysticButton onClick={onUnlock} loading={loading} size="md">
          Phá Ấn — {price}
        </MysticButton>

        <p style={{ color: 'var(--text-ghost)', fontSize: '0.65rem', fontFamily: 'var(--font-mono)', letterSpacing: '0.1em' }}>
          hoặc đăng ký Premium · 49.000đ/tháng
        </p>
      </div>
    </div>
  )
}
