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
  CardHeader,
  CardTitle,
} from '@/shared/ui/card'
import { login } from '../api/login'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'

export function LoginForm() {
  const router = useRouter()
  const form = useForm<LoginFields>({
    resolver: zodResolver(loginSchema),
    defaultValues: { email: '', password: '' },
  })
  const onSubmit = async (data: LoginFields) => {
    try {
      await login(data)
      toast.success('Вы успешно вошли в админ панель')
      router.refresh()
    } catch (error) {
      toast.error('Неверный email или пароль')
    }
  }

  return (
    <Card className='w-[400px] max-w-[95%]'>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <CardHeader className='items-center'>
            <CardTitle className='text-2xl'>Вход в админ панель</CardTitle>
            <CardDescription>Введите email и пароль администратора</CardDescription>
          </CardHeader>
          <CardContent className='space-y-4'>
            <FormField
              control={form.control}
              name='email'
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <Input type='email' placeholder='Email' autoComplete='username' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='password'
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <Input
                      type='password'
                      placeholder='Пароль'
                      autoComplete='current-password'
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </CardContent>
          <div className='px-6 pb-6'>
            <Button disabled={form.formState.isSubmitting} type='submit' className='w-full'>
              Войти
            </Button>
          </div>
        </form>
      </Form>
    </Card>
  )
}
