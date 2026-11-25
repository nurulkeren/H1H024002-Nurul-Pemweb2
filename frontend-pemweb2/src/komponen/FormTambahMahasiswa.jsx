
import { useState } from 'react'
import { simpanMahasiswa } from '../api/layananMahasiswa'

export default function FormTambahMahasiswa({ onBerhasil, onBatal }) {
  const [form, setForm] = useState({
    program_studi_id: '',
    nim: '',
    nama: '',
    email: '',
    angkatan: '',
    ipk: '',
    aktif: true,
  })

  const [galat, setGalat] = useState({})
  const [pesanGalat, setPesanGalat] = useState('')
  const [sedangMenyimpan, setSedangMenyimpan] = useState(false)

  function tanganiPerubahan(e) {
    const { name, value, type, checked } = e.target

    setForm((sebelumnya) => ({
      ...sebelumnya,
      [name]: type === 'checkbox' ? checked : value,
    }))

    setGalat((sebelumnya) => ({
      ...sebelumnya,
      [name]: [],
    }))
  }

  async function tanganiKirim(e) {
    e.preventDefault()

    setGalat({})
    setPesanGalat('')
    setSedangMenyimpan(true)

    try {
      const data = {
        ...form,
        program_studi_id: Number(form.program_studi_id),
        angkatan: Number(form.angkatan),
        ipk: form.ipk === '' ? null : Number(form.ipk),
      }

      await simpanMahasiswa(data)

      alert('Data mahasiswa berhasil ditambahkan!')
      onBerhasil()
    } catch (error) {
      if (error.response?.status === 422) {
        setGalat(error.response.data.errors || {})
        setPesanGalat('Periksa kembali data yang kamu masukkan.')
      } else {
        setPesanGalat(
          error.response?.data?.pesan ||
          'Gagal menyimpan data mahasiswa.'
        )
      }
    } finally {
      setSedangMenyimpan(false)
    }
  }

  return (
    <div style={{ maxWidth: 600, margin: '20px auto' }}>
      <h2>Tambah Mahasiswa</h2>

      {pesanGalat && (
        <p style={{ color: 'crimson' }}>
          {pesanGalat}
        </p>
      )}

      <form onSubmit={tanganiKirim}>
        <div>
          <label>ID Program Studi</label>
          <input
            type="number"
            name="program_studi_id"
            value={form.program_studi_id}
            onChange={tanganiPerubahan}
            required
          />
          {galat.program_studi_id?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <div>
          <label>NIM</label>
          <input
            type="text"
            name="nim"
            maxLength="20"
            value={form.nim}
            onChange={tanganiPerubahan}
            required
          />
          {galat.nim?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <div>
          <label>Nama</label>
          <input
            type="text"
            name="nama"
            maxLength="100"
            value={form.nama}
            onChange={tanganiPerubahan}
            required
          />
          {galat.nama?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <div>
          <label>Email</label>
          <input
            type="email"
            name="email"
            value={form.email}
            onChange={tanganiPerubahan}
            required
          />
          {galat.email?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <div>
          <label>Angkatan</label>
          <input
            type="number"
            name="angkatan"
            min="2000"
            max="2100"
            value={form.angkatan}
            onChange={tanganiPerubahan}
            required
          />
          {galat.angkatan?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <div>
          <label>IPK</label>
          <input
            type="number"
            name="ipk"
            min="0"
            max="4"
            step="0.01"
            value={form.ipk}
            onChange={tanganiPerubahan}
          />
          {galat.ipk?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <div>
          <label>
            <input
              type="checkbox"
              name="aktif"
              checked={form.aktif}
              onChange={tanganiPerubahan}
            />
            Mahasiswa Aktif
          </label>
          {galat.aktif?.map((pesan, i) => (
            <p key={i} style={{ color: 'crimson' }}>{pesan}</p>
          ))}
        </div>

        <button type="submit" disabled={sedangMenyimpan}>
          {sedangMenyimpan ? 'Menyimpan...' : 'Simpan Mahasiswa'}
        </button>

        <button type="button" onClick={onBatal}>
          Batal
        </button>
      </form>
    </div>
  )
}
