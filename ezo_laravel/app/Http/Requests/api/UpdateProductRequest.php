<?php

namespace App\Http\Requests\api;

use Illuminate\Foundation\Http\FormRequest;

class UpdateProductRequest extends FormRequest
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
            'product_image' => 'image|mimes:jpg,jpeg,png,gif,webp|max:2048',
            'product_name' => 'required|string|max:255',
            'product_sell_price' => 'required|decimal:0,2|min:0.0',
            'measuring_unit' => 'nullable|string|max:255',
            'product_category' => 'nullable|string|max:255',
            'product_quantity' => 'required|numeric|min:0',
            'mrp' => 'required|decimal:0,2|min:0.0',
            'purchase_price' => 'nullable|decimal:0,2|min:0.0',
            'ac_sale_price' => 'nullable|decimal:0,2|min:0.0',
            'non_ac_sale_price' => 'nullable|decimal:0,2|min:0.0',
            'online_delivery_sell_price' => 'nullable|decimal:0,2|min:0.0',
            'online_sale_price' => 'nullable|decimal:0,2|min:0.0',
            'tax' => 'nullable|string|max:10',
            'price_with_tax' => 'nullable|in:Y,N',
            'cess' => 'nullable|integer',
            'hsn_code' => 'nullable|string|max:255',
            'product_description' => 'nullable|string|max:255',
            'low_stock_alert' => 'nullable|integer|min:0',
            'product_storage_location' => 'nullable|string|max:255',
            'bulk_purchase_unit' => 'nullable|string|max:255',
            'retail_sale_unit_per_bulk_purchase' => 'nullable|decimal:0,2|min:0.0',
            'bulk_purchase_unit_per_retail_price' => 'nullable|decimal:0,2|min:0.0',
            'expiry_date' => 'nullable|date',
            'show_product_online_store' => 'nullable|string|max:5',
        ];
    }
}
