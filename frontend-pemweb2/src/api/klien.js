
import axios from 'axios'

const backendAktif = import.meta.env.VITE_BACKEND_AKTIF || 'laravel'

const alamatAPI = {
  laravel: import.meta.env.VITE_API_LARAVEL,
  fiber: import.meta.env.VITE_API_FIBER,
}

if (!alamatAPI[backendAktif]) {
  throw new Error(`Backend ${backendAktif} tidak dikenali`)
}

const klien = axios.create({
  baseURL: alamatAPI[backendAktif],
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
    ...(backendAktif === 'fiber'
      ? {
          'X-Api-Key': import.meta.env.VITE_FIBER_API_KEY,
        }
      : {}),
  },
  timeout: 10000,
})

klien.interceptors.request.use((konfigurasi) => {
  const token = sessionStorage.getItem('token')

  if (token) {
    konfigurasi.headers.Authorization = `Bearer ${token}`
  }

  return konfigurasi
})


klien.interceptors.response.use(
    (respons) => respons,
    (galat) => {
      if (
        backendAktif === 'fiber' &&
        galat.response?.data?.galat
      ) {
        galat.response.data.errors = galat.response.data.galat
      }
  
      // Hanya hapus token dan redirect jika bukan proses login
      const sedangLogin = galat.config?.url?.includes('/auth/login')
  
      if (galat.response?.status === 401 && !sedangLogin) {
        sessionStorage.removeItem('token')
        window.location.href = '/login'
      }
  
      return Promise.reject(galat)
    }
  )
  
  export default klien
  