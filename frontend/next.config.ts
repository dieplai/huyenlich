import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
  output: 'standalone',
  experimental: {
    optimizePackageImports: ['three', '@react-three/fiber', '@react-three/drei'],
  },
  images: {
    formats: ['image/webp', 'image/avif'],
  },
  async rewrites() {
    const apiURL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080'
    return [{ source: '/api/:path*', destination: `${apiURL}/api/:path*` }]
  },
}

export default nextConfig
