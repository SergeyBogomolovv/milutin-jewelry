'use client'
import { Form, FormControl, FormField, FormItem, FormMessage } from '@/shared/ui/form'
import { useForm } from 'react-hook-form'
import { LoginFields, loginSchema } from '../model/login-schema'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input } from '@/shared/ui/input'
import { Button } from '@/shared/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/ui/card'
import { sendCode } from '../api/send-code'
import { login } from '../api/login'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'
import { useTransition } from 'react'

export function LoginForm() {
  const [isPending, startTransition] = useTransition()
  const router = useRouter()
  const form = useForm<LoginFields>({
    resolver: zodResolver(loginSchema),
    defaultValues: { code: '' },
  })
  const onSubmit = async (data: LoginFields) => {
    const success = await login(data)
    if (!success) {
      toast.error('Неверный код')
      return
    }
    toast.success('Вы успешно вошли в админ панель')
    router.refresh()
  }

  const handleSendCode = async () => {
    const success = await sendCode()
    if (!success) {
      toast.error('Ошибка отправки кода')
      return
    }
    toast.success('Код отправлен')
  }

  return (
    <Card className='w-[400px] max-w-[95%]'>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <CardHeader className='items-center'>
            <CardTitle className='text-2xl'>Вход в админ панель</CardTitle>
            <CardDescription>Код придет на почту администратора</CardDescription>
          </CardHeader>
          <CardContent>
            <FormField
              control={form.control}
              name='code'
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <Input placeholder='Введите шестизначный код' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </CardContent>
          <CardFooter className='flex items-center gap-2'>
            <Button disabled={form.formState.isSubmitting} type='submit' className='w-full'>
              Подтвердить
            </Button>
            <Button
              onClick={() => startTransition(handleSendCode)}
              variant='outline'
              type='button'
              className='w-full'
              disabled={isPending}
            >
              {isPending ? 'Отправка...' : 'Получить код'}
            </Button>
          </CardFooter>
        </form>
      </Form>
    </Card>
  )
}
