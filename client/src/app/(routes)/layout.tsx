import { Header } from '@/components/header'
import { PropsWithChildren } from 'react'

export default function RootLayout({ children }: PropsWithChildren) {
  return (
    <div className="max-w-4xl mx-auto border-x min-h-screen">
      <Header />
      <div>{children}</div>
    </div>
  )
}
