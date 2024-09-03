import { PaddingContainer } from '@/components/padding-container'
import { Button } from '@/components/ui/button'
import Link from 'next/link'

export function RegisterLink() {
  return (
    <div className="mt-auto w-full fixed bottom-0 z-10">
      <PaddingContainer className="bg-background border-t h-20 flex items-center justify-end">
        <Button className="gap-1" asChild>
          <Link prefetch={false} href="/">
            <span>Registrar</span>
          </Link>
        </Button>
      </PaddingContainer>
    </div>
  )
}
