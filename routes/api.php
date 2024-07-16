<?php

use App\Http\Controllers\API\ExpenseController;
use App\Http\Controllers\API\InvoiceDetailsController;
use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken;
use App\Http\Controllers\API\ProductCategoryController;
use App\Http\Controllers\API\Auth\AuthenticatedSessionController;
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

// Product 
// route for displaying products
Route::get('/product', [ProductController::class, 'index'])->name('product.index')->middleware('auth:api');
// route for creating/adding a new product
Route::get('/product/create', [ProductController::class, 'create'])->withoutMiddleware(VerifyCsrfToken::class)->name('product.create')->middleware('auth:api');
Route::post('/product/store', [ProductController::class, 'store'])->name('product.store')->middleware('auth:api');
// route for editing product
Route::get('/product/{product}/edit', [ProductController::class, 'edit'])->name('product.edit')->middleware('auth:api');
Route::put('/product/{product}/update', [ProductController::class, 'update'])->name('product.update')->middleware('auth:api');
// route for searching product
Route::get('/product/{search}/search', [ProductController::class, 'searchProduct'])->name('product.search')->middleware('auth:api');
// route for deleting product
Route::delete('/product/{product}/delete', [ProductController::class, 'destroy'])->name('product.delete')->middleware('auth:api');
// route for all products
Route::get('/product/all', [ProductController::class, 'allProducts'])->name('product.all')->middleware('auth:api');
// route for displaying products for a certain category
Route::get('/product/{product}/displayProductsForCategory', [ProductController::class, 'displayProductsForCategory'])->name('product.display')->middleware('auth:api');

// route for displaying product sales report
Route::get('/product/report', [ProductController::class, 'productSaleReport'])->middleware('auth:api');
// route for displaying product stock summary report
Route::get('/product/productStockReport', [ProductController::class, 'productStockReport'])->middleware('auth:api');
// route for displaying product detail report
Route::get('/product/productDetailReport', [ProductController::class, 'productDetailReport'])->middleware('auth:api');

// Product category
Route::get('/product/category', [ProductCategoryController::class, 'index'])->name('product.category')->middleware('auth:api');
Route::get('/product/category/create', [ProductCategoryController::class, 'create'])->name('productCategory.create')->middleware('auth:api');
Route::post('/product/category/store', [ProductCategoryController::class, 'store'])->name('productCategory.store')->middleware('auth:api');
Route::get('/product/category/edit', [ProductCategoryController::class, 'edit'])->name('productCategory.edit')->middleware('auth:api');
Route::put('/product/category/{id}/update', [ProductCategoryController::class, 'update'])->name('productCategory.update')->middleware('auth:api');
Route::delete('/product/category/{id}/delete', [ProductCategoryController::class, 'destroy'])->name('productCategory.delete')->middleware('auth:api');

// Product Stock
Route::get('/product/stock', [ProductStockController::class, 'index'])->name('product.stock')->middleware('auth:api');
// route for creating stock
Route::post('/product/stock/store', [ProductStockController::class, 'store'])->name('product.stock.store')->middleware('auth:api');
// route for updating stock
Route::put('/product/stock/update', [ProductStockController::class, 'update'])->name('product.stock.update')->middleware('auth:api');
// route for deleting stock
Route::delete('/product/stock/delete', [ProductStockController::class, 'destroy'])->name('product.stock.delete')->middleware('auth:api');

// Customer 
// route for displaying customers
Route::get('/customer', [CustomerController::class, 'index'])->name('customer.index')->middleware('auth:api');
// route for creating/adding a new customer
Route::get('/customer/create', [CustomerController::class, 'create'])->name('customer.create')->middleware('auth:api');
Route::post('/customer/store', [CustomerController::class, 'store'])->name('customer.store')->middleware('auth:api');
// route for editing a customer
Route::get('/customer/{customer}/edit', [CustomerController::class, 'edit'])->name('customer.edit')->middleware('auth:api');
Route::put('/customer/{customer}/update', [CustomerController::class, 'update'])->name('customer.update')->middleware('auth:api');
// route for searching a customer
Route::get('/customer/search', [CustomerController::class, 'searchCustomer'])->name('customer.search')->middleware('auth:api');
// route for deleting a customer
Route::delete('/customer/{customer}/delete', [CustomerController::class, 'destroy'])->name('customer.delete')->middleware('auth:api');

// route for displaying customer details report
Route::get('/customer/report', [CustomerController::class, 'detailReport'])->name('customer.report')->middleware('auth:api');
// route for displaying customer invoice report
Route::get('/customer/invoiceReport', [CustomerController::class, 'invoiceReport'])->name('customer.invoice.report')->middleware('auth:api');
// route for displaying customer invoice details report
Route::get('/customer/invoiceDetailReport', [CustomerController::class, 'invoiceDetailReport'])->name('customer.invoice.report')->middleware('auth:api');


