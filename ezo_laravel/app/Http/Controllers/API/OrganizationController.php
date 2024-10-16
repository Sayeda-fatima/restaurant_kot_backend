<?php

namespace App\Http\Controllers\API;

use Illuminate\Support\Facades\DB;
use App\Models\Organization;
use App\Http\Controllers\Controller;
use App\Http\Requests\api\StoreOrganizationRequest;
use App\Http\Requests\api\UpdateOrganizationRequest;

class OrganizationController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $organization = DB::table('organizations')
                            ->select('id', 'name', 'access_given')
                            ->get();

        return response()->json([
            'message' => 'success',
            'data' => $organization
        ],200);
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
    public function store(StoreOrganizationRequest $request)
    {
        try{
            $organization = Organization::create([
                'name' => $request->name,
                'access_given' => $request->access_given
            ]);

            return response()->json([
                'message' => 'success',
                'data' => $organization
            ],201);
        }catch(\Exception $e){
            error_log('Error adding organization: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to add organization', 'error' => $e->getMessage()], 500);
        }
    }

    /**
     * Display the specified resource.
     */
    public function show(Organization $organization)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Organization $organization)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateOrganizationRequest $request, Organization $organization)
    {
        try{
            $data = $request->all();
            $organization->update($data);
            return response()->json([
                'message' => 'success',
                'data' => $organization->fresh()
            ],200);
        }catch(\Exception $e){
            error_log('Error updating organization: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update organization', 'error' => $e->getMessage()], 500);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Organization $organization)
    {
        try{
            $organization->delete();

            return response()->json([
                'message' => 'success',
                'data' => $organization
            ],200);
        }catch(\Exception $e){
            error_log('Error deleting organization: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete organization', 'error' => $e->getMessage()], 500);
        }
    }
}
