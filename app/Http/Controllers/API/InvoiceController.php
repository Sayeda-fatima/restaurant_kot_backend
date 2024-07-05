<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Redirect;
use App\Models\Invoice;
use App\Http\Requests\api\StoreInvoiceRequest;
use App\Http\Requests\api\UpdateInvoiceRequest;
use App\Http\Controllers\Controller;

class InvoiceController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $invoice = DB::table('invoices')
                    ->select('id', 'customer_id', 'total_price', 'mode_of_payment')
                    ->get()
                    ->paginate(25);
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
        return view('invoice.create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreInvoiceRequest $request)
    {
        try{
            $invoice = Invoice::create([
                'customer_id' => $request->customer_id,
                'customer_name' => $request->customer_name,
                'total_price' => $request->total_price,
                'billing_address' => $request->billing_address,
                'mode_of_payment' => $request->mode_of_payment
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $invoice
            ]);
        }
        catch(\Exception $e){
            error_log('Error creating invoice: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create invoice', 'error' => $e->getMessage()], 500);
        }
        //return Redirect::route('invoice.index')->with('success', 'Invoice generated successfully');
    }

    /**
     * Display the specified resource.
     */
    public function show(Invoice $invoice)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Invoice $invoice)
    {
        return view('invoice.edit');
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateInvoiceRequest $request, Invoice $invoice)
    {
        try{
            $data = $request->all();
            $invoice->update($data);
            return response()->json([
                'message' => 'success',
                'data' => $invoice->fresh()
            ]);
        }
        catch(\Exception $e){
            error_log('Error updating invoice: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update invoice', 'error' => $e->getMessage()], 500);
        }
        
        //return Redirect::route('invoice.index')->with('success', 'Invoice edited successfully');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Invoice $invoice)
    {
        try{
            $invoice -> delete();
            return response()->json([
                'message' => 'success',
                'data' => $invoice
            ]);
        }
        catch(\Exception $e){
            error_log('Error deleting invoice: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete invoice', 'error' => $e->getMessage()], 500);
        }
        
    }
}
