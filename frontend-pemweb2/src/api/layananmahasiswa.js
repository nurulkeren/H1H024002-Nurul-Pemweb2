
import klien from './klien'

const backendAktif = import.meta.env.VITE_BACKEND_AKTIF || 'laravel'

export async function ambilDaftarMahasiswa(parameter = {}) {
  const parameterAPI = { ...parameter }

  if (backendAktif === 'fiber') {
    parameterAPI.halaman = parameterAPI.page || 1
    delete parameterAPI.page
  }

  const respons = await klien.get('/mahasiswa/', {
    params: parameterAPI,
  })

  const hasil = respons.data

  // Samakan struktur metadata Fiber dengan Laravel
  if (backendAktif === 'fiber' && hasil.meta) {
    hasil.meta = {
      ...hasil.meta,
      current_page: hasil.meta.halaman,
      last_page: hasil.meta.total_halaman,
    }
  }

  return hasil
}

export async function ambilDetailMahasiswa(id) {
  const respons = await klien.get(`/mahasiswa/${id}`)
  return respons.data
}

export async function simpanMahasiswa(data) {
  const respons = await klien.post('/mahasiswa/', data)
  return respons.data
}

export async function hapusMahasiswa(id) {
  const respons = await klien.delete(`/mahasiswa/${id}`)
  return respons.data
}
