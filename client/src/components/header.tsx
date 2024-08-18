import { AudioWaveform } from 'lucide-react'
import { PaddingContainer } from './padding-container'

export function Header() {
  return (
    <div className="border-b">
      <PaddingContainer verticalPadding className="flex items-center gap-2">
        <div>
          <AudioWaveform className="text-primary h-7 w-7" />
        </div>
        <strong className="text-2xl">Densor</strong>
      </PaddingContainer>
    </div>
  )
}
