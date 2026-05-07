'use client'
import { cn } from '@/lib/utils'

interface MysticButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'gold' | 'outline' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  loading?: boolean
}

export function MysticButton({
  children, className, variant = 'gold', size = 'md', loading, disabled, ...props
}: MysticButtonProps) {
  const base = [
    'relative inline-flex items-center justify-center gap-2',
    'font-body font-medium tracking-[0.15em] uppercase text-xs',
    'transition-all duration-300 cursor-none select-none',
    'disabled:opacity-40 disabled:cursor-not-allowed',
  ].join(' ')

  const sizes = {
    sm: 'px-6 py-2.5',
    md: 'px-9 py-3.5',
    lg: 'px-12 py-4',
  }

  const variants = {
    gold: [
      'bg-[var(--gold)] text-[var(--void)]',
      'shadow-[0_0_24px_var(--gold-glow)]',
      'hover:bg-[var(--gold-bright)] hover:shadow-[0_0_40px_rgba(201,162,39,0.55)]',
      'active:scale-[0.97]',
    ].join(' '),
    outline: [
      'border border-[rgba(201,162,39,0.4)] text-[var(--gold)]',
      'hover:border-[var(--gold)] hover:bg-[rgba(201,162,39,0.06)]',
    ].join(' '),
    ghost: 'text-[var(--text-mid)] hover:text-[var(--gold)]',
  }

  return (
    <button
      disabled={disabled || loading}
      className={cn(base, sizes[size], variants[variant], className)}
      {...props}
    >
      {loading ? (
        <>
          <span className="w-3.5 h-3.5 border border-current border-t-transparent rounded-full animate-spin" />
          <span>Đang xử lý</span>
        </>
      ) : children}

      {/* Shimmer on gold variant */}
      {variant === 'gold' && !loading && (
        <span
          aria-hidden
          className="absolute inset-0 overflow-hidden pointer-events-none"
        >
          <span className="absolute inset-0 translate-x-[-100%] hover:translate-x-[100%] bg-gradient-to-r from-transparent via-white/20 to-transparent transition-transform duration-700 ease-in-out" />
        </span>
      )}
    </button>
  )
}
