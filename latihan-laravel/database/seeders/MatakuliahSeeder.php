<?php

namespace Database\Seeders;

use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class MatakuliahSeeder extends Seeder
{
    public function run(): void
    {
        Matakuliah::create(['kode' => 'MK01', 'nama' => 'Pemrograman Web II', 'sks' => 3, 'semester' => 4]);
        Matakuliah::create(['kode' => 'MK02', 'nama' => 'Basis Data', 'sks' => 3, 'semester' => 2]);
        Matakuliah::create(['kode' => 'MK03', 'nama' => 'Jaringan Komputer', 'sks' => 3, 'semester' => 3]);
        Matakuliah::create(['kode' => 'MK04', 'nama' => 'Sistem Operasi', 'sks' => 2, 'semester' => 3]);
    }
}