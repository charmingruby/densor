'use client'

import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { SensorCategory } from '@/data/types/sensor-category'
import { api } from '@/lib/axios'
import { parseFetchSensorCategoriesResult } from '@/data/api/fetch-sensor-categories'
import { AxiosError } from 'axios'

const formSchema = z.object({
  name: z.string(),
  description: z.string(),
  equipmentName: z.string(),
  sectorName: z.string(),
  categoryName: z.string(),
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
      equipmentName: '',
      categoryName: '',
      sectorName: '',
    },
  })

  const bootstrap = async () => {
    const { data } = await api.get('/sensors/categories')

    const result = parseFetchSensorCategoriesResult(data)

    setCategories(result)
  }
  async function onSubmit(data: FormData) {
    console.log(data)

    try {
      const res = await api.post('/sensors', {
        name: data.name,
        description: data.description,
        equipment_name: data.equipmentName,
        sector_name: data.sectorName,
        sensor_category_name: data.categoryName,
        observation: data.observation,
      })
      console.log(res)
    } catch (err) {
      if (err instanceof AxiosError) {
        console.log(err)
      }
    }
  }

  return { form, onSubmit, bootstrap, categories }
}
