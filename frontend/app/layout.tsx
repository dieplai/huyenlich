import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  metadataBase: new URL('https://tarot.local'),
  title: {
    default: 'Tarot — Trải Bài Rider-Waite-Smith',
    template: '%s | Tarot',
  },
  description: 'Trải bài Tarot Rider-Waite-Smith online, rút bài trực quan và luận giải AI bằng tiếng Việt.',
  keywords: ['tarot online', 'trải bài tarot', 'Rider-Waite-Smith', 'bói bài tarot', 'xem tarot miễn phí'],
  openGraph: {
    siteName: 'Tarot',
    locale: 'vi_VN',
    type: 'website',
  },
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="vi">
      <body suppressHydrationWarning>{children}</body>
    </html>
  )
}