// Supplier
// route for displaying supplier
Route::get('/supplier', [SupplierController::class, 'index'])->name('supplier.index')->middleware('auth:api');
// route for creating/adding a new supplier
Route::get('/supplier/create', [SupplierController::class, 'create'])->name('supplier.create')->middleware('auth:api');
Route::post('/supplier/store', [SupplierController::class, 'store'])->name('supplier.store')->middleware('auth:api');
// route for editing a supplier
Route::get('/supplier/{supplier}/edit', [SupplierController::class, 'edit'])->name('supplier.edit')->middleware('auth:api');
Route::put('/supplier/{supplier}/update', [SupplierController::class, 'update'])->name('supplier.update')->middleware('auth:api');
// route for searching a supplier
Route::get('/supplier/search', [SupplierController::class, 'searchSupplier'])->name('supplier.search')->middleware('auth:api');
// route for deleting a supplier
Route::delete('/supplier/{supplier}/delete', [SupplierController::class, 'destroy'])->name('supplier.delete')->middleware('auth:api');

// Invoice
// route for displaying invoice
Route::get('/invoice', [InvoiceController::class, 'index'])->name('invoice.index')->middleware('auth:api');
// route for creating an invoice
Route::get('/invoice/create', [InvoiceController::class, 'create'])->name('invoice.create')->middleware('auth:api');
Route::post('/invoice/store', [InvoiceController::class, 'store'])->name('invoice.store')->middleware('auth:api');
// route for editing an invoice
Route::get('/invoice/{invoice}/edit', [InvoiceController::class, 'edit'])->name('invoice.edit')->middleware('auth:api');
Route::put('/invoice/{invoice}/update', [InvoiceController::class, 'update'])->name('invoice.update')->middleware('auth:api');
// route for deleting an invoice
Route::delete('/invoice/{invoice}/delete', [InvoiceController::class, 'destroy'])->name('invoice.delete')->middleware('auth:api');

// route for displaying sale report
Route::get('/invoice/report', [InvoiceController::class, 'saleReport'])->name('invoice.report')->middleware('auth:api');
// route for displaying sale report profit and loss
Route::get('/invoice/profitreport', [InvoiceController::class, 'saleProfitReport'])->middleware('auth:api');
// route for displaying money in report
Route::get('/invoice/moneyinreport', [InvoiceController::class, 'moneyInReport'])->middleware('auth:api');
// route for displaying end day report
Route::get('/invoice/endDayReport', [InvoiceController::class, 'endDayReport'])->middleware('auth:api');

// Invoice details
// route for displaying invoice details
Route::get('/invoice/{invoice}', [InvoiceDetailsController::class, 'index'])->name('invoice.details.index')->middleware('auth:api');
// route for adding invoice details
Route::post('/invoice/{invoice}/store', [InvoiceDetailsController::class, 'store'])->middleware('auth:api');
// route for editing invoice details
Route::put('/invoice/{invoice}/update', [InvoiceDetailsController::class, 'update'])->name('invoice.details.update')->middleware('auth:api');
// route for deleting invoice details
Route::delete('/invoice/{invoice}/delete', [InvoiceDetailsController::class, 'destroy'])->name('invoice.details.delete')->middleware('auth:api');

// route for displaying sale report 
Route::get('/invoice/{invoice}/report', [InvoiceDetailsController::class, 'saleReport'])->name('inoice.details.report')->middleware('auth:api');

// Transaction
// route for displaying transaction
Route::get('/transaction', [TransactionController::class, 'index'])->name('transaction.index')->middleware('auth:api');
// route for creating transaction 
Route::post('/transaction/store', [TransactionController::class, 'store'])->name('transaction.store')->middleware('auth:api');
// route for updating transaction
Route::put('/transaction/{transaction}/update', [TransactionController::class, 'update'])->name('transaction.update')->middleware('auth:api');
// route for deleting transaction
Route::delete('/transaction/{transaction}/delete', [TransactionController::class, 'destroy'])->name('transaction.delete')->middleware('auth:api');
// route for getting money out report
Route::get('/transaction/moneyOutReport', [TransactionController::class, 'moneyOutReport'])->name('transaction.moneyOutReport')->middleware('auth:api');
//  route for getting product report
Route::get('/transaction/productReport', [TransactionController::class, 'productReport'])->name('transaction.productReport')->middleware('auth:api');

// Expense
// route for displaying expense
Route::get('/expense', [ExpenseController::class, 'index'])->name('expense.index')->middleware('auth:api');
// route for creating expense
Route::post('/expense/store', [ExpenseController::class, 'store'])->name('expense.store')->middleware('auth:api');
// route for updating expense
Route::put('/expense/{expense}/update', [ExpenseController::class, 'update'])->name('expense.update')->middleware('auth:api');
// route for deleting expense
Route::delete('/expense/{expense}/delete', [ExpenseController::class, 'destroy'])->name('expense.delete')->middleware('auth:api');
// route for getting expense report
Route::get('/expense/report', [ExpenseController::class, 'expenseReport'])->name('expense.report')->middleware('auth:api');
