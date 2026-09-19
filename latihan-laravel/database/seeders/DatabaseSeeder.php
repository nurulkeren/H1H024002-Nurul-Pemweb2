<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    public function run(): void
    {
        // 1. Buat Data Program Studi
        $this->call(ProgramStudiSeeder::class);

        // 2. Buat 30 Data Mahasiswa (Wajib ada agar tabel terisi data!)
        Mahasiswa::factory()->count(30)->create();

        // 3. Buat Data Matakuliah
        $this->call(MatakuliahSeeder::class);

        // 4. Hubungkan Mahasiswa dan Matakuliah beserta Nilainya
        $this->call(MahasiswaMatakuliahSeeder::class);

        // (Opsional) Akun User
        User::factory()->create([
            'name' => 'Test User',
            'email' => 'test@example.com',
        ]);
    }
}