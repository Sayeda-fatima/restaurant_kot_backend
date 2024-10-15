<?php

use App\Http\Controllers\API\ProfileController;
use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\API\ProductController;
use App\Http\Controllers\Auth\AuthenticatedSessionController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/dashboard', function () {
    return view('dashboard');
})->middleware(['auth', 'verified'])->name('dashboard');

Route::middleware('auth')->group(function () {
    Route::get('/profile', [ProfileController::class, 'edit'])->name('profile.edit');
    Route::patch('/profile', [ProfileController::class, 'update'])->name('profile.update');
    Route::delete('/profile', [ProfileController::class, 'destroy'])->name('profile.destroy');
});

require __DIR__.'/auth.php';

Route::get('/product', [ProductController::class, 'index'])->name('product.index');
Route::get('/product/create', [ProductController::class, 'create'])->withoutMiddleware(VerifyCsrfToken::class)->name('product.create');
Route::post('/product/store', [ProductController::class, 'store'])->withoutMiddleware(VerifyCsrfToken::class)->name('product.store');
