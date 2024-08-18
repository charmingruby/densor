import { Header } from '@/components/header'
import { PropsWithChildren } from 'react'

export default function RootLayout({ children }: PropsWithChildren) {
  return (
    <div className="max-w-4xl mx-auto h-screen max-h-screen overflow-hidden flex flex-col flex-1">
      <Header />

      <div className="flex flex-col flex-1 mt-20 overflow-auto">{children}</div>
    </div>
  )
}
