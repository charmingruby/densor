'use client'

import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useRouter } from 'next/navigation'

const formSchema = z.object({
  name: z.string(),
  description: z.string(),
  equipmentId: z.string(),
  sectorId: z.string(),
  categoryId: z.string(),
  observation: z.string(),
})

type FormData = z.infer<typeof formSchema>

export function useSensorRegisterController() {
  const { push } = useRouter()

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: '',
      description: '',
      observation: '',
      equipmentId: '',
      categoryId: '',
      sectorId: '',
    },
  })

  function onSubmit(values: FormData) {
    console.log(values)
    push('/sensors')
  }

  return { form, onSubmit }
}
