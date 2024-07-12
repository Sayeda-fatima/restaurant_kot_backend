<?php

namespace App\Http\Requests\api;

use Illuminate\Foundation\Http\FormRequest;

class StoreInvoiceDetailsRequest extends FormRequest
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
            'invoice_id' => 'required|integer|exists:invoices,id',
            'product_id' => 'required|integer|exists:products,id',
            'quantity' => 'required|integer',
            'discount' => 'nullable|decimal:0,2|min:0'
        ];
    }
}
