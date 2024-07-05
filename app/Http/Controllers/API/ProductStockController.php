<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Redirect;
use App\Models\ProductStock;
use App\Http\Requests\api\StoreProductStockRequest;
use App\Http\Requests\api\UpdateProductStockRequest;
use App\Http\Controllers\Controller;

class ProductStockController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        return view('product.stock');
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return view('productStock.create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreProductStockRequest $request)
    {
        $productStock = ProductStock::create([
            'product_id' => $request->product_id,
            'product_quantity' => $request->product_quantity,
            'product_update_type' => $request->product_update_type,
            'product_update_quantity' => $request->product_update_quantity
        ]);

        return Redirect::route('product.stock')->with('success', 'Stock added successfully');
    }

    /**
     * Display the specified resource.
     */
    public function show(ProductStock $productStock)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(ProductStock $productStock)
    {
        return view('productStock.edit');
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateProductStockRequest $request, ProductStock $productStock)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(ProductStock $productStock)
    {
        //
    }

    public function updateStock(Request $request){
        
    }
}
