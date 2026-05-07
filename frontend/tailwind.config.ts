import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./app/**/*.{ts,tsx}', './components/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        void: '#040410',
        deep: '#0a0a1e',
        surface: '#12122a',
        elevated: '#1a1a38',
        gold: '#d4af37',
        'gold-glow': '#f0d060',
        mystic: '#8b5cf6',
        cosmic: '#00d4ff',
        rose: '#f43f5e',
      },
      fontFamily: {
        serif: ['Cormorant Garamond', 'Georgia', 'serif'],
        sans: ['Be Vietnam Pro', 'Inter', 'sans-serif'],
        mono: ['Space Mono', 'monospace'],
      },
      backgroundImage: {
        'cosmic-gradient': 'radial-gradient(ellipse at center, #1a0a2e 0%, #040410 70%)',
        'gold-gradient': 'linear-gradient(135deg, #d4af37 0%, #f0d060 50%, #d4af37 100%)',
        'mystic-gradient': 'linear-gradient(135deg, #8b5cf6 0%, #d4af37 100%)',
      },
      animation: {
        'pulse-slow': 'pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        'float': 'float 6s ease-in-out infinite',
        'glow': 'glow 2s ease-in-out infinite alternate',
      },
      keyframes: {
        float: {
          '0%, 100%': { transform: 'translateY(0px)' },
          '50%': { transform: 'translateY(-20px)' },
        },
        glow: {
          from: { boxShadow: '0 0 10px #d4af37, 0 0 20px #d4af37' },
          to: { boxShadow: '0 0 20px #d4af37, 0 0 40px #d4af37, 0 0 60px #8b5cf6' },
        },
      },
    },
  },
  plugins: [],
}

export default config
