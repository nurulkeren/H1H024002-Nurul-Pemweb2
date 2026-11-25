
import klien from './klien'

const backendAktif = import.meta.env.VITE_BACKEND_AKTIF || 'laravel'

export async function login(email, kataSandi) {
  const dataLogin = {
    email,
    [backendAktif === 'fiber' ? 'kata_sandi' : 'password']: kataSandi,
  }

  const respons = await klien.post('/auth/login', dataLogin)

  return respons.data
}

export async function ambilProfil() {
  const respons = await klien.get('/auth/profil')
  return respons.data
}

export async function logout() {
  const respons = await klien.post('/auth/logout')
  return respons.data
}
