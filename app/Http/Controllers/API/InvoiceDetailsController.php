<?php

namespace App\Http\Controllers\API;

use Database\Factories\InvoiceDetailsFactory;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
use App\Models\InvoiceDetails;
use App\Models\Product;
use App\Models\Invoice;
use App\Http\Requests\api\StoreInvoiceDetailsRequest;
use App\Http\Requests\api\UpdateInvoiceDetailsRequest;
use App\Http\Controllers\Controller;

class InvoiceDetailsController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        Gate::authorize('viewAny', InvoiceDetails::class);
        /*$invoiceDetails = DB::table('invoice_details')
                            ->join('products', 'invoice_details.product_id', '=', 'products.id')
                            ->select('invoice_details.invoice_id', 'products.product_name', 'invoice_details.quantity', 'products.product_sell_price'); //select sum(total_product_price) as total 
        */
        $invoice = DB::raw('select invoice_details.invoice_id, products.product_name, invoice_details.quantity, products.mrp, invoice_details.total_price from invoice_details left join products on invoice_details.product_id=products.id');
        return response()->json([
            'message' => 'success',
            'data' => $invoice
        ]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreInvoiceDetailsRequest $request)
    {
        // access type -> ADMIN, SALES, STAFF
        Gate::authorize('create', InvoiceDetails::class);
        $product = Product::find($request->product_id);
        $invoice = Invoice::find($request->invoice_id);
        try{
            $invoiceDetails = InvoiceDetails::create([
                'invoice_id' => $request->invoice_id,
                'product_id' => $request->product_id,
                'product_name' => $product->product_name,
                'quantity' => $request->quantity,
                'unit_product_price' => $product->mrp,
                'discount' =>$request->discount,
                'total_product_price' => ($product->mrp * $request->quantity * (1-$request->discount/100))
            ]);
            // update the total_price in invoice table 
            $invoice->total_price += $invoiceDetails->total_product_price;
            $invoice->save();
            return response()->json([
                'message' => 'success',
                'data' => $invoiceDetails
            ],201);
        }
        catch(\Exception $e){
            error_log('Error creating invoice details: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create invoice details', 'error' => $e->getMessage()], 500);
        }
    }

    /**
     * Display the specified resource.
     */
    public function show(InvoiceDetails $invoiceDetails)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(InvoiceDetails $invoiceDetails)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateInvoiceDetailsRequest $request, InvoiceDetails $invoiceDetails, int $id)
    {
        Gate::authorize('update', $invoiceDetails);
        try{
            $product = Product::find($request->product_id);
            $invoice = Invoice::find($request->invoice_id);
            $invoiceDetails = InvoiceDetails::find($id);

            $invoiceDetails->update([
                'invoice_id' => $request->invoice_id,
                'product_id' => $request->product_id,
                'product_name' => $product->product_name,
                'quantity' => $request->quantity,
                'unit_product_price' => $product->mrp,
                'total_product_price' => ($product->mrp * $request->quantity)
            ]);
            // update total price in invoice table
            $invoice->total_price = $invoice->invoiceDetails->sum('total_product_price');
            $invoice->save();

            return response()->json([
                'message' => 'success',
                'data' => $invoiceDetails->fresh()
            ],200);
        }catch (\Exception $e) {
            error_log('Error updating invoice details: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update invoice details', 'error' => $e->getMessage()], 500);
        }

    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(InvoiceDetails $invoiceDetails, Request $request, int $id)
    {
        Gate::authorize('delete', $invoiceDetails);
        $invoice = Invoice::find($request->invoice_id);
        try{
            $invoiceDetails = InvoiceDetails::find($id);
            $deletedPrice = $invoiceDetails->total_product_price;
            $invoiceDetails->delete();
            // update the total price in the invoice table
            $invoice->total_price -= $deletedPrice;
            $invoice->save();
            return response()->json([
                'message' => 'success',
                'data' => $invoiceDetails
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting invoice details: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete invoice details', 'error' => $e->getMessage()], 500);
        }
    }

    // transaction report -> invoice details -> sale report
    public function saleReport(Request $request){
        // access -> ADMIN
        Gate::authorize('view', InvoiceDetails::class);
        //$invoice = Invoice::find($id);
        $id = $request->id;
        $date1 = $request->date_from;
        $date2 = $request->date_to;
        // select * from invoice_details where invoice_id = $id and created_by between $date1 and $date2;
        $query = DB::select('SELECT date(created_at) as date, product_name, quantity, unit_product_price, total_product_price from invoice_details where invoice_id=? and created_at between ? and ? order by created_at', [$id, $date1, $date2]);
        //print_r($query);
        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }

    public function estimatePrice(Request $request){
        //$product = Product::find($request->product_id);
        $products = $request->input('products');
        $estimateData = [];
        $totalCost = 0;
        foreach($products as $productData){
            $product = Product::find($productData['product_id']);
            $data = [
                'product_id' => $request->product_id,
                'product_name' => $product->product_name,
                'quantity' => $request->quantity,
                'unit_product_price' => $product->mrp,
                'discount' => $request->discount,
                'total_product_price' => ($product->mrp * $request->quantity * $request->discount/100),
            ];
        }
        $estimateData[] = $data;
        $totalCost += $data['total_product_price'];
        return response()->json([
            'message' => 'success',
            'data' => $data,
            'total_cost' => $totalCost
        ]);
    }
}
