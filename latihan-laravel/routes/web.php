<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MahasiswaController;
use App\Http\Controllers\MatakuliahController;
use App\Http\Controllers\MahasiswaWebController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/salam', function () {
    return 'Selamat datang di Pemrograman Web II';
});

Route::get('/semester/{angka}', function (int $angka) {
    return 'Semester ke ' . $angka;
})->whereNumber('angka');

// Route Praktikum & Tugas (MahasiswaWebController)
Route::get('/mahasiswa-data', [MahasiswaWebController::class, 'index'])->name('mahasiswa.data');
Route::get('/mahasiswa-top-tekkom', [MahasiswaWebController::class, 'topTeknikKomputer'])->name('mahasiswa.top');
Route::get('/mahasiswa/{id}', [MahasiswaWebController::class, 'show'])->name('mahasiswa.show');

// Route Lama (MahasiswaController & MatakuliahController)
Route::get('/data-mahasiswa', [MahasiswaController::class, 'index'])->name('mahasiswa.index');
Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class, 'show'])->name('mahasiswa.detail_lama');
Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']);

Route::get('/matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');
Route::get('/matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');
Route::get('/cari-matakuliah', [MatakuliahController::class, 'cari'])->name('matakuliah.cari');