<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Redirect;
use App\Models\Product;
use App\Http\Requests\api\StoreProductRequest;
use App\Http\Requests\api\UpdateProductRequest;
use App\Http\Controllers\Controller;


class ProductController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {


        // display list of products
        $product = DB::table('products')
            ->select('product_name', 'product_sell_price', 'product_quantity')
            ->orderby('product_name')
            ->orderby('product_category')
            ->get();

        //return view('product.index', ['product'=> $product]);
        return response()->json(['products' => $product]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        //creating a new product

        return view('product.create');

    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreProductRequest $request)
    {
        try {
            $product = Product::create([
                'product_name' => $request->input('product_name'),
                'product_sell_price' => $request->input('product_sell_price'),
                'measuring_unit' => $request->input('measuring_unit'),
                'product_category' => $request->input('product_category'),
                'product_quantity' => $request->input('product_quantity'),
                'mrp' => $request->input('mrp'),
                'purchase_price' => $request->input('purchase_price'),
                'ac_sale_price' => $request->input('ac_sale_price'),
                'non_ac_sale_price' => $request->input('non_ac_sale_price'),
                'online_delivery_sell_price' => $request->input('online_delivery_sell_price'),
                'online_sell_price' => $request->input('online_sell_price'),
                'tax' => $request->input('tax'),
                'price_with_tax' => $request->input('price_with_tax'),
                'cess' => $request->input('cess'),
                'hsn_code' => $request->input('hsn_code'),
                'product_description' => $request->input('product_description'),
                'low_stock_alert' => $request->input('low_stock_alert'),
                'product_storage_location' => $request->input('product_storage_location'),
                'bulk_purchase_unit' => $request->input('bulk_purchase_unit'),
                'retail_sale_unit_per_bulk_purchase' => $request->input('retail_sale_unit_per_bulk_purchase'),
                'bulk_purchase_unit_per_retail_sale' => $request->input('bulk_purchase_unit_per_retail_sale'),
                'expiry_date' => $request->input('expiry_date'),
                'show_product_online_store' => $request->input('show_product_online_store')
            ]);
            return response()->json(['product' => $product, 'message' => 'Product created successfully'], 201);
        } catch (\Exception $e) {
            error_log('Error creating product: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create product', 'error' => $e->getMessage()], 500);
        }
    }
    /**
     * Display the specified resource.
     */
    public function show(Product $product)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Product $product)
    {
        return view('product.edit', ['product' => $product]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateProductRequest $request, Product $product)
    {
        try {
            $data = $request->all();
            print_r($product);
            $product->update($data);
            return response()->json([
                'message' => 'success',
                'data' => $product->fresh()
            ],201);
        }
        catch (\Exception $e) {
            error_log('Error updating product: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update product', 'error' => $e->getMessage()], 500);
        }

        
        //return Redirect::route('product.index')->with('success', 'Product updated successfully');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Product $product)
    {
        //delete a product
        try{
            $product->delete();
            return response()->json([
                'message' => 'success',
                'data' => $product
            ],201);
        }
        catch(\Exception $e){
            error_log('Error deleting product: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete product', 'error' => $e->getMessage()], 500);
        }
        //return Redirect::route('products.index')->with('success', 'Product deleted successfully');
    }
    public function allProducts()
    {
        //display all product for invoive
        $product = Product::orderBy('product_name', 'DESC')
            ->orderBy('product_category')
            ->get();
        return response()->json([
            'product' => $product
        ],201);
    }

    public function searchProduct(Request $request)
    {
        $search = $request->get('search_term');
        if ($search != NULL) {
            $product = Product::where('id', 'LIKE', "%$search%")
                ->orwhere('product_name', 'LIKE', "%$search%")
                ->orwhere('product_description', 'LIKE', "%$search%")
                ->get()
                ->paginate(25);
            return response()->json([
                'product' => $product,
                'message' => 'Products found'
            ], 201);
        } else {
            return response()->json([
                'message' => 'No products found'
            ], 404);
        }
    }

    public function updateStock(Request $request)
    {
        $product = $request->get('id');
        $product_quantity = DB::table('products')
            ->where('product_id', 'LIKE', "$product")
            ->update(['product_quantity' => 2]);

    }

    public function displayProductsForCategory(Request $request){
        $search = $request->get('product_category');
        if(!$search){
            $product = DB::table('products')
                        ->select('product_name', 'product_sell_price', 'product_quantity')
                        ->where('product_category', '=', "$search")
                        ->get();
            return response()->json([
                'message' => 'success',
                'data' => $product
            ],201);
        }
        else{
            return response()->json([
                'message' => 'No products in this category'
            ],404);
        }
    }
}
