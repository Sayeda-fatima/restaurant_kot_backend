<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
use App\Models\Supplier;
use App\Http\Requests\api\StoreSupplierRequest;
use App\Http\Requests\api\UpdateSupplierRequest;
use App\Http\Controllers\Controller;

class SupplierController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Request $request)
    {
        Gate::authorize('viewAny', Supplier::class);
        $organization_id = $request->organization_id;
        $supplier = DB::table('suppliers')
                        ->select('name', 'phone_no', 'billing_type')
                        ->whereRaw('organization_id=?', [$organization_id])
                        ->orderby('name')
                        ->orderby('category')
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
        Gate::authorize('create', Supplier::class);
        try{
            $supplier = Supplier::create([
                'organization_id' => $request->organization_id,
                'name' => $request->name,
                'phone_no' => $request->phone_no,
                'category' => $request->category,
                'billing_address' => $request->billing_address,
                'billing_province' => $request->billing_province,
                'billing_postal_code' => $request->billing_postal_code,
                'delivery_address' => $request->delivery_address,
                'delivery_province' => $request->delivery_province,
                'delivery_postal_code' => $request->delivery_postal_code,
                'gst_number' => $request->gst_number,
                'billing_term' => $request->billing_term,
                'billing_type' => $request->billing_type,
                'date_of_birth' => $request->date_of_birth,
                'whatsapp_alert' => $request->whatsapp_alert
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
        Gate::authorize('update', $supplier);
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
        Gate::authorize('update', $supplier);

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
        
        Gate::authorize('view', Supplier::class);

        $organization_id = $request->organization_id;
        $search = $request->get('search_term');
        if($search!=NULL){
            $supplier = Supplier::where('name', 'LIKE', "%$search%")
                                ->orwhere('id', 'LIKE', "%$search")
                                ->orwhere('phone_no', 'LIKE', "$search%")
                                ->havingRaw('organization_id=?', [$organization_id])
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
