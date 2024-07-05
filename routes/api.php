<?php

use App\Http\Controllers\API\InvoiceDetailsController;
use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken;
use App\Http\Controllers\API\ProductCategoryController;
use App\Http\Controllers\Auth\AuthenticatedSessionController;
use App\Http\Controllers\API\ProductStockController;
use App\Http\Controllers\API\ProfileController;
use App\Http\Controllers\API\InvoiceController;
use App\Http\Controllers\API\ProductController;
use App\Http\Controllers\API\CustomerController;
use App\Http\Controllers\API\SupplierController;
use App\Http\Controllers\API\TransactionController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');
Route::get('/dashboard', function () {
    return view('dashboard');
})->middleware(['auth', 'verified'])->name('dashboard');

Route::middleware('auth')->group(function () {
    Route::get('/profile', [ProfileController::class, 'edit'])->name('profile.edit');
    Route::patch('/profile', [ProfileController::class, 'update'])->name('profile.update');
    Route::delete('/profile', [ProfileController::class, 'destroy'])->name('profile.destroy');
});
Route::post('/login',[AuthenticatedSessionController::class ,'store']);

// require __DIR__.'/auth.php';

// request for invoice 
Route::get('/invoice', [InvoiceController::class, 'getInvoice'])->name('getInvoice.index');

// route for displaying transactions
Route::get('/displayTransactions', [TransactionController::class, 'index'])->name('transaction.index');

// Product 
// route for displaying products
Route::get('/product', [ProductController::class, 'index'])->name('product.index');
// route for creating/adding a new product
Route::get('/product/create', [ProductController::class, 'create'])->withoutMiddleware(VerifyCsrfToken::class)->name('product.create');
Route::post('/product/store', [ProductController::class, 'store'])->withoutMiddleware(VerifyCsrfToken::class)->name('product.store');
// route for editing product
Route::get('/product/{product}/edit', [ProductController::class, 'edit'])->name('product.edit');
Route::put('/product/{product}/update', [ProductController::class, 'update'])->name('product.update');
// route for searching product
Route::get('/product/{search}/search', [ProductController::class, 'searchProduct'])->name('product.search');
// route for deleting product
Route::delete('/product/{product}/delete', [ProductController::class, 'destroy'])->name('product.delete');
// route for all products
Route::get('/product/all', [ProductController::class, 'allProducts'])->name('product.all');
// route for displaying products for a certain category
Route::get('/product/{product}/displayProductsForCategory', [ProductController::class, 'displayProductsForCategory'])->name('product.display');

// Product category
Route::get('/product/category', [ProductCategoryController::class, 'index'])->name('product.category');
Route::get('/product/category/create', [ProductCategoryController::class, 'create'])->name('productCategory.create');
Route::post('/product/category/store', [ProductCategoryController::class, 'store'])->name('productCategory.store');
Route::get('/product/category/edit', [ProductCategoryController::class, 'edit'])->name('productCategory.edit');
Route::put('/product/category/{id}/update', [ProductCategoryController::class, 'update'])->name('productCategory.update');
Route::delete('/product/category/{id}/delete', [ProductCategoryController::class, 'destroy'])->name('productCategory.delete');

// Product Stock
Route::get('/product/stock', [ProductStockController::class, 'index'])->name('product.stock');

// Customer 
// route for displaying customers
Route::get('/customer', [CustomerController::class, 'index'])->name('customer.index');
// route for creating/adding a new customer
Route::get('/customer/create', [CustomerController::class, 'create'])->name('customer.create');
Route::post('/customer/store', [CustomerController::class, 'store'])->name('customer.store');
// route for editing a customer
Route::get('/customer/{customer}/edit', [CustomerController::class, 'edit'])->name('customer.edit');
Route::put('/customer/{customer}/update', [CustomerController::class, 'update'])->name('customer.update');
// route for searching a customer
Route::get('/customer/search', [CustomerController::class, 'searchCustomer'])->name('customer.search');
// route for deleting a customer
Route::delete('/customer/{customer}/delete', [CustomerController::class, 'destroy'])->name('customer.delete');

// Supplier
// route for displaying supplier
Route::get('/supplier', [SupplierController::class, 'index'])->name('supplier.index');
// route for creating/adding a new supplier
Route::get('/supplier/create', [SupplierController::class, 'create'])->name('supplier.create');
Route::post('/supplier/store', [SupplierController::class, 'store'])->name('supplier.store');
// route for editing a supplier
Route::get('/supplier/{supplier}/edit', [SupplierController::class, 'edit'])->name('supplier.edit');
Route::put('/supplier/{supplier}/update', [SupplierController::class, 'update'])->name('supplier.update');
// route for searching a supplier
Route::get('/supplier/search', [SupplierController::class, 'searchSupplier'])->name('supplier.search');
// route for deleting a supplier
Route::delete('/supplier/{supplier}/delete', [SupplierController::class, 'destroy'])->name('supplier.delete');

// Invoice
// route for displaying invoice
Route::get('/invoice', [InvoiceController::class, 'index'])->name('invoice.index');
// route for creating an invoice
Route::get('/invoice/create', [InvoiceController::class, 'create'])->name('invoice.create');
Route::post('/invoice/store', [InvoiceController::class, 'store'])->name('invoice.store');
// route for editing an invoice
Route::get('/invoice/{invoice}/edit', [InvoiceController::class, 'edit'])->name('invoice.edit');
Route::put('/invoice/{invoice}/update', [InvoiceController::class, 'update'])->name('invoice.update');
// route for deleting an invoice
Route::delete('/invoice/{invoice}/delete', [InvoiceController::class, 'destroy'])->name('invoice.delete');

// INvoice details
// route for displaying invoice details
Route::get('/invoice/{invoice}', [InvoiceDetailsController::class, 'index'])->name('invoice.details.index');
