<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
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
        // access -> ADMIN, STAFF, SALES
        Gate::authorize('viewAny', Product::class);

        // display list of products
        $product = DB::table('products')
            ->select('image','name', 'sell_price', 'quantity')
            ->orderby('name')
            ->orderby('category')
            ->get();

        return response()->json([
            'message' => 'success',
            'data' => $product
        ],200);
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
        Gate::authorize('create', Product::class);
       
        try {
            // Generate unique filename
            $imageName = time() . '_' . $request->file('image')->getClientOriginalName();  
            $image = $request->file('image')->storeAs('productImages', $imageName);  

            $product = Product::create([
                'organization_id' => $request->organization_id,
                'image' => $image,
                'name' => $request->input('name'),
                'sell_price' => $request->input('sell_price'),
                'measuring_unit' => $request->input('measuring_unit'),
                'category_id' => $request->input('category_id'),
                'quantity' => $request->input('quantity'),
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
                'description' => $request->input('description'),
                'low_stock_alert' => $request->input('low_stock_alert'),
                'storage_location' => $request->input('storage_location'),
                'bulk_purchase_unit' => $request->input('bulk_purchase_unit'),
                'retail_sale_unit_per_bulk_purchase' => $request->input('retail_sale_unit_per_bulk_purchase'),
                'bulk_purchase_unit_per_retail_sale' => $request->input('bulk_purchase_unit_per_retail_sale'),
                'expiry_date' => $request->input('expiry_date'),
                'show_product_online_store' => $request->input('show_product_online_store')
            ]);
            $product->image = $image;
            $product->save();
            return response()->json(['message' => 'Product created successfully', 'product' => $product], 201);
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
        Gate::authorize('update', $product);
        try {
            $data = $request->all();
            $product->update($data);
            return response()->json([
                'message' => 'success',
                'data' => $product->fresh()
            ],200);
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
        Gate::authorize('delete', $product);
        //delete a product
        try{
            $product->delete();
            return response()->json([
                'message' => 'success',
                'data' => $product
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting product: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete product', 'error' => $e->getMessage()], 500);
        }
        //return Redirect::route('products.index')->with('success', 'Product deleted successfully');
    }
    public function allProducts(Request $request)
    {
        Gate::authorize('viewAny', Product::class);
        $organization_id = $request->organization_id;
        //display all product for invoice
        $product = Product::whereRaw('organization_id=?', [$organization_id])
            ->orderBy('name', 'DESC')
            ->orderBy('category_id')
            ->get();
        return response()->json([
            'product' => $product
        ],201);
    }

    public function searchProduct(Request $request)
    {
        Gate::authorize('view', Product::class);

        $organization_id = $request->organization_id;
        $search = $request->get('search_term');
        if ($search != NULL) {
            $product = Product::where('id', 'LIKE', "%$search%")
                ->orwhere('name', 'LIKE', "%$search%")
                ->orwhere('description', 'LIKE', "%$search%")
                ->havingRaw('organization_id=?', [$organization_id])
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
        //Gate::authorize('view', Product::class);
        // adjust stock of a product
        $product = Product::find($request->id);
        $update_type = $request->update_type;
        $update_quantity = $request->update_quantity;
        if($update_type === 'add'){
            $product->quantity += $update_quantity;
            $product->save();
        }
        else{
            $product->quantity -= $update_quantity;
            $product->save();
        }
        return response()->json([
            'message' => 'success',
            'data' => $product->fresh()
        ]);

    }
// 
    public function displayProductsForCategory(Request $request){

        Gate::authorize('view', Product::class);

        $organization_id = $request->organization_id;
        $search = $request->get('category');
        if($search){
            $product = DB::table('products')
                        ->select('name', 'sell_price', 'quantity')
                        ->whereRaw('organization_id=?', [$organization_id])
                        ->where('category_id', '=', "$search")
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

    // product report -> product sale report
    public function productSaleReport(Request $request){
        Gate::authorize('view', Product::class);

        $organization_id = $request->organization_id;
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $query = DB::select('SELECT invoice_details.name, 
                products.category_id, product_categories.product_category,
                sum(invoice_details.quantity) as total_sale_quantity, 
                sum(total_price) as total_sale_amount 
            from invoice_details 
            left join products on invoice_details.id=products.id 
            left join product_catgories on product_categories.id=products.category_id
            where products.organization_id=? and date(invoice_details.created_at) between ? and ? 
            group by invoice_details.id', [$organization_id, $date_from, $date_to]);

        $total_data = DB::select('SELECT sum(invoice_details.quantity) as total_sale_quantity, 
                sum(total_price) as total_sale_amount 
            from invoice_details 
            left join products on invoice_details.id=products.id 
            where products.organization_id=? and date(invoice_details.created_at) between ? and ?', [$organization_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query,
            'total_data' => $total_data
        ],200);
    }

    // product report -> product stock summary report if quantity<=low stock alert
    public function productStockReport(Request $request){

        Gate::authorize('view', Product::class);

        $organization_id = $request->organization_id;
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT name, category, quantity as current_stock, mrp as sale_price, purchase_price, (quantity * mrp) as stock_valuation from products where organization_id=? and date(created_at) between ? and ?;', [$organization_id, $date_from, $date_to]);

        $total_data = DB::select('SELECT count(id) as total_unique_items, (SELECT quantity from products where quantity <= low_stock_alert) as total_low_stock_items from products where organization_id=? and date(created_at) between ? and ? ;', [$organization_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query,
            'total_data' => $total_data
        ]);
    }

    // product report -> product details report
    public function productDetailReport(Request $request){

        Gate::authorize('view', Product::class);

        $organization_id = $request->organization_id;
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $query = DB::select('SELECT * from products where organization_id=? and date(created_at) between ? and ?', [$organization_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]); 
    }
}
