import { $api } from '@/shared/api/api'
import { LoginFields } from '../model/login-schema'
import { toast } from 'sonner'
import { isAxiosError } from 'axios'

export const login = async (data: LoginFields) => {
  try {
    await $api.post('/auth/login', data)
    toast.success('Вы успешно вошли в админ панель')
  } catch (error) {
    if (isAxiosError(error)) {
      if (error.status === 400) {
        toast.error('Неверный код')
      } else {
        toast.error('Ошибка входа')
      }
    }
  }
}
