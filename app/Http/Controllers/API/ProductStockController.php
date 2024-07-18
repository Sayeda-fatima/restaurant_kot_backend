<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
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
        Gate::authorize('viewAny', ProductStock::class);
        $productStock = DB::table('product_stocks')
                        ->select('invoice_details_id', 'product_name', 'product_stock_before_update', 'product_update_quantity', 'product_stock_after_update')
                        ->orderby('invoice_details_id')
                        ->get();
        return response()->json([
            'message' => 'success',
            'data' => $productStock
        ]);
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
        Gate::authorize('create', ProductStock::class);
        try{
            $product = Product::find($request->product_id);

            // Calculate the product stock after update based on the update type
            $productStockAfterUpdate = ($request->product_update_type == 'add')
                ? $product->product_quantity + $request->product_update_quantity
                : $product->product_quantity - $request->product_update_quantity;

            $productStock = ProductStock::create([
                'invoice_details_id' => $request->invoice_details_id,
                'product_id' => $request->product_id,
                'product_name' => $product->product_name,
                'product_stock_before_update' => $product->product_quantity,
                'product_update_quantity' => $request->product_update_quantity,
                'product_update_type' => $request->product_update_type,
                'product_stock_after_update' => $productStockAfterUpdate
            ]);

            // update quantity in the products table
            $product->product_quantity = $productStock->product_stock_after_update;

            return response()->json([
                'message' => 'success',
                'data' => $productStock
            ],201);

        } catch (\Exception $e) {
            error_log('Error creating product stock: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create product stock', 'error' => $e->getMessage()], 500);
        }
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
        Gate::authorize('update', $productStock);
        try{
            $product = Product::find($request->product_id);
            $data = $request->all();
            $productStock -> update([
                'invoice_details_id' => $request->invoice_details_id,
                'product_id' => $request->product_id,
                'product_name' => $product->product_name,
                'product_stock_before_update' => $product->product_quantity,
                'product_update_quantity' => $request->product_update_quantity,
                'product_update_type' => $request->product_update_type,
                'product_stock_after_update' => ($request->product_update_type == 'add') ? $product->product_quantity + $request->product_update_quantity : $product->product_quantity - $request->product_update_quantity
            ]);

            // update quantity in the products table
            $product->product_quantity = $productStock->product_stock_after_update;

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
        Gate::authorize('delete', $productStock);
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


    public function stockReport(Request $request){
        Gate::authorize('view', ProductStock::class);
        
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $query = DB::select('SELECT * from product_stocks where date(created_at) between ? and ?', [$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
}
