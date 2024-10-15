<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
use App\Models\Expense;
use App\Models\Supplier;
use App\Http\Requests\api\StoreExpenseRequest;
use App\Http\Requests\api\UpdateExpenseRequest;
use App\Http\Controllers\Controller;

class ExpenseController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        Gate::authorize('viewAny', Expense::class);
        $expense = DB::table('expenses')
                    ->select('id', 'supplier_name', 'amount_paid')
                    ->get();
        return response()->json([
            'message' => 'success',
            'data' => $expense
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
    public function store(StoreExpenseRequest $request)
    {
        Gate::authorize('create', Expense::class);
        try{
            $supplier = Supplier::find($request->supplier_id);
            $expense = Expense::create([
                'supplier_id' => $request->supplier_id,
                'supplier_name' =>$supplier->supplier_name,
                'expense_category' => $request->expense_category,
                'total_amount' => $request->total_amount,
                'amount_paid' => $request->amount_paid,
                'amount_due' => ($request->total_amount - $request->amount_paid)
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $expense
            ],201);
        }catch(\Exception $e){
            error_log('Error creating expense: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create expense', 'error' => $e->getMessage()], 500);
        }
 
    }

    /**
     * Display the specified resource.
     */
    public function show(Expense $expense)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Expense $expense)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateExpenseRequest $request, Expense $expense)
    {
        Gate::authorize('update', $expense);
        try{
            $expense->update([
                'supplier_id' => $request->supplier_id,
                'expense_category' => $request->expense_category,
                'total_amount' => $request->total_amount,
                'amount_paid' => $request->amount_paid,
                'amount_due' => ($request->total_amount - $request->amount_paid)
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $expense->fresh()
            ]);
        }catch(\Exception $e){
            error_log('Error updating expense: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update expense', 'error' => $e->getMessage()], 500);
        }

    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Expense $expense)
    {
        Gate::authorize('delete', $expense);
        try{
            $expense -> delete();
            return response()->json([
                'message' => 'success',
                'data' => $expense
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting expense: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete expense', 'error' => $e->getMessage()], 500);
        }
    }

    public function expenseReport(Request $request){
        Gate::authorize('view', Expense::class);
        
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT date(created_at), supplier_name, expense_category, id as expnese_no, total_amount, amount_paid, mode_of_payment from expenses where date(created_at) between ? and ?;', [$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
}
