<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index()
    {
        $daftarMatakuliah = [
            [
                'kode' => 'TK101',
                'nama' => 'Pemrograman Dasar',
                'sks' => 3
            ],
            [
                'kode' => 'TK102',
                'nama' => 'Pemrograman Web II',
                'sks' => 3
            ],
            [
                'kode' => 'TK103',
                'nama' => 'Sistem Operasi',
                'sks' => 3
            ],
            [
                'kode' => 'TK104',
                'nama' => 'Jaringan Komputer',
                'sks' => 4
            ],
            [
                'kode' => 'TK105',
                'nama' => 'Basis Data',
                'sks' => 2
            ],
        ];

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah
        ]);
    }

    public function show(string $kode)
    {
        return view('matakuliah.show', [
            'kode' => $kode
        ]);
    }

    public function cari(Request $request)
    {
        $kataKunci = $request->query('q', '');

        $daftarMatakuliah = [
            [
                'kode' => 'TK101',
                'nama' => 'Pemrograman Dasar',
                'sks' => 3
            ],
            [
                'kode' => 'TK102',
                'nama' => 'Pemrograman Web II',
                'sks' => 3
            ],
            [
                'kode' => 'TK103',
                'nama' => 'Sistem Operasi',
                'sks' => 3
            ],
            [
                'kode' => 'TK104',
                'nama' => 'Jaringan Komputer',
                'sks' => 4
            ],
            [
                'kode' => 'TK105',
                'nama' => 'Basis Data',
                'sks' => 2
            ],
        ];

        $hasil = array_filter($daftarMatakuliah, function ($matakuliah) use ($kataKunci) {
            return stripos($matakuliah['nama'], $kataKunci) !== false;
        });

        return view('matakuliah.index', [
            'daftarMatakuliah' => $hasil,
            'kataKunci' => $kataKunci
        ]);
    }
}