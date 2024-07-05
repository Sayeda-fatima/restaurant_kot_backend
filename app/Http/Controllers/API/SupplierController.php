<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Redirect;
use App\Models\Supplier;
use App\Http\Requests\api\StoreSupplierRequest;
use App\Http\Requests\api\UpdateSupplierRequest;
use App\Http\Controllers\Controller;

class SupplierController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $supplier = DB::table('suppliers')
                        ->select('supplier_name', 'supplier_phone_no', 'supplier_billing_type')
                        ->orderby('supplier_name')
                        ->orderby('supplier_category')
                        ->get();
        //return view('supplier.index', ['supplier'=>$supplier]);
        return response()->json([
            'message' => 'success',
            'data' => $supplier
        ]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return view('supplier.index');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreSupplierRequest $request)
    {
        try{
            $supplier = Supplier::create([
                'supplier_name' => $request->supplier_name,
                'supplier_phone_no' => $request->supplier_phone_no,
                'supplier_category' => $request->supplier_category,
                'supplier_billing_address' => $request->supplier_billing_address,
                'supplier_billing_province' => $request->supplier_billing_province,
                'supplier_billing_postal_code' => $request->supplier_billing_postal_code,
                'supplier_delivery_address' => $request->supplier_delivery_address,
                'supplier_delivery_province' => $request->supplier_delivery_province,
                'supplier_delivery_postal_code' => $request->supplier_delivery_postal_code,
                'supplier_gst_number' => $request->supplier_gst_number,
                'supplier_billing_term' => $request->supplier_billing_term,
                'supplier_billing_type' => $request->supplier_billing_type,
                'supplier_date_of_birth' => $request->supplier_date_of_birth,
                'supplier_whatsapp_alert' => $request->supplier_whatsapp_alert
            ]);
            return response()->json([
                'message' => 'supplier added successfully',
                'data' => $supplier
            ]);
        }
        catch(\Exception $e){
            error_log('Error adding supplier: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to add supplier', 'error' => $e->getMessage()], 500);
        }

        //return Redirect::route('supplier.index')->with('success', 'Supplier added successfully');
    }

    /**
     * Display the specified resource.
     */
    public function show(Supplier $supplier)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Supplier $supplier)
    {
        return view('supplier.edit', ['supplier'=>$supplier]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateSupplierRequest $request, Supplier $supplier)
    {
        try{
            $data = $request->all();
            $supplier -> update($data);
            return response()->json([
                'message' => 'success',
                'data' => $supplier->fresh()
            ]);
        }
        catch(\Exception $e){
            error_log('Error updating supplier: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update supplier', 'error' => $e->getMessage()], 500);
        }
        

        //return Redirect::route('supplier.index')->with('success', 'Supplier updated successfully');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Supplier $supplier)
    {
        try{
            $supplier->delete();
            return response()->json([
                'message' => 'success',
                'data' => $supplier
            ]);
        }
        catch(\Exception $e){
            error_log('Error deleting supplier: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete supplier', 'error' => $e->getMessage()], 500);
        }
        //return Redirect::route('supplier.index')->with('success', 'Supplier deleted successfully');
    }

    public function searchSupplier(Request $request){
        $search = $request->get('search_term');
        if($search!=NULL){
            $supplier = Supplier::where('supplier_name', 'LIKE', "%$search%")
                                ->orwhere('id', 'LIKE', "%$search")
                                ->orwhere('supplier_phone_no', 'LIKE', "$search%")
                                ->get();
            return response()->json([
                'supplier' => $supplier,
                'message' => 'Supplier found'
            ],200);
        }
        else{
            return response()->json([
                'message' => 'Supplier not found'
            ],404);
        }
    }
}
