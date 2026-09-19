@extends('layouts.app')

@section('judul', 'Top 10 Mahasiswa Teknik Komputer')

@section('konten')
<h1 class="h3 mb-4">Top 10 Mahasiswa IPK Tertinggi - Teknik Komputer</h1>

<table class="table table-striped table-bordered bg-white">
    <thead class="table-dark">
        <tr>
            <th>Peringkat</th>
            <th>NIM</th>
            <th>Nama</th>
            <th>Program Studi</th>
            <th>Angkatan</th>
            <th>IPK</th>
        </tr>
    </thead>
    <tbody>
        @forelse ($topMahasiswa as $index => $mhs)
        <tr>
            <td>{{ $index + 1 }}</td>
            <td>{{ $mhs->nim }}</td>
            <td>{{ $mhs->nama }}</td>
            <td>{{ $mhs->programStudi->nama }}</td>
            <td>{{ $mhs->angkatan }}</td>
            <td><strong>{{ $mhs->ipk }}</strong></td>
        </tr>
        @empty
        <tr>
            <td colspan="6" class="text-center">Data mahasiswa tidak ditemukan.</td>
        </tr>
        @endforelse
    </tbody>
</table>

<a href="{{ route('mahasiswa.data') }}" class="btn btn-secondary mt-3">Kembali</a>
@endsection