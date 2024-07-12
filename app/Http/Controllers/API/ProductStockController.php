<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Redirect;
use App\Models\Product;
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
        try{
            $product = Product::find($request->product_id);
            $productStock = ProductStock::create([
                'invoice_details_id' => $request->invoice_details_id,
                'product_id' => $request->product_id,
                'product_name' => $product->product_name,
                'product_update_quantity' => $request->product_update_quantity,
                'product_update_type' => $request->product_update_type,
                'product_stock_after_update' => $request->product_stock_after_update
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $productStock
            ],201);
        }
        //return Redirect::route('product.stock')->with('success', 'Stock added successfully');
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
        try{
            $data = $request->all();
            $productStock -> update($data);
            return response()->json([
                'message' => 'success',
                'data' => $productStock->fresh()
            ],200);
        }
        catch(\Exception $e){
            error_log('Error updating product stock: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update product stock', 'error' => $e->getMessage()], 500);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(ProductStock $productStock)
    {
        try{
            $productStock->delete();
            return response()->json([
                'message' => 'success',
                'data' => $productStock
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting product stock: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete product stock', 'error' => $e->getMessage()], 500);
        }
    }

    public function updateStock(Request $request){
        
    }

    public function stockReport(Request $request){
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $query = DB::select('SELECT * from product_stocks where date(created_at) between ? and ?', [$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
}
