import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import '@/styles/globals.css'
import { cn } from '@/lib/cn'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Densor',
  description: 'Sensor monitoration app.',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en">
      <body
        className={cn(
          'min-h-screen bg-background font-sans antialiased',
          inter.className,
        )}
      >
        {children}
      </body>
    </html>
  )
}
