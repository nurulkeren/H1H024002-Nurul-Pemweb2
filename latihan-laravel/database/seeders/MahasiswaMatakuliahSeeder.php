<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class MahasiswaMatakuliahSeeder extends Seeder
{
    public function run(): void
    {
        $mahasiswas = Mahasiswa::take(10)->get();
        $matakuliahs = Matakuliah::all();

        foreach ($mahasiswas as $mhs) {
            foreach ($matakuliahs as $mk) {
                $mhs->matakuliahs()->attach($mk->id, [
                    'nilai' => collect(['A', 'B+', 'B', 'C+'])->random()
                ]);
            }
        }
    }
}