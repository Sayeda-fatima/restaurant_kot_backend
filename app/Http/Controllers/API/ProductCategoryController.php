<?php

namespace App\Http\Controllers\API;

use Illuminate\Support\Facades\Redirect;
use Illuminate\Support\Facades\DB;
use App\Models\ProductCategory;
use App\Http\Requests\api\StoreProductCategoryRequest;
use App\Http\Requests\api\UpdateProductCategoryRequest;
use App\Http\Controllers\Controller;

class ProductCategoryController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $product = DB::table('product_categories')
            ->select('product_category')
            ->get();
        return response()->json([
            'message' => 'success',
            'data' => $product
        ]);
        //return view('productCategory.index', ['product'=>$product]); 
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return view('product.category');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreProductCategoryRequest $request)
    {
        try {
            $productCategory = ProductCategory::create($request->all());

            return response()->json([
                'message' => 'success',
                'data' => $productCategory
            ], 201);
        } catch (\Exception $e) {
            error_log('Error creating product category: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create product category', 'error' => $e->getMessage()], 500);
        }

        //return Redirect::route('productCategory.index')->with('success', 'Product Category added successfully');

    }

    /**
     * Display the specified resource.
     */
    public function show(ProductCategory $productCategory)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(ProductCategory $productCategory)
    {
        return view('productCategory.edit', ['productCategory' => $productCategory]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateProductCategoryRequest $request, ProductCategory $productCategory, int $id)
    {
        try {
            $productCategory = ProductCategory::find($id);
            $data = $request->all();
            $productCategory->update($data);

            return response()->json([
                'message' => 'success',
                'data' => $productCategory->fresh()
            ], 201);
        } catch (\Exception $e) {
            error_log('Error updating product category: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update product category', 'error' => $e->getMessage()], 500);
        }

        //return Redirect::route('productCategory.index')->with('success', 'Product category updated successfully');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(ProductCategory $productCategory, int $id)
    {
        try {
            $productCategory = ProductCategory::find($id);
            $productCategory->delete();

            return response()->json([
                'message' => 'success',
                'data' => $productCategory
            ]);
        } catch (\Exception $e) {
            error_log('Error creating product: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create product', 'error' => $e->getMessage()], 500);
        }

        //return Redirect::route('productCategory.index')->with('success', 'Product category deleted successfully');
    }
}
