<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class PeranAdmin
{
    /**
     * Handle an incoming request.
     *
     * @param  Closure(Request): (Response)  $next
     */
    public function handle($request, \Closure $next)
    {
        if ($request->user()?->peran !== 'admin') {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Akses hanya untuk admin',
            ], 403);
        }

        return $next($request);
    }
}
