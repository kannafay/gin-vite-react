import { createBrowserRouter } from 'react-router'
import Loading from '@/components/Loading'
import Home from '@/pages/Home'

const router = createBrowserRouter([
  {
    path: '/',
    Component: Home,
  },
  {
    path: '/about',
    loader: async () => await fetch('/api/about'),
    HydrateFallback: Loading,
    lazy: async () => {
      const { default: About } = await import('@/pages/About')
      return { Component: About }
    },
  },
])

export default router