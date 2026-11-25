
import {
  BrowserRouter,
  Route,
  Routes,
  Navigate
} from 'react-router-dom'

import Login from './halaman/login.jsx'
import DaftarMahasiswa from './halaman/daftarmahasiswa.jsx'
import RuteTerlindungi from './komponen/RuteTerlindungi.jsx'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={<Navigate to="/mahasiswa" replace />}
        />

        <Route
          path="/login"
          element={<Login />}
        />

        <Route
          path="/mahasiswa"
          element={
            <RuteTerlindungi>
              <DaftarMahasiswa />
            </RuteTerlindungi>
          }
        />
      </Routes>
    </BrowserRouter>
  )
}
