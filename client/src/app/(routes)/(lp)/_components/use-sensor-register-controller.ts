'use client'

import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { api } from '@/lib/axios'
import { useState } from 'react'
import { SensorCategory } from '@/data/types/sensor-category'

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
  const [categories, setCategories] = useState<SensorCategory[]>([])
  const form = useForm<FormData>({
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

  const bootstrap = async () => {
    const { data } = await api.get('/sensors/categories')
    console.log(data)
  }

  function onSubmit(values: FormData) {}

  return { form, onSubmit, bootstrap }
}
