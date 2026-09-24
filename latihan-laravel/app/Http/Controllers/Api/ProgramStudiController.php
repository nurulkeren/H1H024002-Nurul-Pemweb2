<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\MahasiswaResource;
use App\Models\ProgramStudi;
use Illuminate\Http\Request;

class ProgramStudiController extends Controller
{
    public function mahasiswa(Request $request, ProgramStudi $programStudi)
    {
        $perHalaman = min(
            max((int) $request->get('per_halaman', 10), 1),
            100
        );

        $mahasiswa = $programStudi->mahasiswas()
            ->with('programStudi')
            ->paginate($perHalaman);

        return MahasiswaResource::collection($mahasiswa);
    }
}