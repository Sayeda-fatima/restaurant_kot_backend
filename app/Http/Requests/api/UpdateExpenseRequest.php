<?php

namespace App\Http\Requests\api;

use Illuminate\Foundation\Http\FormRequest;

class UpdateExpenseRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'supplier_id' => 'required|integer|exists:suppliers,id',
            'expense_category' => 'required|string|max:255',
            'total_amount' => 'required|decimal:0,2|min:0',
            'amount_paid' => 'required|decimal:0,2|min:0',
            'mode_of_payment' => 'required|in:bank,cash,cheque'
        ];
    }
}
