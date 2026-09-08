@extends('layouts.app')

@section('judul', 'Daftar Matakuliah')

@section('konten')

<div class="container-fluid">

    <!-- Header -->
    <div class="d-flex justify-content-between align-items-center mb-4">
        <div>
            <h1 class="h3 fw-bold mb-1">Daftar Matakuliah</h1>
            <p class="text-muted mb-0">
                Kelola dan lihat daftar matakuliah yang tersedia.
            </p>
        </div>
    </div>

    <!-- Search Card -->
    <div class="card border-0 shadow-sm mb-4">
        <div class="card-body">

            <form action="{{ route('matakuliah.cari') }}" method="GET">

                <div class="row g-2">

                    <div class="col-md-10">
                        <input
                            type="text"
                            name="q"
                            class="form-control"
                            placeholder="Cari nama matakuliah..."
                            value="{{ $kataKunci ?? '' }}"
                        >
                    </div>

                    <div class="col-md-2">
                        <button
                            type="submit"
                            class="btn btn-primary w-100"
                        >
                            🔍 Cari
                        </button>
                    </div>

                </div>

            </form>

        </div>
    </div>

    <!-- Information -->
    <x-kartu-info judul="Informasi">
        Data matakuliah pada halaman ini masih menggunakan array statis.
    </x-kartu-info>

    <!-- Table Card -->
    <div class="card border-0 shadow-sm">

        <div class="card-body">

            <div class="d-flex justify-content-between align-items-center mb-3">

                <div>
                    <h5 class="fw-bold mb-1">Data Matakuliah</h5>

                    @if(isset($kataKunci) && $kataKunci !== '')
                        <small class="text-muted">
                            Hasil pencarian untuk:
                            <strong>"{{ $kataKunci }}"</strong>
                        </small>
                    @else
                        <small class="text-muted">
                            Menampilkan seluruh matakuliah
                        </small>
                    @endif

                </div>

            </div>

            <div class="table-responsive">

                <table class="table table-hover align-middle mb-0">

                    <thead class="table-light">
                        <tr>
                            <th>Kode</th>
                            <th>Nama Matakuliah</th>
                            <th>SKS</th>
                            <th class="text-center">Aksi</th>
                        </tr>
                    </thead>

                    <tbody>

                        @forelse ($daftarMatakuliah as $matakuliah)

                            <tr>

                                <td>
                                    <span class="fw-semibold">
                                        {{ $matakuliah['kode'] }}
                                    </span>
                                </td>

                                <td>
                                    {{ $matakuliah['nama'] }}
                                </td>

                                <td>
                                    <x-badge-sks :sks="$matakuliah['sks']" />
                                </td>

                                <td class="text-center">

                                    <a
                                        href="{{ route('matakuliah.show', $matakuliah['kode']) }}"
                                        class="btn btn-sm btn-outline-primary"
                                    >
                                        Detail
                                    </a>

                                </td>

                            </tr>

                        @empty

                            <tr>
                                <td colspan="4" class="text-center py-5">

                                    <div class="text-muted">

                                        <div class="fs-1 mb-2">
                                            🔍
                                        </div>

                                        <h6 class="fw-bold">
                                            Matakuliah tidak ditemukan
                                        </h6>

                                        <p class="mb-0">
                                            Coba gunakan kata kunci lain.
                                        </p>

                                    </div>

                                </td>
                            </tr>

                        @endforelse

                    </tbody>

                </table>

            </div>

        </div>

    </div>

</div>

@endsection