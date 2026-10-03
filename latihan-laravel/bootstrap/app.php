<?php

use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;

return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        web: __DIR__.'/../routes/web.php',
        api: __DIR__.'/../routes/api.php',
        commands: __DIR__.'/../routes/console.php',
        health: '/up',
    )
    ->withMiddleware(function (Middleware $middleware): void {
        $middleware->alias([
            'ability' => \Laravel\Sanctum\Http\Middleware\CheckAbilities::class,
            'admin' => \App\Http\Middleware\PeranAdmin::class,
        ]);
        $middleware->redirectGuestsTo(function ($request) {
            if ($request->is('api/*')) {
                return null;
            }
        
            return route('login');
        });
    })
    ->withExceptions(function (Exceptions $exceptions): void {
        $exceptions->render(function (
            \Illuminate\Auth\AuthenticationException $e,
            $request
        ) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Token tidak valid atau belum dikirim',
                ], 401);
            }
        });
    })->create();