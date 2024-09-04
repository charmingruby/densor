'use client'

import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { SensorCategory } from '@/data/types/sensor-category'
import { api } from '@/lib/axios'
import { parseFetchSensorCategoriesResult } from '@/data/api/fetch-sensor-categories'
import { AxiosError } from 'axios'
import { useToast } from '../use-toast'
import { useRouter } from 'next/navigation'

const formSchema = z.object({
  name: z
    .string({ required_error: 'Nome obrigatório' })
    .min(1, 'Nome obrigatório'),
  description: z
    .string({ required_error: 'Descrição obrigatória' })
    .min(1, 'Descrição obrigatória'),
  categoryName: z
    .string({ required_error: 'Categoria obrigatória' })
    .min(1, 'Categoria obrigatória'),
  equipmentName: z.string(),
  sectorName: z.string(),
  observation: z.string(),
})

type FormData = z.infer<typeof formSchema>

export function useSensorRegisterController() {
  const [categories, setCategories] = useState<SensorCategory[]>([])
  const { toast } = useToast()
  const { push } = useRouter()
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

      toast({
        title: 'Sensor criado com sucesso.',
        variant: 'success',
      })

      push('/sensors')
    } catch (err) {
      if (err instanceof AxiosError) {
        console.log(err)
        toast({
          title: err.response?.data.message + '.',
          variant: 'destructive',
        })
      }
    }
  }

  const nameFilled = !!form.watch('name')
  const descriptionFilled = !!form.watch('description')
  const categoryNameFilled = !!form.watch('categoryName')
  const equipmentNameFilled = !!form.watch('equipmentName')
  const sectorNameFilled = !!form.watch('sectorName')

  const isFormReady =
    nameFilled &&
    descriptionFilled &&
    categoryNameFilled &&
    (sectorNameFilled || equipmentNameFilled)

  return { form, onSubmit, bootstrap, isFormReady, categories }
}
