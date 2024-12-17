import { NextRequest, NextResponse } from 'next/server'
import { checkAuth } from './shared/utils/auth'

const PUBLIC_ROUTES = ['/login']
const NOT_AUTHETICATED_REDIRECT_URL = '/login'
const AUTHETICATED_REDIRECT_URL = '/'

export default async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname
  const isOnPublicRoute = PUBLIC_ROUTES.includes(path)
  const authenticated = !checkAuth(req.cookies.get('auth_token')?.value || '')

  if (!isOnPublicRoute && !authenticated) {
    return NextResponse.redirect(new URL(NOT_AUTHETICATED_REDIRECT_URL, req.nextUrl))
  }

  if (isOnPublicRoute && authenticated) {
    return NextResponse.redirect(new URL(AUTHETICATED_REDIRECT_URL, req.nextUrl))
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}
