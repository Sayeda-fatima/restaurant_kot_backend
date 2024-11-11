<?php

namespace Database\Factories;

use App\Models\Organization;
use App\Models\ProductCategory;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Product>
 */
class ProductFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'organization_id' => 1,//Organization::factory(),
            'image' => $this->faker->url(),
            'name' => $this->faker->name,
            'sell_price' => $this->faker->numberBetween(2, 24),
            'measuring_unit' => $this->faker->word(),
            'category_id' => 3,//ProductCategory::factory(),
            'quantity' => $this->faker->numberBetween(2, 60000000),
            'mrp' => $this->faker->numberBetween(1, 9999999999),
            'purchase_price' => $this->faker->numberBetween(1, 9999999999),
            'ac_sale_price' => $this->faker->numberBetween(1, 9999999999),
            'non_ac_sale_price' => $this->faker->numberBetween(1, 9999999999),
            'online_delivery_sell_price' => $this->faker->numberBetween(1, 9999999999),
            'online_sell_price' => $this->faker->numberBetween(1, 9999999999),
            'tax' => $this->faker->numberBetween(1, 18),
            'price_with_tax' => 'Y',//$this->faker->words('Y', 'N'),
            'cess' => $this->faker->numberBetween(1, 9999999999),
            'hsn_code' => $this->faker->numberBetween(1, 9999999999),
            'description' => $this->faker->sentence(),
            'low_stock_alert' => 'Y',//$this->faker->words('Y', 'N'),
            'storage_location' => $this->faker->address(),
            'bulk_purchase_unit' => $this->faker->word(),
            'retail_sale_unit_per_bulk_purchase' => $this->faker->numberBetween(1, 9999999999),
            'bulk_purchase_unit_per_retail_sale' => $this->faker->numberBetween(1, 9999999999),
            'expiry_date' => $this->faker->date(),
            'show_product_online_store' => 'Y',
        ];
    }
}
