import { cn } from '@/lib/cn'
import { PropsWithChildren } from 'react'

type PaddingContainerProps = PropsWithChildren & {
  className?: string
  verticalPadding?: boolean
}

export function PaddingContainer({
  children,
  verticalPadding = false,
  className,
}: PaddingContainerProps) {
  const builtClassName = verticalPadding
    ? `py-4 px-8 max-w-4xl border-x ${className}`
    : `px-8 max-w-4xl border-x ${className}`

  return <div className={cn(builtClassName)}>{children}</div>
}
