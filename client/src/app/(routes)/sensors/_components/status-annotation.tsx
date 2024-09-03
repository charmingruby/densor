import { CircleDashed } from 'lucide-react'

interface StatusAnnotationProps {
  status: string
}

export function StatusAnnotation({ status }: StatusAnnotationProps) {
  return (
    <div className="flex items-center gap-1">
      <div className="flex items-center gap-1">
        <CircleDashed width={10} strokeWidth={3} className="text-primary" />
        <strong className="text-xs">Status:</strong>
      </div>
      <span className="text-xs font-medium text-primary">{status}</span>
    </div>
  )
}
