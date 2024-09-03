import { PropsWithChildren } from 'react'

function RequiredFieldIndicator() {
  return (
    <div>
      <span className="text-primary font-medium">*</span>
    </div>
  )
}

export function RequiredFieldContainer({ children }: PropsWithChildren) {
  return (
    <div className="flex items-center gap-1">
      {children}
      <RequiredFieldIndicator />
    </div>
  )
}
