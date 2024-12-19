import { LoginFields } from '../model/login-schema'
import { toast } from 'sonner'
import { fetcher } from '@/shared/api/fetcher'

export const login = async (data: LoginFields) => {
  try {
    const res = await fetcher('/auth/login', { method: 'POST', body: JSON.stringify(data) })
    if (res.status === 400) {
      toast.error('Неверный код')
      return
    }
    toast.success('Вы успешно вошли в админ панель')
  } catch (error) {
    toast.error('Ошибка входа')
  }
}
