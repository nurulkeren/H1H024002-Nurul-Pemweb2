<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMatakuliahRequest;
use App\Http\Requests\UpdateMatakuliahRequest;
use App\Http\Resources\MatakuliahResource;
use App\Models\Matakuliah;
use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $query = Matakuliah::query();

        // Pencarian berdasarkan kode atau nama
        if ($request->filled('cari')) {
            $cari = $request->cari;

            $query->where(function ($q) use ($cari) {
                $q->where('kode', 'like', "%{$cari}%")
                    ->orWhere('nama', 'like', "%{$cari}%");
            });
        }

        // Filter berdasarkan semester
        if ($request->filled('semester')) {
            $query->where('semester', $request->semester);
        }

        // Sorting
        $urut = $request->get('urut', 'kode');
        $arah = $request->get('arah', 'asc');

        $kolomDiizinkan = ['kode', 'nama', 'sks', 'semester'];

        if (in_array($urut, $kolomDiizinkan)) {
            $query->orderBy(
                $urut,
                in_array($arah, ['asc', 'desc']) ? $arah : 'asc'
            );
        }

        // Pagination
        $perHalaman = min(
            max((int) $request->get('per_halaman', 10), 1),
            100
        );

        $matakuliah = $query->paginate($perHalaman);

        return MatakuliahResource::collection($matakuliah);
    }

    public function store(StoreMatakuliahRequest $request)
    {
        $matakuliah = Matakuliah::create($request->validated());

        return (new MatakuliahResource($matakuliah))
            ->additional([
                'sukses' => true,
                'pesan' => 'Data mata kuliah berhasil dibuat',
            ])
            ->response()
            ->setStatusCode(201);
    }

    public function show(Matakuliah $matakuliah)
    {
        return new MatakuliahResource($matakuliah);
    }

    public function update(
        UpdateMatakuliahRequest $request,
        Matakuliah $matakuliah
    ) {
        $matakuliah->update($request->validated());

        return (new MatakuliahResource($matakuliah))
            ->additional([
                'sukses' => true,
                'pesan' => 'Data mata kuliah berhasil diperbarui',
            ]);
    }

    public function destroy(Matakuliah $matakuliah)
    {
        $matakuliah->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dihapus',
        ]);
    }
}