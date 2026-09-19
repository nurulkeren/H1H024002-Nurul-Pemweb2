@extends('layouts.app')

@section('judul', 'Detail Mahasiswa')

@section('konten')
<h1 class="h3 mb-4">Detail Mahasiswa</h1>

<div class="card mb-4">
    <div class="card-body">
        <p><strong>NIM:</strong> {{ $mahasiswa->nim }}</p>
        <p><strong>Nama:</strong> {{ $mahasiswa->nama }}</p>
        <p><strong>Program Studi:</strong> {{ $mahasiswa->programStudi->nama }}</p>
        <p><strong>Angkatan:</strong> {{ $mahasiswa->angkatan }}</p>
        <p class="mb-0"><strong>IPK:</strong> {{ $mahasiswa->ipk }}</p>
    </div>
</div>

<h2 class="h4 mb-3">Daftar Mata Kuliah yang Diambil</h2>

<table class="table table-striped bg-white">
    <thead>
        <tr>
            <th>Kode</th>
            <th>Mata Kuliah</th>
            <th>SKS</th>
            <th>Semester</th>
            <th>Nilai</th>
        </tr>
    </thead>
    <tbody>
        @forelse ($mahasiswa->matakuliahs as $mk)
        <tr>
            <td>{{ $mk->kode }}</td>
            <td>{{ $mk->nama }}</td>
            <td>{{ $mk->sks }}</td>
            <td>{{ $mk->semester }}</td>
            <td>
                <span class="badge bg-primary">{{ $mk->pivot->nilai }}</span>
            </td>
        </tr>
        @empty
        <tr>
            <td colspan="5" class="text-center">Belum mengambil mata kuliah.</td>
        </tr>
        @endforelse
    </tbody>
</table>

<a href="{{ route('mahasiswa.data') }}" class="btn btn-secondary mt-3">
    Kembali
</a>
@endsection