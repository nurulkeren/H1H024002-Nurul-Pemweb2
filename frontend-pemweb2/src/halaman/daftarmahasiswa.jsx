
import { useEffect, useState } from 'react'

import {
  ambilDaftarMahasiswa,
  hapusMahasiswa
} from '../api/layananMahasiswa'

import FormTambahMahasiswa from '../komponen/FormTambahMahasiswa'

export default function DaftarMahasiswa() {
  const [daftar, setDaftar] = useState([])
  const [meta, setMeta] = useState(null)
  const [halaman, setHalaman] = useState(1)
  const [cari, setCari] = useState('')
  const [sedangMemuat, setSedangMemuat] = useState(false)
  const [pesanGalat, setPesanGalat] = useState('')
  const [tampilForm, setTampilForm] = useState(false)
  const [refresh, setRefresh] = useState(0)

  // Memuat daftar mahasiswa
  useEffect(() => {
    let dibatalkan = false

    async function muatData() {
      setSedangMemuat(true)
      setPesanGalat('')

      try {
        const hasil = await ambilDaftarMahasiswa({
          page: halaman,
          cari: cari,
          per_halaman: 10,
        })

        if (!dibatalkan) {
          setDaftar(hasil.data || [])
          setMeta(hasil.meta || null)
        }
      } catch (galat) {
        if (!dibatalkan) {
          setPesanGalat('Gagal memuat data mahasiswa')
        }
      } finally {
        if (!dibatalkan) {
          setSedangMemuat(false)
        }
      }
    }

    const penunda = setTimeout(muatData, 400)

    return () => {
      dibatalkan = true
      clearTimeout(penunda)
    }
  }, [halaman, cari, refresh])

  // Setelah mahasiswa berhasil ditambahkan
  function tanganiBerhasil() {
    setTampilForm(false)
    setHalaman(1)
    setCari('')
    setRefresh((sebelumnya) => sebelumnya + 1)
  }

  // Menghapus mahasiswa
  async function tanganiHapus(id, nama) {
    const konfirmasi = window.confirm(
      `Apakah kamu yakin ingin menghapus mahasiswa ${nama}?`
    )

    if (!konfirmasi) return

    try {
      await hapusMahasiswa(id)

      alert('Data mahasiswa berhasil dihapus!')

      // Jika halaman saat ini kosong setelah penghapusan,
      // kembali ke halaman pertama.
      if (daftar.length === 1 && halaman > 1) {
        setHalaman((sebelumnya) => sebelumnya - 1)
      } else {
        setRefresh((sebelumnya) => sebelumnya + 1)
      }
    } catch (galat) {
      alert(
        galat.response?.data?.pesan ||
        'Gagal menghapus data mahasiswa.'
      )
    }
  }

  return (
    <div style={{ maxWidth: 900, margin: '40px auto' }}>

      {tampilForm ? (
        <FormTambahMahasiswa
          onBerhasil={tanganiBerhasil}
          onBatal={() => setTampilForm(false)}
        />
      ) : (
        <>
          <h1>Daftar Mahasiswa</h1>

          <button
            onClick={() => setTampilForm(true)}
            style={{ marginBottom: 16 }}
          >
            + Tambah Mahasiswa
          </button>

          <br />

          <input
            type="search"
            placeholder="Cari nama atau NIM"
            value={cari}
            onChange={(peristiwa) => {
              setHalaman(1)
              setCari(peristiwa.target.value)
            }}
          />

          {sedangMemuat && (
            <p>Memuat data...</p>
          )}

          {pesanGalat && (
            <p style={{ color: 'crimson' }}>
              {pesanGalat}
            </p>
          )}

          {!sedangMemuat &&
            !pesanGalat &&
            daftar.length === 0 && (
              <p>Data tidak ditemukan</p>
            )}

          <table
            border="1"
            cellPadding="8"
            style={{
              width: '100%',
              marginTop: 16,
              borderCollapse: 'collapse'
            }}
          >
            <thead>
              <tr>
                <th>NIM</th>
                <th>Nama</th>
                <th>Angkatan</th>
                <th>IPK</th>
                <th>Aksi</th>
              </tr>
            </thead>

            <tbody>
              {daftar.map((mahasiswa) => (
                <tr key={mahasiswa.id}>
                  <td>{mahasiswa.nim}</td>
                  <td>{mahasiswa.nama}</td>
                  <td>{mahasiswa.angkatan}</td>
                  <td>{mahasiswa.ipk}</td>

                  <td>
                    <button
                      onClick={() =>
                        tanganiHapus(
                          mahasiswa.id,
                          mahasiswa.nama
                        )
                      }
                    >
                      Hapus
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>

          {meta && (
            <div style={{ marginTop: 16 }}>

              <button
                onClick={() => setHalaman(halaman - 1)}
                disabled={halaman <= 1 || sedangMemuat}
              >
                Sebelumnya
              </button>

              <span style={{ margin: '0 12px' }}>
                Halaman {meta.current_page} dari {meta.last_page}
              </span>

              <button
                onClick={() => setHalaman(halaman + 1)}
                disabled={
                  meta.current_page >= meta.last_page ||
                  sedangMemuat
                }
              >
                Berikutnya
              </button>

            </div>
          )}
        </>
      )}

    </div>
  )
}
